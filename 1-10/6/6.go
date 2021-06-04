package main

import "fmt"

func main() {
	var a, b int = 0, 0
	for i := 1; i <= 100; i++ {
		a += i
		b += i * i
	}
	a = a * a
	fmt.Println(a - b) // 25164150
}
