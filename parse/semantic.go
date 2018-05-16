package parse

import (
	"fmt"
	"os"

	"github.com/scottshotgg/Express/token"
)

// CheckType ...
func (m *Meta) CheckType() {
	fmt.Println("CheckType")
	// previous := m.LastToken
	// current := m.CurrentToken
	// next := m.NextToken

	switch m.NextToken.Type {
	case "IDENT":
		// append the current and next tokens
		// check for init/set/assign
		// check for value/ident

		m.CollectCurrentToken()
		m.Shift()
		m.CollectCurrentToken()

		switch m.NextToken.Type {
		case "INIT":
			fallthrough
		case "SET":
			fallthrough
		case "ASSIGN":
			fmt.Println("found an ASSIGN type")
			m.Shift()
			m.CollectCurrentToken()

			fmt.Println(m.NextToken)
			if m.NextToken.Type != "LITERAL" && m.NextToken.Type != "IDENT" {
				fmt.Println("did not find a literal")
				os.Exit(0)
			}

			m.Shift()
			m.CollectCurrentToken()
		}
	}
}

// CheckBlock ...
func (m *Meta) CheckBlock() {
	fmt.Println("hi")

	for {
		m.Shift()

		current := m.CurrentToken
		switch current.Type {
		case "TYPE":
			// at this point we need to think about the different options
			// int name			: simplest
			// int name = 6 : next simplest
			// name int = 6 : would be taken care of by a random ident
			// we would probably also have other ones for function params and stuff

			m.CheckType()

			if m.NextToken == (token.Token{}) {
				fmt.Println("returning")
				return
			}
		}
	}
}

// Semantic ...
func Semantic(tokens []token.Token) ([]token.Token, error) {
	// Auto inject the brackets to ensure that they are there
	meta := Meta{
		IgnoreWS:         true,
		Tokens:           tokens[0].Value.True.([]token.Token),
		Length:           len(tokens[0].Value.True.([]token.Token)),
		CheckOptmization: true,
	}

	fmt.Println(tokens)

	meta.CheckBlock()
	fmt.Println("tokens", meta.EndTokens)

	// Here we are continuously applying semantic pressure to squash the tokens and furthur
	// simplify the tokens generated
	// for meta.CheckOptmization {
	// 	fmt.Println("Optimizing", meta.OptimizationAttempts)
	// 	// meta.CollectTokens(meta.ParseBlock().Value.True.([]token.Token))
	// 	fmt.Println("endTokens", meta.EndTokens)

	// 	fmt.Println(meta.EndTokens)
	// 	metaTokens := meta.EndTokens[0].Value.True.([]token.Token)
	// 	metaTokens = append(append([]token.Token{token.TokenMap["{"]}, metaTokens...), token.TokenMap["}"])
	// 	fmt.Println("metaTokens", len(metaTokens), len(meta.EndTokens))

	// 	// endTokens = meta.EndTokens

	// 	// TODO: FIXME: w/e this works for now
	// 	// Fix this from pulling off only the top one
	// 	// Only apply SemanticPressure once for now until we figure out the recursion more
	// 	if meta.OptimizationAttempts > 0 {
	// 		break
	// 	}

	// 	// fmt.Println("meta.CheckOptimization", meta.CheckOptmization)

	// 	// if !meta.CheckOptmization {
	// 	// 	break
	// 	// }

	// 	// if len(meta.EndTokens) < len(meta.Tokens) {
	// 	// 	break
	// 	// }

	// 	meta = Meta{
	// 		// FIXME: do we need to fix this hack?
	// 		// Tokens: ,
	// 		Tokens:               metaTokens,
	// 		Length:               len(metaTokens),
	// 		CheckOptmization:     meta.CheckOptmization,
	// 		OptimizationAttempts: meta.OptimizationAttempts + 1,
	// 	}
	// }

	return meta.EndTokens, nil
}
