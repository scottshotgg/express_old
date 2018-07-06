package collector

import (
	"fmt"
	"strconv"
	"unicode"
)

var digitString1 = "5"
var digitString2 = "56"

// Digit attempts to collect the largest string of digits
// from the source text
func Digit(source string) (sourceReturn string, digit int, err error) {
	// var digitString string
	var char rune
	var i int
	for i, char = range source {
		if !unicode.IsDigit(char) {
			break
		}
	}
	digit, err = strconv.Atoi(source[:i+1])
	if err != nil {
		sourceReturn = source
		return
	}
	sourceReturn = source[i:]
	return
}

func DoShit() {
	fmt.Println("init shit bitch")
	fmt.Println("digitString1")
	fmt.Println(Digit(digitString1))
	fmt.Println("digitString2")
	fmt.Println(Digit(digitString2))
}
