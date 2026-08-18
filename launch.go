package main

import (
	"fmt"
	"runtime/metrics"
)

func main() {
	samples := []metrics.Sample{
		{Name: "/gc/heap/live:bytes"},        // 上次GC存活的堆内存
		{Name: "/gc/heap/goal:bytes"},        // 目标GC的堆内存
		{Name: "/gc/gogc:percent"},           // GC的阈值触发时机
		{Name: "/gc/cycles/total:gc-cycles"}, // 周期内GC的总次数
		{Name: "/gc/heap/allocs:bytes"},      // 堆内存分配
	}

	metrics.Read(samples)

	for _, s := range samples {
		fmt.Println(s.Name, s.Value)
	}
}
