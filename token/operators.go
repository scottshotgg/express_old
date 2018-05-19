package token

// OperatorMap holds all defined operator tokens
var OperatorMap = map[string]Token{
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
	".": Token{
		Type:     "ACCESSOR",
		Expected: "EXPR",
		Value: Value{
			Type:   "period",
			String: ".",
		},
	},
}
