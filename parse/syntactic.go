package parse

import (
	"fmt"
	"os"
	"strings"

	"github.com/scottshotgg/Express/token"
)

// Meta holds information about the current parse
type Meta struct {
	AppendVar       bool
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

	CheckOptmization     bool
	OptimizationAttempts int

	// hacky
	LLVMTokens []token.Token

	InheritedMap       map[string]token.Value
	DeclarationMap     map[string]token.Value
	DeclaredType       string
	DeclaredActingType string
	DeclaredName       string
	DeclaredValue      token.Value
	DeclaredAccessType string
}

// ParseVar parses a variable declaration and other statements related to type later on. Anything in the form of <type><ident>
func (m *Meta) ParseVar(t token.Token) error {
	m.Shift()

	switch m.CurrentToken.Type {
	case token.Literal:
		m.CurrentToken.Type = token.Ident

		m.CurrentToken.Expected = token.Assign

		m.CollectCurrentToken()
	}

	return nil
}

// ParseIdent parses an identifier
func (m *Meta) ParseIdent(blockTokens *[]token.Token, peek token.Token) {
	m.CheckOptmization = true

	if blockTokens == nil {
		fmt.Println("ERROR: blockTokens is nil")
		os.Exit(5)
	}

	identSplit := strings.Split(peek.Value.String, ".")
	for i, ident := range identSplit {
		*blockTokens = append(*blockTokens, token.Token{
			ID:   0,
			Type: token.Ident,
			// Expected:
			Value: token.Value{
				Type: func() string {
					if len(ident) > 0 && ident[0] > 64 && ident[0] < 91 {
						return "public"
					}

					return "private"
				}(),
				// True: ,
				String: ident,
			},
		})

		if i < len(identSplit)-1 {
			*blockTokens = append(*blockTokens, token.TokenMap["."])
		}
	}
}

// ParseAttribute parses an attribute as defined by ECMA-335 standards. Anything defined by ECMA-335; #[attribute]
// TODO: full ecma335 implementation will go here when i feel like it
func (m *Meta) ParseAttribute() token.Token {
	var ecmaTokens []token.Token
	ecmaToken := token.Token{
		ID:   6,
		Type: token.Attribute,
		// Expected:
		Value: token.Value{
			Type: "statement",
			// String:
		},
	}

	if m.NextToken.Type == token.Bang {
		fmt.Println("applied to scope")
		ecmaToken.Value.Type = "scope"
		// ecmaTokens = append(ecmaTokens, m.CurrentToken)
		m.Shift()
	}

	for {
		m.Shift()

		current := m.CurrentToken

		switch current.Type {

		// attribute of next item down
		case token.LBracket:
			// TODO: should this just be an array of statements that get parsed like anything else?
			// TODO: programmatic compiler directives/attributes?
			// TODO: at this point we could just call an entirely separate ECMA-335 parser
		case token.Ident:
			ecmaTokens = append(ecmaTokens, current)

		case token.RBracket:
			ecmaToken.Value.True = ecmaTokens
			return ecmaToken

		default:
			fmt.Println("idk wtf to do", m)
			os.Exit(9)
		}
	}
}

// ParseSQL will parse inline SQL statements
func (m *Meta) ParseSQL() token.Token {
	return token.Token{}
}

// ParseHTML will parse inline HTML
func (m *Meta) ParseHTML() token.Token {
	return token.Token{}
}

// ParseCSS will parse inline CSS
func (m *Meta) ParseCSS() token.Token {
	return token.Token{}
}

// ParseMarkup will parse inline Markup
func (m *Meta) ParseMarkup() token.Token {
	return token.Token{}
}

// ParseJSON will parse JSON
func (m *Meta) ParseJSON() token.Token {
	return token.Token{}
}

// ParseXML will parse XML
func (m *Meta) ParseXML() token.Token {
	return token.Token{}
}

// ParseNoSQL will parse NoSQL
// this might use graphQL internally
// idk how this is going to work if at all
func (m *Meta) ParseNoSQL() token.Token {
	return token.Token{}
}

