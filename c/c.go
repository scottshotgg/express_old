package c

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/scottshotgg/Express/token"
)

// Anonymous Struct Declaration
// struct thing{
// 	std::string type;
// 	int amount;
// } thingy = {
// 	"hey",
// 	6,
// };

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

	typeMap = map[string]int{
		"int":    1,
		"string": 2,
		"bool":   3,
		"float":  4,
		"char":   5,
		"object": 6,
		// don't know if I'll need this
		"int[]": 7,
	}
)

// func translateObject(t token.Token) {
// 	trueValue := t.True.(map[string]token.Value)
// 	mapString := "std::map<std::string, Any>" + t.Name + ";\n"
// 	fmt.Println("map!!", trueValue)
// 	f.Write([]byte(mapString))

// 	mapValue := func() (valueString string) {
// 		for k, v := range trueValue {
// 			switch v.Type {
// 			// TODO: dont feel like doing this rn, something with the quotes and checking string type w/e
// 			// case "var":
// 			// // int abc = 5;
// 			// // Any zyx = Any{ "int", &abc };
// 			// varName := k + strconv.Itoa(int(r.Uint32()))
// 			// thing := strings.Join([]string{v.Acting, varName, "=", fmt.Sprintf("%v", v.True)}, " ") + ";\n"
// 			// thing += "Any " + k + " = Any{ \"" + v.Acting + "\", &" + varName + "};\n"
// 			// fmt.Println(thing)
// 			// _, err = f.Write([]byte(thing))
// 			// if err != nil {
// 			// 	fmt.Println("error writing to file")
// 			// 	os.Exit(9)
// 			// }

// 			case "object":
// 				v.Name = k
// 				translateObject(token.Token{
// 					Type:  "VAR",
// 					Value: v,
// 				})
// 				valueString += t.Name + "[\"" + k + "\"] = Any{ \"" + v.Type + "\", &" + v.Name + " };\n"

// 			case "array":
// 				v.Name = k
// 				translateArray(token.Token{
// 					Type:  "VAR",
// 					Value: v,
// 				})
// 				valueString += t.Name + "[\"" + k + "\"] = Any{ \"" + v.Type + "\", &" + v.Name + " };\n"

// 			case "string":
// 				// randomize the name of the integer
// 				varName := k + strconv.Itoa(int(r.Uint32()))
// 				valueString += "std::" + v.Type + " " + varName + " = " + fmt.Sprintf("\"%v\"", v.True) + ";\n"
// 				// valueString += "\n" + t.Name + "[\"" + k + "\"] = Any{ \"" + v.Type + "\", (" + v.Type + "*)" + fmt.Sprintf("%v", v.True) + " };"
// 				valueString += t.Name + "[\"" + k + "\"] = Any{ \"" + v.Type + "\", &" + varName + " };\n"

// 			default:
// 				// randomize the name of the integer
// 				varName := k + strconv.Itoa(int(r.Uint32()))
// 				valueString += v.Type + " " + varName + " = " + fmt.Sprintf("%v", v.True) + ";\n"
// 				// valueString += "\n" + t.Name + "[\"" + k + "\"] = Any{ \"" + v.Type + "\", (" + v.Type + "*)" + fmt.Sprintf("%v", v.True) + " };"
// 				valueString += t.Name + "[\"" + k + "\"] = Any{ \"" + v.Type + "\", &" + varName + " };\n"
// 			}
// 		}

// 		return
// 	}()

// 	fmt.Println("mapValue", mapValue)

// 	f.Write([]byte(mapValue))
// }

func translateArray(t token.Value) {
	fmt.Println(t)
	trueValue := t.True.([]token.Value)
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
	thing := arrayType + " " + t.Name + "[] = { " + arrayValue + " };\n"
	fmt.Println(thing)
	_, err = f.Write([]byte(thing))
	if err != nil {
		fmt.Println("error writing to file")
		os.Exit(9)
	}
}

