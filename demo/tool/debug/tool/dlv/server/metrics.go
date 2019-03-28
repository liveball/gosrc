package main

import (
	"log"
	"time"
)

const (
	second = iota // scopes
	minute
	numScopes
)

// update intervals for different metrics scopes
var updateIntervals = []time.Duration{
	time.Second,
	time.Minute,
}

// Capacity of in/out channels, that is "big enough to never be reached".
// Note, we use intentionally small value to simplify the demo.
// 1100000 could be a good choice but it didn't.
const channelCapacity = 100

type MetricType int8

const (
	TypeUnknown MetricType = iota
	TypeCount
	TypeSet
)

type Metric struct {
	Type  MetricType
	Scope int

	Key       string
	Value     int
	CreatedAt time.Time
}

func NewCountMetric(key string, value, scope int) Metric {
	return Metric{
		Type:  TypeCount,
		Scope: scope,

		Key:       key,
		Value:     value,
		CreatedAt: time.Now(),
	}
}

func NewSetMetric(key string, value, scope int) Metric {
	return Metric{
		Type:  TypeSet,
		Scope: scope,

		Key:       key,
		Value:     value,
		CreatedAt: time.Now(),
	}
}

type Metrics struct {
	inChannel  chan Metric
	outChannel chan Metric

	// one for each scope
	counts      []map[string]int
	nextUpdates []time.Time
}

func NewMetrics() *Metrics {
	metrics := &Metrics{
		inChannel:   make(chan Metric, channelCapacity),
		outChannel:  make(chan Metric, channelCapacity),
		counts:      make([]map[string]int, numScopes),
		nextUpdates: make([]time.Time, numScopes),
	}

	now := time.Now().UTC()
	for scope := 0; scope < numScopes; scope++ {
		metrics.counts[scope] = map[string]int{}
		metrics.nextUpdates[scope] = now.Add(updateIntervals[scope])
	}

	go metrics.startInChannelConsumer()
	go metrics.startOutChannelConsumer()

	return metrics
}

// CountS increments counter per second.
func (m *Metrics) CountS(key string) {
	m.inChannel <- NewCountMetric(key, 1, second)
}

// CountM increments counter per minute.
func (m *Metrics) CountM(key string) {
	m.inChannel <- NewCountMetric(key, 1, minute)
}

// SetM stores last value within minute.
func (m *Metrics) SetM(key string, value int) {
	m.inChannel <- NewSetMetric(key, value, minute)
}

// starts a consumer for inChannel
func (m *Metrics) startInChannelConsumer() {
	for inMetrics := range m.inChannel {
		switch inMetrics.Type {
		case TypeCount:
			if _, ok := m.counts[inMetrics.Scope][inMetrics.Key]; !ok {
				m.counts[inMetrics.Scope][inMetrics.Key] = 0
			}
			m.counts[inMetrics.Scope][inMetrics.Key] += inMetrics.Value
		case TypeSet:
			if _, ok := m.counts[inMetrics.Scope][inMetrics.Key]; !ok {
				m.counts[inMetrics.Scope][inMetrics.Key] = 0
			}
			m.counts[inMetrics.Scope][inMetrics.Key] = inMetrics.Value
		}

		// send a metric for scope if createdAt is after nextUpdate
		nextUpdate := m.nextUpdates[inMetrics.Scope]
		if inMetrics.CreatedAt.After(nextUpdate) {
			m.sendMetricsToOutChannel(inMetrics.Scope)
			updateInterval := updateIntervals[inMetrics.Scope]
			for inMetrics.CreatedAt.After(nextUpdate) {
				nextUpdate = nextUpdate.Add(updateInterval)
			}
			m.nextUpdates[inMetrics.Scope] = nextUpdate
		}
	}
}

func (m *Metrics) sendMetricsToOutChannel(scope int) {
	nextUpdate := m.nextUpdates[scope]
	updateInterval := updateIntervals[scope]
	timestamp := nextUpdate.Truncate(updateInterval)

	m.sendCountMetrics(scope, timestamp)

	// add internal per minute self-statistics for every per second metric
	if scope == second {
		m.SetM("metrics.raw_channel", len(m.inChannel))
		m.SetM("metrics.metric_channel", len(m.outChannel))
	}
}

func (m *Metrics) sendCountMetrics(scope int, timestamp time.Time) {
	for key, value := range m.counts[scope] {
		if value == 0 {
			continue
		}

		m.addMetricsToOutChannel(key, value, timestamp)

		m.counts[scope][key] = 0
	}
}

func (m *Metrics) addMetricsToOutChannel(key string, value int, timestamp time.Time) {
	metric := Metric{
		Key:       key,
		Value:     value,
		CreatedAt: timestamp,
	}

	select {
	case m.outChannel <- metric:
	default:
		// note, a race in len(m.outChannel), let's ignore it for now
		log.Printf("metrics channel too big (%d). Dropping (%s)", len(m.outChannel), metric)
	}
}

func (m *Metrics) startOutChannelConsumer() {
	for metric := range m.outChannel {
		log.Printf("%s %d %d\n", metric.Key, metric.Value, metric.CreatedAt.Unix())
	}
}
