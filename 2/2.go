package main

import "fmt"

func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}
}

func main() {
	f := fibonacci()
	count := 0
	for {
		x := f()
		if x > 4000000 {
			break
		}
		if x%2 == 0 {
			count += x
		}
	}

	fmt.Println(count) // 4613732
}
