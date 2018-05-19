package token

// KeywordMap is a map of all the keywords
var KeywordMap = map[string]Token{
	"let": Token{
		Type:     "TYPE",
		Expected: "IDENT",
		Value: Value{
			Type:   "var",
			String: "let",
		},
	},
	"select": Token{
		ID:       9,
		Type:     "KEYWORD",
		Expected: "BLOCK",
		Value: Value{
			Type:   "keyword", // TODO: what to put here?
			String: "select",
		},
	},
}
