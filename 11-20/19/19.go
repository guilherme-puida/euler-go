package main

import (
	"fmt"
	"time"
)

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func main() {
	start := Date(1901, 1, 1)

	count := 0

	for start != Date(2000, 12, 1) {
		if start.Day() == 1 && start.Weekday() == 0 {
			count++
		}
		start = start.AddDate(0, 0, 1)

	}
	fmt.Println(count) // 171
}
