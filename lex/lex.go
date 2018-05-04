package lex

import (
	"fmt"

	"github.com/sgg7269/tokenizer/token"
)

type lexMeta struct {
	Accumulator string
	Tokens      []token.Token
}

// Lex ...
func Lex(input string) ([]token.Token, error) {
	var meta lexMeta

	fmt.Println("lexing shit", input)

	for _, char := range input {

		fmt.Printf("char \"%s\" %s\n", string(char), meta.Accumulator)

		// TODO: need to decide whether we want to append to the accumulator first or second
		if string(char) == " " || string(char) == "\n" {
			if meta.Accumulator != "" {
				if lexemeToken, ok := token.LexemeMap[meta.Accumulator]; ok {
					fmt.Println("Found char1", meta.Accumulator)
					meta.Tokens = append(meta.Tokens, lexemeToken)
				} else {
					fmt.Println("Found literal1", meta.Accumulator)
					meta.Tokens = append(meta.Tokens, token.Token{
						ID:   0,
						Type: "LITERAL",
						Value: token.Value{
							String: meta.Accumulator,
						},
					})
				}
				// Pull this from the TokenMap because we don't want the space in the LexemeMap
				meta.Tokens = append(meta.Tokens, token.TokenMap[string(char)])
				meta.Accumulator = ""
			} else if string(char) == " " || string(char) == "\n" {
				fmt.Println("i got gere1")
				// Pull this from the TokenMap because we don't want the space in the LexemeMap
				meta.Tokens = append(meta.Tokens, token.TokenMap[string(char)])
				meta.Accumulator = ""
			}
			fmt.Println("continuing")

			continue

		} else {
			if lexemeToken, ok := token.LexemeMap[string(char)]; ok {
				fmt.Println("Found char2", meta.Accumulator)

				if meta.Accumulator != "" {
					fmt.Println("Found literal2", meta.Accumulator)
					meta.Tokens = append(meta.Tokens, token.Token{
						ID:   0,
						Type: "LITERAL",
						Value: token.Value{
							String: meta.Accumulator,
						},
					})
					meta.Accumulator = ""
				}

				meta.Tokens = append(meta.Tokens, lexemeToken)
				meta.Accumulator = ""

				// meta.Tokens = append()
				continue
			} else if string(char) == " " || string(char) == "\n" {
				fmt.Println("i got gere2")
				// Pull this from the TokenMap because we don't want the space in the LexemeMap
				meta.Tokens = append(meta.Tokens, token.TokenMap[string(char)])
				meta.Accumulator = ""
			}
		}

		meta.Accumulator += string(char)
		fmt.Println(meta.Accumulator)
	}

	return meta.Tokens, nil
}
