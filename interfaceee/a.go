package interfaceee

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

// SimplePath 简化路径
func SimplePath() {
	s := "/a/b/c"
	array := make([]string, 0)
	fmt.Println(cap(array))
	for _, v := range strings.Split(s, "/") {
		if v == "" || v == "." {
			continue
		}
		if v != ".." {
			fmt.Printf("%p/%d\n", array, cap(array))
			array = append(array, v) // 为啥发生扩容后地址没变？开始打印做对象了(打印的&array是变量地址肯定不变)，只要新slice需要的容量比原来大，就会发生扩容申请新内存地址的数组来存储新值
			fmt.Printf("%p/%d\n", array, cap(array))
		} else if len(array) > 0 {
			array = array[:len(array)-1]
		}
	}
	fmt.Println("/" + strings.Join(array, "/"))
}

var (
	pool = make(chan struct{}, 10)
)

// msgMode消费组件
func msgMode() {
	fins := make(chan struct{})
	fine := make(chan struct{})
	idx := 0
	for {
		idx++
		select {
		case <-fins:
			fine <- struct{}{}
			return
		default:
			test(idx)
		}
	}
}

func test(i int) {
	pool <- struct{}{}
	go func() {
		defer func() {
			<-pool
		}()
		fmt.Println("test", i)
		time.Sleep(time.Minute)
	}()
}

// evenOddItem 并发打印奇偶数
type evenOddItem struct {
	n             int
	evenCh, oddCh chan struct{}
	wg            *sync.WaitGroup
}

func (e *evenOddItem) odd() {
	defer e.wg.Done()
	for v := range e.n {
		if (v+1)%2 == 0 {
			<-e.oddCh
			fmt.Println("odd", v+1)
			e.evenCh <- struct{}{}
		}
	}
}

func (e *evenOddItem) even() {
	defer e.wg.Done()
	for v := range e.n {
		if (v+1)%2 == 1 {
			<-e.evenCh
			fmt.Println("even", v+1)
			if v+1 == e.n-1 {
				return
			}
			e.oddCh <- struct{}{}
		}
	}
}

func EvenOdd() {
	e := &evenOddItem{
		wg:     &sync.WaitGroup{},
		evenCh: make(chan struct{}),
		oddCh:  make(chan struct{}),
		n:      10,
	}

	e.wg.Add(2)

	go e.odd()
	go e.even()
	e.oddCh <- struct{}{}

	e.wg.Wait()
}

// Candy 分发糖果
func Candy(rating []int) int {
	n := 0

	if len(rating) == 0 {
		return n
	}

	res := make([]int, len(rating))

	for i := range res {
		res[i] = 1
	}

	for i := 1; i < len(rating); i++ {
		if rating[i] > rating[i-1] {
			res[i] = res[i-1] + 1
		}
	}

	for i := len(rating) - 2; i >= 0; i-- {
		if rating[i] > rating[i+1] {
			res[i] = max(res[i+1]+1, res[i])
		}
	}

	for _, v := range res {
		n += v
	}

	return n
}

// Mg 生产/消费者模型
type Mg struct {
	val     chan int
	ctx     context.Context
	workers chan struct{}
}

func (m *Mg) producer() {
	idx := 0
	for {
		idx++
		select {
		case <-m.ctx.Done():
			fmt.Println("producer done")
			// 发送方明确通知发送完成，关闭channel
			// close(m.val)
			return
		case m.val <- idx:
		}
	}
}

func (m *Mg) consumer() {
	go func() {
		for n := range m.val {
			m.workers <- struct{}{}
			go func() {
				defer func() {
					<-m.workers
				}()
				time.Sleep(time.Second)
				fmt.Println("consumer", n)
			}()
		}
	}()
}

func MgLaunch() {
	mg := &Mg{
		val:     make(chan int),
		ctx:     context.TODO(),
		workers: make(chan struct{}, 3),
	}

	ctx, cancel := context.WithTimeout(mg.ctx, time.Second*3)
	defer cancel()
	mg.ctx = ctx

	go mg.producer()
	mg.consumer()

	select {
	case <-mg.ctx.Done():
		fmt.Println("main done")
	}
}

// closedChannel 模拟协程关闭
func closedChannel() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	// 协程关闭后可以接收，但range 会阻塞；需要开启协程，或发送方主动关闭channel
	//go func() {
	//	for v := range ch {
	//		fmt.Println(v)
	//	}
	//}()

	// 协程关闭后可以接收，接受对应类型的零值，但不能发送（此处开协程方便演示，否则循环一直接受0）
	go func() {
		for {
			fmt.Println("for>", <-ch)
		}
	}()

	time.Sleep(time.Microsecond * 100)
	fmt.Println("done")

}

func Trap() {
	waters := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	// 用左右两数组记录当前索引i在左右两边能达到的最大高度
	l, r := make([]int, len(waters)), make([]int, len(waters))
	for i := 1; i < len(l); i++ {
		l[i] = max(l[i-1], waters[i-1])
	}
	for i := len(r) - 2; i >= 0; i-- {
		r[i] = max(r[i+1], waters[i+1])
	}

	// 位置i能的最大水位是左右两边的最小水位，减去当前位置的高度就是当前位置的存水量
	num := 0
	for i, v := range waters {
		m := min(l[i], r[i])
		if m > v {
			num += m - v
		}
	}

	fmt.Println(num)
}
