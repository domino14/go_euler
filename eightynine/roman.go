package eightynine

import "log"

// ToRoman Converts a number to a Roman number. Use rules from
// https://projecteuler.net/about=roman_numerals
// Assume number is no bigger than 5000
func ToRoman(num int) string {
	str := ""
	magnitude := 1
	n := num
	r := 0
	ones, fives, tens := "I", "V", "X"
	for {
		n, r = n/10, n%10
		log.Println(n, r)

		switch magnitude {
		case 10:
			ones = "X"
			fives = "L"
			tens = "C"
		case 100:
			ones = "C"
			fives = "D"
			tens = "M"
		case 1000:
			ones = "M"
			fives = "V_"
			tens = "X_"
		}

		switch r {
		case 4:
			if ones != "M" {
				str = ones + fives + str
			} else {
				str = "MMMM" + str
			}
		case 9:
			str = ones + tens + str

		default:
			tempStr := ""
			if r >= 5 {
				tempStr += fives
				r -= 5
			}
			if r < 4 {
				for i := 0; i < r; i++ {
					tempStr += ones
				}
			}
			str = tempStr + str
		}

		magnitude *= 10
		if n == 0 {
			break
		}
	}

	return str
}

// Convert a Roman number to an integer.
func FromRoman(roman string) int {
	// parse string left to right, keep some sort of state machine.

	curval := 0
	lastSeen := ' '
	for _, c := range roman {

		switch c {
		case 'M':
			curval += 1000
			if lastSeen == 'C' {
				curval -= 200
			}
		case 'D':
			curval += 500
			if lastSeen == 'C' {
				curval -= 200
			}
		case 'C':
			curval += 100
			if lastSeen == 'X' {
				curval -= 20
			}
		case 'L':
			curval += 50
			if lastSeen == 'X' {
				curval -= 20
			}
		case 'X':
			curval += 10
			if lastSeen == 'I' {
				curval -= 2
			}
		case 'V':
			curval += 5
			if lastSeen == 'I' {
				curval -= 2
			}
		case 'I':
			curval++
		}
		lastSeen = c
	}
	return curval
}
