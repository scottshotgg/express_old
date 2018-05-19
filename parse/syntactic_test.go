package parse

import (
	"testing"

	"github.com/scottshotgg/Express/token"
)

func TestShiftWithWS(t *testing.T) {
	tokens := []token.Token{
		token.TokenMap[" "],
	}

	m := Meta{
		Tokens: tokens,
		Length: len(tokens),
	}

	m.ShiftWithWS()
	if m.CurrentToken != token.TokenMap[" "] {
		t.Fatal("Expected space, got", m.CurrentToken)
	}
}

func TestShift(t *testing.T) {
	tokens := []token.Token{
		token.TokenMap["int"],
		token.TokenMap[" "],
	}

	m := Meta{
		Tokens: tokens,
		Length: len(tokens),
	}

	m.Shift()
	m.Shift()

	if m.CurrentToken != token.TokenMap["int"] {
		t.Fatal("Expected int, got", m.CurrentToken)
	}

	m.Shift()

	if m.CurrentToken != (token.Token{}) {
		t.Fatal("Expected int, got", m.CurrentToken)
	}
}
