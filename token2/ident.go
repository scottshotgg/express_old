package token2

type IdentifierToken struct {
	LiteralToken
	name           string
	accessModifier AccessModifierType
}

func (i *IdentifierToken) GetName() string {
	return i.name
}

func (i *IdentifierToken) GetAccessModifier() AccessModifierType {
	return i.accessModifier
}

// func NewIdent() {
// 	i := NewInt()
// 	return &IdentifierToken{
// 		NewLiteral(),
// 	}
// }

// func NewIdentFromBool() {
// 	i := NewBool()
// }

// func NewIdentFromFloat() {
// 	i := NewFloat()
// }

// func NewIdentFromString() {
// 	i := NewString()
// }
