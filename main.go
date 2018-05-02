package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"unicode"

	"github.com/sgg7269/tokenizer/parse"
	"github.com/sgg7269/tokenizer/token"
	//"llvm.org/llvm/bindings/go/llvm"
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

// LexMeta ...
type LexMeta struct {
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
	endTokens  = []token.Token{}
	llvmStart  = "define i32 @main() #0 {\n"
	llvmEnd    = "ret i32 0\n}"
)

func determineToken(meta LexMeta) {
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
	meta := LexMeta{
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
			meta = LexMeta{
				// FIXME: invert this var name
				OnlyNumbers: true,
			}

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
				meta = LexMeta{
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

		case '-':
			if meta.EscapeNext {
				meta.Accumulator += string(char)
				meta.EscapeNext = false
				continue
			}
			determineToken(meta)
			p.Tokens = append(p.Tokens, token.Token{
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
			determineToken(meta)
			p.Tokens = append(p.Tokens, token.Token{
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
			determineToken(meta)
			p.Tokens = append(p.Tokens, token.Token{
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
			determineToken(meta)
			p.Tokens = append(p.Tokens, token.Token{
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
			determineToken(meta)
			p.Tokens = append(p.Tokens, token.Token{
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
			determineToken(meta)
			p.Tokens = append(p.Tokens, token.Token{
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

			// 	p.Tokens = append(p.Tokens, token.Token{
			// 		Type: "LITERAL",
			// 		Value: token.Value{
			// 			Type:   "string",
			// 			String: meta.Accumulator,
			// 		},
			// 	})

			// 	meta.Accumulator = ""
			// 	// meta = LexMeta {
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

// TODO: rename this function and redo all comments/printouts to reflect that
func outputTokens() {
	lexFilename := p.Name + ".lex"

	// For more granular writes, open a file for writing.
	f, err := os.Create(lexFilename)
	defer func() {
		if err = f.Close(); err != nil {
			fmt.Println("ERROR: Could not close file:", lexFilename)
		}
	}()
	if err != nil {
		fmt.Println("ERROR: Could not open token output file:", lexFilename)
		os.Exit(9)
	}
	w := bufio.NewWriter(f)

	fmt.Println()
	fmt.Println("Outputting tokens to:", lexFilename)

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
				fmt.Println("ERROR: Could not write to token output file:", lexFilename)
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
				fmt.Println("ERROR: Could not write to token output file:", lexFilename)
				os.Exit(9)
			}
		}
	}

	err = w.Flush()
	if err != nil {
		fmt.Println("ERROR: Could not flush writer, data may be missing:", lexFilename)
	}
}

// TODO: should look into making english-like inputs for the indentation
// TODO: add flags for verbosity, printouts, whether to not parse the tokens, output file for outputting tokens, output format (xml, json, native) etc
// TODO: check if the file exists and if it does warn them that the output file will be overridden and ask if they still want to go through
// func parseFlags() {
// 	jsonIndentPtr := flag.String("jsonIndent", "blank", "Indent that will be used for the JSON printout of the tokens")
// 	flag.Parse()

// 	fmt.Println("jsonptr", *jsonIndentPctr)

// 	// jsonIndent
// 	jsonIndent = *jsonIndentPtr
// }

// TODO: we need a `getNextNonWSToken`

// func getFactor(i int) bool {
// 	fmt.Println("getting factor")

// 	// FIXME: we should first add a check for parens
// 	// FIXME: this will only work for spaces between them; hence the +2
// 	fmt.Println(p.Tokens[i].Type)
// 	if p.Tokens[i].Type == "LITERAL" {
// 		fmt.Println("Found a literal")
// 		endTokens = append(endTokens, p.Tokens[i])
// 		return true
// 	}

// 	return false
// }

// func llvmConversion(endTokens []token.Token) {
// 	llFilename := p.Name + ".ll"

// 	// For more granular writes, open a file for writing.
// 	f, err := os.Create(llFilename)
// 	defer func() {
// 		if err = f.Close(); err != nil {
// 			fmt.Println("ERROR: Could not close file:", llFilename)
// 		}
// 	}()
// 	if err != nil {
// 		fmt.Println("ERROR: Could not open token output file:", llFilename)
// 		os.Exit(9)
// 	}
// 	w := bufio.NewWriter(f)

// 	llvmInstructionString := ""

// 	// TODO: this needs to be outputted as program.expr.parse
// 	for i := 0; i < len(endTokens); i++ {
// 		t := endTokens[i]

// 		switch t.Type {
// 		case "TYPE":
// 			switch t.Value.String {
// 			case "int":
// 				// TODO: see if the variable declaration is something we already have
// 				llvmInstructionString += "%1 = alloca i32, align 4\n"
// 				// TODO: default value will force-find the next literal
// 			}
// 		case "LITERAL":
// 			llvmInstructionString += "store i32 " + t.Value.String + ", i32* %1, align 4"
// 		}

// 		// if t.Value.String == "int" {
// 		// 	llvmInstructionString += "%1 = alloca i32, align 4\n"
// 		// } else if t.Value.Type == "integer" {
// 		// 	llvmInstructionString += "store i32 5, i32* %1, align 4\n"
// 		// }
// 	}

// 	_, err = w.WriteString(llvmStart)
// 	if err != nil {
// 		fmt.Println("omggg!!!1")
// 		return
// 	}

// 	_, err = w.WriteString(llvmInstructionString + "\n")
// 	if err != nil {
// 		fmt.Println("omggg!!!2")
// 		return
// 	}

// 	_, err = w.WriteString(llvmEnd)
// 	if err != nil {
// 		fmt.Println("omggg!!!4")
// 		return
// 	}

// 	err = w.Flush()
// 	if err != nil {
// 		fmt.Println("ERROR: Could not flush writer, data may be missing:", llFilename)
// 	}
// }

func main() {
	// TODO: add some flags later
	// parseFlags()

	argLen := len(os.Args)

	if argLen < 2 {
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
	// TODO: change the name of this to accurately reflect lex vs parse tokens
	outputTokens()

	endTokens = parse.Parse(p.Tokens, p.Name)
	fmt.Println()
	fmt.Println("Endtokens:")

	for _, et := range endTokens {
		fmt.Println(et)
	}
	// llvmConversion(endTokens)
}
