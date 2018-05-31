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
	// Create a new LLVM IR module.
	m := ir.NewModule()

	// Create a function definition and append it to the module.
	//
	//    int rand(void) { ... }
	mainFunc := m.NewFunction("main", types.I32)
	mainBlock := mainFunc.NewBlock("main")

	returnBlock := ir.NewBlock("")

	fmt.Println("hi", tokens)

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

		case token.BoolType:
			value, ok := t.Value.True.(bool)
			if !ok {
				fmt.Println("ERROR: not able to assert type", t)
				os.Exit(8)
			}
			fmt.Println("found a bool")

			boolValue := 0
			if value {
				boolValue = 1
			}

			mainBlock.NewStore(constant.NewInt(int64(boolValue), types.I8), mainBlock.NewAlloca(types.I8))

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
