package parse_test

import (
	"fmt"
	"testing"

	"github.com/scottshotgg/ExpressRedo/parse"
)

var (
	stack     *parse.Stack
	testScope = parse.Scope{
		"i": parse.NewVariable("i", 6, parse.INT),
		// "a": parse.NewVariable("a", "hey its me", parse.STRING),
	}
)

func TestNewStack(t *testing.T) {
	stack = parse.NewStack()
	fmt.Printf("Stack: %+v\n", stack)
	fmt.Println()
}

func TestPush(t *testing.T) {
	TestNewStack(t)

	for k, v := range []string{"a", "b", "c"} {
		stack.Push(parse.Scope{
			v: parse.NewVariable(v, k, parse.INT),
		})
	}
	// stack.Push(testScope)
	// stack.Push(testScope)
	fmt.Printf("Stack: %+v\n", stack)
	fmt.Println()
}

func TestPop(t *testing.T) {
	TestPush(t)

	fmt.Printf("Stack: %+v\n", stack)

	pop, err := stack.Pop()
	if err != nil {
		fmt.Println("failed")
		t.Fail()
		return
	}

	fmt.Printf("Pop: %+v\n", pop)
	fmt.Printf("Stack: %+v\n", stack)
	fmt.Println()
}

func TestPeek(t *testing.T) {
	TestPush(t)

	peek, err := stack.Peek()
	if err != nil {
		t.Fail()
		return
	}

	fmt.Printf("Peek: %+v\n", peek)
	fmt.Printf("Stack: %+v\n", stack)
	fmt.Println()
}
