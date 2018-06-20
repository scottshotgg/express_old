package parse

// // TODO: FIXME: simplify architecture and requirements
// // TODO: FIXME: code using m.NextToken, not m.CurrentToken

// import (
// 	"fmt"
// 	"os"
// 	"strconv"

// 	"github.com/pkg/errors"
// 	"github.com/scottshotgg/Express/token"
// )

// // TODO: FIXME: add token IDs, make token Types consistent
// // TODO: handle token.Separator to dump the parse

// // EvaluateBinaryOperation ...
// // TODO: add in * and / and <
// func (m *Meta) EvaluateBinaryOperation(left, right, op token.Value) (token.Value, error) {
// 	fmt.Println("hi im EvaluateBinaryOperation")
// 	switch op.Type {
// 	case "add":
// 		addValue, err := m.AddOperands(left, right)
// 		if err != nil {
// 			return token.Value{}, errors.New("Error adding operands")
// 		}
// 		return addValue, nil

// 	case "sub":
// 		subValue, err := m.AddOperands(left, right)
// 		if err != nil {
// 			return token.Value{}, errors.New("Error subtracting operands")
// 		}
// 		return subValue, nil

// 	case "lthan":
// 		// lessValue, err := m.LessThanOperands(left, right)
// 		// if err != nil {
// 		// 	return token.Value{}, errors.New("Error evaluating boolean expression")
// 		// }
// 		// return lessValue, nil
// 		lt := left.True.(int) < right.True.(int)
// 		return token.Value{
// 			Type:       token.BoolType,
// 			True:       lt,
// 			String:     strconv.FormatBool(lt),
// 			AccessType: token.PrivateAccessType,
// 		}, nil

// 	default:
// 		err := errors.Errorf("Undefined operator; left: %+v right: %+v op: %+v", left, right, op)
// 		fmt.Println(err.Error())
// 		return token.Value{}, err
// 	}
// }

// // EvaluateUnaryOperation ...
// // TODO: implement this stuff
// func (m *Meta) EvaluateUnaryOperation(left, op token.Value) { // (token.Value, error) {
// }

// // GetFactor ...
// func (m *Meta) GetFactor() (token.Value, error) {
// 	fmt.Println("hi im GetFactor")
// 	// A factor can be one of three things right now
// 	// - literal value
// 	// - ident value
// 	// - another expression; started by a left paren
// 	// TODO: add in idents
// 	// TODO: add in expression

// 	current := m.CurrentToken
// 	switch current.Type {
// 	case token.Literal:
// 		fmt.Println("literal", current)
// 		return current.Value, nil

// 	case token.Ident:
// 		fmt.Println("ident", current)
// 		if ident, ok := m.DeclarationMap[current.Value.String]; ok {
// 			return ident, nil
// 		}
// 		return token.Value{}, errors.Errorf("Undefined variable reference: %+v", current.Value)

// 	default:
// 		err := errors.New("Error getting factor")
// 		fmt.Println(err.Error(), current)
// 		return token.Value{}, err
// 	}
// }

// // GetTerm ...
// // TODO: recurse if there is another term operator
// // TODO: need to change this to PriOp
// func (m *Meta) GetTerm() (token.Value, error) {
// 	fmt.Println("hi im GetTerm")
// 	factor, err := m.GetFactor()
// 	if err != nil {
// 		fmt.Println("Error getting term", err)
// 		return token.Value{}, err
// 	}
// 	fmt.Println("factor", factor)

// 	for {
// 		m.Shift()
// 		current := m.CurrentToken
// 		switch current.Type {
// 		case token.SecOp:
// 			fmt.Println("secop", current)
// 			// we need to look for another factor
// 			m.Shift()
// 			factor2, err := m.GetFactor()
// 			if err != nil {
// 				return token.Value{}, errors.New("Error finding factor after operator")
// 			}
// 			fmt.Println("factor2", factor2)

// 			value, err := m.EvaluateBinaryOperation(factor, factor2, current.Value)
// 			if err != nil {
// 				return token.Value{}, errors.New("Error evaluating binary operation")
// 			}
// 			fmt.Println("value", value)
// 			return value, nil

// 		case token.LThan:
// 			// we need to look for another factor
// 			m.Shift()
// 			factor2, err := m.GetFactor()
// 			if err != nil {
// 				return token.Value{}, errors.New("Error finding factor after operator")
// 			}
// 			fmt.Println("factor2", factor2)

// 			value, err := m.EvaluateBinaryOperation(factor, factor2, current.Value)
// 			if err != nil {
// 				return token.Value{}, errors.New("Error evaluating binary operation")
// 			}
// 			fmt.Println("value", value)
// 			return token.Value{
// 				Type: "bool_expr",
// 				True: map[string]token.Value{
// 					"left":  factor,
// 					"right": factor2,
// 					"op":    current.Value,
// 					"value": value,
// 				},
// 			}, nil

