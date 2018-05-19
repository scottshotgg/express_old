package token

// SeparatorMap holds all defined statement separators
var SeparatorMap = map[string]Token{
	",": Token{
		Type:     "SEPARATOR",
		Expected: "EXPR",
		Value: Value{
			Type:   "comma",
			String: ",",
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
	" ": Token{
		Type: "WS",
		Value: Value{
			Type:   "space",
			String: " ",
		},
	},
}
