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

// https://en.wikipedia.org/wiki/Lattice_path
// Choose(40, 10) comes from this wikipedia page
// The general formula is Choose(n + k, n) to get
//number of paths from (0, 0) to (n, k)
func main() {
	start := time.Now()
	ans := Choose(40, 20)
	fmt.Println(ans.Text(10), time.Since(start)) // 137846528820, 11.934µs
}
