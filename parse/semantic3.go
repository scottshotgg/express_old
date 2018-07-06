package parse

// import (
// 	"fmt"
// 	"os"
// 	"strconv"

// 	"github.com/pkg/errors"
// 	"github.com/scottshotgg/Express/token"
// )

// // FIXME: move this to its own file
// // LessThanOperands ...
// func (m *Meta) LessThanOperands(left, right token.Value) (token.Value, error) {
// 	// FIXME: this only works for ints right now
// 	return token.Value{
// 		True: left.True.(int) * right.True.(int),
// 		// String: strconv.Itoa(valueToken.True.(int)),
// 	}, nil
// }

// // EvaluateBinaryOperation ...
// // TODO: add in * and / and <
// func (m *Meta) EvaluateBinaryOperation(left, right, op token.Value) (opToken token.Value, err error) {
// 	fmt.Println("EvaluateBinaryOperation")

// 	switch op.Type {
// 	case "add":
// 		opToken, err = m.AddOperands(left, right)
// 		if err != nil {
// 			err = errors.New("Error adding operands")
// 		}

// 	case "sub":
// 		opToken, err = m.SubOperands(left, right)
// 		if err != nil {
// 			err = errors.New("Error subtracting operands")
// 		}

// 	case "mult":
// 		opToken, err = m.MultOperands(left, right)
// 		if err != nil {
// 			err = errors.New("Error multiplying operands")
// 		}

// 	case "div":
// 		opToken, err = m.DivOperands(left, right)
// 		if err != nil {
// 			err = errors.New("Error dividing operands")
// 		}

// 	case "lthan":
// 		opToken, err = m.LessThanOperands(left, right)
// 		if err != nil {
// 			err = errors.New("Error evaluating boolean expression")
// 		}

// 	default:
// 		err = errors.Errorf("Undefined operator; left: %+v right: %+v op: %+v", left, right, op)
// 		fmt.Println(err.Error())
// 	}

// 	// opToken.Name = op.Type + "Op"
// 	// opToken.Type = "OP"
// 	// opToken.OpMap = opMap
// 	// opToken.True = opMap["eval"].(token.Value)
// 	// opToken.String = left.String + op.String + right.String

// 	opToken.OpMap = map[string]interface{}{
// 		"left":  left,
// 		"right": right,
// 		"op":    op,
// 		// "string": left.String + op.String + right.String,
// 	}
// 	if opToken.Type == token.IntType {
// 		opToken.String = strconv.Itoa(opToken.True.(int))
// 	}
// 	return
// }

// // EvaluateUnaryOperation ...
// // TODO: implement this stuff
// func (m *Meta) EvaluateUnaryOperation(left, op token.Value) { // (token.Value, error) {
// }

// // GetFactor ...
// func (m *Meta) GetFactor() (token.Value, error) {
// 	fmt.Println("GetFactor")
// 	next := m.NextToken
// 	fmt.Printf("next %+v\n", next)

// 	var value token.Value
// 	var err error

// 	switch m.NextToken.Type {
// 	case token.Literal:
// 		m.Shift()
// 		value = m.CurrentToken.Value
// 		// FIXME: holy fuck haxorz
// 		if value.Type == token.IntType {
// 			value.String = strconv.Itoa(value.True.(int))
// 		}
// 		fmt.Println("hey its me the value", value)

// 	case token.Ident:
// 		m.Shift()
// 		var ok bool
// 		if value, ok = m.DeclarationMap[m.CurrentToken.Value.String]; !ok {
// 			if m.LastMeta != nil {
// 				fmt.Println(m.DeclarationMap)
// 				if value, ok = (*m.LastMeta).DeclarationMap[m.CurrentToken.Value.String]; !ok {
// 					// FIXME: holy fuck haxorz
// 					if value.Type == token.IntType {
// 						fmt.Printf("fuckthisshit2 %+v\n", m.CurrentToken)
// 						value.String = next.Value.String
// 					}
// 					fmt.Println((*m.LastMeta).DeclarationMap)
// 					return token.Value{}, errors.New("Undefined variable reference")
// 				}
// 			}
// 			// FIXME: holy fuck haxorz
// 			if value.Type == token.IntType {
// 				fmt.Printf("fuckthisshit %+v\n", m.CurrentToken)
// 				value.String = next.Value.String
// 			}
// 		}

