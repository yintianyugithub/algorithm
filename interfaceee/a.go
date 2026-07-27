package interfaceee

import (
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
