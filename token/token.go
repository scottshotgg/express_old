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
// TODO: Verify ALL 'Expected' and 'Type' attributes
// TODO: might look into doing some letters from other natural languages (ie, `forwardtick`)

// TokenMap ...
var TokenMap = map[string]Token{
	"TYPE": Token{
		Type:     "TYPE",
		Expected: "IDENT",
		Value: Value{
			Type:   "var",
			String: "let",
		},
	},

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
		Type:     "SET",
		Expected: "EXPR",
		Value: Value{
			Type:   "set_assign",
			String: ":",
		},
	},
	":=": Token{
		Type:     "INIT",
		Expected: "EXPR",
		Value: Value{
			Type:   "init_assign",
			String: ":=",
		},
	},

	// OPERANDS
	"+": Token{
		Type:     "SEC_OP",
		Expected: "EXPR",
		Value: Value{
			Type:   "add",
			String: "+",
		},
	},
	"-": Token{
		Type:     "SEC_OP",
		Expected: "EXPR",
		Value: Value{
			Type:   "sub",
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
	"%": Token{
		Type:     "MOD",
		Expected: "EXPR",
		Value: Value{
			Type:   "op_3",
			String: "%",
		},
	},
	"^": Token{
		Type:     "EXP",
		Expected: "EXPR",
		Value: Value{
			Type:   "op_3",
			String: "^",
		},
	},
	"!": Token{
		Type:     "NOT",
		Expected: "EXPR",
		Value: Value{
			Type:   "op_3",
			String: "!",
		},
	},
	// TODO: look up what this is called and shit
	"?": Token{
		Type:     "NOT",
		Expected: "EXPR",
		Value: Value{
			Type:   "op_3",
			String: "!",
		},
	},
	"_": Token{
		Type:     "UNDERSCORE",
		Expected: "EXPR",
		Value: Value{
			Type:   "op_3",
			String: "_",
		},
	},
	// FIXME: DOLLA DOLLA BILLS YALL: define this
	"$": Token{
		Type:     "DDBY",
		Expected: "EXPR",
		Value: Value{
			Type:   "op_3",
			String: "$",
		},
	},
	// FIXME: AMP/AND - define this
	"&": Token{
		Type:     "AMP",
		Expected: "EXPR",
		Value: Value{
			Type:   "op_3",
			String: "&",
		},
	},
	// FIXME: PIPE/OR - define what a pipe is/does
	"|": Token{
		Type:     "PIPE",
		Expected: "EXPR",
		Value: Value{
			Type:   "op_3",
			String: "|",
		},
	},
	// FIXME: still need to decide what to do with this one
	"#": Token{
		Type:     "HASH",
		Expected: "EXPR",
		Value: Value{
			Type:   "op_3",
			String: "#",
		},
	},

	// TODO: we should just make the parser look for "." and then "+"
	// VECTOR OPERANDS
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

	// SEPARATORS
	",": Token{
		Type:     "SEPARATOR",
		Expected: "EXPR",
		Value: Value{
			Type:   "comma",
			String: ",",
		},
	},
	" ": Token{
		Type: "WS",
		Value: Value{
			Type:   "space",
			String: " ",
		},
	},
	"\t": Token{
		Type: "WS",
		Value: Value{
			Type:   "tab",
			String: "\t",
		},
	},
	"\n": Token{
		Type: "WS",
		Value: Value{
			Type:   "newline",
			String: "\n",
		},
	},
	";": Token{
		Type: "EOS",
		// Expected: "STATEMENT",
		Value: Value{
			Type:   "semicolon",
			String: ";",
		},
	},

	// ENCLOSERS
	"(": Token{
		Type:     "L_PAREN",
		Expected: "EXPR",
		Value: Value{
			Type:   "op_3", // TODO: check all these
			String: "(",
		},
	},
	")": Token{
		Type:     "R_PAREN",
		Expected: "EXPR",
		Value: Value{
			Type:   "op_3", // TODO: check all these
			String: ")",
		},
	},
	"{": Token{
		Type:     "L_BRACE",
		Expected: "EXPR",
		Value: Value{
			Type:   "op_3", // TODO: check all these
			String: "{",
		},
	},
	"}": Token{
		Type:     "R_BRACE",
		Expected: "EXPR",
		Value: Value{
			Type:   "op_3", // TODO: check all these
			String: "}",
		},
	},
	"[": Token{
		Type:     "L_BRACKET",
		Expected: "EXPR",
		Value: Value{
			Type:   "op_3", // TODO: check all these
			String: "[",
		},
	},
	"]": Token{
		Type:     "R_BRACKET",
		Expected: "EXPR",
		Value: Value{
			Type:   "op_3", // TODO: check all these
			String: "]",
		},
	},
	"`": Token{
		Type:     "GRAVE",
		Expected: "EXPR",
		Value: Value{
			Type:   "op_3",
			String: "`",
		},
	},
	"~": Token{
		Type:     "TILDE",
		Expected: "EXPR",
		Value: Value{
			Type:   "op_3",
			String: "~",
		},
	},
	"'": Token{
		Type:     "S_QUOTE",
		Expected: "EXPR",
		Value: Value{
			Type:   "op_3",
			String: "'",
		},
	},
	"\"": Token{
		Type:     "D_QUOTE",
		Expected: "EXPR",
		Value: Value{
			Type:   "op_3",
			String: "\"",
		},
	},
	"<": Token{
		Type:     "L_THAN",
		Expected: "EXPR",
		Value: Value{
			Type:   "op_3",
			String: "<",
		},
	},
	">": Token{
		Type:     "G_THAN",
		Expected: "EXPR",
		Value: Value{
			Type:   "op_3",
			String: ">",
		},
	},
	// TODO: not sure if this will be an encloser or not
	"@": Token{
		Type:     "AT",
		Expected: "EXPR",
		Value: Value{
			Type:   "op_3",
			String: "@",
		},
	},
}
