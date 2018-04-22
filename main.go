package main

import (
	"fmt"
	"strconv"
	"unicode"

	"github.com/sgg7269/tokenizer/token"
)

// TODO: we currently have to have the space afterwards, need to add the parsing code into the semicolon code
// TODO: we currently only have numbers being parsed within floats, so nothing like '9.9f', etc
var input = "bool a := true .+ 9;"

// Program ...
type Program struct {
	Index  int
	Value  string
	Length int
	EOS    bool
	Tokens []token.Token
}

// ParseMeta ...
type ParseMeta struct {
	Accumulator string
	EscapeNext  bool
	Period      bool
	OnlyNumbers bool
	Enclosed    struct {
		Value   byte
		Matched bool
	}
}

var p Program

func determineToken(meta ParseMeta) {
	if meta.Accumulator != "" {
		if t, ok := token.TokenMap[meta.Accumulator]; ok {
			p.Tokens = append(p.Tokens, t)
		} else {
			// Check if we are enclosed by a " and if so process as a string
			if meta.Enclosed.Value == '"' {
				p.Tokens = append(p.Tokens, token.Token{
					Type: "LITERAL",
					Value: token.Value{
						String: meta.Accumulator,
					},
				})
				// Check if there is a period in the literal, if there is process as a float
			} else if meta.Period {
				// This always parses to 64 bits, downconvert later if needed
				conv, err := strconv.ParseFloat(meta.Accumulator, 64)
				if err != nil {
					fmt.Println("float parse: uh oh spaghetti-o", meta)
				}

				p.Tokens = append(p.Tokens, token.Token{
					Type: "LITERAL",
					Value: token.Value{
						Type:   "float",
						True:   conv,
						String: meta.Accumulator,
					},
				})
				// Otherwise we need to check if the name is only numbers
			} else {
				onlyDigits := true
				for _, char := range meta.Accumulator {
					// If we find a letter and this is not a string value, it must be a variable reference OR a variable declaration
					if !unicode.IsDigit(char) {
						// FIXME: fix this to cut early
						onlyDigits = false
					}
				}
				// The equals is getting here, need to code in all chars that aren't that
				if onlyDigits {
					conv, err := strconv.Atoi(meta.Accumulator)
					if err != nil {
						fmt.Println("int parse: uh oh spaghetti-o", meta)
					}
					p.Tokens = append(p.Tokens, token.Token{
						Type: "LITERAL",
						Value: token.Value{
							Type:   "integer",
							True:   conv,
							String: meta.Accumulator,
						},
					})

				} else {
					p.Tokens = append(p.Tokens, token.Token{
						Type:     "IDENT",
						Expected: "ASSIGN",
						Value: token.Value{
							Type:   "ident",
							String: meta.Accumulator,
						},
					})
				}
			}
		}
	}
}

func lex() {
	meta := ParseMeta{
		// FIXME: invert this var name
		OnlyNumbers: true,
	}

	for _, char := range input {
		// FIXME: make a map for token delmiters
		switch char {
		case ' ':
			if meta.Enclosed.Value != 0 && meta.Enclosed.Matched != true {
				meta.Accumulator += string(char)
				continue
			}

			determineToken(meta)

			p.Tokens = append(p.Tokens, token.Token{
				Type: "SPACE",
				Value: token.Value{
					String: " ",
				},
			})
			meta.Accumulator = ""

		case ';':
			if meta.Enclosed.Value != 0 && meta.Enclosed.Matched != true {
				meta.Accumulator += string(char)
				continue
			}

			determineToken(meta)

			p.EOS = true
			p.Tokens = append(p.Tokens, token.Token{
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

				p.Tokens = append(p.Tokens, token.Token{
					Type: "LITERAL",
					Value: token.Value{
						Type:   "string",
						String: meta.Accumulator,
					},
				})

				meta.Accumulator = ""
				// meta = ParseMeta{
				// 	OnlyNumbers: true,
				// }
			}

		case '\\':
			meta.EscapeNext = true

		case '.':
			// if meta.Enclosed.Value != "" {

			// }
			meta.Accumulator += string(char)
			meta.Period = true

		default:
			// if unicode.IsLetter(char) {
			// 	meta.OnlyNumbers = false
			// }
			// fmt.Println(meta)
			// fmt.Println("IsLetter", unicode.IsLetter(char))
			meta.Accumulator += string(char)
		}
	}

	if meta.Accumulator != "" {
		fmt.Printf("Accumulator not empty %#v\n", meta)
	}
}

func printTokens() {
	for _, token := range p.Tokens {
		fmt.Printf("%#v\n", token)
	}
}

func main() {
	p = Program{
		Value:  input,
		Length: len(input),
	}

	fmt.Println(input)
	fmt.Println()

	lex()

	// if meta.Enclosed.Value != 0 && !meta.Enclosed.Matched {
	// 	fmt.Println("Enclosing not matched")
	// 	printTokens()
	// 	os.Exit(5)
	// }

	// uncomment to start checking for EOS again
	// if !p.EOS {
	// 	fmt.Println("Statement not ended")
	// 	printTokens()
	// 	os.Exit(4)
	// }

	p.Tokens = append(p.Tokens, token.Token{
		Type: "EOF",
		Value: token.Value{
			String: string(0),
		},
	})

	printTokens()
}
