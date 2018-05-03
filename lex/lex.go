package lex

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"

	"github.com/sgg7269/tokenizer/token"
)

type EnclosedMeta struct {
	Value   byte
	Matched bool
}

// Meta ...
type Meta struct {
	Tokens      []token.Token
	EOS         bool
	Accumulator string
	EscapeNext  bool
	Period      bool
	OnlyNumbers bool
	Enclosed    EnclosedMeta
}

// TODO: chagne the name of this later
// var lexMeta Meta

// TODO: this will work for now, but we really need to return an error
// FIXME: don't know if we need this but fuck it, I'll just create a new one later
func determineToken(accumulator string) (token.Token, error) {
	if accumulator != "" {
		if t, ok := token.TokenMap[accumulator]; ok {
			return t, nil
		} else {
			// Check if we are enclosed by a " and if so process as a string
			if meta.Enclosed.Value == '"' {
				return token.Token{
					Type: "LITERAL",
					Value: token.Value{
						String: accumulator,
					},
				}, nil

				// Check if there is a period in the literal, if there is process as a float
			} else if meta.Period {
				// This always parses to 64 bits, downconvert later if needed
				conv, err := strconv.ParseFloat(accumulator, 64)
				if err != nil {
					fmt.Println("float parse: uh oh spaghetti-o", meta)
				}

				return token.Token{
					Type: "LITERAL",
					Value: token.Value{
						Type:   "float",
						True:   conv,
						String: accumulator,
					},
				}, nil
				// Otherwise we need to check if the name is only numbers
			} else {
				onlyDigits := true
				for _, char := range accumulator {
					// If we find a letter and this is not a string value, it must be a variable reference OR a variable declaration
					if !unicode.IsDigit(char) {
						// FIXME: fix this to cut early
						onlyDigits = false
					}
				}
				// The equals is getting here, need to code in all chars that aren't that
				if onlyDigits {
					conv, err := strconv.Atoi(accumulator)
					if err != nil {
						fmt.Println("int parse: uh oh spaghetti-o", meta)
					}
					return token.Token{
						Type: "LITERAL",
						Value: token.Value{
							Type:   "integer",
							True:   conv,
							String: accumulator,
						},
					}, nil

				} else {
					return token.Token{
						Type:     "IDENT",
						Expected: "ASSIGN",
						Value: token.Value{
							Type:   "ident",
							String: accumulator,
						},
					}, nil
				}
			}
		}
	}

	return token.Token{}, errors.New("Accumulator is empty")
}

