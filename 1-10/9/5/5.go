package main

import "fmt"

func main() {
	for i := 20; ; i += 20 {
		v := true
		for j := 2; j <= 20; j++ {
			if i%j != 0 {
				v = false
			}
		}
		if v {
			fmt.Println(i) // 232792560
			break
		}
	}
}
