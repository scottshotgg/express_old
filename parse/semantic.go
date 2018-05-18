package parse

import (
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/pkg/errors"
	"github.com/scottshotgg/Express/token"
)

var (
	declaredType  string
	declaredName  string
	declaredValue token.Value
	declarations  = map[string]token.Value{}
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

// CheckIdent ...
func (m *Meta) CheckIdent() {
	fmt.Println("ident")

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

		// if the form is [ident] [= | : | :=] then expect an expression
		// m.GetExpr()

	default:
		fmt.Println("vert da ferk")
		os.Exit(9)
	}
}

// GetFactor ...
func (m *Meta) GetFactor() {
	m.Shift()
	switch m.CurrentToken.Type {
	case "IDENT":
		fmt.Println("found an ident")
		tValue, ok := declarations[m.CurrentToken.Value.String]
		if !ok {
			fmt.Println("Undefined variable reference")
			os.Exit(9)
		}
		if tValue.Type != declaredType {
			fmt.Println("Variable type mismatch")
			fmt.Println("Expected", declaredType, "got", m.CurrentToken.Value.Type)
			os.Exit(9)
		}
		// TODO: we may actually want to say that this is in actuallality the variable and not the value so that it can do optimizations
		declaredValue = tValue
		m.CollectToken(token.Token{
			ID: 1,
			// Type:, // TODO: not sure what to put here
			Value: tValue,
		})

	case "LITERAL":
		fmt.Println("found a literal")
		if m.CurrentToken.Value.Type != declaredType {
			fmt.Println("Variable type mismatch")
			fmt.Println("Expected", declaredType, "got", m.CurrentToken.Value.Type)
			os.Exit(9)
		}
		declaredValue = m.CurrentToken.Value
		m.CollectCurrentToken()

	case "L_PAREN":
		fmt.Println("found an expr")
		m.GetExpr()

	default:
		fmt.Println("ERROR getting factor")
		fmt.Println("Expected factor, got", m.CurrentToken)
		os.Exit(9)
	}
}

// GetTerm ...
func (m *Meta) GetTerm() {
	m.GetFactor()
}

// GetSecOp ...
func (m *Meta) GetSecOp() {
	fmt.Println("current", m.CurrentToken)
}

// AddOperands ...
func (m *Meta) AddOperands(op1, op2 interface{}) (token.Value, error) {
	// total := int(left.Value.True.(int64) + right.Value.True.(int64))
	// valueToken = token.Value{
	// 	Type:   reflect.TypeOf(total).String(),
	// 	True:   total,
	// 	String: strconv.Itoa(total),
	// }
	fmt.Println(op1, op2)

	rightType := reflect.TypeOf(op1)
	leftType := reflect.TypeOf(op2)
	fmt.Println("typeof", leftType, rightType)

	var valueToken token.Value

	if leftType == rightType {
		switch leftType.Kind() {
		case reflect.Int:
			fallthrough
		case reflect.Int8:
			fallthrough
		case reflect.Int16:
			fallthrough
		case reflect.Int32:
			fallthrough
		case reflect.Int64:
			value := int(op1.(int64) + op2.(int64))
			valueToken.Type = "int"
			valueToken.True = value
			valueToken.String = strconv.Itoa(value)

		case reflect.String:
			valueToken.Type = "string"
			valueToken.True = op1.(string) + op2.(string)
			valueToken.String = valueToken.True.(string)

		default:
			fmt.Println("Type not declared for AddOperands", op1, op2, leftType, rightType)
			os.Exit(9)
		}

		return valueToken, nil

	}

	err := errors.New("Could not perform AddOperand on operands")
	fmt.Println(err, op1, op2, leftType, rightType)
	return token.Value{}, err
}