// Lex is the external function for lexing a set of characters
// TODO: this should return the lexems and an error
// FIXME: make a flag to skip WS
func Lex(input string) ([]token.Token, error) {
	var meta = Meta{
		// FIXME: invert this var name
		OnlyNumbers: true,
	}

	// hack in 'p' variable for now

	for _, char := range input {

		// FIXME: make a map for token delmiters
		switch char {
		case ' ':
			if meta.Enclosed.Value != 0 && meta.Enclosed.Matched != true {
				meta.Accumulator += string(char)
				continue
			}

			determineToken(&meta)

			// FIXME: convert this to read from the map
			meta.Tokens = append(meta.Tokens, token.Token{
				Type: "WS",
				Value: token.Value{
					String: " ",
				},
			})
			meta = Meta{
				Tokens: meta.Tokens,
				// FIXME: invert this var name
				OnlyNumbers: true,
			}

		case ';':
			if meta.Enclosed.Value != 0 && meta.Enclosed.Matched != true {
				meta.Accumulator += string(char)
				continue
			}

			determineToken(&meta)

			meta.EOS = true
			meta.Tokens = append(meta.Tokens, token.Token{
				Type: "EOS",
				Value: token.Value{
					String: ";",
				},
			})
			meta.Accumulator = ""

		case '"':
			if meta.EscapeNext {
				meta.Accumulator += string(char)
				meta.EscapeNext = false
				continue
			}

			// This first if block controls whether quotes are included in the value of a string literal
			if meta.Enclosed.Value == 0 {
				meta.Enclosed.Value = '"'
				meta.Accumulator += string(char)
			} else if meta.Enclosed.Value == '"' && meta.Enclosed.Matched == false {
				meta.Enclosed.Matched = true
				meta.Accumulator += string(char)

				meta.Tokens = append(meta.Tokens, token.Token{
					Type: "LITERAL",
					Value: token.Value{
						Type:   "string",
						String: meta.Accumulator,
					},
				})

				meta.Accumulator = ""
				meta = Meta{
					Tokens:      meta.Tokens,
					OnlyNumbers: true,
				}
			}

		case ':':
			if meta.EscapeNext {
				meta.Accumulator += string(char)
				meta.EscapeNext = false
				continue
			}
			determineToken(&meta)
			meta.Accumulator = ":"
			determineToken(&meta)
			meta.Accumulator = ""

		case ',':
			if meta.EscapeNext {
				meta.Accumulator += string(char)
				meta.EscapeNext = false
				continue
			}
			determineToken(&meta)
			meta.Accumulator = ","
			determineToken(&meta)
			meta.Accumulator = ""

		case '=':
			if meta.EscapeNext {
				meta.Accumulator += string(char)
				meta.EscapeNext = false
				continue
			}
			determineToken(&meta)
			meta.Accumulator = "="
			determineToken(&meta)
			meta.Accumulator = ""

		case '{':
			if meta.EscapeNext {
				meta.Accumulator += string(char)
				meta.EscapeNext = false
				continue
			}
			determineToken(&meta)
			meta.Tokens = append(meta.Tokens, token.Token{
				Type: "L_BRACE",
				Value: token.Value{
					Type:   "L_BRACE",
					String: "{",
				},
			})
			meta.Accumulator = ""

		case '}':
			if meta.EscapeNext {
				meta.Accumulator += string(char)
				meta.EscapeNext = false
				continue
			}
			determineToken(&meta)
			meta.Tokens = append(meta.Tokens, token.Token{
				Type: "R_BRACE",
				Value: token.Value{
					Type:   "R_BRACE",
					String: "}",
				},
			})
			meta.Accumulator = ""

		case '[':
			if meta.EscapeNext {
				meta.Accumulator += string(char)
				meta.EscapeNext = false
				continue
			}
			determineToken(&meta)
			meta.Tokens = append(meta.Tokens, token.Token{
				Type: "L_BRACKET",
				Value: token.Value{
					Type:   "L_BRACKET",
					String: "[",
				},
			})
			meta.Accumulator = ""

		case ']':
			if meta.EscapeNext {
				meta.Accumulator += string(char)
				meta.EscapeNext = false
				continue
			}
			determineToken(&meta)
			meta.Tokens = append(meta.Tokens, token.Token{
				Type: "R_BRACKET",
				Value: token.Value{
					Type:   "R_BRACKET",
					String: "]",
				},
			})
			meta.Accumulator = ""

		case '-':
			if meta.EscapeNext {
				meta.Accumulator += string(char)
				meta.EscapeNext = false
				continue
			}
			determineToken(&meta)
			meta.Tokens = append(meta.Tokens, token.Token{
				Type:     "SEC_OP",
				Expected: "EXPR",
				Value: token.Value{
					Type:   "sub",
					String: "-",
				},
			})
			meta.Accumulator = ""

		case '+':
			if meta.EscapeNext {
				meta.Accumulator += string(char)
				meta.EscapeNext = false
				continue
			}
			determineToken(&meta)
			meta.Tokens = append(meta.Tokens, token.Token{
				Type:     "SEC_OP",
				Expected: "EXPR",
				Value: token.Value{
					Type:   "add",
					String: "+",
				},
			})
			meta.Accumulator = ""

		case '/':
			if meta.EscapeNext {
				meta.Accumulator += string(char)
				meta.EscapeNext = false
				continue
			}
			determineToken(&meta)
			meta.Tokens = append(meta.Tokens, token.Token{
				Type:     "PRI_OP",
				Expected: "EXPR",
				Value: token.Value{
					Type:   "div",
					String: "/",
				},
			})
			meta.Accumulator = ""

		case '*':
			if meta.EscapeNext {
				meta.Accumulator += string(char)
				meta.EscapeNext = false
				continue
			}
			determineToken(&meta)
			meta.Tokens = append(meta.Tokens, token.Token{
				Type:     "PRI_OP",
				Expected: "EXPR",
				Value: token.Value{
					Type:   "mult",
					String: "*",
				},
			})
			meta.Accumulator = ""

		case '(':
			if meta.EscapeNext {
				meta.Accumulator += string(char)
				meta.EscapeNext = false
				continue
			}
			determineToken(&meta)
			meta.Tokens = append(meta.Tokens, token.Token{
				Type:     "L_PAREN",
				Expected: "EXPR",
				Value: token.Value{
					Type:   "op_3", // TODO: check all these
					String: "(",
				},
			})
			meta.Accumulator = ""

		case ')':
			if meta.EscapeNext {
				meta.Accumulator += string(char)
				meta.EscapeNext = false
				continue
			}
			determineToken(&meta)
			meta.Tokens = append(meta.Tokens, token.Token{
				Type:     "R_PAREN",
				Expected: "EXPR",
				Value: token.Value{
					Type:   "op_3", // TODO: check all these
					String: ")",
				},
			})
			meta.Accumulator = ""

			// This first if block controls whether quotes are included in the value of a string literal
			// if meta.Enclosed.Value == '{' && meta.Enclosed.Matched == false {
			// 	meta.Enclosed.Matched = true
			// 	meta.Accumulator += string(char)

			// 	meta.Tokens = append(meta.Tokens, token.Token{
			// 		Type: "LITERAL",
			// 		Value: token.Value{
			// 			Type:   "string",
			// 			String: meta.Accumulator,
			// 		},
			// 	})

			// 	meta.Accumulator = ""
			// 	// meta = Meta {
			// 	// 	OnlyNumbers: true,
			// 	// }
			// }

		case '\\':
			if meta.Enclosed.Value == '"' || meta.Enclosed.Value == '\'' {
				// TODO: could do a get next here or something and just escape it at that point
				meta.EscapeNext = true
				continue
			}
			// TODO: define what we should do if it is not enclosed
			fmt.Println("Backslash out of place")

		case '.':
			// if meta.Enclosed.Value != "" {

			// }
			meta.Accumulator += string(char)
			meta.Period = true

		case '\n':
			if meta.EscapeNext {
				meta.Accumulator += string(char)
				meta.EscapeNext = false
				continue
			}
			determineToken(&meta)
			meta.Tokens = append(meta.Tokens, token.Token{
				Type: "NEWLINE",
				Value: token.Value{
					Type:   "newline",
					String: "\n",
				},
			})
			meta.Accumulator = ""

		case '\t':
			meta.Tokens = append(meta.Tokens, token.Token{
				Type: "WS",
				Value: token.Value{
					Type:   "tab",
					String: "\t",
				},
			})

		default:
			// FIXME: need to add *unallowed* characters in here or in the determineToken function

			// if unicode.IsLetter(char) {
			// 	meta.OnlyNumbers = false
			// }
			// fmt.Println(meta)
			// fmt.Println("IsLetter", unicode.IsLetter(char))
			meta.Accumulator += string(char)
		}
	}

	if meta.Accumulator != "" {
		// TODO: might need to make this return something so that we can determine what we get back
		determineToken(&meta)
		// if we found a token, which we should atleast get a literal back, then clear the accumulator, otherwise print an error
		// TODO: make the determineMeta do this function automatically
		meta.Accumulator = ""

		// fmt.Printf("Accumulator not empty %#v\n", meta)
	}

	// Append EOF (since this is not a token that we will recieve) and return
	// return append(meta.Tokens, token.Token{
	// 	Type: "EOF",
	// 	Value: token.Value{
	// 		String: string(0),
	// 	},
	// }), nil

	return meta.Tokens, nil
}
