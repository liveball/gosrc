package main

import (
	"gosrc/module/concurrentspider/engine"
	"gosrc/module/concurrentspider/model"
	"gosrc/module/concurrentspider/parser"
	"gosrc/module/concurrentspider/scheduler"
)

const (
	zhenaiURL = "http://www.zhenai.com/zhenghun"
)

func main() {
	ce := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	ce.Run(model.Request{
		URL:       zhenaiURL,
		ParseFunc: parser.ParseCityList,
	})
	return
}
