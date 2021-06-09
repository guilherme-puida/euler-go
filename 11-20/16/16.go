package main

import (
	"fmt"
	"math/big"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	var i big.Int
	i.Exp(big.NewInt(2), big.NewInt(1000), nil)

	ans := 0
	for _, e := range i.Text(10) {
		n, _ := strconv.Atoi(string(e))
		ans += n
	}

	fmt.Println(ans, time.Since(start)) // 1366, 17.698µs
}
