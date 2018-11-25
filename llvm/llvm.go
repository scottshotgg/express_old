package llvm

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/scottshotgg/Express/token"
)

func boolToInt(value bool) int64 {
	if value {
		return 1
	}

	return 0
}

func stringToArray(value string) *constant.CharArray {
	return constant.NewCharArray([]byte(value))
}

// Translate ...
func Translate(tokens []token.Token) {
	// TODO: FIXME: need to fix the unique token names coming from the semantic parser
	fmt.Println("tokens", tokens)

	// Create a new LLVM IR module.
	m := ir.NewModule()

	// TODO: FIXME: this needs to be based on the flag or w/e to turn off dynamic types
	// m.NewType("var", types.NewStruct([]types.Type{
	// 	types.NewInt(8),
	// 	types.NewPointer(types.NewInt(32)),
	// }...))

	mainFunc := m.NewFunc("main", types.I32)
	mainBlock := mainFunc.NewBlock("main")

	// iPtr := mainBlock.NewAlloca(types.NewPointer(types.I32))
	// mainBlock.NewStore(constant.NewInt(10, types.I32), iPtr)

	// fmt.Println("iPtr.Typ", iPtr.Typ)
	// iPtr.Typ = types.NewPointer(types.Float)
	// fmt.Println("iPtr.Typ2", iPtr.Typ)

	// ir.NewAddrSpaceCast()

	returnBlock := ir.NewBlock("return")

	for _, t := range tokens {
		switch t.Value.Type {
		case token.IntType:
			value, ok := t.Value.True.(int)
			if !ok {
				fmt.Println("ERROR: not able to assert type")
				os.Exit(8)
			}
			fmt.Println("found an int")
			alloc := mainBlock.NewAlloca(types.I32)
			// alloc.SetName(t.Value.Name)
			mainBlock.NewStore(constant.NewInt(types.I32, int64(value)), alloc)

		case token.FloatType:
			value, ok := t.Value.True.(float64)
			if !ok {
				fmt.Println("ERROR: not able to assert type")
				os.Exit(8)
			}
			fmt.Println("found a float")
			alloc := mainBlock.NewAlloca(types.Double)
			// alloc.SetName(t.Value.Name)
			mainBlock.NewStore(constant.NewFloat(types.Double, value), alloc)

		case token.BoolType:
			value, ok := t.Value.True.(bool)
			if !ok {
				fmt.Println("ERROR: not able to assert type", t)
				os.Exit(8)
			}
			fmt.Println("found a bool")
			alloc := mainBlock.NewAlloca(types.I8)
			// alloc.SetName(t.Value.Name)
			mainBlock.NewStore(constant.NewInt(types.I8, boolToInt(value)), alloc)

		case token.StringType:
			value, ok := t.Value.True.(string)
			if !ok {
				fmt.Println("ERROR: not able to assert type", t)
				os.Exit(8)
			}
			fmt.Println("found a string")
			alloc := mainBlock.NewAlloca(types.NewArray(uint64(len(value)), types.I32))
			// alloc.SetName(t.Value.Name)
			mainBlock.NewStore(stringToArray(value), alloc)

		case token.VarType:
			fmt.Println("hmm... not sure if varType should be in here", t)

		case token.ObjectType:
			value, ok := t.Value.True.(map[string]token.Value)
			if !ok {
				fmt.Println("ERROR: not able to assert type", t)
				os.Exit(8)
			}
			fmt.Println("found an object", value)

			var fields []constant.Constant
			var fieldTypes []types.Type
			for _, field := range value {
				// ideally we should call this function recursively
				switch field.Type {
				case token.IntType:
					value, ok := field.True.(int)
					if !ok {
						fmt.Println("ERROR: not able to assert type")
						os.Exit(8)
					}
					fields = append(fields, constant.NewInt(types.I32, int64(value)))
					fieldTypes = append(fieldTypes, types.I32)

				case token.FloatType:
					value, ok := field.True.(float64)
					if !ok {
						fmt.Println("ERROR: not able to assert type")
						os.Exit(8)
					}
					fields = append(fields, constant.NewFloat(types.Double, value))
					fieldTypes = append(fieldTypes, types.Double)

				case token.BoolType:
					value, ok := field.True.(bool)
					if !ok {
						fmt.Println("ERROR: not able to assert type")
						os.Exit(8)
					}
					fields = append(fields, constant.NewInt(types.I8, boolToInt(value)))
					fieldTypes = append(fieldTypes, types.I8)

				case token.StringType:
					value, ok := field.True.(string)
					if !ok {
						fmt.Println("ERROR: not able to assert type", t)
						os.Exit(8)
					}
					fields = append(fields, stringToArray(value))
					fieldTypes = append(fieldTypes, types.NewArray(uint64(len(value)), types.I32))

				case token.VarType:
					fmt.Println("hmm... not sure if varType should be in here", t)

				default:
					fmt.Println("wtf is this token", t)
					continue
				}
			}
			alloc := mainBlock.NewAlloca(types.NewStruct(fieldTypes...))
			// alloc.SetName(t.Value.Name)
			mainBlock.NewStore(constant.NewStruct(fields...), alloc)

		default:
			fmt.Println("ERROR: did not know what to do with token", t)
			fmt.Println("token was of type", t.Value.Type)
			fmt.Println()
		}
	}

	// TODO: will have to do something to figure out where this goes next
	mainBlock.NewBr(returnBlock)

	returnBlock.NewRet(constant.NewInt(types.I32, 0))
	returnBlock.Parent = mainFunc
	mainFunc.Blocks = append(mainFunc.Blocks, returnBlock)

	fmt.Println()
	fmt.Println(m)
	err := ioutil.WriteFile("main.expr.ll", []byte(m.String()), 0644)
	if err != nil {
		fmt.Println("ERROR: Could not write file", err)
		os.Exit(8)
	}
}
