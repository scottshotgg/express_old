package token2

type DefaultToken struct {
	id        int
	tokenType TokenType
	expected  TokenType
}

func (dt *DefaultToken) GetID() int {
	return dt.id
}

func (dt *DefaultToken) GetTokenType() TokenType {
	return dt.tokenType
}

func (dt *DefaultToken) GetExpected() TokenType {
	return dt.expected
}

func (dt *DefaultToken) SetID(id int) {
	dt.id = id
}

func (dt *DefaultToken) SetTokenType(tokenType TokenType) {
	dt.tokenType = tokenType
}

func (dt *DefaultToken) SetExpected(expected TokenType) {
	dt.expected = expected
}
