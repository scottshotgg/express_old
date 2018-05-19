package token

// OperatorMap holds all defined operator tokens
var OperatorMap = map[string]Token{
	"+": Token{
		Type: "SEC_OP",
		Value: Value{
			Type:   "add",
			String: "+",
		},
	},
	"-": Token{
		Type: "SEC_OP",
		Value: Value{
			Type:   "sub",
			String: "-",
		},
	},
	"*": Token{
		Type: "PRI_OP",
		Value: Value{
			Type:   "mult",
			String: "*",
		},
	},
	"/": Token{
		Type: "PRI_OP",
		Value: Value{
			Type:   "div",
			String: "/",
		},
	},
	"%": Token{
		Type: "PRI_OP",
		Value: Value{
			Type:   "mod",
			String: "%",
		},
	},
	"^": Token{
		Type: "PRI_OP",
		Value: Value{
			Type:   "exp",
			String: "^",
		},
	},
	"!": Token{
		Type: "BANG",
		Value: Value{
			Type:   "bang",
			String: "!",
		},
	},
	"?": Token{
		Type: "QM",
		Value: Value{
			Type:   "qm",
			String: "!",
		},
	},
	"_": Token{
		Type: "UNDERSCORE",
		Value: Value{
			Type:   "underscore",
			String: "_",
		},
	},
	// FIXME: DOLLA DOLLA BILLS YALL: define this
	"$": Token{
		Type: "DDBY",
		Value: Value{
			Type:   "ddby",
			String: "$",
		},
	},
	"&": Token{
		Type: "AMP",
		Value: Value{
			Type:   "op_3",
			String: "&",
		},
	},
	"|": Token{
		Type: "PIPE",
		Value: Value{
			Type:   "op_3",
			String: "|",
		},
	},
	"#": Token{
		Type: "HASH",
		Value: Value{
			Type:   "op_3",
			String: "#",
		},
	},

	// TODO: add the templated operators ability to the parser and remove the tokens completely
	// VECTOR OPERANDS
	".+": Token{
		Type: "VEC_ADD",
		Value: Value{
			Type:   "op_3",
			String: ".+",
		},
	},
	".-": Token{
		Type: "VEC_SUB",
		Value: Value{
			Type:   "op_4",
			String: ".-",
		},
	},
	".*": Token{
		Type: "VEC_MULT",
		Value: Value{
			Type:   "op_3",
			String: ".*",
		},
	},
	"./": Token{
		Type: "VEC_DIV",
		Value: Value{
			Type:   "op_3",
			String: "./",
		},
	},
	".": Token{
		Type: "ACCESSOR",
		Value: Value{
			Type:   "period",
			String: ".",
		},
	},
}
