package program

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/scottshotgg/Express/token"
)

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

// New returns a new Express program struct with initialized values
func New(programName string, compileStages map[string][]token.Token) (Program, error) {
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

// PrintTokens ...
func (p *Program) PrintTokens(stage, jsonIndent string) {
	for _, t := range p.Tokens[stage] {
		if t.Type == "BLOCK" || t.Type == "ARRAY" || t.Type == "GROUP" || t.Type == "FUNCTION" || t.Type == "ATTRIBUTE" {
			jsonIndent += "\t"

			po := Program{
				Tokens: map[string][]token.Token{
					"parse": t.Value.True.([]token.Token),
				},
			}

			fmt.Println()
			fmt.Println(jsonIndent[0:len(jsonIndent)-1] + t.Type)
			po.PrintTokens("parse", jsonIndent)

			jsonIndent = jsonIndent[0 : len(jsonIndent)-1]
			continue
		}

		tokenJSON, err := json.Marshal(t)
		if err != nil {
			fmt.Printf("\nERROR: Could not marshal JSON from token: %#v\n", t)
			os.Exit(9)
		}
		fmt.Println(jsonIndent + string(tokenJSON))
	}
}
