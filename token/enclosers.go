package token

// EncloserMap holds all valid encloser tokens
var EncloserMap = map[string]Token{
	"(": Token{
		Type: "L_PAREN",
		Value: Value{
			Type:   "op_3",
			String: "(",
		},
	},
	")": Token{
		Type: "R_PAREN",
		Value: Value{
			Type:   "op_3",
			String: ")",
		},
	},

	"{": Token{
		Type: "L_BRACE",
		Value: Value{
			Type:   "op_3",
			String: "{",
		},
	},
	"}": Token{
		Type: "R_BRACE",
		Value: Value{
			Type:   "op_3",
			String: "}",
		},
	},

	"[": Token{
		Type: "L_BRACKET",
		Value: Value{
			Type:   "op_3",
			String: "[",
		},
	},
	"]": Token{
		Type: "R_BRACKET",
		Value: Value{
			Type:   "lthan",
			String: "]",
		},
	},

	"<": Token{
		Type: "L_THAN",
		Value: Value{
			Type:   "lthan",
			String: "<",
		},
	},
	">": Token{
		Type: "G_THAN",
		Value: Value{
			Type:   "rthan",
			String: ">",
		},
	},

	"`": Token{
		Type: "GRAVE",
		Value: Value{
			Type:   "op_3",
			String: "`",
		},
	},
	"~": Token{
		Type: "TILDE",
		Value: Value{
			Type:   "op_3",
			String: "~",
		},
	},
	"'": Token{
		Type: SQuote,
		Value: Value{
			Type:   "squote",
			String: "'",
		},
	},
	"\"": Token{
		Type: DQuote,
		Value: Value{
			Type:   "dquote",
			String: "\"",
		},
	},
	"@": Token{
		Type: At,
		Value: Value{
			Type:   "op_3",
			String: "@",
		},
	},
}