// 	case token.Group:
// 		meta := Meta{
// 			AppendDeclarations: true,
// 			IgnoreWS:           true,
// 			Tokens:             m.NextToken.Value.True.([]token.Token),
// 			Length:             len(m.NextToken.Value.True.([]token.Token)),
// 			CheckOptmization:   true,
// 			LastMeta:           m,
// 			DeclarationMap:     map[string]token.Value{},
// 		}
// 		meta.Shift()
// 		// Might have to change this to GetExpression
// 		value, err = meta.GetExpression()
// 		if err != nil {
// 			return token.Value{}, err
// 		}
// 		// FIXME: holy fuck haxorz
// 		if value.Type == token.IntType {
// 			value.String = strconv.Itoa(value.True.(int))
// 		}
// 		m.Shift()
// 		// os.Exit(9)

// 	// case "":
// 	// 	fmt.Println("we at the end?")
// 	// 	os.Exit(8)

// 	default:
// 		fmt.Println("next2", m.NextToken)
// 		return token.Value{}, errors.Errorf("default %+v", m.NextToken)
// 	}
// 	fmt.Println("value thing again", value)

// 	switch m.NextToken.Type {
// 	case token.PriOp:
// 		m.Shift()
// 		op := m.CurrentToken
// 		value2, verr := m.GetFactor()
// 		if verr != nil {
// 			return token.Value{}, verr
// 		}
// 		fmt.Println("value2thing", value2)

// 		value, err = m.EvaluateBinaryOperation(value, value2, op.Value)
// 		if err != nil {
// 			return token.Value{}, err
// 		}
// 		// FIXME: holy fuck haxorz
// 		if value.Type == token.IntType {
// 			value.String = ""
// 		}

// 	case token.Increment:
// 		value, err = m.AddOperands(value, token.Value{
// 			Type: token.IntType,
// 			True: 1,
// 		})
// 		if err != nil {
// 			return token.Value{}, err
// 		}
// 	}

// 	// FIXME: holy fuck haxorz
// 	if value.Type == token.IntType {
// 		value.String = next.Value.String
// 	}
// 	fmt.Println("returning")
// 	return value, nil
// }

// // GetTerm ...
// func (m *Meta) GetTerm() (token.Value, error) {
// 	fmt.Println("GetTerm")

// 	totalTerm, err := m.GetFactor()
// 	if err != nil {
// 		return token.Value{}, err
// 	}

// 	for {
// 		switch m.NextToken.Type {
// 		case token.SecOp:
// 			m.Shift()
// 			fmt.Println("woah i got a secop")
// 			op := m.CurrentToken
// 			factor2, ferr := m.GetFactor()
// 			if ferr != nil {
// 				return token.Value{}, ferr
// 			}
// 			fmt.Println("factor2", factor2)

// 			totalTerm, err = m.EvaluateBinaryOperation(totalTerm, factor2, op.Value)
// 			if err != nil {
// 				return token.Value{}, err
// 			}
// 			// FIXME: holy fuck haxorz
// 			if totalTerm.Type == token.IntType {
// 				totalTerm.String = strconv.Itoa(totalTerm.True.(int))
// 			}

