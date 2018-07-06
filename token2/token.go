package token2

type (
	Token interface {
		GetID() int
		GetTokenType() TokenType
		GetExpected() TokenType

		SetID(int)
		SetTokenType(TokenType)
		SetExpected(TokenType)
	}

	Value interface {
		// GetName() string
		GetValueType() ValueType
		// GetActingType() ValueType // TODO: I think this would be an extra method on var?
		GetTrueValue() interface{}
		// String() string
		GetStringValue() string
		// GetAccessModifier() AccessModifierType
	}

	TokenType          string
	ValueType          string
	AccessModifierType string
)

// FIXME: downcase all of these
const (
	Custom       TokenType = "CUSTOM"
	Literal      TokenType = "LITERAL"
	Var          TokenType = "VAR"
	Ident        TokenType = "IDENT"
	Type         TokenType = "TYPE"
	Whitespace   TokenType = "WS"
	Literal      TokenType = "LITERAL"
	Attribute    TokenType = "ATTRIBUTE"
	Keyword      TokenType = "KEYWORD"
	SQL          TokenType = "SQL"
	Comma        TokenType = "COMMA"
	EOS          TokenType = "EOS"
	Separator    TokenType = "SEPARATOR"
	Bang         TokenType = "BANG"
	At           TokenType = "AT"
	Hash         TokenType = "HASH"
	Block        TokenType = "BLOCK"
	Function     TokenType = "FUNCTION"
	Group        TokenType = "GROUP"
	Array        TokenType = "ARRAY"
	Set          TokenType = "SET"
	Assign       TokenType = "ASSIGN"
	Init         TokenType = "INIT"
	PriOp        TokenType = "PRI_OP"
	SecOp        TokenType = "SEC_OP"
	Mult         TokenType = "MULT"
	LBrace       TokenType = "L_BRACE"
	LBracket     TokenType = "L_BRACKET"
	LParen       TokenType = "L_PAREN"
	LThan        TokenType = "L_THAN"
	RBrace       TokenType = "R_BRACE"
	RBracket     TokenType = "R_BRACKET"
	RParen       TokenType = "R_PAREN"
	GThan        TokenType = "G_THAN"
	DQuote       TokenType = "D_QUOTE"
	SQuote       TokenType = "S_QUOTE"
	Pipe         TokenType = "PIPE"
	Ampersand    TokenType = "AMPERSAND"
	Dollar       TokenType = "DOLLARa"
	Underscore   TokenType = "UNDERSCORE"
	QuestionMark TokenType = "QM"
	Accessor     TokenType = "ACCESSOR"
	Increment    TokenType = "INCREMENT"
)

const (
	VarValue      ValueType = "var"
	IntValue      ValueType = "int"
	FloatValue    ValueType = "float"
	StringValue   ValueType = "string"
	BoolValue     ValueType = "bool"
	CharValue     ValueType = "char"
	ObjectValue   ValueType = "object"
	ArrayValue    ValueType = "array"
	SetValue      ValueType = "set"
	IntArrayValue ValueType = "int[]"
)

const (
	PublicAccessModifier  AccessModifierType = "public"
	PrivateAccessModifier AccessModifierType = "private"
)

func New() Token {
	return NewDefaultToken()
}

// func NewValue() Value {
// 	return &DefaultValue{}
// }
