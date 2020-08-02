package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) { // 并发执行
			for {
				//fmt.Printf("Hello from goroutine %d\n", i)
				a[i]++
				runtime.Gosched()
			}

		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}

// 协程 Coroutine
/*
1. 轻量级“线程”
2. 非抢占式多任务处理， 由协程主动交出控制权
3. 编译器/解释器/虚拟机层面的多任务
4. 多个协程可能在一个或多个线程上运行
*/

// goroutine可能的切换点
/*
1. I/O,select
2. channel
3. 等待锁
4. 函数调用（有时）
5. runtime.Gosched()
*/
