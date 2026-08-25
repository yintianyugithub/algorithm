package interfaceee

import (
	"context"
	"fmt"
	"net/http"
	"runtime/metrics"
	"time"
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

// go内存泄漏：分配的内存不能被GC,导致内存持续增长。pprof heap持续增长 + goroutine profile 协程泄漏
// go内存逃逸：分配的内存超过了自身的作用域，栈上的空间被分配到堆上。导致GC压力，cpu吞吐降低；go build -gcflag="m" [program]
// 生产环境：
// Prometheus + Grafana + Pyroscope
// 网络轮询器epoll和kqueue实现原理
// gc初始化、开始、标记、清理、结束
// 内存泄漏
var (
	oom1Slice = make([][]int, 0)
)

// oom1 无限写入
func oom1() {
	tmp := make([]int, 0)
	for {
		tmp = append(tmp, 1)
		oom1Slice = append(oom1Slice, tmp)
		time.Sleep(time.Microsecond * 20)
	}
}

// oom2 http响应body未关闭
func oom2() {
	rsp, _ := http.Get("https://www.google.com")

	//defer rsp.Body.Close()

	fmt.Println(rsp.StatusCode)
}

// 协程泄漏
// oom3 goroutine阻塞导致协程泄漏(连带内存泄漏)
func oom3() {
	ch := make(chan struct{})
	go func() {
		ch <- struct{}{}
	}()
}

// oom3Fix 使用带缓冲的channel 或 context限制超时时间 或 有人接收
func oom3Fix(ctx context.Context) {
	ch := make(chan struct{}, 1)
	go func() {
		select {
		case ch <- struct{}{}:
		case <-ctx.Done():
		}
	}()
	<-ch
}

// oom4 slice引用底层大数组，内存不能被GC释放
func oom4() []byte {
	big := make([]byte, 1024*1024*100)

	small := big[:10]
	small[0] = 1

	fmt.Println(big[0], big[1], big[2], big[3])

	return small
}

// 内存逃逸

// escape1 返回局部变量的指针
func escape1() *int {
	x := 1
	return &x
}

type es struct {
	Name string
}

func escape2() *es {
	return &es{"name is escape"}
}
