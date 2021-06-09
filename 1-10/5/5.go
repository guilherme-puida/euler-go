package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 20; ; i += 20 {
		v := true
		for j := 2; j <= 20; j++ {
			if i%j != 0 {
				v = false
			}
		}
		if v {
			fmt.Println(i, time.Since(start)) // 232792560, 1.702244523s
			break
		}
	}
}
