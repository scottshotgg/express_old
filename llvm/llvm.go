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

// Translate ...
func Translate(tokens []token.Token) {
	fmt.Println("tokens", tokens)

	// Create a new LLVM IR module.
	m := ir.NewModule()

	// TODO: FIXME: this needs to be based on the flag or w/e to turn off dynamic types
	// m.NewType("var", types.NewStruct([]types.Type{
	// 	types.NewInt(8),
	// 	types.NewPointer(types.NewInt(32)),
	// }...))

	mainFunc := m.NewFunction("main", types.I32)
	mainBlock := mainFunc.NewBlock("main")

	// iPtr := mainBlock.NewAlloca(types.NewPointer(types.I32))
	// mainBlock.NewStore(constant.NewInt(10, types.I32), iPtr)

	// fmt.Println("iPtr.Typ", iPtr.Typ)
	// iPtr.Typ = types.NewPointer(types.Float)
	// fmt.Println("iPtr.Typ2", iPtr.Typ)

	// ir.NewAddrSpaceCast()

	returnBlock := ir.NewBlock("")

	for _, t := range tokens {
		// fmt.Println(t)

		switch t.Value.Type {
		case token.IntType:
			value, ok := t.Value.True.(int)
			if !ok {
				fmt.Println("ERROR: not able to assert type")
				os.Exit(8)
			}
			fmt.Println("found an int")
			mainBlock.NewStore(constant.NewInt(int64(value), types.I32), mainBlock.NewAlloca(types.I32))

		case token.FloatType:
			value, ok := t.Value.True.(float64)
			if !ok {
				fmt.Println("ERROR: not able to assert type")
				os.Exit(8)
			}
			fmt.Println("found a float")
			mainBlock.NewStore(constant.NewFloat(value, types.Double), mainBlock.NewAlloca(types.Double))

		case token.BoolType:
			value, ok := t.Value.True.(bool)
			if !ok {
				fmt.Println("ERROR: not able to assert type", t)
				os.Exit(8)
			}
			fmt.Println("found a bool")

			boolValue := int64(0)
			if value {
				boolValue = 1
			}
			mainBlock.NewStore(constant.NewInt(boolValue, types.I8), mainBlock.NewAlloca(types.I8))

		// store i8* getelementptr inbounds ([12 x i8], [12 x i8]* @.str, i32 0, i32 0), i8** %2, align 8
		case token.StringType:
			value, ok := t.Value.True.(string)
			if !ok {
				fmt.Println("ERROR: not able to assert type", t)
				os.Exit(8)
			}
			fmt.Println("found a string")

			var vArray []constant.Constant
			for _, char := range value {
				vArray = append(vArray, constant.NewInt(int64(char), types.I32))
			}
			mainBlock.NewStore(constant.NewArray(vArray...), mainBlock.NewAlloca(types.NewArray(types.I32, int64(len(value)))))

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
					fields = append(fields, constant.NewInt(int64(value), types.I32))
					fieldTypes = append(fieldTypes, types.I32)

				case token.FloatType:
					value, ok := field.True.(float64)
					if !ok {
						fmt.Println("ERROR: not able to assert type")
						os.Exit(8)
					}
					fields = append(fields, constant.NewFloat(value, types.Double))
					fieldTypes = append(fieldTypes, types.Double)

				case token.BoolType:
					value, ok := field.True.(bool)
					if !ok {
						fmt.Println("ERROR: not able to assert type")
						os.Exit(8)
					}
					boolValue := int64(0)
					if value {
						boolValue = 1
					}
					fields = append(fields, constant.NewInt(boolValue, types.I8))
					fieldTypes = append(fieldTypes, types.I8)

				case token.StringType:
					value, ok := field.True.(string)
					if !ok {
						fmt.Println("ERROR: not able to assert type", t)
						os.Exit(8)
					}
					vArray := []constant.Constant{}
					for _, char := range value {
						vArray = append(vArray, constant.NewInt(int64(char), types.I32))
					}
					fields = append(fields, constant.NewArray(vArray...))
					fieldTypes = append(fieldTypes, types.NewArray(types.I32, int64(len(value))))
					//constant.NewArray(vArray...), mainBlock.NewAlloca()

				case token.VarType:

				default:
					fmt.Println("wtf is this token", t)
					continue
				}
			}

			mainBlock.NewStore(constant.NewStruct(fields...), mainBlock.NewAlloca(types.NewStruct(fieldTypes...)))

		default:
			fmt.Println("ERROR: did not know what to do with token", t)
			fmt.Println("token was of type", t.Value.Type)
			fmt.Println()
		}
	}

	// TODO: will have to do something to figure out where this goes next
	mainBlock.NewBr(returnBlock)

	returnBlock.NewRet(constant.NewInt(0, types.I32))
	mainFunc.AppendBlock(returnBlock)

	fmt.Println()
	fmt.Println(m)
	err := ioutil.WriteFile("main.expr.ll", []byte(m.String()), 0644)
	if err != nil {
		fmt.Println("ERROR: Could not write file", err)
		os.Exit(8)
	}
}
