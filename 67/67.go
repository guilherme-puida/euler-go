package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	fileBuf, err := ioutil.ReadFile("/home/gui/euler-go/67/p067_triangle.txt")
	if err != nil {
		panic(err)
	}

	ot := string(fileBuf)

	var triangle [][]int

	s := strings.Split(ot, "\n")
	for _, line := range s { // Making the triangle grid
		var newLine []int
		temp := strings.Split(line, " ")
		for _, e := range temp {
			n, _ := strconv.Atoi(string(e))
			newLine = append(newLine, n)
		}
		triangle = append(triangle, newLine)
	}

	for i := len(triangle) - 3; i >= 0; i-- { // Bottom down approach
		previousLine := triangle[i+1]
		currentLine := triangle[i]
		var newLine []int
		for j, e := range currentLine {
			var element int

			a := previousLine[j+1]
			b := previousLine[j]
			element = e + Max(a, b)

			newLine = append(newLine, element)
		}
		triangle[i] = newLine
	}

	fmt.Println(triangle[0][0]) // 7273

}
