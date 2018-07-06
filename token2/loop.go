package token2

// FIXME: come back to this
// We should change this to be a boolean expression
type Check struct {
	left  Value
	op    Value
	right Value
}

type Loop struct {
	DefaultToken
	start int
	end   int
	step  int
	check Check
	// TODO: this should change to an identifier
	valueVar Value
}

func (loop *Loop) TokenID() int {
	return loop.id
}

func (loop *Loop) TokenType() TokenType {
	return loop.tokenType
}

func (loop *Loop) Expected() TokenType {
	return loop.expected
}

func (loop *Loop) Value() Value {
	return loop.value
}

func NewLoop() Token {
	return &Loop{}
}
