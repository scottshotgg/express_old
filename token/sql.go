package token

// SQLMap is a map of all the SQL specific tokens
var SQLMap = map[string]Token{
	"SELECT": Token{
		ID:       9,
		Type:     "KEYWORD",
		Expected: "EXPR",
		Value: Value{
			Type:   "SQL", // TODO: what to put here?
			String: "SELECT",
		},
	},
	"FROM": Token{
		ID:       9,
		Type:     "KEYWORD",
		Expected: "EXPR",
		Value: Value{
			Type:   "SQL", // TODO: what to put here?
			String: "FROM",
		},
	},
	"WHERE": Token{
		ID:       9,
		Type:     "KEYWORD",
		Expected: "EXPR",
		Value: Value{
			Type:   "SQL", // TODO: what to put here?
			String: "WHERE",
		},
	},
}
