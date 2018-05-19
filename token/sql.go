package token

// FIXME: leave this for fixing until we need it

// SQLMap is a map of all the SQL specific tokens
var SQLMap = map[string]Token{
	"SELECT": Token{
		ID:   9,
		Type: Keyword,
		Value: Value{
			Type:   "SQL", // TODO: what to put here?
			String: "SELECT",
		},
	},
	"FROM": Token{
		ID:   9,
		Type: Keyword,
		Value: Value{
			Type:   "SQL", // TODO: what to put here?
			String: "FROM",
		},
	},
	"WHERE": Token{
		ID:   9,
		Type: Keyword,
		Value: Value{
			Type:   "SQL", // TODO: what to put here?
			String: "WHERE",
		},
	},
}
