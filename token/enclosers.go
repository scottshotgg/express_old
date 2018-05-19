package token

// EncloserMap holds all valid encloser tokens
var EncloserMap = map[string]Token{
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
