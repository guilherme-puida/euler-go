package main

import (
	"fmt"
)

func isAbundant(n int) bool {
	sum := 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			if n/i != i {
				sum += n / i
			}
		}
	}

	return sum > n
}

func main() {
	var abundant []int
	for i := 1; i <= 20161; i++ {
		if isAbundant(i) {
			abundant = append(abundant, i)
		}
	}

	abundantSum := make(map[int]bool)
	for i := 0; i < len(abundant); i++ {
		for j := i; i < len(abundant); j++ {
			if abundant[i]+abundant[j] <= 20161 {
				abundantSum[abundant[i]+abundant[j]] = true
			} else {
				break
			}
		}
	}

	ans := 0
	for i := 1; i <= 20161; i++ {
		_, ok := abundantSum[i]
		if !ok {
			ans += i
		}
	}

	fmt.Println(ans) // 4179871

}
