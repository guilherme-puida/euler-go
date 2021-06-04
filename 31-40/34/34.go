package main

import (
	"fmt"
	"strconv"
)

func Factorial(n int) int {
	ans := 1
	for i := 2; i <= n; i++ {
		ans *= i
	}
	return ans
}

func DigitFactorials(n int) int {
	sum := 0
	str := strconv.Itoa(n)
	for _, e := range str {
		digit, _ := strconv.Atoi(string(e))
		sum += Factorial(digit)
	}

	return sum
}

func main() {
	ans := 0
	for i := 3; ; i++ {
		if i == DigitFactorials(i) {
			ans += i
			fmt.Println(i, ans) // 40730
		}
	}
}
