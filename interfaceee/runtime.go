package interfaceee

import (
	"fmt"
	"runtime/metrics"
)

// GC调优
// GC的5个阶段 1.准备 2.开始 3.标记 4.清理 5.结束
// 那个阶段发生STW ？
// 开启写屏障时发生STW
// 插入写屏障(A->B，新增A->C，把B标记成灰色)
// 删除写屏障(A->B，新增A->C，把B标记成灰色)
// 混合写屏障
// cpu占用率很高 go进程占用很少内存，确频繁触发GC，是什么原因？cpu是干嘛的？什么时候触发GC？如何发现频繁触发GC？
// 内存泄漏
// 内存逃逸

// ForQuote 循环引用
func ForQuote() {

}

func readGC() {
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