// 		// TODO: need to fix this....
// 		case token.LThan:
// 			// ident := m.LastToken
// 			nextTokenOpString := m.NextToken.Value.String
// 			m.Shift()
// 			// op := m.CurrentToken
// 			factor2, ferr := m.GetTerm()
// 			if ferr != nil {
// 				return token.Value{}, ferr
// 			}
// 			fmt.Println("lthan totalTerm", totalTerm)
// 			fmt.Println("lthan factor2", factor2)
// 			// totalTerm, err = m.EvaluateBinaryOperation(totalTerm, factor2, op.Value)
// 			// if err != nil {
// 			// 	return token.Value{}, err
// 			// }
// 			// FIXME: holy fuck haxorz
// 			// if totalTerm.Type == token.IntType {
// 			factor2.String = totalTerm.String + nextTokenOpString + factor2.String
// 			// }
// 			fmt.Println("totalBoolTerm", totalTerm)
// 			return factor2, nil

// 		case token.Separator:
// 			m.Shift()
// 			// FIXME: holy fuck haxorz
// 			if totalTerm.Type == token.IntType {
// 				totalTerm.String = strconv.Itoa(totalTerm.True.(int))
// 			}
// 			return totalTerm, nil

// 		default:
// 			// FIXME: holy fuck haxorz
// 			if totalTerm.Type == token.IntType {
// 				totalTerm.String = strconv.Itoa(totalTerm.True.(int))
// 			}
// 			fmt.Println("i am here", m.NextToken)
// 			return totalTerm, nil
// 		}
// 	}
// }

// // GetExpression ...
// func (m *Meta) GetExpression() (token.Value, error) {
// 	fmt.Println("GetExpression")
// 	fmt.Println("m.NextToken", m.NextToken)

// 	switch m.NextToken.Type {
// 	// Assignment Expression
// 	case token.Assign:
// 		m.DeclaredName = m.CurrentToken.Value.String
// 		m.DeclaredAccessType = m.CurrentToken.Value.Type
// 		switch m.NextToken.Value.Type {
// 		case "init":
// 			if m.DeclaredType != "" {
// 				return token.Value{}, errors.New("Type with init is not valid")
// 			}
// 			m.DeclaredType = token.SetType
// 			fallthrough

// 		case "assign":
// 			m.Shift()
// 			expr, err := m.GetExpression()
// 			if err != nil {
// 				return token.Value{}, err
// 			}
// 			fmt.Println("expr", expr)
// 			fmt.Println("m.DeclaredType", m.DeclaredType)
// 			if m.DeclaredType == token.SetType {
// 				m.DeclaredType = expr.Type
// 			} else if m.DeclaredType != expr.Type {
// 				// TODO: implicit type casting here
// 				return token.Value{}, errors.New("No implicit type casting as of now")
// 			}
// 			m.DeclarationMap[m.DeclaredName] = token.Value{
// 				Name:       m.DeclaredName,
// 				Type:       m.DeclaredType,
// 				True:       expr.True,
// 				AccessType: m.DeclaredAccessType,
// 			}
// 			fmt.Println(m.DeclarationMap)
// 			return m.DeclarationMap[m.DeclaredName], nil
// 		}

// 	// case token.LThan:
// 	// 	fmt.Println("wtf")
// 	// 	fmt.Println("current", m.CurrentToken)
// 	// 	fmt.Println("next", m.NextToken)
// 	// 	m.Shift()
// 	// 	term, err := m.GetTerm()
// 	// 	if err != nil {
// 	// 		return token.Value{}, err
// 	// 	}
// 	// 	return term, nil

// 	case token.Increment:
// 		fmt.Println("woah increment brah")
// 		// term, err := m.AddOperands()

// 	default:
// 		return m.GetTerm()
// 	}

// 	return token.Value{}, errors.Errorf("default %+v", m.NextToken)
// }

// // GetKeyword ...
// func (m *Meta) GetKeyword() (token.Value, error) {
// 	fmt.Println("GetKeyword")

// 	switch m.NextToken.Value.String {
// 	// TODO: this needs to be reworked
// 	case token.For:
// 		fmt.Println("formap", m.DeclarationMap)
// 		fmt.Println("found a for loop22")
// 		temp := *m
// 		meta := &temp
// 		meta.LastMeta = m
// 		meta.DeclarationMap = map[string]token.Value{}
// 		meta.Shift()

