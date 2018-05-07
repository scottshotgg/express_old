package parse

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/sgg7269/tokenizer/token"
)

// Meta ...
type Meta struct {
	IgnoreWS        bool
	ParseIndex      int
	Length          int
	Tokens          []token.Token
	SemanticTokens  []token.Token
	SyntacticTokens []token.Token
	EndTokens       []token.Token

	LastToken    token.Token
	CurrentToken token.Token
	NextToken    token.Token

	LastCollectedToken token.Token
	//TokenLock sync.Mutex
}

// CollectToken ...
func (m *Meta) CollectToken(token token.Token) {
	m.LastCollectedToken = token
	m.EndTokens = append(m.EndTokens, token)
}

// CollectCurrentToken ...
func (m *Meta) CollectCurrentToken() {
	m.CollectToken(m.CurrentToken)
}

// CollectLastToken ...
func (m *Meta) CollectLastToken() {
	m.CollectToken(m.LastToken)
}

// GetNextNonWSToken ...
func (m *Meta) GetNextNonWSToken() (token.Token, error) {
	for {
		t, err := m.GetNextToken()
		if err != nil {
			werr := errors.Wrap(err, "m.GetNextToken()")
			fmt.Println("ERROR:", werr)
			return token.Token{}, werr
		}

		if t.Type == "WS" {
			continue
		}

		return t, nil
	}
}

// GetLastNonWSToken ...
func (m *Meta) GetLastNonWSToken() (token.Token, error) {
	for {
		t, err := m.GetLastToken()
		if err != nil {
			werr := errors.Wrap(err, "m.GetLastToken()")
			fmt.Println("ERROR:", werr)
			return token.Token{}, werr
		}

		if t.Type == "WS" {
			continue
		}

		return t, nil
	}
}

// PeekNextNonWSToken ...
func (m *Meta) PeekNextNonWSToken() (token.Token, error) {
	index := m.ParseIndex
	for {
		if index > -1 && index < m.Length {
			t, err := m.PeekTokenAtIndex(index)
			if err != nil {
				werr := errors.Wrap(err, "m.PeekNextNonWSToken()")
				fmt.Println("ERROR:", werr)
				return token.Token{}, werr
			}

			if t.Type == "WS" {
				continue
			}

			return t, nil
		}

		return token.Token{}, errors.New("Out of tokens")
	}
}

// PeekLastNonWSToken ...
func (m *Meta) PeekLastNonWSToken() (token.Token, error) {
	for {
		// t, err := m.PeekLastToken()
		// if err != nil {
		// 	werr := errors.Wrap(err, "m.PeekLastNonWSToken()")
		// 	fmt.Println("ERROR:", werr)
		// 	return token.Token{}, werr
		// }

		// if t.Type == "WS" {
		// 	continue
		// }

		return token.Token{}, nil
	}
}

// Create these if we need them
// GetNextNonWSTokenFromIndex ...
// GetLastNonWSTokenFromIndex ...

// GetNextToken ...
func (m *Meta) GetNextToken() (token.Token, error) {
	fmt.Println("get index", m.ParseIndex)
	if m.ParseIndex < m.Length {
		m.ParseIndex++
		return m.Tokens[m.ParseIndex], nil
	}

	return token.Token{}, errors.New("Out of tokens")
}

// PeekNextToken ...
func (m *Meta) PeekNextToken() token.Token {
	return m.NextToken
}

// GetLastToken ...
func (m *Meta) GetLastToken() (token.Token, error) {
	switch {
	case m.ParseIndex > 0:
		m.ParseIndex--
		fallthrough
	case m.ParseIndex == 0:
		return m.Tokens[m.ParseIndex], nil
	default:
		return token.Token{}, errors.New("Already at last token")
	}
}

// PeekLastToken ...
func (m *Meta) PeekLastToken() token.Token {
	return m.LastToken
}

// PeekLastToken ...
func (m *Meta) PeekLastCollectedToken() token.Token {
	return m.LastCollectedToken
}

// GetCurrentToken ...
func (m *Meta) GetCurrentToken() (token.Token, error) {
	if m.ParseIndex > -1 && m.ParseIndex < m.Length {
		return m.Tokens[m.ParseIndex], nil
	}

	return token.Token{}, errors.New("Current parseIndex outside of token range")
}

// GetTokenAtIndex ...
func (m *Meta) PeekTokenAtIndex(index int) (token.Token, error) {
	if index > -1 && index < m.Length {
		return m.Tokens[index], nil
	}

	return token.Token{}, errors.New("Current parseIndex outside of token range")
}

