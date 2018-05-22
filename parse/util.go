package parse

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/scottshotgg/Express/token"
)

// CollectTokens appends an array of tokens passed in to the EndTokens attribute of Meta
func (m *Meta) CollectTokens(tokens []token.Token) {
	m.LastCollectedToken = tokens[len(tokens)-1]
	m.EndTokens = append(m.EndTokens, tokens...)
}

// CollectToken appends a single token to the EndTokens attribute of Meta
func (m *Meta) CollectToken(token token.Token) {
	m.LastCollectedToken = token
	m.EndTokens = append(m.EndTokens, token)
}

// RemoveLastCollectedToken removes the last token put into EndTokens
func (m *Meta) RemoveLastCollectedToken() {
	m.LastCollectedToken = m.EndTokens[len(m.EndTokens)-1]
	m.EndTokens = m.EndTokens[:len(m.EndTokens)-1]
}

// PopLastCollectedToken removes the last token put into EndTokens
func (m *Meta) PopLastCollectedToken() token.Token {
	m.LastCollectedToken = m.EndTokens[len(m.EndTokens)-2]
	m.EndTokens = m.EndTokens[:len(m.EndTokens)-1]

	return m.EndTokens[len(m.EndTokens)-1]
}

// CollectCurrentToken appends the token held in the CurrentToken attribute to the EndTokens array
func (m *Meta) CollectCurrentToken() {
	m.CollectToken(m.CurrentToken)
}

// CollectLastToken appends the token held in the LastToken attribute to the EndTokens array
func (m *Meta) CollectLastToken() {
	m.CollectToken(m.LastToken)
}

// GetLastToken returns the LastToken attribute
func (m *Meta) GetLastToken() token.Token {
	return m.LastToken
}

// PeekLastCollectedToken returns the last token appended to the EndTokens array
func (m *Meta) PeekLastCollectedToken() token.Token {
	return m.LastCollectedToken
}

// GetCurrentToken returns the CurrentToken attribute
func (m *Meta) GetCurrentToken() token.Token {
	return m.CurrentToken
}

// PeekTokenAtIndex returns the token at that ParseIndex if valid
func (m *Meta) PeekTokenAtIndex(index int) (token.Token, error) {
	if index > -1 && index < m.Length {
		return m.Tokens[index], nil
	}

	return token.Token{}, errors.New("Current parseIndex outside of token range")
}

// Shift operates the parses like a 3-bit (3 token) SIPO shift register consuming the tokens until the end of the line
func (m *Meta) Shift() {
	m.LastToken = m.CurrentToken
	m.CurrentToken = m.NextToken

	for {
		if m.ParseIndex < m.Length {
			if m.Tokens[m.ParseIndex].Type == token.Whitespace {
				m.ParseIndex++
				continue
			}

			m.NextToken = m.Tokens[m.ParseIndex]
			m.ParseIndex++
			return
		}

		m.NextToken = token.Token{}
		return
	}
}

// ShiftWithWS operates the parses like a 3-bit (3 token) SIPO shift register consuming the tokens until the end of the line
func (m *Meta) ShiftWithWS() {
	m.LastToken = m.CurrentToken
	m.CurrentToken = m.Tokens[m.ParseIndex]

	for {
		if m.ParseIndex+1 < m.Length {
			m.ParseIndex++

			m.NextToken = m.Tokens[m.ParseIndex]
			return
		}

		m.NextToken = token.Token{}
		return
	}
}

// TokenToString marshals a token into it's JSON representation
func TokenToString(t token.Token) string {
	jsonToken, err := json.Marshal(t)
	if err != nil {
		return err.Error()
	}

	return string(jsonToken)
}
