package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func DigitCount(n int) int {
	v := math.Log10(float64(n))
	return int(v) + 1
}

func main() {
	start := time.Now()
	v := [7]int{1, 10, 100, 1000, 10000, 100000, 1000000}
	curV, curPos := 0, 0
	ans := 1

	for i := 1; curPos < 7; i++ {
		if DigitCount(i)+curV >= v[curPos] {
			n := strconv.Itoa(i)
			pos := v[curPos] - curV - 1
			n = string(n[pos])
			a, _ := strconv.Atoi(n)
			ans *= a

			curPos++
		}
		curV += DigitCount(i)
	}
	fmt.Println(ans, time.Since(start)) // 210, 5.060911ms
}