// Shift ...
func (m *Meta) Shift() {
	m.LastToken = m.CurrentToken

	m.CurrentToken = m.Tokens[m.ParseIndex]

	for {
		fmt.Println("m.ParseIndex", m.ParseIndex+1, "m.Length", m.Length)
		if m.ParseIndex+1 < m.Length {
			m.ParseIndex++
			if m.Tokens[m.ParseIndex].Type == "WS" {
				continue
			}

			m.NextToken = m.Tokens[m.ParseIndex]
			return
		}

		m.NextToken = token.Token{}
		return
	}
}

// ParseType ...
func (m *Meta) ParseType(token token.Token, index int) error {
	m.Shift()

	switch m.CurrentToken.Type {
	case "LITERAL":
		m.CurrentToken.Type = "IDENT"

		// TODO: this will prevent us from doing var declarations
		m.CurrentToken.Expected = "ASSIGN"

		m.CollectCurrentToken()
	}

	return nil
}

// GetFactor ...
func (m *Meta) GetFactor() {
	fmt.Println("getting ze factor", m.CurrentToken)

	switch m.CurrentToken.Type {
	case "LITERAL":
		m.CollectCurrentToken()
	default:
		fmt.Println("got something other than LITERAL at GetExpr")
	}
}

// GetTerm ...
func (m *Meta) GetTerm() {
	fmt.Println("getting ze term", m.CurrentToken)

	switch m.CurrentToken.Type {
	case "LITERAL":
		m.GetFactor()
	default:
		fmt.Println("got something other than LITERAL at GetExpr")
	}
}

// GetExpr ...
func (m *Meta) GetExpr() {
	fmt.Println("getting ze expr", m.CurrentToken)

	switch m.CurrentToken.Type {
	case "LITERAL":
		m.GetTerm()
	default:
		fmt.Println("got something other than LITERAL at GetExpr")
	}
}

// GetStatement ...
func (m *Meta) GetStatement() {
	fmt.Println("getting ze statement", m.CurrentToken)

	switch m.CurrentToken.Type {
	case "TYPE":
		m.CollectCurrentToken()
		m.Shift()
		m.GetExpr()
	case "LITERAL":
		m.GetExpr()
	default:
		fmt.Println("got something other than LITERAL at GetStatement")
	}

	m.Shift()

	if m.CurrentToken.Type == "EOS" {
		m.CollectCurrentToken()
	}
}

// ParseIdent ...
func (m *Meta) ParseIdent(blockTokens *[]token.Token, peek token.Token) {
	if blockTokens == nil {
		fmt.Println("ERROR: blockTokens is nil")
		os.Exit(5)
	}

	identSplit := strings.Split(peek.Value.String, ".")
	for i, ident := range identSplit {
		fmt.Println("IDENTSPLITFOR", ident)
		*blockTokens = append(*blockTokens, token.Token{
			ID:   0,
			Type: "IDENT",
			// Expected:
			Value: token.Value{
				// Type: ,
				// True: ,
				String: ident,
			},
		})

		if i < len(identSplit)-1 {
			*blockTokens = append(*blockTokens, token.TokenMap["."])
		}
	}
}

// TokenToString ...
func TokenToString(t token.Token) string {
	jsonToken, err := json.Marshal(t)
	if err != nil {
		return err.Error()
	}

	return string(jsonToken)
}

