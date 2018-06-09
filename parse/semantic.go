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
		// fmt.Printf("tValue %+v, ok %b\n", tValue, ok)
		if !ok {
			fmt.Println("Undefined variable reference")
			os.Exit(9)
		}
		if m.DeclaredType != token.SetType && m.DeclaredType != token.VarType && tValue.Type != m.DeclaredType {
			fmt.Println("Variable type mismatch")
			fmt.Println("Expected", m.DeclaredType, "got", tValue.Type)
			fmt.Println(tValue)
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
		if m.DeclaredType != token.SetType && m.DeclaredType != token.VarType && m.CurrentToken.Value.Type != m.DeclaredType {
			fmt.Println("Variable type mismatch")
			fmt.Println("Expected", m.DeclaredType, "got", m.CurrentToken.Value.Type)
			// fmt.Println()
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
		fmt.Println("declarationMap", m.DeclarationMap)

		meta := Meta{
			IgnoreWS:         true,
			Tokens:           m.CurrentToken.Value.True.([]token.Token),
			Length:           len(m.CurrentToken.Value.True.([]token.Token)),
			CheckOptmization: true,
			DeclarationMap:   map[string]token.Value{},
			// InheritedMap:     m.DeclarationMap, // TODO: anonymous scopes will not need this, if check it?
		}
		meta.Shift()
		dMap := meta.CheckBlock()

		// TODO: this will probably need to change when start doing functions but this is fine for now
		// Filter out all private declared entites
		// Only publicly declared entities should be return from a scope/object
		// FIXME: do not filter out yet
		// for key, value := range dMap {
		// 	if value.AccessType != token.PublicAccessType {
		// 		delete(dMap, key)
		// 	}
		// }

		m.DeclaredValue = token.Value{
			Type: token.ObjectType,
			True: dMap,
		}
		m.DeclaredValue.Name = m.DeclaredName
		m.CollectToken(token.Token{
			ID:    0,
			Type:  token.Block,
			Value: m.DeclaredValue,
		})

	case token.Array:
		fmt.Println(m.DeclaredType, m.DeclaredValue)
		fmt.Println("found an array current", m.CurrentToken)
		m.DeclaredValue = token.Value{
			Type:   token.ArrayType,
			Acting: m.DeclaredValue.Acting,
			True:   m.CheckArray(),
		}
		fmt.Printf("collecting this token %+v\n", m.DeclaredValue)
		m.CollectToken(token.Token{
			ID:    0,
			Type:  token.Array,
			Value: m.DeclaredValue,
		})

		// TODO: this needs an error
		// _, arrayValue := m.CheckArray()
		// if err != nil {
		// 	// TODO:
		// 	fmt.Println("ERROR: need to handle this", err)
		// }
		// m.DeclaredValue = token.Value{
		// 	Type: token.ArrayType,
		// 	// Acting: arrayType,
		// 	True: arrayValue,
		// }
		// os.Exit(1)
		// look through the entire array and analyze the types
		// check for type declaration
		// if no type defined, array of vars (can be multitype)
		// follow normal assignment rules, init, set, type implications, etc
		// on and off switch for static array type implication

	case token.Type:
		m.Shift()
		if m.NextToken.Type == token.Array {
			// arrayType := m.CurrentToken.Value.Type
			m.Shift()
			// TODO: this needs an error
			// _, arrayValue := m.CheckArray()
			m.CheckArray()
			// if err != nil {
			// 	// TODO:
			// 	fmt.Println("ERROR: need to handle this", err)
			// }
			// m.DeclaredValue = token.Value{
			// 	Type:   token.ArrayType,
			// 	Acting: arrayType,
			// 	True:   arrayValue,
			// }
		}

	default:
		fmt.Println("ERROR getting factor")
		fmt.Println("Expected factor, got", m.CurrentToken)
		os.Exit(9)
	}

	// idk why this was here but it was breaking addition with idents in it...
	// if m.DeclaredType != token.SetType && m.DeclaredType != token.VarType && m.CurrentToken.Value.Type != m.DeclaredType {
	// 	fmt.Println("Variable type mismatch")
	// 	fmt.Println("Expected", m.DeclaredType, "got", m.CurrentToken.Value.Type)
	// 	os.Exit(9)
	// }
}

