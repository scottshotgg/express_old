package token

import "reflect"

// Value ...
type Value struct {
	Type   string
	True   interface{}
	String string
}

// Token ...
type Token struct {
	ID       int
	Type     string
	Expected string
	Value    Value
}

// TODO: make a map of all enclosers, map[string]Token

// TokenMap ...
var TokenMap = map[string]Token{
	// TYPES
	"let": Token{
		Type:     "TYPE",
		Expected: "IDENT",
		Value: Value{
			Type:   "var",
			String: "let",
		},
	},
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

	// ASSIGN
	"=": Token{
		Type:     "ASSIGN",
		Expected: "EXPR",
		Value: Value{
			Type:   "assign",
			String: "=",
		},
	},
	":": Token{
		Type:     "ASSIGN",
		Expected: "EXPR",
		Value: Value{
			Type:   "set_assign",
			String: ":",
		},
	},
	":=": Token{
		Type:     "ASSIGN",
		Expected: "EXPR",
		Value: Value{
			Type:   "init_assign",
			String: ":=",
		},
	},

	// OPERANDS
	"+": Token{
		Type:     "ADD",
		Expected: "EXPR",
		Value: Value{
			Type:   "op_4",
			String: "+",
		},
	},
	"-": Token{
		Type:     "SUB",
		Expected: "EXPR",
		Value: Value{
			Type:   "op_4",
			String: "-",
		},
	},
	"*": Token{
		Type:     "MULT",
		Expected: "EXPR",
		Value: Value{
			Type:   "op_3",
			String: "*",
		},
	},
	"/": Token{
		Type:     "DIV",
		Expected: "EXPR",
		Value: Value{
			Type:   "op_3",
			String: "/",
		},
	},
	".+": Token{
		Type:     "VEC_ADD",
		Expected: "EXPR",
		Value: Value{
			Type:   "op_3",
			String: ".+",
		},
	},
	".-": Token{
		Type:     "VEC_SUB",
		Expected: "EXPR",
		Value: Value{
			Type:   "op_4",
			String: ".-",
		},
	},
	".*": Token{
		Type:     "VEC_MULT",
		Expected: "EXPR",
		Value: Value{
			Type:   "op_3",
			String: ".*",
		},
	},
	"./": Token{
		Type:     "VEC_DIV",
		Expected: "EXPR",
		Value: Value{
			Type:   "op_3",
			String: "./",
		},
	},

	// LITERALS
	"true": Token{
		Type: "LITERAL",
		Value: Value{
			Type:   "bool",
			True:   true,
			String: "true",
		},
	},
	"false": Token{
		Type: "LITERAL",
		Value: Value{
			Type:   "bool",
			True:   false,
			String: "false",
		},
	},
}