// ParseBlock ..
func (m *Meta) ParseBlock() token.Token {
	// FIXME: could do something fancy with another meta and then use that but w/e
	blockTokens := []token.Token{}

	for {
		m.Shift()

		current := m.CurrentToken

		switch current.Type {
		case "SEPARATOR":
			fallthrough
		case "EOS":
			// TODO: this will need to check the last and next token type later to determine wtf to do
			blockTokens = append(blockTokens, m.CurrentToken)

		case "WS":
			continue

		case "TYPE":
			peek := m.PeekNextToken()
			switch peek.Type {
			case "IDENT":
				m.ParseIdent(&blockTokens, m.CurrentToken)

			case "LITERAL":
				blockTokens = append(blockTokens, m.CurrentToken)

				m.Shift()
				m.CurrentToken.Type = "IDENT"
				blockTokens = append(blockTokens, m.CurrentToken)
			default:
				fmt.Println("wtf", m)
				os.Exit(7)
			}

		case "ASSIGN":
			blockTokens = append(blockTokens, m.CurrentToken)

		case "SET":
			peek := m.PeekNextToken()
			switch peek.Type {
			case "ASSIGN":
				if t, ok := token.TokenMap[current.Value.String+peek.Value.String]; ok {
					blockTokens = append(blockTokens, t)
					m.Shift()
				}
			// TODO: AHA! we need to make a parse expr and stuff here
			// FIXME: we need to make something to take care of "SET", just check the next token for now
			// case "IDENT":

			// 	// FIXME: this is a hacked in thing; REALLY need to get blocks bootstrapped
			// 	m.ParseIdent(&blockTokens, m.CurrentToken)
			// 	// for _, t := range blockTokens {
			// 	// 	m.CollectToken(t)
			// 	// }
			default:
				blockTokens = append(blockTokens, current)
				fmt.Println("WTF", peek)
				continue
			}

		case "IDENT":
			m.ParseIdent(&blockTokens, m.CurrentToken)
			// TODO: this case might need to move to the Syntactic part of the parser
		case "LITERAL":
			// TODO: this may cause some problems
			switch m.PeekLastCollectedToken().Type {
			case "SET":
				fallthrough
			case "ASSIGN":
				fallthrough
			case "INIT":
				blockTokens = append(blockTokens, m.CurrentToken)
			}

		case "L_BRACKET":
			blockTokens = append(blockTokens, m.ParseArray())

		case "L_BRACE":
			// blockTokens = append(blockTokens, m.CurrentToken)
			blockTokens = append(blockTokens, m.ParseBlock())

		case "R_BRACE":
			// blockTokens = append(blockTokens, m.CurrentToken)
			return token.Token{
				ID:   0,
				Type: "BLOCK",
				// Expected: TODO: do the same thing that we did on the array but use the meta tokens
				Value: token.Value{
					Type: "block",
					True: blockTokens,
					// String: TODO: do the same thing that we did on array
				},
			}

		case "D_QUOTE":
			blockTokens = append(blockTokens, m.ParseString())

		default:
			fmt.Println("IDK WTF TO DO with this token", m.CurrentToken)
			os.Exit(6)
		}
	}
}

// // TODO: I don't think we'll need this
// func (m *Meta) ParseMap() token.Token {}

// ParseArray ...
// TODO: we could make an array a BLOCK of statements using a separator ",", thus we wouldn't have to do anything special for an array
func (m *Meta) ParseArray() token.Token {
	// arrayToken := token.Token{
	// 	ID:   1,
	// 	Type: "ARRAY",
	// 	// Expected: TODO: calc this later
	// 	Value: token.Value{
	// 		Type: "ARRAY",
	// 		True: ,
	// 	},
	// }

	arrayTokens := []token.Token{}

	for {
		m.Shift()

		switch m.CurrentToken.Type {
		case "SEPARATOR":
			continue

		case "D_QUOTE":
			arrayTokens = append(arrayTokens, m.ParseString())
		// case "LITERAL":

		case "LITERAL":
			fmt.Println("hi", m.CurrentToken)
			arrayTokens = append(arrayTokens, m.CurrentToken)

		case "L_BRACE":
			arrayTokens = append(arrayTokens, m.ParseBlock())

		case "L_BRACKET":
			arrayTokens = append(arrayTokens, m.ParseArray())

		case "R_BRACKET":
			return token.Token{
				ID:   1,
				Type: "ARRAY",
				// Expected: TODO: calc this later
				Value: token.Value{
					Type: "array",
					True: arrayTokens,
					// String: func() (arrayTokensString string) {
					// 	for _, t := range arrayTokens {
					// 		arrayTokensString += TokenToString(t)
					// 	}

					// 	return
					// }(),
				},
			}

		default:
			fmt.Println("ERROR: Unrecognized array token\n", m)
			os.Exit(8)
		}

		// FIXME: This should throw an error
		if m.NextToken == (token.Token{}) {
			return token.Token{}
		}
	}
}

// ParseString ...
func (m *Meta) ParseString() token.Token {
	stringLiteral := ""
	for {
		stringLiteral += m.CurrentToken.Value.String
		m.Shift()

		if m.NextToken.Value.String == "\"" {
			stringLiteral += m.CurrentToken.Value.String
			m.Shift()
			stringLiteral += m.CurrentToken.Value.String

			return token.Token{
				Type: "LITERAL",
				// Expected: func() string {
				// 	switch m.PeekLastCollectedToken().Type {
				// 	case "SET":
				// 		fallthrough
				// 	case "ASSIGN":
				// 		fallthrough
				// 	case "INIT":
				// 		return "EOS" // it would be cool if we could make expected an array
				// 	default:
				// 		fmt.Println("WTF TO DO", m)
				// 		os.Exit(9)
				// 	}

				// 	return ""
				// }(),
				Value: token.Value{
					Type:   "string",
					True:   stringLiteral[1 : len(stringLiteral)-1],
					String: stringLiteral,
				},
				// Expected TODO: maybe do this later based on the last one
			}
		}
	}
}

