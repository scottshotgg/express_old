package main

import (
	"fmt"
	"os"

	c "github.com/scottshotgg/Express/c"
	"github.com/scottshotgg/Express/lex"
	"github.com/scottshotgg/Express/parse"
	program "github.com/scottshotgg/Express/program"
	"github.com/scottshotgg/Express/token"
	//"llvm.org/llvm/bindings/go/llvm"
)

// TODO: we currently have to have the space afterwards, need to add the parsing code into the semicolon code
// TODO: we currently only have numbers being parsed within floats, so nothing like '9.9f', etc

var (
	jsonIndent = "\t"
	// endTokens  = []token.Token{}
	// llvmStart  = "define i32 @main() #0 {\n"
	// llvmEnd    = "ret i32 0\n}"
	// compileStages = "lex", "sytactic", "semantic"

	// TODO: we should include filenames here and some metadata about the stages? Maybe we make the parse data here as well and pass in a pointer
	compileStages = map[string][]token.Token{
		"lex":   {},
		"parse": {},
	}
)

func main() {
	// TODO: add some flags later
	// parseFlags()

	argLen := len(os.Args)

	if argLen < 2 {
		fmt.Println("ERROR: You must provide an input program")
		return
	}

	p, err := program.New(os.Args[argLen-1], compileStages)
	if err != nil {
		fmt.Println("ERROR: Could not instantiate program structure", err)
		return
	}

	// too lazy to do a cool print out right now
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

	p.PrintTokens("lex", jsonIndent)

	// TODO: always output tokens right now
	// TODO: change the name of this to accurately reflect lex vs parse tokens
	// outputTokens()

	// TODO: fix this to return an err
	fmt.Println()

	// Syntactic parse time
	p.Tokens["parse"], err = parse.Parse(p.Tokens["lex"])
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	fmt.Println("\nSemantic tokens:")
	p.PrintTokens("parse", jsonIndent)
	fmt.Println()

	// Semantic parse time
	p.Tokens["semantic"], err = parse.Semantic(p.Tokens["parse"])
	if err != nil {
		fmt.Println("wat dat err do brah", err)
		os.Exit(1)
	}

	// llvm.Translate(p.Tokens["semantic"])
	c.Translate(p.Tokens["semantic"])

	// fmt.Println(string(semanticTokensJSON))

	// err = ioutil.WriteFile("main.tokens.json", []byte(string(semanticTokensJSON)), 0644)
	// if err != nil {
	// 	fmt.Println("ERROR", err)
	// 	os.Exit(9)
	// }
}
