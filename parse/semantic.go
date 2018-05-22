package parse

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/scottshotgg/Express/token"
)

// CheckType checks the usage of the type
func (m *Meta) CheckType() {
	fmt.Println("CheckType")

	switch m.NextToken.Type {
	case token.Ident:

		m.CollectCurrentToken()
		m.Shift()
		m.CollectCurrentToken()

		//TODO: change all of these to be assign
		switch m.NextToken.Type {
		case token.Init:
			fallthrough
		case token.Set:
			fallthrough
		case token.Assign:
			fmt.Println("found an ASSIGN type")
			m.Shift()
			m.CollectCurrentToken()

			fmt.Println(m.NextToken)
			if m.NextToken.Type != token.Literal && m.NextToken.Type != token.Ident {
				fmt.Println("did not find a literal")
				os.Exit(0)
			}

			m.Shift()
			m.CollectCurrentToken()
		}
	}
}

// CheckIdent check the usage of the ident
func (m *Meta) CheckIdent() {
	fmt.Println(token.Ident)

	m.CollectCurrentToken()

	switch m.NextToken.Type {
	case token.Init:
		fallthrough
	case token.Set:
		fallthrough
	case token.Assign:
		fmt.Println("found an ASSIGN type")
		m.Shift()
		m.CollectCurrentToken()

		// if the form is [ident] [= | : | :=] then expect an expression
		m.GetExpression()

	default:
		fmt.Println("vert da ferk")
		os.Exit(9)
	}
}

// GetFactor returns the next factor in the sequence
func (m *Meta) GetFactor() {
	m.Shift()
	switch m.CurrentToken.Type {
	case token.Ident:
		fmt.Println("found an ident")
		tValue, ok := m.DeclarationMap[m.CurrentToken.Value.String]
		if !ok {
			fmt.Println("Undefined variable reference")
			os.Exit(9)
		}
		if m.DeclaredType != token.VarType && tValue.Type != m.DeclaredType {
			fmt.Println("Variable type mismatch")
			fmt.Println("Expected", m.DeclaredType, "got", tValue.Type)
			os.Exit(9)
		}

		// TODO: we may actually want to say that this is in actuallality the variable and not the value so that it can do optimizations
		m.DeclaredValue = tValue
		fmt.Println("declareds", m.DeclaredName, m.DeclaredType, m.DeclaredValue)
		m.CollectToken(token.Token{
			ID: 1,
			// Type:, // TODO: not sure what to put here
			Value: tValue,
		})

	case token.Literal:
		fmt.Println("found a literal")
		if m.DeclaredType != token.VarType && m.CurrentToken.Value.Type != m.DeclaredType {
			fmt.Println("Variable type mismatch")
			fmt.Println("Expected", m.DeclaredType, "got", m.CurrentToken.Value.Type)
			os.Exit(9)
		}

		m.DeclaredValue = m.CurrentToken.Value
		m.CollectCurrentToken()

	case token.LParen:
		fmt.Println("found an expr")
		m.GetExpression()

	case token.Block:
		// FIXME: remove this hack shit later
		fmt.Println("m.DeclaredName", m.DeclaredName)
		fmt.Println("found ze bracket")
		meta := Meta{
			IgnoreWS:         true,
			Tokens:           m.CurrentToken.Value.True.([]token.Token),
			Length:           len(m.CurrentToken.Value.True.([]token.Token)),
			CheckOptmization: true,
			DeclarationMap:   map[string]token.Value{},
		}
		fmt.Println(meta)
		fmt.Println(len(meta.Tokens))
		meta.Shift()
		dMap := meta.CheckBlock()
		fmt.Println("dMap", dMap)
		fmt.Println("m.DeclaredName", m.DeclaredName)

		// TODO: this will probably need to change when start doing functions but this is fine for now
		// Filter out all private declared entites
		// Only publicly declared entities should be return from a scope/object
		for key, value := range dMap {
			if value.AccessType != token.PublicAccessType {
				delete(dMap, key)
			}
		}

		m.DeclaredValue = token.Value{
			Type: token.ObjectType,
			True: dMap,
		}

	default:
		fmt.Println("ERROR getting factor")
		fmt.Println("Expected factor, got", m.CurrentToken)
		os.Exit(9)
	}
}

