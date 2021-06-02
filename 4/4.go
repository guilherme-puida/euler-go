package main

import (
	"fmt"
	"strconv"
)

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(rs)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}

	return string(rs)
}

func main() {
	ans := 0

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			cur := i * j
			strCur := strconv.Itoa(cur)
			if strCur == reverse(strCur) && cur > ans {
				ans = cur
			}
		}
	}

	fmt.Println(ans) // 906609
}
