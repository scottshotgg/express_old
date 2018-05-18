package token

import (
	"fmt"
	"os"
)

// Lexemes ...
var Lexemes = []string{
	"var",
	"int",
	"string",
	"select",
	"SELECT",
	"FROM",
	"WHERE",
	":",
	"=",
	"+",
	"-",
	"*",
	"/",
	"(",
	")",
	"{",
	"}",
	"[",
	"]",
	"\"",
	";",
	",",
	"#",
	"!",
	"<",
	">",
	"@",
	// "„",
}

// LexemeMap is used for holding the lexemes that will be used to identify tokens in the lexer
var LexemeMap = map[string]Token{}

func init() {
	for _, lexeme := range Lexemes {
		if lexToken, ok := TokenMap[lexeme]; !ok {
			fmt.Println("ERROR: Lexeme not found in TokenMap: ", lexeme)
			os.Exit(7)
		} else {
			LexemeMap[lexeme] = lexToken
		}
	}
}
