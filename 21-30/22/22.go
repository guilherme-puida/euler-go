package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func LetterValue(letter string) int {
	alph := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return strings.Index(alph, letter) + 1
}

func NameValue(name string) int {
	value := 0
	for _, letter := range name {
		value += LetterValue(string(letter))
	}
	return value
}

func main() {
	ot, err := ioutil.ReadFile("/home/gui/euler-go/21-30/22/p022_names.txt")
	if err != nil {
		panic(err)
	}

	s := strings.ReplaceAll(string(ot), `"`, "")
	names := strings.Split(s, ",")
	sort.Strings(names)

	score := 0

	for i, name := range names {

		value := NameValue(name)
		score += (i + 1) * value
	}

	fmt.Println(score) // 871198282

}
