package token

import "reflect"

// TypeMap holds all defined type tokens
var TypeMap = map[string]Token{
	"var": Token{
		Type: "TYPE",
		Value: Value{
			Type:   "var",
			String: "var",
		},
	},
	"int": Token{
		Type: "TYPE",
		Value: Value{
			Type:   "int",
			True:   reflect.Int,
			String: "int",
		},
	},
	"float": Token{
		Type: "TYPE",
		Value: Value{
			Type:   "float",
			True:   reflect.Float64,
			String: "float",
		},
	},
	"char": Token{
		Type: "TYPE",
		Value: Value{
			Type:   "char",
			True:   reflect.Int,
			String: "char",
		},
	},
	"string": Token{
		Type: "TYPE",
		Value: Value{
			Type:   "string",
			True:   reflect.String,
			String: "string",
		},
	},
	"bool": Token{
		Type: "TYPE",
		Value: Value{
			Type:   "bool",
			True:   reflect.Bool,
			String: "bool",
		},
	},
}
