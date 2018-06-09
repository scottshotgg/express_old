package c

import (
	"fmt"
	"strings"

	"github.com/scottshotgg/Express/token"
)

// This is a placeholder for the C converter package that will be used to convert Express -> C
// Doing so will allow Express to leverage all available C tools

// CAssignmentStatement ...
type CAssignmentStatement struct {
	Type  string
	Name  string
	Value string
}

// Translate ...
func Translate(tokens []token.Token) {
	// fmt.Println("tokens", tokens)

	for _, t := range tokens {
		fmt.Println(t)

		// if the token type is var make a var statement in C
		if t.Type == "VAR" {
			switch t.Value.Type {
			case "array":
				trueValue := t.Value.True.([]token.Value)

				// assuming only single type arrays until I have time to do multi type arrays in C
				arrayType := trueValue[0].Type
				arrayValue := func() (valueString string) {
					for i, v := range trueValue {
						valueString += fmt.Sprintf("%v", v.True)
						if i != len(trueValue)-1 {
							valueString += ", "
						}
					}

					return
				}()

				fmt.Println(arrayType + " " + t.Value.Name + " = { " + arrayValue + " }")

			case "object":
				// In the case of the object we need to essentially instantiate a struct that will be used even if only temporarily
				// could just use that json library for now but wtf

			default:
				fmt.Println(strings.Join([]string{t.Value.Type, t.Value.Name, "=", fmt.Sprintf("%v", t.Value.True)}, " ") + ";")
			}
		}
	}
}

// class MyFieldInterface
// {
//     int m_Size; // of course use appropriate access level in the real code...
//     ~MyFieldInterface() = default;
// }

// template <typename T>
// class MyField : public MyFieldInterface {
//     T m_Value;
// }

// struct MyClass {
//     std::map<string, MyFieldInterface* > fields;
// }
