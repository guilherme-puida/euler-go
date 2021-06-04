package main

import (
	"fmt"
	"math/big"
)

func main() {
	m := make(map[string]bool)

	for a := 2; a <= 100; a++ {
		bigA := big.NewInt(int64(a))
		for b := 2; b <= 100; b++ {
			bigB := big.NewInt(int64(b))
			var n big.Int
			n.Exp(bigA, bigB, nil)
			m[n.Text(10)] = true

		}
	}

	fmt.Println(len(m)) // 9183
}
