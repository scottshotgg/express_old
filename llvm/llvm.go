package llvm

import (
	"fmt"
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
		fmt.Println(t)

		if t.Value.Type == token.IntType {
			value, ok := t.Value.True.(int)
			if !ok {
				fmt.Println("ERROR: not able to assert type")
				os.Exit(8)
			}

			fmt.Println("found an int")
			mainBlock.NewStore(constant.NewInt(int64(value), types.I32), mainBlock.NewAlloca(types.I32))
		}
	}

	// TODO: will have to do something to figure out where this goes next
	mainBlock.NewBr(returnBlock)

	returnBlock.NewRet(constant.NewInt(0, types.I32))
	mainFunc.AppendBlock(returnBlock)

	fmt.Println()
	fmt.Println(m)
}
