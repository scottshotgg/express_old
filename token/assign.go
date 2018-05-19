package token

// AssignMap holds every assignment operator
var AssignMap = map[string]Token{
	"=": Token{
		Type:     "ASSIGN",
		Expected: "EXPR",
		Value: Value{
			Type:   "assign",
			String: "=",
		},
	},
	":": Token{
		Type:     "SET",
		Expected: "EXPR",
		Value: Value{
			Type:   "set_assign",
			String: ":",
		},
	},
	":=": Token{
		Type:     "INIT",
		Expected: "EXPR",
		Value: Value{
			Type:   "init_assign",
			String: ":=",
		},
	},
}
