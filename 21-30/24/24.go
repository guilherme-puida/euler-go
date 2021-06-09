package main

import (
	"fmt"
	"time"
)

func Factoradic(n int) [10]int {
	var factoradic [10]int
	i := 1
	for n != 0 {
		factoradic[10-i] = n % i
		n /= i
		i++
	}
	return factoradic
}

func removeElement(slice []int, element int) []int {
	var index int

	for i, e := range slice {
		if e == element {
			index = i
			break
		}
	}

	copy(slice[index:], slice[index+1:])
	slice[len(slice)-1] = 0
	return slice[:len(slice)-1]
}

func main() {
	start := time.Now()
	digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	factoradic := Factoradic(999999)

	var perm []int
	for _, e := range factoradic {
		perm = append(perm, digits[e])
		digits = removeElement(digits, digits[e])
		fmt.Println(perm, digits, e)
	}

	fmt.Println(perm, time.Since(start)) // 2783915460, 62.37µs
}
