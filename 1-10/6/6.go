package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	var a, b int = 0, 0
	for i := 1; i <= 100; i++ {
		a += i
		b += i * i
	}
	a = a * a
	fmt.Println(a-b, time.Since(start)) // 25164150, 134ns
}
