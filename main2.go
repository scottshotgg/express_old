// generic determine token
// vector operations
// regex for literals
// [math][assign] operations
// keywords

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang-collections/collections/stack"
	"github.com/sgg7269/tokenizer/token"
)

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
	HasLetters  bool
	Enclosed    struct {
		Value   byte
		Matched bool
	}
}

// ParseMeta ...
type ParseMeta struct {
	Accumulator string
	EscapeNext  bool
	Period      bool
	HasLetters  bool
	Expect      string
	Enclosed    struct {
		Value   byte
		Matched bool
		Stack   stack.Stack
	}
}

var (
	p          Program
	jsonIndent string
)

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

// TODO: optimize out this stupid hack that we have because we have whitespace
// TODO: we could keep track of the tokens that we see
func getLastNonWSToken(index int) int {
	index--
	for ; index > -1; index-- {
		// fmt.Println("getLastNonWSToken", index, p.Tokens[index])
		if p.Tokens[index].Type == "WS" {
			continue
		}
		// fmt.Println("this is what we are returning", index)
		return index
	}

	return -1
}

// TODO: Maybe we can somehow make a double map?
func getExpect(lastToken, currentToken *token.Token) string {
	fmt.Println("getExpect", lastToken, currentToken)

	// TODO: don't know if we need to do this
	// TODO: there are other times that we will need this but as of now, the tokens are set to the expected os we dont
	if lastToken.Expected == currentToken.Type {
		switch lastToken.Type {
		case "TYPE":
			// FIXME: something needs to be done about this, hashmaps will be too much I think
			if currentToken.Type == "IDENT" {
				// TODO: this is where we could use the multple expectations
				// TODO: either that or make an "expectation map", but this would have to take into account the token type
				// 			 before hand and would encompass what we are doing right now into a data structure
				currentToken.Expected = "ASSIGN" // FIXME: this will mean that we can only accept equality statements for variable declarations
			}
			/*	Don't know what to do with this for now
				else if currentToken.Type == "EXPR" {
					fmt.Println("looking for an expressiong ")
				}
			*/
		}
	}

	return ""
}

func getExpr(index int) {
	fmt.Println(index)

	// for ; index < len(p.Tokens); index++ {
	// 	switch p.Tokens[index]
	// }
}

func parse() {
	meta := ParseMeta{}
	// s := stack.New()
	// fmt.Println(s, meta)

	// To parse:
	// get a token
	// if it has an expectation - look for that

	for i, t := range p.Tokens {
		// fmt.Println("meta.Expected", meta.Expect)

		if t.Type == "WS" {
			continue
		}

		if to, ok := token.TokenMap[t.Value.String]; ok {
			fmt.Println("to.Expected", to.Expected)

			fmt.Println("to", to)

			if meta.Expect == "" {
				// fmt.Println("meta.Expect", to.Value.String)

				meta.Expect = to.Expected
			} else if to.Type == meta.Expect {
				fmt.Println("hi")
			} else if to.Type == "LITERAL" {
				// p.Tokens[i] =
				// fmt.Println("hi2")
				switch meta.Expect {
				case "IDENT":
					fmt.Println("found an ident")
				}
			} else if to.Expected == "EXPR" {
				fmt.Println("looking for an expressiong ")
				getExpr(i)
			}
			// FIXME: we need to put logic to determine the next one based on the last and current one
		} else {
			fmt.Println("token that didnt make it", t)

			if meta.Expect != "" {
				// fmt.Println("1token that didnt make it", t)
				// fmt.Println(t, "hey its me")
				if t.Type == meta.Expect {
					// fmt.Println("equal", t)
				} else if t.Type == "LITERAL" {
					p.Tokens[i].Type = meta.Expect

					switch meta.Expect {
					case "IDENT":
						fmt.Println("found55 an ident")
					}
				}

				lastIndex := getLastNonWSToken(i)
				if lastIndex == -1 {
					fmt.Println("got a -1")
					return
				}

				meta.Expect = getExpect(&p.Tokens[lastIndex], &p.Tokens[i])
				fmt.Println(meta.Expect)

			} else {
				fmt.Println("i dont fucking know")
			}
		}
	}
}

func lex() {
	meta := LexMeta{}

	for _, c := range p.Value {
		// FIXME: TODO: AHA! We need to make SEPARATE TokenMap and SymbolMap. The token map is for parsing, symbol map is for lexing
		if t, ok := token.TokenMap[string(c)]; ok {
			if meta.Accumulator != "" {
				if to, ok := token.TokenMap[meta.Accumulator]; ok {
					p.Tokens = append(p.Tokens, to)
				} else {
					p.Tokens = append(p.Tokens, token.Token{
						Type: "LITERAL",
						Value: token.Value{
							Type:   "unknown",
							String: meta.Accumulator,
						},
					})
				}
			}

			p.Tokens = append(p.Tokens, t)
			meta = LexMeta{}
			continue
		}
		meta.Accumulator += string(c)
	}
}

func main() {
	programName := "program.expr"

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

	p.Tokens = append(p.Tokens, token.Token{
		Type: "EOF",
		Value: token.Value{
			String: string(0),
		},
	})

	printTokens()

	fmt.Println()
	fmt.Println()
	fmt.Println("=======================")
	fmt.Println()
	fmt.Println("Parsing:")
	fmt.Println()
	fmt.Println("=======================")
	fmt.Println()
	fmt.Println()

	parse()

	printTokens()

	// floatLiteral := "[0-9]*.[0-9]+"
	// intLiteral := "[0-9]+"
	// stringLiteral := "\".*\""

	// match, _ := regexp.MatchString(stringLiteral, "\".9\"")
	// fmt.Println(match)

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

	// // TODO: always output tokens right now
	// outputTokens()
}
