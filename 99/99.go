package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("/home/gui/euler-go/99/p099_base_exp.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line := 1
	var maxValue float64 = 0
	maxLine := line
	for scanner.Scan() {
		numbers := strings.Split(scanner.Text(), ",")
		base, _ := strconv.Atoi(numbers[0])
		exponent, _ := strconv.Atoi(numbers[1])

		number := math.Log(float64(base)) * float64(exponent)

		if number > maxValue {
			maxValue = number
			maxLine = line
		}
		line++
	}

	fmt.Println(maxLine) // 709

}
