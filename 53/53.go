package main

import (
	"fmt"
	"math/big"
	"time"
)

func Choose(n, r int) big.Int {
	var ans, up, down, down2 big.Int
	up.MulRange(1, int64(n))
	down.MulRange(1, int64(n-r))
	down2.MulRange(1, int64(r))
	down.Mul(&down, &down2)
	ans.Div(&up, &down)

	return ans
}

func main() {
	start := time.Now()
	ans := 0
	for n := 1; n <= 100; n++ {
		for r := 1; r <= n; r++ {
			a := Choose(n, r)
			if a.Cmp(big.NewInt(1000000)) == 1 {
				ans += 1
			}
		}
	}

	fmt.Println(ans, time.Since(start)) // 4075, 37.660958ms
}