// 		value, err := meta.GetStatement()
// 		if err != nil {
// 			return token.Value{}, err
// 		}
// 		fmt.Println("value11", value)
// 		fmt.Println("last", meta.LastToken)
// 		fmt.Println("current", meta.CurrentToken)
// 		fmt.Println("next", meta.NextToken)

// 		value2, err := meta.GetExpression()
// 		if err != nil {
// 			return token.Value{}, err
// 		}
// 		fmt.Println("value22", value2)
// 		fmt.Println("last", meta.LastToken)
// 		fmt.Println("current", meta.CurrentToken)
// 		fmt.Println("next", meta.NextToken)
// 		// m.Shift()

// 		value3, err := meta.GetExpression()
// 		if err != nil {
// 			return token.Value{}, err
// 		}
// 		fmt.Println("value3", value3)

// 		stepAmount, err := m.SubOperands(value3, value)
// 		if err != nil {
// 			return token.Value{}, err
// 		}
// 		fmt.Println("step", stepAmount)

// 		// Need to open up the block
// 		// we might try doing something
// 		// where the new meta stuff is in the function
// 		// block, err := meta.CheckBlock()
// 		// if err != nil {
// 		// 	return token.Value{}, err
// 		// }
// 		// fmt.Println("block", block)
// 		// os.Exit(9)
// 		meta.Shift()
// 		fmt.Println(meta.NextToken)
// 		block, err := Semantic([]token.Token{
// 			meta.NextToken,
// 		})
// 		if err != nil {
// 			return token.Value{}, err
// 		}
// 		fmt.Println("body", block)

// 		// Swap the scopes back when the for loop is out of execution
// 		meta.DeclarationMap = m.DeclarationMap
// 		meta.InheritedMap = m.InheritedMap
// 		*m = *meta

// 		// fmt.Println("value2againboi", value2)
// 		mapThing := value2.OpMap.(map[string]interface{})
// 		fmt.Println("mapThing", mapThing)

// 		return token.Value{
// 			Type: token.For,
// 			True: map[string]token.Value{
// 				"start": value,
// 				"end":   value2,
// 				"step":  stepAmount,
// 				"body":  block[0],
// 				"check": token.Value{
// 					String: value2.String,
// 				},
// 			},
// 		}, nil

// 	case "if":
// 		temp := *m
// 		meta := &temp
// 		meta.LastMeta = m
// 		meta.DeclarationMap = map[string]token.Value{}
// 		meta.Shift()
// 		fmt.Println("declaredMap", m.DeclarationMap)
// 		fmt.Println("inheritedMap", m.InheritedMap)

// 		value2, err := meta.GetExpression()
// 		if err != nil {
// 			return token.Value{}, err
// 		}
// 		fmt.Println("value22", value2)
// 		fmt.Println("last", meta.LastToken)
// 		fmt.Println("current", meta.CurrentToken)
// 		fmt.Printf("next %+v\n", meta.NextToken)

// 		fmt.Println(meta.NextToken.Value.True)
// 		block, err := Semantic([]token.Token{
// 			meta.NextToken,
// 		})
// 		if err != nil {
// 			return token.Value{}, err
// 		}
// 		fmt.Println("body", block)

// 		// Swap the scopes back when the for loop is out of execution
// 		meta.DeclarationMap = m.DeclarationMap
// 		meta.InheritedMap = m.InheritedMap
// 		*m = *meta

// 		return token.Value{
// 			Type:   token.If,
// 			String: value2.String,
// 			// True: map[string]token.Value{
// 			// 	"check": token.Value{
// 			// 		String: value2.String,
// 			// 	},
// 			// },
// 			// True: // body would go here
// 		}, nil

// 		os.Exit(9)
// 		// TODO: there would be some composition of blocks here and shit

// 	default:
// 		fmt.Println("keyword not recognized", m.NextToken)
// 		os.Exit(9)
// 	}