// 		// TODO: we could look at simplifying this to i = i + 1 before it even gets here
// 		case token.Increment:
// 			return m.AddOperands(factor, token.Value{
// 				Type: token.IntType,
// 				True: 1,
// 			})

// 		case token.Separator:
// 			m.Shift()
// 			return factor, nil

// 			// default:
// 			// 	return token.Value{}, errors.Errorf("Undefined operator: %+v", current)
// 		}
// 	}
// }

// // GetExpression ...
// func (m *Meta) GetExpression() (token.Value, error) {
// 	fmt.Println("hi im GetExpression")
// 	term, err := m.GetTerm()
// 	if err != nil {
// 		fmt.Println("Error getting expression", err)
// 		return token.Value{}, err
// 	}
// 	fmt.Println("term", term)

// 	// if m.NextToken.Type == token.LThan {
// 	// 	fmt.Println("woah we got another expression")
// 	// }

// 	return term, nil
// }

// // func (m *Meta) GetBooleanExpression() (token.Value, error) {
// // 	// Expression
// // 	// Boolean Operator
// // 	// Expression
// // 	expr1, err := m.GetExpression()
// // 	if err != nil {
// // 		fmt.Println("Error getting expression", err)
// // 		return token.Value{}, err
// // 	}
// // 	fmt.Println("expression1", expr1)

// // 	fmt.Println("current token", m.CurrentToken)

// // 	expr2, err := m.GetExpression()
// // 	if err != nil {
// // 		fmt.Println("Error getting expression", err)
// // 		return token.Value{}, err
// // 	}
// // 	fmt.Println("expression2", expr2)

// // 	return token.Value{}, nil
// // }

// // GetAssignmentStatement ...
// // TODO: should we move ident discovery to GetExpression ?
// // TODO: need to somehow do < and >
// // TODO: need to include SecOp
// // TODO: recurse GetExpression call if another op
// func (m *Meta) GetAssignmentStatement() (token.Value, error) {
// 	fmt.Println("hi im GetAssignmentStatement")

// 	m.Shift()

// 	for {
// 		current := m.CurrentToken

// 		switch current.Type {
// 		case token.Type:
// 			m.DeclaredType = current.Value.Type
// 			m.Shift()
// 			fallthrough

// 		case token.Ident:
// 			fmt.Println("ident", current)
// 			// FIXME: fix the name
// 			m.DeclaredName = current.Value.String
// 			// FIXME: fix the accessType
// 			m.DeclaredAccessType = current.Value.Type

// 			if m.NextToken.Type == token.Assign {
// 				m.Shift()
// 				fmt.Println("m.Shift m.NextToken", m.NextToken)
// 				switch m.CurrentToken.Value.Type {
// 				case "init":
// 					// Moving this down to the switch
// 					// if m.DeclaredType != "" {
// 					// 	err := errors.New("Error: found type accompanying init operator")
// 					// 	fmt.Println(err.Error())
// 					// 	return err
// 					// }
// 					m.DeclaredType = token.SetType

// 					// check the map for the !var
// 					// error if it is there
// 					fallthrough

// 				case "assign":
// 					// check the map for the var
// 					// error if NOT there
// 					m.Shift()
// 					expr, err := m.GetExpression()
// 					if err != nil {
// 						fmt.Println("Error getting expression", err)
// 						return token.Value{}, err
// 					}
// 					fmt.Println("expression", expr)
// 					m.DeclaredValue = expr
// 					if m.DeclaredType != m.DeclaredValue.Type {
// 						// at this point we need to do something else to ensure that the types are the same
// 						// if it is a
// 						// - var: can be anything so it doesn't matter
// 						// - set: interpreting the type from the value, so set the DeclaredType
// 						// - other: TODO: need to have some way of figuring this out....
// 						switch m.DeclaredType {
// 						case token.SetType:
// 							m.DeclaredType = m.DeclaredValue.Type

// 						case token.VarType:
// 							// TODO: need to do something with the var here

// 						case "":
// 							err := errors.New("Error: no type found")
// 							fmt.Println(err.Error())
// 							return token.Value{}, err

// 						default:
// 							// TODO: call some other function
// 						}
// 					}

// 					// I don't think we need to do this
// 					// at this point we have everything we need to insert the variable
// 					if m.AppendDeclarations {
// 						m.DeclarationMap[m.DeclaredName] = token.Value{
// 							Name:       m.DeclaredName,
// 							Type:       m.DeclaredType,
// 							True:       m.DeclaredValue.True,
// 							AccessType: m.DeclaredAccessType,
// 						}
// 					}
// 					fmt.Println("m.DeclarationMap", m.DeclarationMap)
// 					return token.Value{
// 						Name:       m.DeclaredName,
// 						Type:       m.DeclaredType,
// 						True:       m.DeclaredValue.True,
// 						AccessType: m.DeclaredAccessType,
// 					}, nil

// 				case "set":
// 					// need to define the standard usage of this
// 					// should probably be used for setting object vars or something idk

