package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"
)

func main() {
	start := time.Now()
	file, err := os.Open("/home/gui/Programming/Go/euler-go/11-20/13/13.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	var a big.Int
	for scanner.Scan() {
		n := new(big.Int)
		n, ok := n.SetString(scanner.Text(), 10)

		if !ok {
			log.Fatal(err)
		}

		a.Add(&a, n)
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(a.Text(10)[0:10], time.Since(start)) // 5537376230, 83.861µs
}
