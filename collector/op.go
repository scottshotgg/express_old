package collector

import (
	"errors"
	"fmt"
	"os"
)

var opString1 = "+"
var opString2 = "-"

var ops = map[string]bool{
	"+": true,
}

// Op attempts to collect a single op from the source text
func Op(source string) (op string, err error) {
	if len(source) > 0 {
		if op := ops[string(source[0])]; op {
			return string(source[0]), nil
		} // could make some sort of index here
		return "", errors.New("shit")
	}
	return "", errors.New("shit2")
}

func DoShit2() {
	fmt.Println("init shit bitch")
	fmt.Println("opString1")
	fmt.Println(Op(opString1))
	fmt.Println("opString2")
	fmt.Println(Op(opString2))
	os.Exit(9)
}