func translateVariableStatement(t token.Value) error {
	// if the token type is var make a var statement in C

	switch t.Type {
	case "var":
		// int abc = 5;
		// Any zyx = Any{ "int", &abc };
		varName := t.Name + strconv.Itoa(int(r.Uint32()))
		thing := strings.Join([]string{t.Acting, varName, "=", fmt.Sprintf("%v", t.True)}, " ") + ";\n"
		thing += "Any " + t.Name + " = Any{ \"" + t.Acting + "\", &" + varName + " };\n"
		fmt.Println(thing)
		_, err = f.Write([]byte(thing))
		if err != nil {
			fmt.Println("error writing to file")
			os.Exit(9)
		}
		return nil

	case "object":
		// translateObject(t)

	case "array":
		translateArray(t)
		return nil

	// In the case of the object we need to essentially instantiate a struct that will be used even if only temporarily
	// could just use that json library for now but wtf
	// fmt.Println("std::map<string, " + +"> " + t.Name)
	case "string":
		thing := "std::" + strings.Join([]string{t.Type, t.Name, "=", fmt.Sprintf("\"%v\"", t.True)}, " ") + ";\n"
		fmt.Println(thing)
		_, err = f.Write([]byte(thing))
		if err != nil {
			fmt.Println("error writing to file")
			os.Exit(9)
		}
		return nil

	case "int":
		fallthrough
	case "bool":
		thing := strings.Join([]string{t.Type, t.Name, "=", fmt.Sprintf("%v", t.True)}, " ") + ";\n"
		fmt.Println(thing)
		_, err = f.Write([]byte(thing))
		if err != nil {
			fmt.Println("error writing to file")
			os.Exit(9)
		}
		return nil

	default:
		fmt.Println("am i an error ???")
		return errors.New("i am not nil")
	}
	return errors.New("why am i here")
}

func translateIf(t token.Value) {
	fmt.Println("wtf")
	fmt.Printf("t %+v\n", t)

	_, err = f.Write([]byte(fmt.Sprintf("if (%s) {\n", t.String)))
	if err != nil {
		fmt.Println("error writing to file")
		os.Exit(9)
	}

	opMap, ok := t.True.(map[string]token.Value)
	if !ok {
		fmt.Println("omfg error")
		os.Exit(9)
	}

	// body, ok := opMap["body"].True.([]token.Value)
	// if !ok {
	// 	fmt.Println("omfg error")
	// 	os.Exit(9)
	// }
	TranslateBlock(opMap["body"])

	_, err = f.Write([]byte("\n}\n"))
	if err != nil {
		fmt.Println("error writing to file")
		os.Exit(9)
	}
}

func translateLoop(t token.Value) error {
	if t.Type != "for" {
		return errors.New("blah")
	}

	fmt.Printf("t %+v\n", t)
	tValue, ok := t.True.(map[string]token.Value)
	if !ok {
		return errors.New("not the type")
	}
	fmt.Println(tValue)

	loop := fmt.Sprintf("{\nint %s=%d;\nwhile (%s<%d) {\n",
		tValue["start"].Name, tValue["start"].True.(int), tValue["start"].Name,
		tValue["end"].True.(int))
	fmt.Println(loop)
	_, err = f.Write([]byte(loop))
	if err != nil {
		fmt.Println("error writing to file")
		os.Exit(9)
	}

	TranslateBlock(tValue["body"])

	loopEnding := fmt.Sprintf("%s+=%d;\n}\n}\n", tValue["start"].Name, tValue["step"].True.(int))
	fmt.Println(loopEnding)
	_, err = f.Write([]byte(loopEnding))
	if err != nil {
		fmt.Println("error writing to file")
		os.Exit(9)
	}

	return nil
}

func TranslateBlock(tv token.Value) {
	_, err = f.Write([]byte("{\n"))
	if err != nil {
		fmt.Println("error writing to file")
		os.Exit(9)
	}

	insideBlock := tv.True.([]token.Value)
	fmt.Println("insideBlock", insideBlock[0])

	for _, t := range insideBlock {
		fmt.Println("t", t)

		if err = translateVariableStatement(t); err != nil {
			fmt.Println("i am here translateVariableStatement", err)
			if err = translateLoop(t); err != nil {
				translateIf(t)
			}
		}
	}

	f.Write([]byte("}"))
}

// Translate ...
// TODO: FIXME: this needs to be someone modularized or recursive to support nested structures
func Translate(tokens []token.Value) {
	fmt.Println("\n\n\ntokens", tokens)

	r = rand.New(rand.NewSource(time.Now().UnixNano()))

	f, err = os.Create("main.expr.cpp")
	if err != nil {
		fmt.Println("got an err creating file")
		os.Exit(9)
	}

	// TODO: check all f.Write errors I guess
	f.Write([]byte("#include <map>\n#include <string>\n"))
	f.Write([]byte("struct Any { std::string type; void* data; };\n"))
	f.Write([]byte("int main()"))

	TranslateBlock(tokens[0])
}
