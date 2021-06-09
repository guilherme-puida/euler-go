package main

import (
	"fmt"
	"time"
)

func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func main() {
	start := time.Now()
	ch := make(chan int)
	go Generate(ch)
	for i := 1; i <= 10000; i++ {
		prime := <-ch
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}

	prime := <-ch
	fmt.Println(prime, time.Since(start)) // 104743, 8.85683667s
}
