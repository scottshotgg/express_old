package parse

import (
	"fmt"
	"testing"

	"github.com/scottshotgg/Express/lex"
	"github.com/scottshotgg/Express/program"
	"github.com/scottshotgg/Express/token"
)

var (
	compileStages = map[string][]token.Token{
		"lex":   {},
		"parse": {},
	}
)

func TestCheckArray(t *testing.T) {
	p, err := program.New("array.expr", compileStages)
	if err != nil {
		fmt.Println("ERROR: Could not instantiate program structure", err)
		return
	}

	// Lex time
	p.Tokens["lex"], err = lex.Lex(p.Value)
	if err != nil {
		fmt.Println("ERROR: Could not lex input:", err)
		return
	}

	// Syntactic parse time
	p.Tokens["parse"], err = Parse(p.Tokens["lex"])
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	p.PrintTokens("parse", "\t")

	arrayTokens := p.Tokens["parse"][0].Value.True.([]token.Token)[2:]

	meta := Meta{
		IgnoreWS:         true,
		Tokens:           arrayTokens,
		Length:           len(arrayTokens),
		CheckOptmization: true,
		DeclarationMap:   map[string]token.Value{},
	}
	meta.Shift()
	meta.Shift()

	fmt.Println(meta.Tokens)

	meta.CheckArray()
}
