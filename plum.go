package main

import (
	"fmt"
	"plum/readline"
	"strings"
)

func READ(str string) string {
	return str
}

func EVAL(ast, env string) string {
	return ast
}

func PRINT(exp string) string {
	return exp
}

// repl
func repl(str string) string {
	return PRINT(EVAL(READ(str), ""))
}

func main() {
	for {
		text, err := readline.Readline("user> ")
		text = strings.TrimRight(text, "\n")
		if err != nil {
			return
		}
		fmt.Println(repl(text))
	}
}
