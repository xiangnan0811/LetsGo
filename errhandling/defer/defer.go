package main

import (
	"bufio"
	"fmt"
	"os"

	"imooc.com/xiangnan/learngo/functional/fib"
)

func tryDefer() {
	defer fmt.Println(1) // 栈
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occurred")
	fmt.Println(4)
}

func tryDefer2() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile(filename string) { // 资源管理与异常处理
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n", pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	// defer调用
	/*
		1. 确保调用在函数结束时发生
		2. 参数在defer语句时计算
		3. defer列表为后进先出
	*/

	// 何时使用defer调用
	/*
		1. Open/Close
		2. Lock/Unlock
		3. PrintHeader/PrintFooter
	*/
	writeFile("fib.txt")
	//tryDefer2()
}
