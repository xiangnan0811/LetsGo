package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)
	var values []int

	tm := time.After(time.Second * 10)
	tick := time.Tick(time.Second)
	for {
		var activateWorker chan<- int
		var activateValue int
		if len(values) > 0 {
			activateWorker = worker
			activateValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activateWorker <- activateValue:
			values = values[1:]

		case <-time.After(time.Millisecond * 800):
			fmt.Println("Timeout")
		case <-tick:
			fmt.Println("queue len =", len(values))
		case <-tm:
			fmt.Println("Bye")
			return
		}
	}
}