// GetOperationValue ...
func (m *Meta) GetOperationValue(left token.Token, right token.Token, op token.Token) token.Value {
	// token.Value {
	// 	// Type: m.DetermineTypeFromOperation(), // TODO: save this until later when we want to support multi type operations
	// 	Type:
	// },
	// // valueToken.True
	// // valueToken.String
	// m.CollectToken(valueToken)

	// fmt.Println(left.Value.True + right.Value.True)

	// if leftType == rightType {
	// switch on the op
	switch op.Value.String {
	case "+":
		value, err := m.AddOperands(left.Value.True, right.Value.True)
		if err != nil {
			fmt.Println("could not add operands idk wtf happened", left.Value.True, right.Value.True)
			os.Exit(9)
		}
		return value

	default:
		fmt.Println("Invalid operand", op)
	}

	// } else {
	// 	// were gonna have to do something else
	// 	// usupported for now
	// 	fmt.Println("Unsupported operation detected")
	// 	fmt.Println("typeof", leftType, op, rightType)
	// 	os.Exit(9)
	// }

	return token.Value{}
}

// GetExpr ...
func (m *Meta) GetExpr() {
	m.GetTerm()
	value1 := m.LastCollectedToken
	m.RemoveLastCollectedToken()
	fmt.Println("last", m.LastCollectedToken)

	// FIXME: need to make something to evaluate the statement
	if m.NextToken.Type == "SEC_OP" {
		m.Shift()
		op := m.CurrentToken

		m.GetTerm()
		value2 := m.LastCollectedToken
		m.RemoveLastCollectedToken()

		// FIXME: TODO: really should do this with some sort of eval or using reflection
		switch op.Value.String {
		case "+":
			fmt.Println("value1, value2", value1, value2)

			valueToken := token.Token{
				ID:   1,
				Type: "LITERAL",
				// Expected: "",
				Value: m.GetOperationValue(value1, value2, op),
			}
			fmt.Println(valueToken)
			m.CollectToken(valueToken)

		default:
			fmt.Println("Operator not defined")
			fmt.Println("Found operator:", op)
			os.Exit(9)
		}
	}
}

// GetAssignmentStatement ...
func (m *Meta) GetAssignmentStatement() {
	m.Shift()
	switch m.CurrentToken.Type {
	// Get the TYPE
	case "TYPE":
		fmt.Println("found a type")
		// switch m.CurrentToken.Value.String {
		// 	case "int"
		// }
		declaredType = m.CurrentToken.Value.String
		m.CollectCurrentToken()

		// Get the IDENT
		m.Shift()
		if m.CurrentToken.Type != "IDENT" {
			fmt.Println("Syntax error getting assignment_stmt")
			fmt.Println("Expected IDENT, got", m.CurrentToken)
			os.Exit(9)
		}
		_, ok := declarations[m.CurrentToken.Value.String]
		if ok {
			fmt.Println("Variable already declared")
			os.Exit(9)
		}
		declaredName = m.CurrentToken.Value.String
		m.CollectCurrentToken()

		// Get the assignemnt operator
		m.Shift()
		if m.CurrentToken.Type != "ASSIGN" && m.CurrentToken.Type != "INIT" && m.CurrentToken.Type != "SET" {
			fmt.Println("Syntax error getting assignment_stmt")
			fmt.Println("Expected assign_op, got", m.CurrentToken)
			os.Exit(9)
		}
		m.CollectCurrentToken()

		// FIXME: this should return an error that we can check
		m.GetExpr()

		declarations[declaredName] = declaredValue
		declaredType = ""
		declaredName = ""
		declaredValue = token.Value{}
		fmt.Println(declarations)

	default:
		fmt.Println("ERROR getting statement")
		fmt.Println("expected statement beginning, got", m.CurrentToken)
		os.Exit(9)
	}

	// m.GetExpr()
}

// GetStatement ...
func (m *Meta) GetStatement() {
	m.GetAssignmentStatement()
}

// CheckBlock ...
func (m *Meta) CheckBlock() {
	fmt.Println("hi")

	for {
		m.GetStatement()

		// current := m.CurrentToken
		// switch current.Type {
		// case "TYPE":
		// 	// at this point we need to think about the different options
		// 	// int name			: simplest
		// 	// int name = 6 : next simplest
		// 	// name int = 6 : would be taken care of by a random ident
		// 	// we would probably also have other ones for function params and stuff
		// 	m.CheckType()

		// case "IDENT":
		// 	m.CheckIdent()
		// }

		if m.NextToken == (token.Token{}) {
			fmt.Println("returning")
			return
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
