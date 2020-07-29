package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sumArgs(numbers ...int) int {
	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}

func main() {
	if result, err := eval(3, 4, "l"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	q, r := div(19, 5)
	fmt.Printf("19 div 5, q is %d, r is %d\n", q, r)
	result := apply(pow, 3, 5)
	fmt.Println(result)
	fmt.Println(sumArgs(1, 2, 3, 4, 5))
}
