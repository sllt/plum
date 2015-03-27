package main

import (
	"errors"
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

// eval
func eval_ast(ast PlumType, env map[string]PlumType) (PlumType, error) {
	//fmt.Printf("eval_ast: %#v\n", ast)
	if Symbol_Q(ast) {
		k := ast.(Symbol).Val
		exp, ok := env[k]
		if !ok {
			return nil, errors.New("'" + k + "' not found")
		}
		return exp, nil
	} else if List_Q(ast) {
		lst := []PlumType{}
		for _, a := range ast.(List).Val {
			exp, e := EVAL(a, env)
			if e != nil {
				return nil, e
			}
			lst = append(lst, exp)
		}
		return List{lst, nil}, nil
	} else if Vector_Q(ast) {
		lst := []PlumType{}
		for _, a := range ast.(Vector).Val {
			exp, e := EVAL(a, env)
			if e != nil {
				return nil, e
			}
			lst = append(lst, exp)
		}
		return Vector{lst, nil}, nil
	} else if HashMap_Q(ast) {
		m := ast.(HashMap)
		new_hm := HashMap{map[string]PlumType{}, nil}
		for k, v := range m.Val {
			ke, e1 := EVAL(k, env)
			if e1 != nil {
				return nil, e1
			}
			if _, ok := ke.(string); !ok {
				return nil, errors.New("non string hash-map key")
			}
			kv, e2 := EVAL(v, env)
			if e2 != nil {
				return nil, e2
			}
			new_hm.Val[ke.(string)] = kv
		}
		return new_hm, nil
	} else {
		return ast, nil
	}
}

func EVAL(ast PlumType, env map[string]PlumType) (PlumType, error) {
	//fmt.Printf("EVAL: %v\n", printer.Pr_str(ast, true))
	switch ast.(type) {
	case List: // continue
	default:
		return eval_ast(ast, env)
	}

	// apply list
	el, e := eval_ast(ast, env)
	if e != nil {
		return nil, e
	}
	f, ok := el.(List).Val[0].(func([]PlumType) (PlumType, error))
	if !ok {
		return nil, errors.New("attempt to call non-function")
	}
	return f(el.(List).Val[1:])
}

func PRINT(exp PlumType) (string, error) {
	return printer.Pr_str(exp, true), nil
}

var repl_env = map[string]PlumType{
	"+": func(a []PlumType) (PlumType, error) {
		return a[0].(int) + a[1].(int), nil
	},
	"-": func(a []PlumType) (PlumType, error) {
		return a[0].(int) - a[1].(int), nil
	},
	"*": func(a []PlumType) (PlumType, error) {
		return a[0].(int) * a[1].(int), nil
	},
	"/": func(a []PlumType) (PlumType, error) {
		return a[0].(int) / a[1].(int), nil
	},
}

// repl
func repl(str string) (PlumType, error) {
	var exp PlumType
	var res string
	var e error
	if exp, e = READ(str); e != nil {
		return nil, e
	}
	if exp, e = EVAL(exp, repl_env); e != nil {
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
