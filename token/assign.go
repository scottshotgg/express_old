package token

// AssignMap holds every assignment operator
var AssignMap = map[string]Token{
	"=": Token{
		Type: "ASSIGN",
		Value: Value{
			Type:   "assign",
			String: "=",
		},
	},
	":": Token{
		Type: "ASSIGN",
		Value: Value{
			Type:   "set",
			String: ":",
		},
	},
	":=": Token{
		Type: "ASSIGN",
		Value: Value{
			Type:   "init",
			String: ":=",
		},
	},
}
