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
	var err error

	// Make a token and set the default value to bool; this is just because its the
	// first case in the switch and everything below sets it, so it makes the code a bit
	// cleaner
	t := token.Token{
		ID:   0,
		Type: "LITERAL",
		Value: token.Value{
			True: false,
			Type: "bool",
		},
	}

	fmt.Println("acc", meta.Accumulator)

	switch meta.Accumulator {
	// Check if its false, we dont have to set anything for that
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
				t.Value.True, err = strconv.ParseInt(t.Value.String, 2, 64)
				if err != nil {
					fmt.Println("ERROR", err)
				}
				t.Value.Type = "binary"

			// Octal
			case "0o":
				t.Value.True, err = strconv.ParseInt(t.Value.String, 8, 64)
				if err != nil {
					fmt.Println("ERROR", err)
				}
				t.Value.Type = "octal"

			// Hexadecimal
			case "0x":
				t.Value.True, err = strconv.ParseInt(t.Value.String, 16, 64)
				if err != nil {
					fmt.Println("ERROR", err)
				}
				t.Value.Type = "hexadecimal"

			// Else it must be either an int, float, or an ident
			default:
				// Clear the string value
				t.Value.String = ""

				// Attempt to parse an int from the accumulator
				t.Value.True, err = strconv.ParseInt(meta.Accumulator, 0, 0)
				t.Value.Type = "int"

				// If it errors, check to see if it is an int
				if err != nil {
					// Attempt to parse a float from the accumulator
					t.Value.True, err = strconv.ParseFloat(meta.Accumulator, 0)
					t.Value.Type = "float"
					if err != nil {
						// leave this checking for the semantic
						// 	identSplit := strings.Split(meta.Accumulator, ".")
						// 	if len(identSplit) > 1 {
						// 		for _, ident := range identSplit {

						// 		}
						// 	}

						// If it errors, assume that it is an ident (for now)
						t.Type = "IDENT"
						t.Value = token.Value{
							String: meta.Accumulator,
						}
					}
				}
			}
		} else {
			// Clear the string value
			t.Value.String = ""

			// Attempt to parse an int from the accumulator
			t.Value.True, err = strconv.ParseInt(meta.Accumulator, 0, 0)
			t.Value.Type = "int"

			// If it errors, check to see if it is an int
			if err != nil {
				// Attempt to parse a float from the accumulator
				t.Value.True, err = strconv.ParseFloat(meta.Accumulator, 0)
				t.Value.Type = "float"
				if err != nil {
					// leave this checking for the semantic
					// 	identSplit := strings.Split(meta.Accumulator, ".")
					// 	if len(identSplit) > 1 {
					// 		for _, ident := range identSplit {

					// 		}
					// 	}

					// If it errors, assume that it is an ident (for now)
					t.Type = "IDENT"
					t.Value = token.Value{
						String: meta.Accumulator,
					}
				}
			}
		}
	}

	return t
}

// Lex ...
func Lex(input string) ([]token.Token, error) {
	var meta lexMeta

	fmt.Println("lexing shit", input)

	for index := 0; index < len(input); index++ {
		char := input[index]

		fmt.Printf("char \"%s\" %s\n", string(char), meta.Accumulator)

		// TODO: need to decide whether we want to append to the accumulator first or second
		if string(char) == " " || string(char) == "\n" {
			if meta.Accumulator != "" {
				if lexemeToken, ok := token.LexemeMap[meta.Accumulator]; ok {
					fmt.Println("Found char1", meta.Accumulator)
					meta.Tokens = append(meta.Tokens, lexemeToken)
				} else {
					fmt.Println("Found literal1", meta.Accumulator)
					meta.Tokens = append(meta.Tokens, meta.LexLiteral())
				}
				// Pull this from the TokenMap because we don't want the space in the LexemeMap
				meta.Tokens = append(meta.Tokens, token.TokenMap[string(char)])
				meta.Accumulator = ""
			} else if string(char) == " " || string(char) == "\n" {
				fmt.Println("i got gere1")
				// Pull this from the TokenMap because we don't want the space in the LexemeMap
				meta.Tokens = append(meta.Tokens, token.TokenMap[string(char)])
				meta.Accumulator = ""
			}
			fmt.Println("continuing")

			continue

		} else {
			if lexemeToken, ok := token.LexemeMap[string(char)]; ok {
				fmt.Println("Found char2", meta.Accumulator)

				// String out the comments
				switch lexemeToken.Type {
				case "DIV":
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
								fmt.Println(input[index], input[index+1])
								if index == len(input) || (input[index] == '*' && input[index+1] == '/') {
									break
								}
							}
						}
					}
					continue
				}

				if meta.Accumulator != "" {
					fmt.Println("Found literal2", meta.Accumulator)

					meta.Tokens = append(meta.Tokens, meta.LexLiteral())
					meta.Accumulator = ""
				}

				meta.Tokens = append(meta.Tokens, lexemeToken)
				meta.Accumulator = ""

				// meta.Tokens = append()
				continue
			} else if string(char) == " " || string(char) == "\n" {
				fmt.Println("i got gere2")
				// Pull this from the TokenMap because we don't want the space in the LexemeMap
				meta.Tokens = append(meta.Tokens, token.TokenMap[string(char)])
				meta.Accumulator = ""
			}
		}

		meta.Accumulator += string(char)
		fmt.Println(meta.Accumulator)

		if char == 0 {
			break
		}
	}

	return meta.Tokens, nil
}
