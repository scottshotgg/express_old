package parse

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/scottshotgg/Express/token"
)

// GetFactor ...
func (m *Meta) GetFactor() (token.Value, error) {
	// A factor can be one of three things right now
	// - literal value
	// - ident value
	// - another expression; started by a left paren

	current := m.CurrentToken
	switch current.Type {
	case token.Literal:
		fmt.Println("literal", current)
		return current.Value, nil

	default:
		err := errors.New("Error getting factor")
		fmt.Println(err.Error(), current)
		return token.Value{}, err
	}

	return token.Value{}, nil
}

// GetBinaryOperationValue ...
func (m *Meta) GetBinaryOperationValue(left, right, op token.Value) (token.Value, error) {
	switch op.Type {
	case "add":
		addValue, err := m.AddOperands(left, right)
		if err != nil {
			return token.Value{}, errors.New("Error adding operands")
		}
		return addValue, nil

	case "sub":
		subValue, err := m.AddOperands(left, right)
		if err != nil {
			return token.Value{}, errors.New("Error subtracting operands")
		}
		return subValue, nil

		// case "add":
		// 	addValue, err := m.AddOperands(factor, factor2)
		// 	if err != nil {
		// 		return token.Value{}, errors.New("Error adding operands", factor, factor2)
		// 	}
		// 	return addValue, nil

		// case "add":
		// 	addValue, err := m.AddOperands(factor, factor2)
		// 	if err != nil {
		// 		return token.Value{}, errors.New("Error adding operands", factor, factor2)
		// 	}
		// 	return addValue, nil

	}

	return token.Value{}, nil
}

// GetTerm ...
func (m *Meta) GetTerm() (token.Value, error) {
	factor, err := m.GetFactor()
	if err != nil {
		fmt.Println("Error getting term", err)
		return token.Value{}, err
	}
	fmt.Println("factor", factor)

	for {
		m.Shift()
		switch m.CurrentToken.Type {
		case token.SecOp:
			secop := m.CurrentToken.Value
			fmt.Println("secop", secop)
			// we need to look for another factor
			m.Shift()
			factor2, err := m.GetFactor()
			if err != nil {
				return token.Value{}, errors.New("Error finding factor after operator")
			}
			fmt.Println("factor2", factor2)

			value, err := m.GetBinaryOperationValue(factor, factor2, secop)
			if err != nil {
				return token.Value{}, errors.New("Error adding operands")
			}
			fmt.Println("value", value)
			return value, nil

		case token.Separator:
			m.Shift()
			return factor, nil

		default:
			return token.Value{}, errors.New("wtf")
		}
	}

	return factor, nil
}

// GetExpression ...
func (m *Meta) GetExpression() (token.Value, error) {
	term, err := m.GetTerm()
	if err != nil {
		fmt.Println("Error getting expression", err)
		return token.Value{}, err
	}
	fmt.Println("term", term)

	return term, nil
}

// GetAssignmentStatement ...
func (m *Meta) GetAssignmentStatement() (token.Value, error) {
	fmt.Println("hi im getassignmentstatement")

	m.Shift()

	for {
		current := m.CurrentToken

		switch current.Type {
		case token.Type:
			m.DeclaredType = current.Value.Type
			m.Shift()
			fallthrough

		case token.Ident:
			fmt.Println("ident", current)
			// FIXME: fix the name
			m.DeclaredName = current.Value.String
			// FIXME: fix the accessType
			m.DeclaredAccessType = current.Value.Type

			if m.NextToken.Type == token.Assign {
				m.Shift()
				switch m.CurrentToken.Value.Type {
				case "init":
					// Moving this down to the switch
					// if m.DeclaredType != "" {
					// 	err := errors.New("Error: found type accompanying init operator")
					// 	fmt.Println(err.Error())
					// 	return err
					// }
					m.DeclaredType = token.SetType

					// check the map for the !var
					// error if it is there
					fallthrough

				case "assign":
					// check the map for the var
					// error if NOT there
					m.Shift()
					expr, err := m.GetExpression()
					if err != nil {
						fmt.Println("Error getting expression", err)
						return token.Value{}, err
					}
					fmt.Println("expression", expr)
					m.DeclaredValue = expr
					if m.DeclaredType != m.DeclaredValue.Type {
						// at this point we need to do something else to ensure that the types are the same
						// if it is a
						// - var: can be anything so it doesn't matter
						// - set: interpreting the type from the value, so set the DeclaredType
						// - other: TODO: need to have some way of figuring this out....
						switch m.DeclaredType {
						case token.SetType:
							m.DeclaredType = m.DeclaredValue.Type

						case token.VarType:
							// TODO: need to do something with the var here

						case "":
							err := errors.New("Error: no type found")
							fmt.Println(err.Error())
							return token.Value{}, err

						default:
							// TODO: call some other function
						}
					}

					// I don't think we need to do this
					// at this point we have everything we need to insert the variable
					// if m.AppendDeclarations {
					// 	m.DeclarationMap[m.DeclaredName] = token.Value{
					// 		Name:       m.DeclaredName,
					// 		Type:       m.DeclaredType,
					// 		True:       m.DeclaredValue.True,
					// 		AccessType: m.DeclaredAccessType,
					// 	}
					// }
					// fmt.Println("m.DeclarationMap", m.DeclarationMap)
					return token.Value{
						Name:       m.DeclaredName,
						Type:       m.DeclaredType,
						True:       m.DeclaredValue.True,
						AccessType: m.DeclaredAccessType,
					}, nil

				case "set":
					// need to define the standard usage of this
					// should probably be used for setting object vars or something idk

				default:
					fmt.Println("Not an assignment operator", m.CurrentToken)
				}
			}

			return token.Value{}, errors.New("Invalid assignment statement")

		case "":
			return token.Value{}, errors.New("EOF")

		default:
			fmt.Println("hey its me, the default", current)
		}

		m.Shift()
	}

	return token.Value{}, nil
}