// ParseFunctionDef parses a function definition. Anything in the form <ident><group><group><block> or <ident><group><block>
func (m *Meta) ParseFunctionDef(current token.Token) token.Token {
	m.CheckOptmization = true

	var functionTokens []token.Token

	m.ParseIdent(&functionTokens, current)
	m.Shift()
	argumentTokens := m.ParseGroup()
	functionTokens = append(functionTokens, argumentTokens.Value.True.([]token.Token)...)

	if m.NextToken.Type == token.LParen {
		fmt.Println("Found return tokens")
		argumentTokens = m.ParseGroup()
		functionTokens = append(functionTokens, argumentTokens.Value.True.([]token.Token)...)
	}

	// add these tokens to the function tokens and return that token
	// return append(functionTokens, groupToken.Value.True.([]token.Token)...)
	// functionTokens = append(functionTokens, argumentTokens.Value.True.([]token.Token)...)
	return token.Token{
		ID:   4,
		Type: token.Function,
		// Expected: //TODO:
		Value: token.Value{
			Type: "def",
			True: functionTokens,
			// String: //TODO:
		},
	}
}

// ParseFunctionCall parses a function call. Anything in the form <ident><group>
func (m *Meta) ParseFunctionCall(current token.Token) token.Token {
	m.CheckOptmization = true

	// var functionTokens [][]token.Token
	var functionTokens []token.Token
	// var returnTokens []token.Token

	m.ParseIdent(&functionTokens, current)
	m.Shift()
	// FIXME: TODO: these should all return errors
	argumentTokens := m.ParseGroup()

	// add these tokens to the function tokens and return that token
	// return append(functionTokens, groupToken.Value.True.([]token.Token)...)
	return token.Token{
		ID:   4,
		Type: token.Function,
		// Expected: //TODO:
		Value: token.Value{
			Type: "call",
			True: append(functionTokens, argumentTokens.Value.True.([]token.Token)...),
			// String: //TODO:
		},
	}
}

// ParseGrave will parse the graves (backticks)
func (m *Meta) ParseGrave() token.Token {
	return token.Token{}
}

// ParseCharOrEscapedString will parse chars and escaped strings. Anything encapsulated in singular quotes.
func (m *Meta) ParseCharOrEscapedString() token.Token {
	return token.Token{}
}

// ParseString parses a string literal. Anything surrounded by quotes.
func (m *Meta) ParseString() token.Token {
	m.CheckOptmization = true

	stringLiteral := ""
	for {
		m.ShiftWithWS()
		fmt.Println("current", m.CurrentToken)

		// FIXME: stop doing hacky shit, purge this shit, need to preserve whitespaces in the lexer
		stringLiteral += m.CurrentToken.Value.String
		if m.NextToken.Value.String == "\"" {

			m.ShiftWithWS()

			return token.Token{
				Type: token.Literal,
				Value: token.Value{
					Type:   "string",
					True:   stringLiteral,
					String: stringLiteral,
				},
			}
		}
		// Getting the last 'separating' character; aka a whitespace that was separating the tokens
	}
}

// ParseGroup parses a grouping of items; tuple, function arguments, function returns. Anything encapsulated in parenthesis.
func (m *Meta) ParseGroup() token.Token {
	m.CheckOptmization = true

	groupTokens := []token.Token{}

	for {
		m.Shift()

		current := m.CurrentToken

		switch current.Type {
		case token.RParen:
			return token.Token{
				ID:   1,
				Type: token.Group,
				Value: token.Value{
					Type: token.Group,
					True: groupTokens,
				},
			}

		case token.Literal:
			groupTokens = append(groupTokens, current)

		case token.Type:
			peek := m.NextToken
			switch peek.Type {
			case token.LBracket:
				fmt.Println("found array")
				os.Exit(8)

			case token.Ident:
				m.ParseIdent(&groupTokens, m.CurrentToken)

			case token.Literal:
				groupTokens = append(groupTokens, m.CurrentToken)

				m.Shift()
				m.CurrentToken.Type = token.Ident
				groupTokens = append(groupTokens, m.CurrentToken)
			default:
				os.Exit(7)
			}

		case token.Ident:
			m.ParseIdent(&groupTokens, m.CurrentToken)

		case token.Separator:
			continue

		case token.DQuote:
			groupTokens = append(groupTokens, m.ParseString())

		case token.LBrace:
			groupTokens = append(groupTokens, m.ParseBlock())

		case token.LBracket:
			groupTokens = append(groupTokens, m.ParseArray())

		default:
			fmt.Println("ERROR: Unrecognized group token\n", current, m)
			os.Exit(8)
		}
	}
}

