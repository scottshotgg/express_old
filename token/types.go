package token

import "reflect"

// TypeMap holds all defined type tokens
var TypeMap = map[string]Token{
	"var": Token{
		Type:     "TYPE",
		Expected: "IDENT",
		Value: Value{
			Type:   "var",
			String: "var",
		},
	},
	"int": Token{
		Type: "TYPE",
		// TODO: we should expect a token with this "Value.Type"
		Expected: "IDENT",
		Value: Value{
			Type:   "var",
			True:   reflect.Int,
			String: "int",
		},
	},
	"float": Token{
		Type:     "TYPE",
		Expected: "IDENT",
		Value: Value{
			Type:   "var",
			True:   reflect.Float64,
			String: "float",
		},
	},
	"char": Token{
		Type:     "TYPE",
		Expected: "IDENT",
		Value: Value{
			Type:   "var",
			True:   reflect.Int,
			String: "char",
		},
	},
	"string": Token{
		Type:     "TYPE",
		Expected: "IDENT",
		Value: Value{
			Type:   "var",
			True:   reflect.String,
			String: "string",
		},
	},
	"bool": Token{
		Type:     "TYPE",
		Expected: "IDENT",
		Value: Value{
			Type:   "var",
			True:   reflect.Bool,
			String: "bool",
		},
	},
}
