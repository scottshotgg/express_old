package lex

import (
	"fmt"
	"strconv"

	"github.com/scottshotgg/Express/token"
)

type lexMeta struct {
	Accumulator string
	Tokens      []token.Token
	LastToken   token.Token
}

func (meta *lexMeta) LexLiteral() token.Token {
	// var err error

	// Make a token and set the default value to bool; this is just because its the
	// first case in the switch and everything below sets it, so it makes the code a bit
	// cleaner
	// We COULD do this with tokens in the tokenMap for true and false
	t := token.Token{
		ID:   0,
		Type: token.Literal,
		Value: token.Value{
			True: false,
			Type: token.BoolType,
		},
	}

	switch meta.Accumulator {
	// Default value is false, we only need to catch the case
	case "false":

	// Check if its true
	case "true":
		t.Value.True = true

	// Else move on and figure out what kind of number it is (or an ident)
	default:
		// Figure out from the two starting characters
		if len(meta.Accumulator) > 2 {
			t.Value.String = meta.Accumulator[2:]
			switch meta.Accumulator[:2] {
			// Binary
			case "0b":
				value, err := strconv.ParseInt(t.Value.String, 2, 64)
				if err != nil {
					fmt.Println("ERROR", err)
				}
				// t.Value.Type = "binary"
				t.Value.True = int(value)
				t.Value.Type = token.IntType
				return t

			// Octal
			case "0o":
				value, err := strconv.ParseInt(t.Value.String, 8, 64)
				if err != nil {
					fmt.Println("ERROR", err)
				}
				// t.Value.Type = "octal"
				t.Value.True = int(value)
				t.Value.Type = token.IntType
				return t

			// Hexadecimal
			case "0x":
				value, err := strconv.ParseInt(t.Value.String, 16, 64)
				if err != nil {
					fmt.Println("ERROR", err)
				}
				// t.Value.Type = "hexadecimal"
				t.Value.True = int(value)
				t.Value.Type = token.IntType
				return t
			}
		}
		// Clear the string value
		t.Value.String = ""

		// Attempt to parse an int from the accumulator
		value, err := strconv.ParseInt(meta.Accumulator, 0, 0)
		if err != nil {
			// TODO:
		}
		t.Value.True = int(value)
		t.Value.Type = token.IntType

		// TODO: need to make something for scientific notation with carrots and e
		// If it errors, check to see if it is an int
		if err != nil {
			// Attempt to parse a float from the accumulator
			t.Value.True, err = strconv.ParseFloat(meta.Accumulator, 0)
			t.Value.Type = token.FloatType
			if err != nil {
				// leave this checking for the semantic
				// 	identSplit := strings.Split(meta.Accumulator, ".")
				// 	if len(identSplit) > 1 {
				// 		for _, ident := range identSplit {

				// 		}
				// 	}

				// If it errors, assume that it is an ident (for now)
				t.Type = token.Ident
				t.Value = token.Value{
					String: meta.Accumulator,
				}
			}
		}
	}

	return t
}

// Lex attemps to lex the token
func Lex(input string) ([]token.Token, error) {
	var meta lexMeta

	for index := 0; index < len(input); index++ {
		char := input[index]
		if string(char) == " " || string(char) == "\n" {
			if meta.Accumulator != "" {
				if lexemeToken, ok := token.LexemeMap[meta.Accumulator]; ok {
					meta.Tokens = append(meta.Tokens, lexemeToken)
				} else {
					meta.Tokens = append(meta.Tokens, meta.LexLiteral())
				}
				// Pull this from the TokenMap because we don't want the space in the LexemeMap
				meta.Tokens = append(meta.Tokens, token.TokenMap[string(char)])
				meta.Accumulator = ""
			} else if string(char) == " " || string(char) == "\n" {
				// Pull this from the TokenMap because we don't want the space in the LexemeMap
				meta.Tokens = append(meta.Tokens, token.TokenMap[string(char)])
				meta.Accumulator = ""
			}

			continue

		} else {
			if lexemeToken, ok := token.LexemeMap[string(char)]; ok {
				// Filter out the comments
				switch lexemeToken.Value.Type {
				case "div":
					index++
					if index < len(input)-1 {
						switch input[index] {
						case '/':
							for {
								index++
								if index == len(input) || input[index] == '\n' {
									break
								}
							}

						case '*':
							for {
								index++
								if index == len(input) || (input[index] == '*' && input[index+1] == '/') {
									break
								}
							}
						}
					}
					continue

				// Use the lexer to parse strings
				case "squote":
					fmt.Println("found an squote")
					fallthrough
				case "dquote":
					stringLiteral := ""

					index++
					for string(input[index]) != lexemeToken.Value.String {
						stringLiteral += string(input[index])
						index++
					}

					varType := token.StringType
					if len(stringLiteral) < 2 {
						varType = token.CharType
					}

					meta.Tokens = append(meta.Tokens, token.Token{
						ID:   0,
						Type: token.Literal,
						Value: token.Value{
							Type:   varType,
							True:   stringLiteral,
							String: stringLiteral,
						},
					})

					continue
				}

				if meta.Accumulator != "" {
					meta.Tokens = append(meta.Tokens, meta.LexLiteral())
					meta.Accumulator = ""
				}

				meta.Tokens = append(meta.Tokens, lexemeToken)
				meta.Accumulator = ""

				continue
			} else if string(char) == " " || string(char) == "\n" {
				// Pull this from the TokenMap because we don't want the space in the LexemeMap
				meta.Tokens = append(meta.Tokens, token.TokenMap[string(char)])
				meta.Accumulator = ""
			}
		}

		meta.Accumulator += string(char)

		if char == 0 {
			break
		}
	}

	return meta.Tokens, nil
}
