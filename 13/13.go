package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
)

func main() {
	file, err := os.Open("/home/gui/euler-go/23/13.txt")
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

	fmt.Println(a.Text(10)) // 5537376230390876637302048746832985971773659831892672
}
