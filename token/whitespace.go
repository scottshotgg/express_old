package token

// WhitespaceMap holds all defined Whitespace tokens
var WhitespaceMap = map[string]Token{
	" ": Token{
		Type: Whitespace,
		Value: Value{
			Type:   "space",
			String: " ",
		},
	},
	"\t": Token{
		Type: Whitespace,
		Value: Value{
			Type:   "tab",
			String: "\t",
		},
	},
	"\n": Token{
		Type: Whitespace,
		Value: Value{
			Type:   "newline",
			String: "\n",
		},
	},
}