// ParseArray parses an array of items. Anything encapulated in square brackets except for attributes.
func (m *Meta) ParseArray() token.Token {
	m.CheckOptmization = true

	arrayTokens := []token.Token{}

	for {
		m.Shift()

		switch m.CurrentToken.Type {
		case token.Separator:
			continue

		case token.Ident:
			m.ParseIdent(&arrayTokens, m.CurrentToken)

		case token.DQuote:
			arrayTokens = append(arrayTokens, m.ParseString())

		case token.Literal:
			arrayTokens = append(arrayTokens, m.CurrentToken)

		case token.LParen:
			arrayTokens = append(arrayTokens, m.ParseGroup())

		case token.LBrace:
			arrayTokens = append(arrayTokens, m.ParseBlock())

		case token.LBracket:
			arrayTokens = append(arrayTokens, m.ParseArray())

		case token.RBracket:
			return token.Token{
				ID:   1,
				Type: token.Array,
				Value: token.Value{
					Type: token.ArrayType,
					True: arrayTokens,
				},
			}

		case token.SecOp:
			arrayTokens = append(arrayTokens, m.CurrentToken)

		case "":
			fmt.Println("we got nothing")

		default:
			fmt.Println("ERROR: Unrecognized array token\n", m.CurrentToken, m)
			os.Exit(8)
		}

		// // FIXME: This should throw an error
		// if m.NextToken == (token.Token{}) {
		// 	fmt.Println("nextToken array", arrayTokens)
		// 	return token.Token{}
		// }
	}
}

