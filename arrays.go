package main

import "fmt"

func initialArray() {
	// 数组定义
	var arr1 [5]int
	arr2 := [3]int{1, 5, 8}
	arr3 := [...]int{2, 5, 7, 8, 0} //  编译器决定数组长度
	var grid [4][5]int              // 二维数组
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)
	// 数值遍历
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	for i := range arr3 {
		fmt.Println(arr3[i])
	}
	maxi := -1
	maxValue := -1
	for i, v := range arr3 {
		if maxValue < v {
			maxi, maxValue = i, v
		}
		fmt.Println(maxi, maxValue)
	}

	// [3]int和[5]int是不同数据类型
	// 数组是值类型
	fmt.Println("printArray(arr1)")
	printArray(&arr1)

	fmt.Println("printArray(arr3)")
	printArray(&arr3)

	fmt.Println("arr1 and arr3")
	fmt.Println(arr1, arr3)

	// 切片后
	fmt.Println("printArray2(arr1)")
	printArray2(arr1)

	fmt.Println("printArray2(arr3)")
	printArray2(arr3)

	fmt.Println("arr1 and arr3")
	fmt.Println(arr1, arr3)
}

func printArray(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArray2(arr [5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	initialArray()
}