// GetStatement gets the next statement in the sequence
func (m *Meta) GetStatement() error {
	fmt.Println("hi im getstatement")

	// switch m.CurrentToken.Type {
	// case token.Block:
	// 	fmt.Println("m.CurrentToken GetStatement()", m.CurrentToken)
	// 	err := m.CheckBlock()
	// 	if err == nil {
	// 		return nil
	// 	}
	// }
	// os.Exit(9)
	// err := m.GetAnonymousScope()
	// if err == nil {
	// 	return nil
	// }

	as, err := m.GetAssignmentStatement()
	if err == nil {
		return nil
	}
	fmt.Println("as", as)

	// FIXME: ideally we should do a switch on the error but w/e for now
	// TODO: we need some way to backtrack to before this operation ...
	// if err != nil {
	// TODO: woah we could do partial compilation
	return errors.New("error getting statement")
	// }
}

// CheckFor ...
func (m *Meta) CheckFor() (token.Value, error) {
	fmt.Println("hi im checkfor")

	// At this point we are parsing a for loop and
	// creating a for loop token
	// For the basic for loop we should expect:
	//		an integer instantiation
	//		a boolean expression
	//		something modifying the integer
	//		for now just ++

	// Get the starting statement of the for loop
	start, err := m.GetAssignmentStatement()
	if err != nil {
		return token.Value{}, err
	}
	fmt.Println("start", start)

	// TODO: need to add boolean operations to GetExpression
	// FIXME: separator is blocking this, but ideally they would be taken
	// care of in the switch above to force the end of the sentence
	end, err := m.GetExpression()
	if err != nil {
		fmt.Println("error", err)
		return token.Value{}, err
	}
	fmt.Println("end", end)

	step, err := m.GetExpression()
	if err != nil {
		fmt.Println("error", err)
		return token.Value{}, err
	}
	fmt.Println("step", step)

	// TODO:
	return token.Value{
		Type: token.For,
		True: map[string]token.Value{
			"start": start,
			"end":   end,
			"step":  step,
		},
	}, nil
}

// CheckKeyword ...
func (m *Meta) CheckKeyword() {
	fmt.Println("hi im checkkeyword")

	for {
		current := m.CurrentToken
		switch current.Value.String {

		case token.For:
			m.AppendDeclarations = false
			forLoop, err := m.CheckFor()
			if err != nil {
				fmt.Println("error", err)
				return
			}
			m.AppendDeclarations = true
			fmt.Println("forLoop", forLoop)

		case "":
			return

		default:
			fmt.Println("hey its me, the default", current)
		}

		m.Shift()
	}
}

// CheckBlock ...
func (m *Meta) CheckBlock() {
	fmt.Println("hi im checkblock")

	for {
		current := m.CurrentToken
		switch current.Type {

		case token.Keyword:
			m.CheckKeyword()

		case "":
			return

		default:
			fmt.Println("hey its me, the default", current)
		}

		m.Shift()
	}
}

// Semantic ...
func Semantic(tokens []token.Token) ([]token.Token, error) {
	fmt.Println("hi im semantic")

	meta := Meta{
		AppendDeclarations: true,
		IgnoreWS:           true,
		Tokens:             tokens[0].Value.True.([]token.Token),
		Length:             len(tokens[0].Value.True.([]token.Token)),
		CheckOptmization:   true,
		DeclarationMap:     map[string]token.Value{},
	}
	meta.Shift()
	meta.Shift()

	meta.CheckBlock()

	return []token.Token{}, nil
}
