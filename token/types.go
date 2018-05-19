package token

// TypeMap holds all defined type tokens
var TypeMap = map[string]Token{
	"var": Token{
		Type: Type,
		Value: Value{
			Type:   "var",
			String: "var",
		},
	},
	"int": Token{
		Type: Type,
		Value: Value{
			Type:   "int",
			String: "int",
		},
	},
	"float": Token{
		Type: Type,
		Value: Value{
			Type:   "float",
			String: "float",
		},
	},
	"char": Token{
		Type: Type,
		Value: Value{
			Type:   "char",
			String: "char",
		},
	},
	"string": Token{
		Type: Type,
		Value: Value{
			Type:   "string",
			String: "string",
		},
	},
	"bool": Token{
		Type: Type,
		Value: Value{
			Type:   "bool",
			String: "bool",
		},
	},
}
