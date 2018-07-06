// This is a simple stack package used _specifically_ for the parser

package parse

import (
	"errors"
	"fmt"
	"sync"
)

type Stack struct {
	length int
	stack  []interface{}
	top    *interface{}
	sync.Mutex
}

var (
	ErrEmptyStack = errors.New("Stack is empty")
)

func NewStack() *Stack {
	return &Stack{
		stack: []interface{}{},
	}
}

func (s *Stack) Length() int {
	return s.length
}

func (s *Stack) Peek() (top interface{}, err error) {
	s.Lock()
	defer s.Unlock()

	err = ErrEmptyStack
	if s.top != nil {
		top = *s.top
		err = nil
	}
	return
}

func (s *Stack) Pop() (pop interface{}, err error) {
	s.Lock()
	defer s.Unlock()

	if s.length < 1 {
		err = ErrEmptyStack
		return
	}

	pop = *s.top
	s.length--
	s.stack = s.stack[:s.length]

	s.top = nil
	if s.length > 0 {
		s.top = &s.stack[s.length-1]
	}

	return
}

func (s *Stack) Push(push interface{}) {
	s.Lock()
	defer s.Unlock()

	s.stack = append(s.stack, push)
	s.top = &s.stack[s.length]
	fmt.Println("s.top push", s.top)
	s.length++
}
