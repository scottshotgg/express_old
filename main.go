package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sgg7269/tokenizer/lex"
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
	Tokens map[string][]token.Token
	// Tokens []token.Token
}

var (
	jsonIndent string
	endTokens  = []token.Token{}
	// llvmStart  = "define i32 @main() #0 {\n"
	// llvmEnd    = "ret i32 0\n}"
	// compileStages = "lex", "sytactic", "semantic"

	// TODO: we should include filenames here and some metadata about the stages? Maybe we make the parse data here as well and pass in a pointer
	compileStages = map[string][]token.Token{
		"lex":   {},
		"parse": {},
	}
)

// NewProgram returns a new Express program struct with initialized values
func NewProgram(programName string) (Program, error) {
	input, err := ioutil.ReadFile(programName)
	if err != nil {
		fmt.Printf("ERROR: Cannot read input program: %s\n", programName)
		return Program{}, err
	}

	return Program{
		Value:  string(input),
		Name:   programName,
		Length: len(input),
		Tokens: compileStages,
	}, nil
	// might do this later, figure this out later
	// Tokens: func() {

	// 	for _, stage := range compileStages {

	// 	}
	// }(),
}

// TODO: make this take a compile stage
func (p *Program) PrintTokens(stage string) {

	if jsonIndent != "" {
		for _, token := range p.Tokens[stage] {
			tokenJSON, err := json.MarshalIndent(token, "", jsonIndent)
			if err != nil {
				fmt.Printf("\nERROR: Could not marshal JSON from token: %#v\n", token)
				os.Exit(9)
			}
			fmt.Println(string(tokenJSON))
		}
	} else {
		for _, token := range p.Tokens[stage] {
			tokenJSON, err := json.Marshal(token)
			if err != nil {
				fmt.Printf("\nERROR: Could not marshal JSON from token: %#v\n", token)
				os.Exit(9)
			}
			fmt.Println(string(tokenJSON))
		}
	}
}

func main() {
	// TODO: add some flags later
	// parseFlags()

	argLen := len(os.Args)

	if argLen < 2 {
		fmt.Println("ERROR: You must provide an input program")
		return
	}

	p, err := NewProgram(os.Args[argLen-1])
	if err != nil {
		fmt.Println("ERROR: Could not instantiate program structure", err)
		return
	}

	// TODO: this should go in the actual stage
	fmt.Println("=======================")
	fmt.Println()
	fmt.Println("Tokenizing:")
	fmt.Println()
	fmt.Println(p.Value)
	fmt.Println()
	fmt.Println("=======================")
	fmt.Println()

	p.Tokens["lex"], err = lex.Lex(p.Value)
	if err != nil {
		fmt.Println("ERROR: Could not lex input:", err)
		return
	}
	// fmt.Println(p.Tokens["lex"])
	// lex.Lex(p.Value)

	// TODO: this should go in the actual lex stage
	// p.Tokens = append(p.Tokens, token.Token{
	// 	Type: "EOF",
	// 	Value: token.Value{
	// 		String: string(0),
	// 	},
	// })

	p.PrintTokens("lex")

	// TODO: always output tokens right now
	// TODO: change the name of this to accurately reflect lex vs parse tokens
	// outputTokens()

	// TODO: fix this
	p.Tokens["parse"] = parse.Parse(p.Tokens["lex"], p.Name)
	fmt.Println()
	fmt.Println("Endtokens:")

	p.PrintTokens("parse")
}