// ApplyPressure ...
func (m *Meta) ApplyPressure() {
	current := m.CurrentToken

	switch current.Type {
	case "SEPARATOR":
		fallthrough
	case "EOS":
		// TODO: this will need to check the last and next token type later to determine wtf to do
		m.CollectCurrentToken()

	case "WS":
		return

	case "TYPE":
		fmt.Println("GETTING TYPE", m.CurrentToken)

		switch m.PeekNextToken().Type {
		case "IDENT":
			m.CollectCurrentToken()

			m.Shift()

			// FIXME: this is a hacked in thing; REALLY need to get blocks bootstrapped
			var blockTokens []token.Token
			m.ParseIdent(&blockTokens, m.CurrentToken)
			for _, t := range blockTokens {
				m.CollectToken(t)
			}

		case "LITERAL":
			m.CollectCurrentToken()

			m.Shift()
			m.CurrentToken.Type = "IDENT"
			m.CollectCurrentToken()

		default:
			fmt.Println("wtf", m)
			os.Exit(7)
		}

	case "ASSIGN":
		m.CollectCurrentToken()

	case "SET":
		peek := m.PeekNextToken()
		switch peek.Type {
		case "ASSIGN":
			if t, ok := token.TokenMap[current.Value.String+peek.Value.String]; ok {
				m.CollectToken(t)
				m.Shift()
			}
		// TODO: AHA! we need to make a parse expr and stuff here
		// FIXME: we need to make something to take care of "SET", just check the next token for now
		default:
			m.CollectCurrentToken()
		}

		// TODO: this case might need to move to the Syntactic part of the parser
	case "LITERAL":
		// TODO: this may cause some problems
		switch m.PeekLastCollectedToken().Type {
		case "SET":
			fallthrough
		case "ASSIGN":
			fallthrough
		case "INIT":
			m.CollectCurrentToken()
		}

	case "IDENT":
		fmt.Println("IDENT", m.CurrentToken)
		// FIXME: this is a hacked in thing; REALLY need to get blocks bootstrapped
		var blockTokens []token.Token
		m.ParseIdent(&blockTokens, m.CurrentToken)
		for _, t := range blockTokens {
			m.CollectToken(t)
		}

	case "L_BRACKET":
		m.CollectToken(m.ParseArray())

	case "L_BRACE":
		// m.CollectCurrentToken()
		m.CollectToken(m.ParseBlock())

	case "R_BRACE":
		// m.CollectCurrentToken()
		return

	case "D_QUOTE":
		m.CollectToken(m.ParseString())

	default:
		fmt.Println("IDK WTF TO DO with this token", m.CurrentToken)
		os.Exit(6)
	}
}

// SemanticPressure ...
func (m *Meta) SemanticPressure() {
	for {
		m.Shift()

		m.ApplyPressure()

		if m.NextToken == (token.Token{}) {
			return
		}
	}
}

// Parse ...
func Parse(tokens []token.Token, name string) ([]token.Token, error) {
	fmt.Println("parsing")

	meta := Meta{
		IgnoreWS: true,
		// ParseIndex: -1,
		Tokens: tokens,
		Length: len(tokens),
	}

	meta.SemanticPressure()
	// meta.SyntacticParse()

	// meta.Shift()
	// meta.GetStatement()
	// fmt.Println("meta", meta)

	// for {
	// 	meta.Shift()

	// 	switch meta.CurrentToken.Type {
	// 	case "EOS":
	// 		// TODO: meta.GetStatement() here
	// 	case "TYPE":
	// 		meta.CollectToken(meta.CurrentToken)
	// 		meta.ParseType(meta.CurrentToken, 0)
	// 	default:
	// 	}

	// 	if meta.NextToken == (token.Token{}) {
	// 		return meta.EndTokens, nil
	// 	}
	// 	// if err != nil {
	// 	// 	werr := errors.Wrap(err, "meta.GetNextToken()")
	// 	// 	fmt.Println("ERROR:", werr)

	// 	// 	break
	// 	// }

	// 	// fmt.Println(t)

	// 	// switch t.Type {
	// 	// case "WS":
	// 	// 	continue
	// 	// case "TYPE":
	// 	// 	fmt.Println("found a type")
	// 	// 	err := meta.ParseType(t, meta.ParseIndex)
	// 	// 	if err != nil {
	// 	// 		fmt.Println("ERROR:", err)
	// 	// 	}
	// 	// }
	// }

	return meta.EndTokens, nil
}
