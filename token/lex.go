package token

import (
	"fmt"
	"os"
)

// var lexems = []token.Token{
// 	"=": Token{
// 		Type:     "ASSIGN",
// 		Expected: "EXPR",
// 		Value: Value{
// 			Type:   "assign",
// 			String: "=",
// 		},
// 	},
// }

// use this token: "„"

// LexemeMap ... fk u go
// TODO: make an accessor function for this var
// TODO: just do this for now

// Lexemes ...
var Lexemes = []string{
	"var",
	"int",
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
}

// LexemeMap ...
var LexemeMap = map[string]Token{}

func init() {
	for _, lexeme := range Lexemes {
		if lexToken, ok := TokenMap[lexeme]; !ok {
			fmt.Println("ERROR: Lexeme not found in TokenMap: ", lexeme)
			// fmt.Println("TokenMap: %#v")
			os.Exit(7)
		} else {
			LexemeMap[lexeme] = lexToken
		}
	}

	// ADDITIONS
	// LexemeMap["SELECT"] = LexemeMap["select"]
}
