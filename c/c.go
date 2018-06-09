package c

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

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

var (
	r   *rand.Rand
	f   *os.File
	err error
)

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

func translateObject(t token.Token) {
	trueValue := t.Value.True.(map[string]token.Value)
	mapString := "std::map<std::string, Any>" + t.Value.Name + ";\n"
	fmt.Println("map!!", trueValue)
	f.Write([]byte(mapString))

	mapValue := func() (valueString string) {
		for k, v := range trueValue {

			switch v.Type {
			case "string":
				// randomize the name of the integer
				varName := k + strconv.Itoa(int(r.Uint32()))
				valueString += "std::" + v.Type + " " + varName + " = " + fmt.Sprintf("\"%v\"", v.True) + ";\n"
				// valueString += "\n" + t.Value.Name + "[\"" + k + "\"] = Any{ \"" + v.Type + "\", (" + v.Type + "*)" + fmt.Sprintf("%v", v.True) + " };"
				valueString += t.Value.Name + "[\"" + k + "\"] = Any{ \"" + v.Type + "\", &" + varName + " };\n"
			default:
				// randomize the name of the integer
				varName := k + strconv.Itoa(int(r.Uint32()))
				valueString += v.Type + " " + varName + " = " + fmt.Sprintf("%v", v.True) + ";\n"
				// valueString += "\n" + t.Value.Name + "[\"" + k + "\"] = Any{ \"" + v.Type + "\", (" + v.Type + "*)" + fmt.Sprintf("%v", v.True) + " };"
				valueString += t.Value.Name + "[\"" + k + "\"] = Any{ \"" + v.Type + "\", &" + varName + " };\n"
			}
		}

		return
	}()

	fmt.Println("mapValue", mapValue)

	f.Write([]byte(mapValue))
}

func translateVariableStatement(t token.Token) {
	// if the token type is var make a var statement in C
	if t.Type == "VAR" {
		switch t.Value.Type {
		case "var":
			// int abc = 5;
			// Any zyx = Any{ "int", &abc };
			varName := t.Value.Name + strconv.Itoa(int(r.Uint32()))
			thing := strings.Join([]string{t.Value.Acting, varName, "=", fmt.Sprintf("%v", t.Value.True)}, " ") + ";\n"
			thing += "Any " + t.Value.Name + " = Any{ \"" + t.Value.Acting + "\", &" + varName + "};\n"
			fmt.Println(thing)
			_, err = f.Write([]byte(thing))
			if err != nil {
				fmt.Println("error writing to file")
				os.Exit(9)
			}

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

			thing := arrayType + " " + t.Value.Name + "[] = { " + arrayValue + " };\n"
			fmt.Println(thing)
			_, err = f.Write([]byte(thing))
			if err != nil {
				fmt.Println("error writing to file")
				os.Exit(9)
			}

		case "object":
			translateObject(t)

		// In the case of the object we need to essentially instantiate a struct that will be used even if only temporarily
		// could just use that json library for now but wtf
		// fmt.Println("std::map<string, " + +"> " + t.Value.Name)
		case "string":
			thing := "std::" + strings.Join([]string{t.Value.Type, t.Value.Name, "=", fmt.Sprintf("\"%v\"", t.Value.True)}, " ") + ";\n"
			fmt.Println(thing)
			_, err = f.Write([]byte(thing))
			if err != nil {
				fmt.Println("error writing to file")
				os.Exit(9)
			}

		default:
			thing := strings.Join([]string{t.Value.Type, t.Value.Name, "=", fmt.Sprintf("%v", t.Value.True)}, " ") + ";\n"
			fmt.Println(thing)
			_, err = f.Write([]byte(thing))
			if err != nil {
				fmt.Println("error writing to file")
				os.Exit(9)
			}
		}
	}
}

// Translate ...
// TODO: FIXME: this needs to be someone modularized or recursive to support nested structures
func Translate(tokens []token.Token) {
	// fmt.Println("tokens", tokens)

	r = rand.New(rand.NewSource(time.Now().UnixNano()))

	f, err = os.Create("main.expr.cpp")
	if err != nil {
		fmt.Println("got an err creating file")
		os.Exit(9)
	}

	// TODO: check all f.Write errors I guess
	f.Write([]byte("#include <map>\n#include <string>\n"))
	f.Write([]byte("struct Any { std::string type; void* data; };\n"))
	f.Write([]byte("int main() {\n"))

	for _, t := range tokens {
		fmt.Println(t)

		translateVariableStatement(t)
	}

	f.Write([]byte("}"))
}
