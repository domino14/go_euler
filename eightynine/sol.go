package eightynine

import (
	"io/ioutil"
	"log"
	"strings"
)

// Solve the actual problem, son.
func Solve() {
	bs, err := ioutil.ReadFile("eightynine/numbers.txt")
	if err != nil {
		log.Fatalln(err)
	}
	s := string(bs)
	numerals := strings.Split(s, "\n")
	saved := 0
	for _, n := range numerals {
		number := FromRoman(n)
		tr := ToRoman(number)

		saved += (len(n) - len(tr))
	}
	log.Println("Saved", saved)
}
