package token

import (
	"fmt"
	"os"
)

// Value ...
type (
	Value struct {
		Type   string
		True   interface{}
		String string
	}

	// Token ...
	Token struct {
		ID       int
		Type     string
		Expected string
		Value    Value
	}
)

// TokenMap ...
var (
	mapArray = []map[string]Token{
		AssignMap,
		EncloserMap,
		KeywordMap,
		OperatorMap,
		SeparatorMap,
		SQLMap,
		TypeMap,
		WhitespaceMap,
	}

	TokenMap = map[string]Token{}
)

// These public consts are to make the entire compiler consistent without having to use
// string literals. These may be changed to ints in the future
const (
	Var          = "VAR"
	Ident        = "IDENT"
	Type         = "TYPE"
	Whitespace   = "WS"
	Literal      = "LITERAL"
	Attribute    = "ATTRIBUTE"
	Keyword      = "KEYWORD"
	SQL          = "SQL"
	Comma        = "COMMA"
	EOS          = "EOS"
	Separator    = "SEPARATOR"
	Bang         = "BANG"
	At           = "AT"
	Hash         = "HASH"
	Block        = "BLOCK"
	Function     = "FUNCTION"
	Group        = "GROUP"
	Array        = "ARRAY"
	Set          = "SET"
	Assign       = "ASSIGN"
	Init         = "INIT"
	PriOp        = "PRI_OP"
	SecOp        = "SEC_OP"
	Mult         = "MULT"
	LBrace       = "L_BRACE"
	LBracket     = "L_BRACKET"
	LParen       = "L_PAREN"
	LThan        = "L_Than"
	RBrace       = "R_BRACE"
	RBracket     = "R_BRACKET"
	RParen       = "R_PAREN"
	GThan        = "G_THAN"
	DQuote       = "D_QUOTE"
	SQuote       = "S_QUOTE"
	Pipe         = "PIPE"
	Ampersand    = "AMPERSAND"
	DDBY         = "DDBY"
	Underscore   = "UNDERSCORE"
	QuestionMark = "QM"
	Accessor     = "Accessor"

	VarType    = "var"
	IntType    = "int"
	FloatType  = "float"
	StringType = "string"
	BoolType   = "bool"
	CharType   = "char"
)

func init() {
	// Load all maps in
	for _, tMap := range mapArray {
		for key, value := range tMap {
			TokenMap[key] = value
		}
	}

	// Load the lexeme map and ensure that all tokens are defined
	for _, lexeme := range Lexemes {
		if lexToken, ok := TokenMap[lexeme]; !ok {
			fmt.Println("ERROR: Lexeme not found in TokenMap: ", lexeme)
			os.Exit(7)
		} else {
			LexemeMap[lexeme] = lexToken
		}
	}
}
