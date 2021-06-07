package main

import "fmt"

func SortedTuple(a, b int) [2]int {
	var t [2]int

	if a > b {
		t[0], t[1] = a, b
	} else {
		t[0], t[1] = b, a
	}
	return t
}

func IsInt(n float64) bool {
	return n == float64(int(n))
}

func RightTriangles(p int) int {
	store := make(map[[2]int]bool)

	if p%2 != 0 {
		return 0
	}

	count := 0
	for b := 1; b < p/2; b++ {
		pf, bf := float64(p), float64(b)
		a := pf / 2 * ((pf - 2*bf) / (pf - bf))

		if IsInt(a) {
			ab := SortedTuple(int(a), b)
			_, ok := store[ab]
			if !ok {
				count++
				store[ab] = true
			}
		}
	}
	return count
}

func main() {
	max, maxValue := 0, 0
	for i := 1; i <= 1000; i++ {
		cur := RightTriangles(i)
		if cur > max {
			max = cur
			maxValue = i
		}
	}

	fmt.Println(maxValue) // 840
}