// 				default:
// 					fmt.Println("Not an assignment operator", m.CurrentToken)
// 				}
// 			}

// 			return token.Value{}, errors.New("Invalid assignment statement")

// 		case "":
// 			return token.Value{}, errors.New("EOF")

// 		default:
// 			fmt.Println("hey its me, the default", current)
// 		}

// 		m.Shift()
// 	}
// }

// // CheckFor ...
// func (m *Meta) CheckFor() (token.Value, error) {
// 	fmt.Println("hi im CheckFor")

// 	// At this point we are parsing a for loop and
// 	// creating a for loop token
// 	// For the basic for loop we should expect:
// 	//		an integer instantiation
// 	//		a boolean expression
// 	//		something modifying the integer
// 	//		for now just ++

// 	// Get the starting statement of the for loop
// 	start, err := m.GetAssignmentStatement()
// 	if err != nil {
// 		return token.Value{}, err
// 	}
// 	fmt.Println("start", start)

// 	// TODO: need to add boolean operations to GetExpression
// 	// FIXME: separator is blocking this, but ideally they would be taken
// 	// care of in the switch above to force the end of the sentence
// 	// TODO: FIXME: GetExpression should return a token.Value with multiple token.Values in it essentially making up an expression.Value
// 	boolExpr, err := m.GetExpression()
// 	if err != nil {
// 		// TODO:
// 	}
// 	fmt.Println("boolExpr", boolExpr)

// 	// FIXME: shift twice for now as we are expecting there to be <ident><relational_op> for the standard
// 	// m.Shift()
// 	// m.Shift()
// 	step, err := m.GetExpression()
// 	if err != nil {
// 		// TODO:
// 	}
// 	fmt.Println("step", step)
// 	step, err = m.SubOperands(step, start)
// 	if err != nil {
// 		// TODO:
// 	}

// 	fmt.Println("boolExpr", boolExpr)

// 	// FIXME: remove with the above
// 	m.Shift()
// 	m.Shift()

// 	body := token.Value{}
// 	// m.CheckBlock()

// 	// TODO:
// 	return token.Value{
// 		Type: token.For,
// 		True: map[string]token.Value{
// 			"start": start,
// 			"end":   boolExpr.True.(map[string]token.Value)["right"],
// 			"step":  step,
// 			"body":  body,
// 		},
// 	}, nil
// }

// // CheckKeyword ...
// func (m *Meta) CheckKeyword() (token.Value, error) {
// 	fmt.Println("hi im CheckKeyword")

// 	for {
// 		current := m.CurrentToken
// 		switch current.Value.String {
// 		case token.For:
// 			// make a new meta to give the for loop a new context for parsing; a new scope
// 			m.InheritedMap = m.DeclarationMap
// 			m.DeclarationMap = map[string]token.Value{}

// 			forLoop, err := m.CheckFor()
// 			if err != nil {
// 				fmt.Println("error", err)
// 				return token.Value{}, err
// 			}
// 			fmt.Printf("forLoop %+v\n", forLoop)
// 			return forLoop, nil

// 		case "":
// 			// might have to fix this
// 			return token.Value{}, nil

// 		default:
// 			fmt.Println("hey its me, the default", current)
// 		}

// 		m.Shift()
// 	}
// }

// // GetStatement ...
// func (m *Meta) GetStatement() (token.Value, error) {
// 	fmt.Println("hi im GetStatement")
// 	fmt.Println("m.NextToken", m.NextToken)
// 	// for {
// 	switch m.NextToken.Type {
// 	case token.Keyword:
// 		keyword, err := m.CheckKeyword()
// 		if err != nil {
// 			// TODO:
// 			return token.Value{}, err
// 		}
// 		return keyword, nil

// 	case token.Ident:
// 		as, err := m.GetAssignmentStatement()
// 		if err != nil {
// 			return token.Value{}, err
// 		}
// 		return as, nil

// 	default:
// 		fmt.Println("hey its me, the default", m.NextToken)
// 		os.Exit(9)
// 	}

// 	return token.Value{}, nil
// }

// // CheckBlock ...
// func (m *Meta) CheckBlock() (token.Value, error) {
// 	fmt.Println("hi im CheckBlock")
// 	blockTokens := []token.Value{}
// 	for {
// 		statement, err := m.GetStatement()
// 		if err != nil {
// 			// TODO:
// 		}
// 		blockTokens = append(blockTokens, statement)

// 		if m.NextToken == (token.Token{}) {
// 			return token.Value{
// 				Type: token.Block,
// 				True: blockTokens,
// 			}, nil
// 		}
// 	}
// }

// // Semantic ...
// // TODO: make this take a map[string]token.Value that has a variable map that it can inherit
// func Semantic(tokens []token.Token) (token.Value, error) {
// 	fmt.Println("hi im Semantic")

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
// 		return token.Value{}, err
// 	}

// 	return block, nil
// }