// 	return token.Value{}, nil
// }

// // GetStatement ...
// func (m *Meta) GetStatement() (token.Value, error) {
// 	fmt.Println("GetStatement")

// 	switch m.NextToken.Type {
// 	case token.Type:
// 		m.DeclaredType = m.NextToken.Value.Type
// 		m.Shift()
// 		// TODO: could either recurse here, or fallthrough
// 		if m.NextToken.Type != token.Ident {
// 			break
// 		}
// 		fallthrough

// 	// TODO: will have to consider declarations too
// 	case token.Ident:
// 		fmt.Println("ident", m.NextToken)
// 		fmt.Println("declaredMap", m.DeclarationMap)
// 		fmt.Println("inheritedMap", m.InheritedMap)
// 		if m.DeclaredType == "" {
// 			if declaredType, ok := m.DeclarationMap[m.NextToken.Value.String]; ok {
// 				m.DeclaredType = declaredType.Type
// 			} else if declaredType, ok := m.InheritedMap[m.NextToken.Value.String]; ok {
// 				m.DeclaredType = declaredType.Type
// 			}
// 		}
// 		fmt.Println("ASSIGNMENT DECLARED VALUE", m.DeclaredValue)
// 		m.Shift()
// 		return m.GetExpression()

// 	case token.Keyword:
// 		keyword, err := m.GetKeyword()
// 		if err != nil {
// 			return token.Value{}, err
// 		}
// 		m.Shift()
// 		return keyword, nil

// 	case token.Separator:
// 		fmt.Println("should we have gotten this here?")
// 		os.Exit(9)

// 	case token.SecOp:
// 		switch m.CurrentToken.Value.Type {
// 		case "sub":
// 			// TODO: need to do something here for negative expression

// 		default:
// 			return token.Value{}, errors.New("Unrecognized position for operator")
// 		}

// 	default:
// 		// TODO: this causes infinite loops when you cant parse
// 		fmt.Println("hey its me, the default", m.NextToken)
// 	}

// 	return token.Value{}, nil
// }

// // CheckBlock ...
// func (m *Meta) CheckBlock() (token.Value, error) {
// 	fmt.Println("CheckBlock")
// 	blockTokens := []token.Value{}

// 	for {
// 		stmt, err := m.GetStatement()
// 		if err != nil {
// 			fmt.Println("err", err)
// 			os.Exit(9)
// 		}
// 		blockTokens = append(blockTokens, stmt)

// 		m.DeclaredName = ""
// 		m.DeclaredType = ""
// 		m.DeclaredAccessType = ""
// 		m.DeclaredActingType = ""
// 		m.DeclaredValue = token.Value{}
// 		fmt.Println("m.DeclarationMap", m.DeclarationMap)

// 		if m.NextToken == (token.Token{}) {
// 			return token.Value{
// 				Type: token.Block,
// 				True: blockTokens,
// 			}, nil
// 		}
// 	}
// }

// // Semantic ...
// func Semantic(tokens []token.Token) ([]token.Value, error) {
// 	fmt.Println("Semantic")

// 	meta := Meta{
// 		AppendDeclarations: true,
// 		IgnoreWS:           true,
// 		Tokens:             tokens[0].Value.True.([]token.Token),
// 		Length:             len(tokens[0].Value.True.([]token.Token)),
// 		CheckOptmization:   true,
// 		DeclarationMap:     map[string]token.Value{},
// 	}
// 	meta.Shift()

// 	block, err := meta.CheckBlock()
// 	if err != nil {
// 		// TODO:
// 		return []token.Value{}, err
// 	}
// 	fmt.Println("block", block)

// 	fmt.Println("declarationMap", meta.DeclarationMap)

// 	return []token.Value{block}, nil
// }

// // TODO: start here
// // TODO: use next token
// // TODO: start very simply with the definition in documentation/notes_about_shit
// // TODO: VERY SIMPLE requirements parsing vars with the return architecture of semantic2
