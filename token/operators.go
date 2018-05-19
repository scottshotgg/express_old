package token

// OperatorMap holds all defined operator tokens
var OperatorMap = map[string]Token{
	"+": Token{
		Type: SecOp,
		Value: Value{
			Type:   "add",
			String: "+",
		},
	},
	"-": Token{
		Type: SecOp,
		Value: Value{
			Type:   "sub",
			String: "-",
		},
	},
	"*": Token{
		Type: PriOp,
		Value: Value{
			Type:   "mult",
			String: "*",
		},
	},
	"/": Token{
		Type: PriOp,
		Value: Value{
			Type:   "div",
			String: "/",
		},
	},
	"%": Token{
		Type: PriOp,
		Value: Value{
			Type:   "mod",
			String: "%",
		},
	},
	"^": Token{
		Type: PriOp,
		Value: Value{
			Type:   "exp",
			String: "^",
		},
	},
	"!": Token{
		Type: Bang,
		Value: Value{
			Type:   "bang",
			String: "!",
		},
	},
	"?": Token{
		Type: QuestionMark,
		Value: Value{
			Type:   "qm",
			String: "!",
		},
	},
	"_": Token{
		Type: Underscore,
		Value: Value{
			Type:   "underscore",
			String: "_",
		},
	},
	// FIXME: DOLLA DOLLA BILLS YALL: define this
	"$": Token{
		Type: DDBY,
		Value: Value{
			Type:   "ddby",
			String: "$",
		},
	},
	"&": Token{
		Type: Ampersand,
		Value: Value{
			Type:   "op_3",
			String: "&",
		},
	},
	"|": Token{
		Type: Pipe,
		Value: Value{
			Type:   "op_3",
			String: "|",
		},
	},
	"#": Token{
		Type: Hash,
		Value: Value{
			Type:   "op_3",
			String: "#",
		},
	},
	".": Token{
		Type: Accessor,
		Value: Value{
			Type:   "period",
			String: ".",
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
}
