package main

import (
	"fmt"
	"time"
)

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func main() {
	start := time.Now()
	date := Date(1901, 1, 1)

	count := 0

	for date != Date(2000, 12, 1) {
		if date.Day() == 1 && date.Weekday() == 0 {
			count++
		}
		date = date.AddDate(0, 0, 1)

	}
	fmt.Println(count, time.Since(start)) // 171, 2.993229ms
}
