package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	var max, maxN int = 0, 0

	for i := 1; i < 1000000; i++ {
		var cur, counter int = i, 1
		for cur != 1 {
			if cur%2 == 0 {
				cur /= 2
			} else {
				cur = 3*cur + 1
			}
			counter++
		}

		if counter > max {
			max = counter
			maxN = i
		}
	}

	fmt.Println(maxN, time.Since(start)) // 837799, 224.189988ms
}
