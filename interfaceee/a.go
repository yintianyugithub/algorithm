package interfaceee

import (
	"fmt"
	"strings"
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