// ParseBlock parses the center piece of the language; the block. Anything encapulated in curly braces.
func (m *Meta) ParseBlock() token.Token {

	m.CheckOptmization = true

	// FIXME: could do something fancy with another meta and then use that but w/e
	blockTokens := []token.Token{}

	for {
		m.Shift()

		current := m.CurrentToken

		switch current.Type {
		// TODO: this needs to change to PRI_OP
		case token.PriOp:
			fmt.Println("found a pri_op")
			blockTokens = append(blockTokens, current)

		case token.SecOp:
			fmt.Println("found a sec_op")
			// if m.NextToken.Type == current.Type {
			// 	m.Shift()
			// 	if t, ok := token.TokenMap[current.Value.String+m.CurrentToken.Value.String]; ok {
			// 		blockTokens = append(blockTokens, t)
			// 	} else {
			// 		fmt.Println("wtf happened here: ", current.Value.String+m.CurrentToken.Value.String)
			// 		os.Exit(9)
			// 	}
			// } else {
			// 	blockTokens = append(blockTokens, current)
			// }
			blockTokens = append(blockTokens, current)

		case token.Array:
			fmt.Println("found an array")
			blockTokens = append(blockTokens, current)

		case token.Keyword:
			fmt.Println("we are here at the keyword thing")
			blockTokens = append(blockTokens, current)
			// switch current.Value.Type {
			// case token.SQL:
			// 	fmt.Println("found a sql keyword")
			// }
			// os.Exit(9)

		case token.GThan:
			fmt.Println("found a greater than")
			blockTokens = append(blockTokens, current)

		case token.LThan:
			fmt.Println("found a greater than")
			blockTokens = append(blockTokens, current)

		case token.At:
			fmt.Println("found an at")
			blockTokens = append(blockTokens, current)

		// TODO: put all of these at the bottom
		// Don't do anything with these for now except append them
		// FIXME: hack to fix the repitition
		case token.Block:
			// blockTokens = append(blockTokens, m.ParseBlock())
			blockTokens = append(blockTokens, current)
		case token.Init:
			fallthrough
		case token.Attribute:
			fallthrough
		case token.Function:
			blockTokens = append(blockTokens, current)
			// fmt.Println(token.Function)

		case token.Group:
			fmt.Println("\nGOTAGROUP")
			fmt.Println()

			functionTokens := []token.Token{current}

			peek := m.NextToken
			// TODO: FIXME: for now we are going to assume that two groups only appear in sequence for a function
			switch peek.Type {
			case token.Group:
				// blockTokens = append(blockTokens, m.ParseFunctionDef(current))
				m.Shift()
				functionTokens = append(functionTokens, m.CurrentToken)

				if m.NextToken.Type == token.Block {
					m.Shift()
					blockTokens = append(blockTokens, token.Token{
						ID:   4,
						Type: token.Function,
						Value: token.Value{
							Type: "def",
							True: append(functionTokens, m.CurrentToken),
						},
					})
				}

			case token.Block:
				m.Shift()

				// TODO: could make a change here to instead just put it as a group but w/e
				// if m.LastCollectedToken.Type == token.Keyword {

				// }

				blockTokens = append(blockTokens, token.Token{
					ID:   4,
					Type: token.Function,
					Value: token.Value{
						Type: "def",
						True: append(functionTokens, m.CurrentToken),
					},
				})

			default:
				fmt.Println("wtf peek following group", peek, m)
				os.Exit(8)
			}

		case token.Hash:
			blockTokens = append(blockTokens, m.ParseAttribute())

		case token.Separator:
			fallthrough

		case token.EOS:
			// TODO: this will need to check the last and next token type later to determine wtf to do
			blockTokens = append(blockTokens, m.CurrentToken)

		case token.Whitespace:
			continue

		case token.Type:
			blockTokens = append(blockTokens, m.CurrentToken)
			peek := m.NextToken
			switch peek.Type {
			case token.Array:
				blockTokens = append(blockTokens, peek)

			case token.Ident:
				m.Shift()
				m.ParseIdent(&blockTokens, m.CurrentToken)

			case token.Literal:
				blockTokens = append(blockTokens, m.CurrentToken)
				m.Shift()
				m.CurrentToken.Type = token.Ident
				blockTokens = append(blockTokens, m.CurrentToken)

			case token.LBracket:
				fmt.Println("found array", current)
				m.Shift()
				if m.NextToken.Type != token.RBracket {
					fmt.Println("syntax ERROR: missing ] after type declaration")
					os.Exit(8)
				}

				// FIXME: fix this and make the ok check
				arrayToken, ok := token.TokenMap[current.Value.String+peek.Value.String+m.NextToken.Value.String]
				m.Shift()
				fmt.Println(arrayToken, ok)
				// blockTokens = append(blockTokens, m.ParseArray())
				blockTokens[len(blockTokens)-1] = arrayToken

				fmt.Println()
				fmt.Println("blockTokens", blockTokens)
				fmt.Println()
				// m.Shift()
				// blockTokens = append(blockTokens, m.ParseArray())
				// m.Shift()
				// fmt.Println("m.Current shit", m.CurrentToken)

				// if m.CurrentToken.Type != token.Ident {
				// 	fmt.Println("syntax error: no ident after array type declaration")
				// 	os.Exit(8)
				// }
				// m.ParseIdent(&blockTokens, m.CurrentToken)

			default:
				fmt.Printf("meta %+v\n", m)
				fmt.Println("ERROR after type declaration: peek, current", peek, current)
				os.Exit(77)
			}

		case token.Assign:
			fmt.Println("ASSIGN", current)
			fmt.Printf("CURRENTVALUETYPE %+v\n", current)
			switch current.Value.Type {
			case "set":
				peek := m.NextToken
				fmt.Println("PEEK", peek)
				switch peek.Type {
				case token.Assign:
					fmt.Println("FOUND :=", current.Value.String+peek.Value.String)
					if t, ok := token.TokenMap[current.Value.String+peek.Value.String]; ok {
						blockTokens = append(blockTokens, t)
						m.Shift()
					}
				default:
					blockTokens = append(blockTokens, m.CurrentToken)
				}

			case "assign":
				fallthrough
			case "init":
				blockTokens = append(blockTokens, current)

			default:
				// blockTokens = append(blockTokens, current)
				// continue
				fmt.Println("ERROR, how did we get in here without an assign type token", current)
				os.Exit(9)
			}

		case token.Ident:
			peek := m.NextToken

			if peek.Type == token.LParen {
				blockTokens = append(blockTokens, m.ParseFunctionCall(m.CurrentToken))
			} else {
				m.ParseIdent(&blockTokens, m.CurrentToken)
			}

			// TODO: this case might need to move to the Syntactic part of the parser
		case token.Literal:
			// TODO: this may cause some problems
			// TODO: this is causing some problems
			// switch m.PeekLastCollectedToken().Type {
			// case "SET":
			// 	fallthrough

			// case token.Assign:
			// 	fallthrough

			// case token.Init:
			// 	blockTokens = append(blockTokens, m.CurrentToken)
			// }
			blockTokens = append(blockTokens, m.CurrentToken)

		case token.LParen:
			blockTokens = append(blockTokens, m.ParseGroup())

		case token.RParen:
			// FIXME: why

		case token.LBracket:
			blockTokens = append(blockTokens, m.ParseArray())

		case token.LBrace:
			blockTokens = append(blockTokens, m.ParseBlock())

		case token.RBrace:
			return token.Token{
				ID:   0,
				Type: token.Block,
				// Expected: TODO: do the same thing that we did on the array but use the meta tokens
				Value: token.Value{
					Type: token.Block,
					True: blockTokens,
					// String: TODO: do the same thing that we did on array
				},
			}

		case token.DQuote:
			blockTokens = append(blockTokens, m.ParseString())

		case "":
			fmt.Println("got nothing")

		default:
			fmt.Println("IDK WTF TO DO with this token", m.CurrentToken)
			os.Exit(6)
		}
		fmt.Println(current, m.NextToken)

		if m.NextToken == (token.Token{}) {
			fmt.Println()
			fmt.Println("nextToken block", blockTokens)
			fmt.Println()
			// fmt.Println("blockTokens", blockTokens)
			return token.Token{
				ID:   0,
				Type: token.Block,
				// Expected: TODO: do the same thing that we did on the array but use the meta tokens
				Value: token.Value{
					Type: token.Block,
					True: blockTokens,
					// String: TODO: do the same thing that we did on array
				},
			}
		}
	}
}

