package parse

import (
	"fmt"
	"os"
	"strconv"

	"github.com/pkg/errors"
	"github.com/scottshotgg/Express/token"
)

// EvaluateBinaryOperation ...
// TODO: add in * and / and <
func (m *Meta) EvaluateBinaryOperation(left, right, op token.Value) (token.Value, error) {
	fmt.Println("hi im EvaluateBinaryOperation")
	switch op.Type {
	case "add":
		addValue, err := m.AddOperands(left, right)
		if err != nil {
			return token.Value{}, errors.New("Error adding operands")
		}
		return addValue, nil

	case "sub":
		subValue, err := m.SubOperands(left, right)
		if err != nil {
			return token.Value{}, errors.New("Error subtracting operands")
		}
		return subValue, nil

	case "lthan":
		// lessValue, err := m.LessThanOperands(left, right)
		// if err != nil {
		// 	return token.Value{}, errors.New("Error evaluating boolean expression")
		// }
		// return lessValue, nil
		lt := left.True.(int) < right.True.(int)
		return token.Value{
			Type:       token.BoolType,
			True:       lt,
			String:     strconv.FormatBool(lt),
			AccessType: token.PrivateAccessType,
		}, nil

	default:
		err := errors.Errorf("Undefined operator; left: %+v right: %+v op: %+v", left, right, op)
		fmt.Println(err.Error())
		return token.Value{}, err
	}
}

// EvaluateUnaryOperation ...
// TODO: implement this stuff
func (m *Meta) EvaluateUnaryOperation(left, op token.Value) { // (token.Value, error) {
}

// GetFactor ...
func (m *Meta) GetFactor() (token.Value, error) {
	fmt.Println("GetFactor")
	fmt.Println("m.NextToken", m.NextToken)

	switch m.NextToken.Type {
	case token.Literal:
		m.Shift()
		return m.CurrentToken.Value, nil

	case token.Ident:
		m.Shift()
		if identValue, ok := m.DeclarationMap[m.CurrentToken.Value.String]; ok {
			return identValue, nil
		}
		return token.Value{}, errors.New("Undefined variable reference")

	default:
		return token.Value{}, errors.Errorf("default %+v", m.NextToken)
	}
}

// GetTerm ...
func (m *Meta) GetTerm() (token.Value, error) {
	fmt.Println("GetTerm")

	totalTerm, err := m.GetFactor()
	if err != nil {
		return token.Value{}, err
	}

	for {
		switch m.NextToken.Type {
		case token.SecOp:
			m.Shift()
			fmt.Println("woah i got a secop")
			op := m.CurrentToken
			factor2, ferr := m.GetFactor()
			if ferr != nil {
				return token.Value{}, ferr
			}
			fmt.Println("factor2", factor2)

			totalTerm, err = m.EvaluateBinaryOperation(totalTerm, factor2, op.Value)
			if err != nil {
				return token.Value{}, err
			}

		default:
			return totalTerm, err
		}
	}
}

// GetExpression ...
func (m *Meta) GetExpression() (token.Value, error) {
	fmt.Println("GetExpression")

	switch m.NextToken.Type {
	// Assignment Expression
	case token.Assign:
		m.DeclaredName = m.CurrentToken.Value.String
		m.DeclaredAccessType = m.CurrentToken.Value.Type
		switch m.NextToken.Value.Type {
		case "init":
			if m.DeclaredType != "" {
				return token.Value{}, errors.New("Type with init is not valid")
			}
			m.DeclaredType = token.SetType
			fallthrough

		case "assign":
			m.Shift()
			expr, err := m.GetExpression()
			if err != nil {
				return token.Value{}, err
			}
			fmt.Println("expr", expr)
			if m.DeclaredType == token.SetType {
				m.DeclaredType = expr.Type
			} else if m.DeclaredType != expr.Type {
				// TODO: implicit type casting here
				return token.Value{}, errors.New("No implicit type casting as of now")
			}
			m.DeclarationMap[m.DeclaredName] = token.Value{
				Name:       m.DeclaredName,
				Type:       m.DeclaredType,
				True:       expr.True,
				AccessType: m.DeclaredAccessType,
			}
			return m.DeclarationMap[m.DeclaredName], nil
		}

	default:
		return m.GetTerm()
	}

	return token.Value{}, errors.Errorf("default %+v", m.NextToken)
}

// GetStatement ...
func (m *Meta) GetStatement() (token.Value, error) {
	fmt.Println("GetStatement")

	switch m.NextToken.Type {
	case token.Type:
		m.DeclaredType = m.NextToken.Value.Type
		m.Shift()
		// TODO: could either recurse here, or fallthrough
		if m.NextToken.Type != token.Ident {
			break
		}
		fallthrough

	// TODO: will have to consider declarations too
	case token.Ident:
		fmt.Println("ident", m.NextToken)
		m.Shift()
		return m.GetExpression()

	default:
		fmt.Println("hey its me, the default", m.NextToken)
	}

	return token.Value{}, nil
}

// CheckBlock ...
func (m *Meta) CheckBlock() (token.Value, error) {
	fmt.Println("CheckBlock")
	blockTokens := []token.Value{}

	for {
		stmt, err := m.GetStatement()
		if err != nil {
			fmt.Println("err", err)
			os.Exit(9)
		}
		blockTokens = append(blockTokens, stmt)

		m.DeclaredName = ""
		m.DeclaredType = ""
		m.DeclaredAccessType = ""
		m.DeclaredActingType = ""
		m.DeclaredValue = token.Value{}

		if m.NextToken == (token.Token{}) {
			return token.Value{
				Type: token.Block,
				True: blockTokens,
			}, nil
		}
	}
}

// Semantic ...
func Semantic(tokens []token.Token) ([]token.Value, error) {
	fmt.Println("Semantic")

	meta := Meta{
		AppendDeclarations: true,
		IgnoreWS:           true,
		Tokens:             tokens[0].Value.True.([]token.Token),
		Length:             len(tokens[0].Value.True.([]token.Token)),
		CheckOptmization:   true,
		DeclarationMap:     map[string]token.Value{},
	}
	meta.Shift()

	block, err := meta.CheckBlock()
	if err != nil {
		// TODO:
		return []token.Value{}, err
	}
	fmt.Println("block", block)

	fmt.Println("declarationMap", meta.DeclarationMap)

	return []token.Value{block}, nil
}

// TODO: start here
// TODO: use next token
// TODO: start very simply with the definition in documentation/notes_about_shit
// TODO: VERY SIMPLE requirements parsing vars with the return architecture of semantic2
