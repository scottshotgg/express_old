package token

// WhitespaceMap holds all defined Whitespace tokens
var WhitespaceMap = map[string]Token{
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
}
