package token

// SeparatorMap holds all defined statement separators
var SeparatorMap = map[string]Token{
	",": Token{
		Type: "SEPARATOR",
		Value: Value{
			Type:   "comma",
			String: ",",
		},
	},
	";": Token{
		Type: "SEPARATOR",
		Value: Value{
			Type:   "semicolon",
			String: ";",
		},
	},
}
