package main

import (
	"fmt"
	"plum/printer"
	"plum/reader"
	"plum/readline"
	. "plum/types"
	"strings"
)

func READ(str string) (PlumType, error) {
	return reader.Read_str(str)
}

func EVAL(ast PlumType, env string) (PlumType, error) {
	return ast, nil
}

func PRINT(exp PlumType) (string, error) {
	return printer.Pr_str(exp, true), nil
}

// repl
func repl(str string) (PlumType, error) {
	var exp PlumType
	var res string
	var e error
	if exp, e = READ(str); e != nil {
		return nil, e
	}
	if exp, e = EVAL(exp, ""); e != nil {
		return nil, e
	}
	if res, e = PRINT(exp); e != nil {
		return nil, e
	}
	return res, nil
}

func main() {
	for {
		text, err := readline.Readline("user> ")
		text = strings.TrimRight(text, "\n")
		if err != nil {
			return
		}
		var out PlumType
		var e error
		if out, e = repl(text); e != nil {
			if e.Error() == "<empty line>" {
				continue
			}
			fmt.Printf("Error: %v\n", e)
			continue
		}
		fmt.Printf("%v\n", out)
	}
}
