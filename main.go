package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"unicode"

	"github.com/sgg7269/tokenizer/token"
)

// TODO: we currently have to have the space afterwards, need to add the parsing code into the semicolon code
// TODO: we currently only have numbers being parsed within floats, so nothing like '9.9f', etc

// Program ...
type Program struct {
	Index  int
	Name   string
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

var (
	p          Program
	jsonIndent string
)

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

	for _, char := range p.Value {
		// FIXME: make a map for token delmiters
		switch char {
		case ' ':
			if meta.Enclosed.Value != 0 && meta.Enclosed.Matched != true {
				meta.Accumulator += string(char)
				continue
			}

			determineToken(meta)

			// FIXME: convert this to read from the map
			p.Tokens = append(p.Tokens, token.Token{
				Type: "WS",
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
				meta = ParseMeta{
					OnlyNumbers: true,
				}
			}

		case ':':
			if meta.EscapeNext {
				meta.Accumulator += string(char)
				meta.EscapeNext = false
				continue
			}
			determineToken(meta)
			meta.Accumulator = ":"
			determineToken(meta)
			meta.Accumulator = ""

		case ',':
			if meta.EscapeNext {
				meta.Accumulator += string(char)
				meta.EscapeNext = false
				continue
			}
			determineToken(meta)
			meta.Accumulator = ","
			determineToken(meta)
			meta.Accumulator = ""

		case '=':
			if meta.EscapeNext {
				meta.Accumulator += string(char)
				meta.EscapeNext = false
				continue
			}
			determineToken(meta)
			meta.Accumulator = "="
			determineToken(meta)
			meta.Accumulator = ""

		case '{':
			if meta.EscapeNext {
				meta.Accumulator += string(char)
				meta.EscapeNext = false
				continue
			}
			determineToken(meta)
			p.Tokens = append(p.Tokens, token.Token{
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
			determineToken(meta)
			p.Tokens = append(p.Tokens, token.Token{
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
			determineToken(meta)
			p.Tokens = append(p.Tokens, token.Token{
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
			determineToken(meta)
			p.Tokens = append(p.Tokens, token.Token{
				Type: "R_BRACKET",
				Value: token.Value{
					Type:   "R_BRACKET",
					String: "]",
				},
			})
			meta.Accumulator = ""

			// This first if block controls whether quotes are included in the value of a string literal
			// if meta.Enclosed.Value == '{' && meta.Enclosed.Matched == false {
			// 	meta.Enclosed.Matched = true
			// 	meta.Accumulator += string(char)

			// 	p.Tokens = append(p.Tokens, token.Token{
			// 		Type: "LITERAL",
			// 		Value: token.Value{
			// 			Type:   "string",
			// 			String: meta.Accumulator,
			// 		},
			// 	})

			// 	meta.Accumulator = ""
			// 	// meta = ParseMeta{
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
			determineToken(meta)
			p.Tokens = append(p.Tokens, token.Token{
				Type: "NEWLINE",
				Value: token.Value{
					Type:   "newline",
					String: "\n",
				},
			})
			meta.Accumulator = ""

		case '\t':
			p.Tokens = append(p.Tokens, token.Token{
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
		determineToken(meta)
		// if we found a token, which we should atleast get a literal back, then clear the accumulator, otherwise print an error
		// TODO: make the determineMeta do this function automatically
		meta.Accumulator = ""

		// fmt.Printf("Accumulator not empty %#v\n", meta)
	}
}

func printTokens() {
	fmt.Println()

	if jsonIndent != "" {
		for _, token := range p.Tokens {
			tokenJSON, err := json.MarshalIndent(token, "", jsonIndent)
			if err != nil {
				fmt.Printf("\nERROR: Could not marshal JSON from token: %#v\n", token)
				os.Exit(9)
			}
			fmt.Println(string(tokenJSON))
		}
	} else {
		for _, token := range p.Tokens {
			tokenJSON, err := json.Marshal(token)
			if err != nil {
				fmt.Printf("\nERROR: Could not marshal JSON from token: %#v\n", token)
				os.Exit(9)
			}
			fmt.Println(string(tokenJSON))
		}
	}
}

func outputTokens() {
	tokenFilename := p.Name + ".tokens"

	// For more granular writes, open a file for writing.
	f, err := os.Create(tokenFilename)
	defer f.Close()
	if err != nil {
		fmt.Println("ERROR: Could not open token output file:", tokenFilename)
		os.Exit(9)
	}
	w := bufio.NewWriter(f)

	fmt.Println()
	fmt.Println("Outputting tokens to:", tokenFilename)

	var tokenJSON []byte
	if jsonIndent != "" {
		for index, token := range p.Tokens {
			tokenJSON, err = json.MarshalIndent(token, "", jsonIndent)
			if err != nil {
				fmt.Printf("\nERROR: Could not marshal JSON from token: %#v\n", token)
				os.Exit(9)
			}
			if index < len(p.Tokens)-1 {
				tokenJSON = append(tokenJSON, '\n')
			}
			// TODO: we should check the amount later
			_, err = w.Write(tokenJSON)
			if err != nil {
				fmt.Println("ERROR: Could not write to token output file:", tokenFilename)
				os.Exit(9)
			}
		}
	} else {
		for index, token := range p.Tokens {
			tokenJSON, err = json.Marshal(token)
			if err != nil {
				fmt.Printf("\nERROR: Could not marshal JSON from token: %#v\n", token)
				os.Exit(9)
			}
			if index < len(p.Tokens)-1 {
				tokenJSON = append(tokenJSON, '\n')
			}
			_, err = w.Write(tokenJSON)
			if err != nil {
				fmt.Println("ERROR: Could not write to token output file:", tokenFilename)
				os.Exit(9)
			}
		}
	}

	err = w.Flush()
	if err != nil {
		fmt.Println("ERROR: Could not flush writer, data may be missing:", tokenFilename)
	}
}

// TODO: should look into making english-like inputs for the indentation
// TODO: add flags for verbosity, printouts, whether to not parse the tokens, output file for outputting tokens, output format (xml, json, native) etc
// TODO: check if the file exists and if it does warn them that the output file will be overridden and ask if they still want to go through
// func parseFlags() {
// 	jsonIndentPtr := flag.String("jsonIndent", "blank", "Indent that will be used for the JSON printout of the tokens")
// 	flag.Parse()

// 	fmt.Println("jsonptr", *jsonIndentPtr)

// 	// jsonIndent
// 	jsonIndent = *jsonIndentPtr
// }

func main() {
	// TODO: add some flags later
	// parseFlags()

	argLen := len(os.Args)

	if argLen < 3 {
		fmt.Println("ERROR: You must provide an input program")
		return
	}

	programName := os.Args[argLen-1]

	input, err := ioutil.ReadFile(programName)
	if err != nil {
		fmt.Printf("ERROR: Cannot read input program: %s\n", programName)
		return
	}

	p = Program{
		Value:  string(input),
		Name:   programName,
		Length: len(input),
	}

	fmt.Println("=======================")
	fmt.Println()
	fmt.Println("Tokenizing:")
	fmt.Println()
	fmt.Println(p.Value)
	fmt.Println()
	fmt.Println("=======================")
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

	// TODO: always output tokens right now
	outputTokens()
}
