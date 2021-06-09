package main

import (
	"fmt"
	"math/big"
	"time"
)

func isPalindrome(str string) bool {
	for i := 0; i < len(str)/2+1; i++ {
		j := len(str) - i - 1
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

func main() {
	start := time.Now()
	n := big.NewInt(1)
	ans := big.NewInt(0)

	for n.Cmp(big.NewInt(1000000)) == -1 {
		str := n.Text(10)
		bStr := n.Text(2)

		if isPalindrome(str) && isPalindrome(bStr) {
			ans.Add(ans, n)
		}

		n.Add(n, big.NewInt(1))
	}

	fmt.Println(ans.Text(10), time.Since(start)) // 872187, 230.151237ms
}
