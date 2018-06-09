package c

import (
	"fmt"

	"github.com/scottshotgg/Express/token"
)

// This is a placeholder for the C converter package that will be used to convert Express -> C
// Doing so will allow Express to leverage all available C tools

// Translate ...
func Translate(tokens []token.Token) {
	fmt.Println("tokens", tokens)

	for _, t := range tokens {
		fmt.Println(t)
	}
}
