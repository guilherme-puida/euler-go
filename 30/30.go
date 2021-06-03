package main

import (
	"fmt"
	"math"
	"strconv"
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
	sum := 0
	for i := 2; ; i++ {
		if i == SumFifthPowerDigits(i) {
			sum += i
			fmt.Println(sum) // 443839
		}
	}
}