// GetTerm gets the next term in the sequence
// TODO: need to check for statement seaparators
func (m *Meta) GetTerm() {
	// m.GetFactor()
	m.GetFactor()

	// FIXME: need to make something to evaluate the statement
	if m.NextToken.Type == token.PriOp {
		value1 := m.LastCollectedToken
		m.RemoveLastCollectedToken()
		fmt.Println("factor value1", value1)
		fmt.Println("last", m.LastCollectedToken)

		m.Shift()
		op := m.CurrentToken

		m.GetTerm()
		value2 := m.LastCollectedToken
		m.RemoveLastCollectedToken()

		// FIXME: TODO: really should do this with some sort of eval or using reflection
		// switch op.Value.String {
		// case "+":
		fmt.Println("value1, value2, op", value1, value2, op)

		valueToken := token.Token{
			ID:   1,
			Type: token.Literal,
			// Expected: "",
			Value: m.GetOperationValue(value1, value2, op),
		}
		// FIXME: ??? fix this
		fmt.Println("ACCESS", valueToken.Value.AccessType)
		valueToken.Value.AccessType = m.DeclaredAccessType

		fmt.Println("valueToken", valueToken)
		m.CollectToken(valueToken)

		m.DeclaredValue = valueToken.Value

		// default:
		// 	fmt.Println("Operator not defined")
		// 	fmt.Println("Found operator:", op)
		// 	os.Exit(9)
		// }
	} else if m.NextToken.Type == token.LThan {
		// FIXME: this needs to handle boolean and other things
		// just ints for now though, we should just put this up there with token.PriOp
		// or call the same function, too lazy right now though
		// TODO: here is where we can have boolean exppressions

		value1 := m.LastCollectedToken
		m.RemoveLastCollectedToken()
		fmt.Println("factor value1", value1)
		fmt.Println("last", m.LastCollectedToken)

		m.Shift()
		// op := m.CurrentToken

		m.GetTerm()
		value2 := m.LastCollectedToken
		m.RemoveLastCollectedToken()

		fmt.Println(m.DeclaredType)
		fmt.Println(m.DeclaredValue)
		fmt.Println(m.CurrentToken)

		m.DeclaredType = token.BoolType
		m.DeclaredValue = token.Value{
			Type: token.BoolType,
			True: value1.Value.True.(int) < value2.Value.True.(int),
		}
	}
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

	case "-":
		value, err := m.SubOperands(left.Value, right.Value)
		if err != nil {
			fmt.Println("could not sub operands idk wtf happened", left.Value, right.Value)
			os.Exit(9)
		}
		return value

	case "*":
		value, err := m.MultOperands(left.Value, right.Value)
		if err != nil {
			fmt.Println("could not mult operands idk wtf happened", left.Value, right.Value)
			os.Exit(9)
		}
		return value

	case "/":
		value, err := m.DivOperands(left.Value, right.Value)
		if err != nil {
			fmt.Println("could not mult operands idk wtf happened", left.Value, right.Value)
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

	// FIXME: need to make something to evaluate the statement
	if m.NextToken.Type == token.SecOp {
		value1 := m.LastCollectedToken
		m.RemoveLastCollectedToken()
		fmt.Println("last", m.LastCollectedToken)

		m.Shift()
		op := m.CurrentToken

		opValue := token.Value{}
		if m.CurrentToken.Value.String == "+" && m.NextToken.Value.String == "+" {
			// TODO: FIXME: this needs to be fixed and shouldn't even really be done here
			opValue = token.Value{
				Type: token.IntType,
				True: value1.Value.True.(int) + 1,
			}
			m.Shift()
		} else {
			m.GetExpression()
			value2 := m.LastCollectedToken
			m.RemoveLastCollectedToken()
			fmt.Println("value1, value2, op", value1, value2, op)
			ov := m.GetOperationValue(value1, value2, op)
			opValue = ov
			fmt.Println("opValue", opValue)
		}

		// FIXME: TODO: really should do this with some sort of eval or using reflection
		// switch op.Value.String {
		// case "+":

		valueToken := token.Token{
			ID:   1,
			Type: token.Literal,
			// Expected: "",
			Value: opValue,
		}
		// FIXME: ??? fix this
		fmt.Println("ACCESS", valueToken.Value.AccessType)
		valueToken.Value.AccessType = m.DeclaredAccessType

		fmt.Println("valueToken", valueToken)
		m.CollectToken(valueToken)

		fmt.Println("DECLAREDVLUE", valueToken.Value)
		m.DeclaredValue = valueToken.Value

		// default:
		// 	fmt.Println("Operator not defined")
		// 	fmt.Println("Found operator:", op)
		// 	os.Exit(9)
		// }
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
		fmt.Println(m.CurrentToken.Value.String)
		fmt.Println("HEY PICK ME BRAH")
		m.DeclaredType = m.CurrentToken.Value.String
		m.DeclaredActingType = m.CurrentToken.Value.Acting
		fmt.Println(m.DeclaredActingType)

		m.CollectCurrentToken()

		// Get the IDENT
		m.Shift()
		// switch m.CurrentToken.Type {
		// case token.Ident:
		// }
		if m.CurrentToken.Type != token.Ident && m.CurrentToken.Type != token.Array {
			// TODO: logic ist very fucky, find better way
			// if m.CurrentToken.Type == token.Array {
			// 	fmt.Println("ARRAY I CHOOSE UUUUUU")
			// 	m.GetFactor()
			// meta := Meta{
			// 	IgnoreWS:         true,
			// 	Tokens:           m.CurrentToken.Value.True.([]token.Token),
			// 	Length:           len(m.CurrentToken.Value.True.([]token.Token)),
			// 	CheckOptmization: true,
			// 	DeclarationMap:   map[string]token.Value{},
			// }
			// meta.Shift()
			// meta.CheckArray()
			// fmt.Println(meta.DeclarationMap)
			// } else {
			fmt.Println("Syntax error getting assignment_stmt")
			fmt.Println("Expected IDENT, got", m.CurrentToken)
			os.Exit(9)
			// }
		}
		if _, ok := m.DeclarationMap[m.CurrentToken.Value.String]; ok {
			fmt.Println("Variable already declared")
			os.Exit(9)
		}
		fmt.Println("m.CurrentToken.Value.Type", m.CurrentToken.Value.Type)
		fmt.Println("m.NextToken.Value.Type", m.NextToken.Value.Type)
		m.DeclaredAccessType = m.CurrentToken.Value.Type
		m.DeclaredName = m.CurrentToken.Value.String
		m.CollectCurrentToken()

		// Get the assignment operator
		m.Shift()
		fmt.Println("m.CurrentToken.Value.Type", m.CurrentToken.Value.Type)
		fmt.Println("m.NextToken.Value.Type", m.NextToken.Value.Type)
		switch m.CurrentToken.Type {
		case token.Assign:
			fallthrough
		case token.Init:
			fallthrough
		case token.Set:
			m.CollectCurrentToken()

			// case token.Array:
			// 	m.Shift()
		}
		// if m.CurrentToken.Type != token.Assign && m.CurrentToken.Type != token.Init && m.CurrentToken.Type != token.Set {
		// 	// switch m.CurrentToken.Value.Type {
		// 	// case "init":
		// 	// case "set":
		// 	// case "assign":
		// 	// default:
		// 	// 	fmt.Println("ERROR how did we get in here", m.CurrentToken)
		// 	// }

		// 	fmt.Println("Syntax error getting assignment_stmt")
		// 	fmt.Println("Expected assign_op, got", m.CurrentToken)
		// 	os.Exit(9)
		// }
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
		m.DeclaredValue.Name = m.DeclaredName
		m.CollectToken(token.Token{
			ID: 0,
			// Type:  m.DeclaredName, // TODO: not sure about this yet
			Value: m.DeclaredValue,
		})

		m.LLVMTokens = append(m.LLVMTokens, token.Token{
			ID:    0,
			Type:  "VAR",
			Value: m.DeclaredValue,
		})

		m.DeclaredType = ""
		m.DeclaredName = ""
		m.DeclaredAccessType = ""
		m.DeclaredValue = token.Value{}
		fmt.Println(m.DeclarationMap)

	case token.Ident:
		fmt.Println("i spy an ident")
		currentIdent := m.CurrentToken
		fmt.Println(currentIdent)

		m.Shift()
		if m.CurrentToken.Type == token.Assign {
			fmt.Println("CURRENTIDENT", currentIdent.Value.Type)
			m.DeclaredAccessType = currentIdent.Value.Type
			m.DeclaredName = currentIdent.Value.String
			current, ok := m.DeclarationMap[currentIdent.Value.String]
			if m.CurrentToken.Value.Type == "assign" {
				if ok {
					m.DeclaredAccessType = current.AccessType
					m.DeclaredType = current.Type
				} else {
					fmt.Println("Variable reference not found", currentIdent)
					os.Exit(9)
				}
			} else if m.CurrentToken.Value.Type == "set" {
				// check that the var is NOT there already
				if !ok {
					m.DeclaredAccessType = current.AccessType
					m.DeclaredType = token.SetType
				} else {
					fmt.Println("Variable reference already declared", currentIdent)
					os.Exit(9)
				}
			} else if m.CurrentToken.Value.Type == "init" {
				// if the var is there then set it to the value (check types)
				// else make the var
				if ok {
					m.DeclaredAccessType = current.AccessType
					m.DeclaredType = current.Type
				} else {
					// FIXME: will have to look at this
					m.DeclaredAccessType = currentIdent.Value.Type
					m.DeclaredType = token.SetType
				}
			} else {
				fmt.Println("something happened")
				os.Exit(8)
			}

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
			m.DeclaredValue.Name = m.DeclaredName
			m.CollectToken(token.Token{
				ID:    0,
				Type:  m.DeclaredName,
				Value: m.DeclaredValue,
			})

			m.LLVMTokens = append(m.LLVMTokens, token.Token{
				ID:    0,
				Type:  "VAR",
				Value: m.DeclaredValue,
			})

			m.DeclaredType = ""
			m.DeclaredName = ""
			m.DeclaredAccessType = ""
			m.DeclaredValue = token.Value{}
			fmt.Println(m.DeclarationMap)
		}

	case token.Separator:
		m.Shift()

		// TODO: Whatever processing we need to do needs to be added on here
		m.DeclarationMap[m.DeclaredName] = m.DeclaredValue
		m.DeclaredValue.Name = m.DeclaredName
		m.CollectToken(token.Token{
			ID:    0,
			Type:  m.DeclaredName,
			Value: m.DeclaredValue,
		})

		m.LLVMTokens = append(m.LLVMTokens, token.Token{
			ID:    0,
			Type:  "VAR",
			Value: m.DeclaredValue,
		})

		m.DeclaredType = ""
		m.DeclaredName = ""
		m.DeclaredAccessType = ""
		m.DeclaredValue = token.Value{}
		fmt.Println(m.DeclarationMap)

	default:
		fmt.Println("ERROR getting assignement statement")
		fmt.Println("expected assignment statement beginning, got", m.CurrentToken)
		return errors.New("blah")
	}

	// m.GetExpression()
	return nil
}

// GetAnonymousScope ...
func (m *Meta) GetAnonymousScope() error { //(token.Value, error) {
	meta := Meta{
		IgnoreWS:         true,
		Tokens:           m.CurrentToken.Value.True.([]token.Token),
		Length:           len(m.CurrentToken.Value.True.([]token.Token)),
		CheckOptmization: true,
		DeclarationMap:   map[string]token.Value{},
	}
	meta.Shift()

	fmt.Println("m.CURRENTTT", m.CurrentToken)

	dMap := meta.CheckBlock()

	// TODO: this will probably need to change when start doing functions but this is fine for now
	// Filter out all private declared entites
	// Only publicly declared entities should be return from a scope/object
	// FIXME: do not filter out yet
	// for key, value := range dMap {
	// 	if value.AccessType != token.PublicAccessType {
	// 		delete(dMap, key)
	// 	}
	// }

	m.DeclaredValue = token.Value{
		Type: token.ObjectType,
		True: dMap,
	}
	m.CollectToken(token.Token{
		ID:    0,
		Type:  token.Block,
		Value: m.DeclaredValue,
	})

	return nil

	// return token.Value{
	// 	// Type:   token.ArrayType,
	// 	// Acting: token.VarType,
	// 	Type: token.ObjectType,
	// 	True: meta.CheckBlock(),
	// }, nil

	// return token.Value
}

// GetStatement gets the next statement in the sequence
func (m *Meta) GetStatement() error {
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

	err := m.GetAssignmentStatement()
	if err == nil {
		return nil
	}

	// FIXME: ideally we should do a switch on the error but w/e for now
	// TODO: we need some way to backtrack to before this operation ...
	// if err != nil {
	// TODO: woah we could do partial compilation
	return errors.New("error getting statement")
	// }
}

// CheckArray builds an array of token.Values and deduces the type that they need to be
func (m *Meta) CheckArray() []token.Value {
	fmt.Println("checking array", m.CurrentToken)
	fmt.Printf("stufffzzz %+v\n", m.DeclaredValue)
	// os.Exit(9)

	arrayType := m.DeclaredActingType

	fmt.Println("ARRAY TYPE", arrayType)

	arrayTokens := m.CurrentToken.Value.True.([]token.Token)

	meta := Meta{
		IgnoreWS:         true,
		Tokens:           arrayTokens,
		Length:           len(arrayTokens),
		CheckOptmization: true,
		DeclarationMap:   m.DeclarationMap,
		InheritedMap:     m.DeclarationMap,
	}

	fmt.Println("declarationMap", meta.DeclarationMap)

	meta.Shift()

	fmt.Printf("%+v\n", meta)

	aTokens := []token.Value{}
	for {
		fmt.Println("stuff for doing stuff 11")
		fmt.Println(arrayType, m.DeclaredValue, meta.CurrentToken.Value.Type, m.DeclaredType, m.DeclaredActingType)
		if meta.NextToken == (token.Token{}) {
			break
		}

		if meta.NextToken.Type == token.Ident {
			meta.DeclaredType = token.SetType

		} else {
			meta.DeclaredType = meta.NextToken.Value.Type
		}

		// meta.DeclaredType = meta.NextToken.Value.Type

		meta.GetExpression()

		if arrayType == "" { // && arrayType != "var" {
			if meta.CurrentToken.Type == token.Ident {
				arrayType = meta.DeclaredValue.Type
			} else {
				arrayType = meta.CurrentToken.Value.Type
			}
			fmt.Println("SETTING arrayType", arrayType)
		} else {
			fmt.Println("stuff for doing stuff")
			fmt.Println(arrayType, meta.DeclaredValue.Type, meta.CurrentToken.Value.Type, meta.DeclaredType, m.DeclaredActingType)

			if m.DeclaredActingType != meta.CurrentToken.Value.Type {
				if meta.DeclaredType == token.SetType {
					arrayType = token.VarType
				} else {
					fmt.Println("ERROR: cannot use", meta.DeclaredValue.Type, "as type", arrayType, "in", meta.DeclaredName)
					os.Exit(8)
				}
			} else {
				if meta.CurrentToken.Type == token.Ident {
					arrayType = meta.DeclaredValue.Type
				} else {
					arrayType = meta.CurrentToken.Value.Type
				}
			}
		}

		aTokens = append(aTokens, meta.DeclaredValue)
	}

	fmt.Println("arrayType, m.DeclaredType, m.DeclaredValue.Type", arrayType, m.DeclaredType, m.DeclaredValue.Type)

	if m.DeclaredType == token.SetType {
		// m.DeclaredType = arrayType
		m.DeclaredValue.Acting = arrayType
	}

	fmt.Println(aTokens)
	return aTokens
}

// CheckGroup ...
func (m *Meta) CheckGroup() []token.Value {
	m.Shift()
	fmt.Println("checking group")
	fmt.Println(m.CurrentToken)

	groupTokens := []token.Value{}

	for {
		// TODO: dont do any checking for now, idc
		// FIXME: need to do checking for bool expressions, maybe we should just change to get expression?

		groupTokens = append(groupTokens, m.CurrentToken.Value)

		if m.NextToken == (token.Token{}) {
			break
		}
	}

	return groupTokens
}

// CheckFunctionDef ...
func (m *Meta) CheckFunctionDef() []token.Value {
	m.Shift()
	fmt.Println("checking function d	efinition", m.Tokens)
	fmt.Println(m.CurrentToken)

	if len(m.Tokens) > 3 {
		fmt.Println("ERROR: More than three parts for function definition")
		os.Exit(9)
	}

	meta := Meta{
		IgnoreWS:         true,
		Tokens:           m.CurrentToken.Value.True.([]token.Token),
		Length:           len(m.CurrentToken.Value.True.([]token.Token)),
		CheckOptmization: true,
		DeclarationMap:   m.DeclarationMap,
	}
	meta.Shift()

	fmt.Println(meta.CheckGroup())

	// if m.NextToken == (token.Token{}) {
	// 	break
	// }

	return []token.Value{}
}

// CheckBlock check the usage of the block
func (m *Meta) CheckBlock() map[string]token.Value {
	var err error
	for {
		fmt.Printf("CURRENT %+v\n", m.NextToken.Type)

		switch m.NextToken.Type {
		case token.Keyword:
			fmt.Println("found a keyword")
			m.Shift()

			switch m.CurrentToken.Value.String {
			case "for":
				fmt.Println("found a for")

				// m.Shift()

				fmt.Println("current", m.CurrentToken)

				// START of the for loop
				as := m.GetAssignmentStatement()

				fmt.Println("as", as)
				m.Shift()
				fmt.Println("getexpr", m.CurrentToken)

				// TODO: FIXME: major hack, take the logic to check the type of vars out of GetFactor
				m.DeclaredType = m.LastCollectedToken.Value.Type

				// END of the for loop
				m.GetExpression()
				fmt.Println(m.DeclaredValue)

				// TODO: at this point the loop isn't needed at all, don't even emit the tokens
				if m.DeclaredValue.True.(bool) == false {
					fmt.Println("loop not needed")
					// TODO: we still need to clear the loop tokens, just don't generate the end token
				} else {
					// FIXME: TODO: major hack, take the logic to check the type of vars out of GetFactor
					m.DeclaredType = token.IntType
					m.Shift()

					// STEP of the for loop
					m.GetExpression()
					m.Shift()

					meta := Meta{
						IgnoreWS:         true,
						Tokens:           m.CurrentToken.Value.True.([]token.Token),
						Length:           len(m.CurrentToken.Value.True.([]token.Token)),
						CheckOptmization: true,
						DeclarationMap:   map[string]token.Value{},
					}
					meta.Shift()

					// BODY of the loop
					// TODO: this needs to be fixed because we only get back variables.... ?
					meta.CheckBlock()

					m.CollectToken(token.Token{
						ID:    0,
						Type:  "FOR",
						Value: token.Value{
						// True:
						// The true value should be a map with three parts:
						// 1. the range; start and end
						// 2. the steps; how to get from the start to the end
						// 3. the body; what are we doing at each iteration
						},
					})
				}

				// expect FUNCTION definition
				// expect grouping
				// grouping has 0-3 expressions within it
				// expect block

				// maybe instead of that -
				// check for
				// 1. assignment statement (standard for)
				// 2. boolean qualifier
				// 3. add statement
				// OR
				// 1. boolean qualifer (while)
				// ORf
				// 1. nothing (for-ever)

				// fmt.Println(m.CurrentToken.Value)
				// meta := Meta{
				// 	IgnoreWS:         true,
				// 	Tokens:           m.CurrentToken.Value.True.([]token.Token),
				// 	Length:           len(m.CurrentToken.Value.True.([]token.Token)),
				// 	CheckOptmization: true,
				// 	DeclarationMap:   m.DeclarationMap,
				// }
				// meta.Shift()

				// theRest := meta.CheckFunctionDef()

				// check that we have up to 3 arguments

				fmt.Println()
				continue

			case "if":
				fmt.Println("found an if")
				// expect grouping
				// grouping size of 0-3
				// expect block
				// evaluate based on group of block

			case "else":
				// expect block

			default:
				fmt.Println("hey its me", m.CurrentToken)
			}
		}

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

		err = m.GetStatement()
		if err != nil {
			fmt.Println("ERROR: could not get statement", err)
			os.Exit(8)
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

	// return meta.EndTokens, nil

	// FIXME: TODO: hacking this in for the C++ translator for now, this is only variable statements
	tks := []token.Token{}
	for k, v := range meta.DeclarationMap {
		v.Name = k
		tks = append(tks, token.Token{
			ID:    0,
			Type:  "VAR",
			Value: v,
		})
	}

	// return meta.LLVMTokens, nil
	return tks, nil
}
