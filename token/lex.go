package token

// Lexemes ...
var (
	Lexemes = []string{
		"var",
		"int",
		"string",
		"select",
		"SELECT",
		"FROM",
		"WHERE",
		":",
		"=",
		"+",
		"-",
		"*",
		"/",
		"(",
		")",
		"{",
		"}",
		"[",
		"]",
		"\"",
		";",
		",",
		"#",
		"!",
		"<",
		">",
		"@",
		// "„",
	}

	// LexemeMap is used for holding the lexemes that will be used to identify tokens in the lexer
	LexemeMap = map[string]Token{}
)
