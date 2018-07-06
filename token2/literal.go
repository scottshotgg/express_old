package token2

import "strconv"

type LiteralToken struct {
	DefaultToken
	value LiteralValue
}

type LiteralValue struct {
	valueType   ValueType
	actingType  ValueType
	trueValue   interface{}
	stringValue string
}

func (dv *LiteralValue) GetValueType() ValueType {
	return dv.valueType
}

func (dv *LiteralValue) GetActingType() ValueType {
	return dv.actingType
}

func (dv *LiteralValue) GetTrueValue() interface{} {
	return dv
}

func (dv *LiteralValue) GetStringValue() string {
	return dv.stringValue
}

func (l *LiteralToken) SetValue(value LiteralValue) {
	l.value = value
}

func NewLiteral() *LiteralToken {
	return &LiteralToken{
		DefaultToken: DefaultToken{
			id:        0,
			tokenType: Literal,
		},
	}
}

func NewLiteralFromValue(value LiteralValue) *LiteralToken {
	l := NewLiteral()
	l.SetValue(value)

	return l
}

func NewInt() *LiteralToken {
	return NewIntFromInt(0)
}

func NewIntFromInt(value int) *LiteralToken {
	return NewLiteralFromValue(LiteralValue{
		value:       IntValue,
		trueValue:   value,
		stringValue: strconv.Itoa(value),
	})
}

func NewBool() *LiteralToken {
	return NewBoolFromBool(false)
}

func NewBoolFromBool(value bool) *LiteralToken {
	return NewLiteralFromValue(LiteralValue{
		value:       BoolValue,
		trueValue:   value,
		stringValue: strconv.FormatBool(value),
	})
}

func NewFloat() *LiteralToken {
	return NewFloatFromFloat(0.0)
}

func NewFloatFromFloat(value float64) *LiteralToken {
	return NewLiteralFromValue(LiteralValue{
		value:       FloatValue,
		trueValue:   value,
		stringValue: strconv.FormatFloat(value, 'f', -1, 64),
	})
}

func NewString() *LiteralToken {
	return NewStringFromString("")
}

func NewStringFromString(value string) *LiteralToken {
	return NewLiteralFromValue(LiteralValue{
		valueType:   StringValue,
		trueValue:   value,
		stringValue: value,
	})
}