// Symtactic begins the parsing process for a passes set of tokens
func Symtactic(tokens []token.Token) ([]token.Token, error) {
	// Auto inject the brackets to ensure that they are there
	meta := Meta{
		IgnoreWS:         true,
		Tokens:           append(append([]token.Token{token.TokenMap["{"]}, tokens...), token.TokenMap["}"]),
		Length:           len(tokens) + 2,
		CheckOptmization: true,
	}
	meta.Shift()

	// Here we are continuously applying semantic pressure to squash the tokens and furthur
	// simplify the tokens generated
	for meta.CheckOptmization {
		fmt.Println("Optimizing", meta.OptimizationAttempts)
		meta.CollectTokens(meta.ParseBlock().Value.True.([]token.Token))
		fmt.Println("endTokens", meta.EndTokens)

		fmt.Println(meta.EndTokens)
		metaTokens := meta.EndTokens[0].Value.True.([]token.Token)
		metaTokens = append(append([]token.Token{token.TokenMap["{"]}, metaTokens...), token.TokenMap["}"])
		fmt.Println("metaTokens", len(metaTokens), len(meta.EndTokens))

		// endTokens = meta.EndTokens

		// TODO: FIXME: w/e this works for now
		// Fix this from pulling off only the top one
		// Only apply SemanticPressure once for now until we figure out the recursion more
		if meta.OptimizationAttempts > 0 {
			break
		}

		// fmt.Println("meta.CheckOptimization", meta.CheckOptmization)

		// if !meta.CheckOptmization {
		// 	break
		// }

		// if len(meta.EndTokens) < len(meta.Tokens) {
		// 	break
		// }

		meta = Meta{
			// FIXME: do we need to fix this hack?
			// Tokens: ,
			Tokens:               metaTokens,
			Length:               len(metaTokens),
			CheckOptmization:     meta.CheckOptmization,
			OptimizationAttempts: meta.OptimizationAttempts + 1,
		}
	}

	return meta.EndTokens, nil
}
