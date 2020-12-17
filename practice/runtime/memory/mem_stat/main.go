package main

import (
	"fmt"
	"runtime"
	"sort"
)

func main() {
	ms := getMemStats()

	fmt.Println("mem.heap_objects", int64(ms.HeapObjects))
	fmt.Println("mem.heap_idle_bytes", int64(ms.HeapIdleBytes))
	fmt.Println("mem.heap_in_use_bytes", int64(ms.HeapInUseBytes))
	fmt.Println("mem.heap_released_bytes", int64(ms.HeapReleasedBytes))
	fmt.Println("mem.gc_pause_usec_100", int64(ms.GCPauseUsec100))
	fmt.Println("mem.gc_pause_usec_99", int64(ms.GCPauseUsec99))
	fmt.Println("mem.gc_pause_usec_95", int64(ms.GCPauseUsec95))
	fmt.Println("mem.next_gc_bytes", int64(ms.NextGCBytes))

	//client.Incr("mem.gc_runs", int64(ms.GCTotalRuns-lastMemStats.GCTotalRuns))
}

type memStats struct {
	HeapObjects       uint64 `json:"heap_objects"`
	HeapIdleBytes     uint64 `json:"heap_idle_bytes"`
	HeapInUseBytes    uint64 `json:"heap_in_use_bytes"`
	HeapReleasedBytes uint64 `json:"heap_released_bytes"`
	GCPauseUsec100    uint64 `json:"gc_pause_usec_100"`
	GCPauseUsec99     uint64 `json:"gc_pause_usec_99"`
	GCPauseUsec95     uint64 `json:"gc_pause_usec_95"`
	NextGCBytes       uint64 `json:"next_gc_bytes"`
	GCTotalRuns       uint32 `json:"gc_total_runs"`
}

func getMemStats() memStats {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)

	// sort the GC pause array
	length := len(ms.PauseNs)
	if int(ms.NumGC) < length {
		length = int(ms.NumGC)
	}

	gcPauses := make([]uint64, length)
	copy(gcPauses, ms.PauseNs[:length])

	sort.Slice(gcPauses, func(i, j int) bool {
		return gcPauses[i] > gcPauses[j]
	})

	return memStats{
		HeapObjects:       ms.HeapObjects,
		HeapIdleBytes:     ms.HeapIdle,
		HeapInUseBytes:    ms.HeapInuse,
		HeapReleasedBytes: ms.HeapReleased,
		//GCPauseUsec100:percentile(100.0, gcPauses, len(gcPauses)) / 1000,
		//GCPauseUsec99:percentile(99.0, gcPauses, len(gcPauses)) / 1000,
		//GCPauseUsec95:percentile(95.0, gcPauses, len(gcPauses)) / 1000,
		NextGCBytes: ms.NextGC,
		GCTotalRuns: ms.NumGC,
	}

}
