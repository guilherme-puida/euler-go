package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func SumFifthPowerDigits(n int) int {
	s := strconv.Itoa(n)
	sum := 0
	for _, e := range s {
		digit, _ := strconv.Atoi(string(e))
		sum += int(math.Pow(float64(digit), 5))
	}
	return sum
}

func main() {
	start := time.Now()
	sum := 0
	for i := 2; ; i++ {
		if i == SumFifthPowerDigits(i) {
			sum += i
		}
		if i > 500000 {
			break
		}
	}
	fmt.Println(sum, time.Since(start)) // 443839, 143.44952ms
}
