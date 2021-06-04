package main

import "fmt"

func BiggestPrimeFactor(n int) int {
	p := 0
	for n%2 == 0 {
		p = 2
		n /= 2
	}
	for i := 3; i*i < n+1; i += 2 {
		for n%i == 0 {
			p = i
			n /= i
		}
	}

	if n > 2 {
		p = n
	}

	return p
}

func main() {
	fmt.Println(BiggestPrimeFactor(600851475143)) // 6857
}