// GetTerm gets the next term in the sequence
func (m *Meta) GetTerm() {
	m.GetFactor()
}

// GetSecOp gets a secondary operation; + and -
func (m *Meta) GetSecOp() {
	fmt.Println("current", m.CurrentToken)
}

// GetOperationValue returns the value of the operation being performed in the statement
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
		value, err := m.AddOperands(left.Value, right.Value)
		if err != nil {
			fmt.Println("could not add operands idk wtf happened", left.Value, right.Value)
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

// GetExpression gets the next expression
func (m *Meta) GetExpression() {
	m.GetTerm()
	value1 := m.LastCollectedToken
	m.RemoveLastCollectedToken()
	fmt.Println("last", m.LastCollectedToken)

	// FIXME: need to make something to evaluate the statement
	if m.NextToken.Type == token.SecOp {
		m.Shift()
		op := m.CurrentToken

		m.GetExpression()
		value2 := m.LastCollectedToken
		m.RemoveLastCollectedToken()

		// FIXME: TODO: really should do this with some sort of eval or using reflection
		switch op.Value.String {
		case "+":
			fmt.Println("value1, value2", value1, value2)

			valueToken := token.Token{
				ID:   1,
				Type: token.Literal,
				// Expected: "",
				Value: m.GetOperationValue(value1, value2, op),
			}
			valueToken.Value.AccessType = m.DeclaredAccessType

			fmt.Println("valueToken", valueToken)
			m.CollectToken(valueToken)

			m.DeclaredValue = valueToken.Value

		default:
			fmt.Println("Operator not defined")
			fmt.Println("Found operator:", op)
			os.Exit(9)
		}
	}
}

// GetAssignmentStatement gets the next assignment statement in the sequence
func (m *Meta) GetAssignmentStatement() error {
	m.Shift()
	switch m.CurrentToken.Type {
	// Get the TYPE
	case token.Type:
		fmt.Println("found a type")
		// switch m.CurrentToken.Value.String {
		// 	case "int"
		// }
		m.DeclaredType = m.CurrentToken.Value.String

		m.CollectCurrentToken()

		// Get the IDENT
		m.Shift()
		if m.CurrentToken.Type != token.Ident {
			fmt.Println("Syntax error getting assignment_stmt")
			fmt.Println("Expected IDENT, got", m.CurrentToken)
			os.Exit(9)
		}
		if _, ok := m.DeclarationMap[m.CurrentToken.Value.String]; ok {
			fmt.Println("Variable already declared")
			os.Exit(9)
		}
		fmt.Println("m.CurrentToken.Value.Type", m.CurrentToken.Value.Type)
		m.DeclaredAccessType = m.CurrentToken.Value.Type
		m.DeclaredName = m.CurrentToken.Value.String
		m.CollectCurrentToken()

		// Get the assignment operator
		m.Shift()
		if m.CurrentToken.Type != token.Assign && m.CurrentToken.Type != token.Init && m.CurrentToken.Type != token.Set {
			fmt.Println("Syntax error getting assignment_stmt")
			fmt.Println("Expected assign_op, got", m.CurrentToken)
			os.Exit(9)
		}
		m.CollectCurrentToken()

		// FIXME: this should return an error that we can check
		m.GetExpression()

		// FIXME: this is changing the variable type to 'var', should probably have a 'realType' and an 'actingType'
		if m.DeclaredType == token.VarType {
			m.DeclaredValue.Acting = m.DeclaredValue.Type
			m.DeclaredValue.Type = token.VarType
			fmt.Println("wtf", m.DeclaredValue)
		}
		m.DeclaredValue.AccessType = m.DeclaredAccessType
		fmt.Printf("DECLARED %+v\n", m.DeclaredValue)

		m.DeclarationMap[m.DeclaredName] = m.DeclaredValue
		m.DeclaredType = ""
		m.DeclaredName = ""
		m.DeclaredAccessType = ""
		m.DeclaredValue = token.Value{}
		fmt.Println(m.DeclarationMap)

	case token.Ident:
		fmt.Println("i spy an ident")
		currentIdent := m.CurrentToken

		if current, ok := m.DeclarationMap[currentIdent.Value.String]; ok {
			m.DeclaredAccessType = current.AccessType
			m.DeclaredName = currentIdent.Value.String
			m.DeclaredType = current.Type
		} else {
			fmt.Println("Variable reference not found", currentIdent)
			os.Exit(9)
		}

		m.Shift()
		if m.CurrentToken.Type == token.Assign {

			fmt.Println(m.CurrentToken)

			// FIXME: this should return an error that we can check
			m.GetExpression()

			// FIXME: this is changing the variable type to 'var', should probably have a 'realType' and an 'actingType'
			if m.DeclaredType == token.VarType {
				m.DeclaredValue.Acting = m.DeclaredValue.Type
				m.DeclaredValue.Type = token.VarType
			}
			m.DeclaredValue.AccessType = m.DeclaredAccessType
			fmt.Printf("DECLARED %+v\n", m.DeclaredValue)

			m.DeclarationMap[m.DeclaredName] = m.DeclaredValue
			m.DeclaredType = ""
			m.DeclaredName = ""
			m.DeclaredAccessType = ""
			m.DeclaredValue = token.Value{}
			fmt.Println(m.DeclarationMap)
		}

	default:
		fmt.Println("ERROR getting assignement statement")
		fmt.Println("expected assignement statement beginning, got", m.CurrentToken)
		return errors.New("blah")
	}

	// m.GetExpression()
	return nil
}

// GetStatement gets the next statement in the sequence
func (m *Meta) GetStatement() {
	err := m.GetAssignmentStatement()
	// FIXME: ideally we should do a switch on the error but w/e for now
	// TODO: we need some way to backtrack to before this operation ...
	if err != nil {
		fmt.Println("error getting statement", err)
		// TODO: woah we could do partial compilation
		os.Exit(9)
	}
}

// CheckBlock check the usage of the block
func (m *Meta) CheckBlock() map[string]token.Value {
	for {
		m.GetStatement()

		// current := m.CurrentToken
		// switch current.Type {
		// case token.Type:
		// 	// at this point we need to think about the different options
		// 	// int name			: simplest
		// 	// int name = 6 : next simplest
		// 	// name int = 6 : would be taken care of by a random ident
		// 	// we would probably also have other ones for function params and stuff
		// 	m.CheckType()

		// case token.Ident:
		// 	m.CheckIdent()
		// }

		if m.NextToken == (token.Token{}) {
			fmt.Println("returning")
			return m.DeclarationMap
		}
	}
}

// Semantic runs a semantic parse on the tokens
func Semantic(tokens []token.Token) ([]token.Token, error) {
	// Auto inject the brackets to ensure that they are there
	meta := Meta{
		IgnoreWS:         true,
		Tokens:           tokens[0].Value.True.([]token.Token),
		Length:           len(tokens[0].Value.True.([]token.Token)),
		CheckOptmization: true,
		DeclarationMap:   map[string]token.Value{},
	}
	meta.Shift()

	fmt.Println(tokens)

	meta.CheckBlock()
	fmt.Println("tokens", meta.EndTokens)
	fmt.Println()
	fmt.Println("DECLARATION MAP:")
	declarationMapJSON, err := json.MarshalIndent(meta.DeclarationMap, "", "\t")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(declarationMapJSON))

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
