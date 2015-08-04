package core

import (
	"errors"
	"fmt"
	"io/ioutil"
	"time"
)

import (
	"plum/printer"
	"plum/reader"
	"plum/readline"
	. "plum/types"
)

// Errors/Exceptions
func throw(a []PlumType) (PlumType, error) {
	return nil, PlumError{a[0]}
}

// String functions

func pr_str(a []PlumType) (PlumType, error) {
	return printer.Pr_list(a, true, "", "", " "), nil
}

func str(a []PlumType) (PlumType, error) {
	return printer.Pr_list(a, false, "", "", ""), nil
}

func prn(a []PlumType) (PlumType, error) {
	fmt.Println(printer.Pr_list(a, true, "", "", " "))
	return nil, nil
}

func println(a []PlumType) (PlumType, error) {
	fmt.Println(printer.Pr_list(a, false, "", "", " "))
	return nil, nil
}

func slurp(a []PlumType) (PlumType, error) {
	switch fileName := a[0].(type) {
	case string:
		b, e := ioutil.ReadFile(fileName)
		if e != nil {
			return nil, e
		}
		return string(b), nil

	}
	return nil, nil

}

// Number functions
func time_ms(a []PlumType) (PlumType, error) {
	return int(time.Now().UnixNano() / int64(time.Millisecond)), nil
}

// Hash Map functions
func copy_hash_map(hm HashMap) HashMap {
	new_hm := HashMap{map[string]PlumType{}, nil}
	for k, v := range hm.Val {
		new_hm.Val[k] = v
	}
	return new_hm
}

func assoc(a []PlumType) (PlumType, error) {
	if len(a) < 3 {
		return nil, errors.New("assoc requires at least 3 arguments")
	}
	if len(a)%2 != 1 {
		return nil, errors.New("assoc requires odd number of arguments")
	}
	if !HashMap_Q(a[0]) {
		return nil, errors.New("assoc called on non-hash map")
	}
	new_hm := copy_hash_map(a[0].(HashMap))
	for i := 1; i < len(a); i += 2 {
		key := a[i]
		if !String_Q(key) {
			return nil, errors.New("assoc called with non-string key")
		}
		new_hm.Val[key.(string)] = a[i+1]
	}
	return new_hm, nil
}

func dissoc(a []PlumType) (PlumType, error) {
	if len(a) < 2 {
		return nil, errors.New("dissoc requires at least 3 arguments")
	}
	if !HashMap_Q(a[0]) {
		return nil, errors.New("dissoc called on non-hash map")
	}
	new_hm := copy_hash_map(a[0].(HashMap))
	for i := 1; i < len(a); i += 1 {
		key := a[i]
		if !String_Q(key) {
			return nil, errors.New("dissoc called with non-string key")
		}
		delete(new_hm.Val, key.(string))
	}
	return new_hm, nil
}

func get(a []PlumType) (PlumType, error) {
	if len(a) != 2 {
		return nil, errors.New("get requires 2 arguments")
	}
	if Nil_Q(a[0]) {
		return nil, nil
	}
	if !HashMap_Q(a[0]) {
		return nil, errors.New("get called on non-hash map")
	}
	if !String_Q(a[1]) {
		return nil, errors.New("get called with non-string key")
	}
	return a[0].(HashMap).Val[a[1].(string)], nil
}

func contains_Q(hm PlumType, key PlumType) (PlumType, error) {
	if Nil_Q(hm) {
		return false, nil
	}
	if !HashMap_Q(hm) {
		return nil, errors.New("get called on non-hash map")
	}
	if !String_Q(key) {
		return nil, errors.New("get called with non-string key")
	}
	_, ok := hm.(HashMap).Val[key.(string)]
	return ok, nil
}

func keys(a []PlumType) (PlumType, error) {
	if !HashMap_Q(a[0]) {
		return nil, errors.New("keys called on non-hash map")
	}
	slc := []PlumType{}
	for k, _ := range a[0].(HashMap).Val {
		slc = append(slc, k)
	}
	return List{slc, nil}, nil
}
func vals(a []PlumType) (PlumType, error) {
	if !HashMap_Q(a[0]) {
		return nil, errors.New("keys called on non-hash map")
	}
	slc := []PlumType{}
	for _, v := range a[0].(HashMap).Val {
		slc = append(slc, v)
	}
	return List{slc, nil}, nil
}

// Sequence functions

func cons(a []PlumType) (PlumType, error) {
	val := a[0]
	lst, e := GetSlice(a[1])
	if e != nil {
		return nil, e
	}

	return List{append([]PlumType{val}, lst...), nil}, nil
}

func concat(a []PlumType) (PlumType, error) {
	if len(a) == 0 {
		return List{}, nil
	}
	slc1, e := GetSlice(a[0])
	if e != nil {
		return nil, e
	}
	for i := 1; i < len(a); i += 1 {
		slc2, e := GetSlice(a[i])
		if e != nil {
			return nil, e
		}
		slc1 = append(slc1, slc2...)
	}
	return List{slc1, nil}, nil
}

func nth(a []PlumType) (PlumType, error) {
	slc, e := GetSlice(a[0])
	if e != nil {
		return nil, e
	}
	idx := a[1].(int)
	if idx < len(slc) {
		return slc[idx], nil
	} else {
		return nil, errors.New("nth: index out of range")
	}
}

func first(a []PlumType) (PlumType, error) {
	if len(a) == 0 {
		return nil, nil
	}
	slc, e := GetSlice(a[0])
	if e != nil {
		return nil, e
	}
	if len(slc) == 0 {
		return nil, nil
	}
	return slc[0], nil
}

func rest(a []PlumType) (PlumType, error) {
	slc, e := GetSlice(a[0])
	if e != nil {
		return nil, e
	}
	if len(slc) == 0 {
		return List{}, nil
	}
	return List{slc[1:], nil}, nil
}

func empty_Q(a []PlumType) (PlumType, error) {
	switch obj := a[0].(type) {
	case List:
		return len(obj.Val) == 0, nil
	case Vector:
		return len(obj.Val) == 0, nil
	case nil:
		return true, nil
	default:
		return nil, errors.New("Count called on non-sequence")
	}
}

func count(a []PlumType) (PlumType, error) {
	switch obj := a[0].(type) {
	case List:
		return len(obj.Val), nil
	case Vector:
		return len(obj.Val), nil
	case map[string]PlumType:
		return len(obj), nil
	case nil:
		return 0, nil
	default:
		return nil, errors.New("Count called on non-sequence")
	}
}

func apply(a []PlumType) (PlumType, error) {
	if len(a) < 2 {
		return nil, errors.New("apply requires at least 2 args")
	}
	f := a[0]
	args := []PlumType{}
	for _, b := range a[1 : len(a)-1] {
		args = append(args, b)
	}
	last, e := GetSlice(a[len(a)-1])
	if e != nil {
		return nil, e
	}
	args = append(args, last...)
	return Apply(f, args)
}

func funcall(a []PlumType) (PlumType, error) {
	if len(a) < 1 {
		return nil, errors.New("funcall requires at least 1 args")
	}
	f := a[0]
	args := []PlumType{}
	for _, b := range a[1:len(a)] {
		args = append(args, b)
		// fmt.Println(b)
	}
	return Apply(f, args)
}

func do_map(a []PlumType) (PlumType, error) {
	if len(a) != 2 {
		return nil, errors.New("map requires 2 args")
	}
	f := a[0]
	results := []PlumType{}
	args, e := GetSlice(a[1])
	if e != nil {
		return nil, e
	}
	for _, arg := range args {
		res, e := Apply(f, []PlumType{arg})
		results = append(results, res)
		if e != nil {
			return nil, e
		}
	}
	return List{results, nil}, nil
}

func conj(a []PlumType) (PlumType, error) {
	if len(a) < 2 {
		return nil, errors.New("conj requires at least 2 arguments")
	}
	switch seq := a[0].(type) {
	case List:
		new_slc := []PlumType{}
		for i := len(a) - 1; i > 0; i -= 1 {
			new_slc = append(new_slc, a[i])
		}
		return List{append(new_slc, seq.Val...), nil}, nil
	case Vector:
		new_slc := seq.Val
		for _, x := range a[1:] {
			new_slc = append(new_slc, x)
		}
		return Vector{new_slc, nil}, nil
	}

	if !HashMap_Q(a[0]) {
		return nil, errors.New("dissoc called on non-hash map")
	}
	new_hm := copy_hash_map(a[0].(HashMap))
	for i := 1; i < len(a); i += 1 {
		key := a[i]
		if !String_Q(key) {
			return nil, errors.New("dissoc called with non-string key")
		}
		delete(new_hm.Val, key.(string))
	}
	return new_hm, nil
}

// Metadata functions
func with_meta(a []PlumType) (PlumType, error) {
	if len(a) != 2 {
		return nil, errors.New("with-meta requires 2 args")
	}
	obj := a[0]
	m := a[1]
	switch tobj := obj.(type) {
	case List:
		return List{tobj.Val, m}, nil
	case Vector:
		return Vector{tobj.Val, m}, nil
	case HashMap:
		return HashMap{tobj.Val, m}, nil
	case Func:
		return Func{tobj.Fn, m}, nil
	case PlumFunc:
		fn := tobj
		fn.Meta = m
		return fn, nil
	default:
		return nil, errors.New("with-meta not supported on type")
	}
}

func meta(a []PlumType) (PlumType, error) {
	obj := a[0]
	switch tobj := obj.(type) {
	case List:
		return tobj.Meta, nil
	case Vector:
		return tobj.Meta, nil
	case HashMap:
		return tobj.Meta, nil
	case Func:
		return tobj.Meta, nil
	case PlumFunc:
		return tobj.Meta, nil
	default:
		return nil, errors.New("meta not supported on type")
	}
}

// Atom functions
func deref(a []PlumType) (PlumType, error) {
	if !Atom_Q(a[0]) {
		return nil, errors.New("deref called with non-atom")
	}
	return a[0].(*Atom).Val, nil
}

func reset_BANG(a []PlumType) (PlumType, error) {
	if !Atom_Q(a[0]) {
		return nil, errors.New("reset! called with non-atom")
	}
	a[0].(*Atom).Set(a[1])
	return a[1], nil
}

func swap_BANG(a []PlumType) (PlumType, error) {
	if !Atom_Q(a[0]) {
		return nil, errors.New("swap! called with non-atom")
	}
	if len(a) < 2 {
		return nil, errors.New("swap! requires at least 2 args")
	}
	atm := a[0].(*Atom)
	args := []PlumType{atm.Val}
	f := a[1]
	args = append(args, a[2:]...)
	res, e := Apply(f, args)
	if e != nil {
		return nil, e
	}
	atm.Set(res)
	return res, nil
}

// core namespace
var NS = map[string]PlumType{
	"=": func(a []PlumType) (PlumType, error) {
		return Equal_Q(a[0], a[1]), nil
	},
	"throw": throw,
	"nil?": func(a []PlumType) (PlumType, error) {
		return Nil_Q(a[0]), nil
	},
	"true?": func(a []PlumType) (PlumType, error) {
		return True_Q(a[0]), nil
	},
	"false?": func(a []PlumType) (PlumType, error) {
		return False_Q(a[0]), nil
	},
	"symbol": func(a []PlumType) (PlumType, error) {
		return Symbol{a[0].(string)}, nil
	},
	"symbol?": func(a []PlumType) (PlumType, error) {
		return Symbol_Q(a[0]), nil
	},
	"keyword": func(a []PlumType) (PlumType, error) {
		if Keyword_Q(a[0]) {
			return a[0], nil
		} else {
			return NewKeyword(a[0].(string))
		}
	},
	"keyword?": func(a []PlumType) (PlumType, error) {
		return Keyword_Q(a[0]), nil
	},

	"pr-str":  func(a []PlumType) (PlumType, error) { return pr_str(a) },
	"str":     func(a []PlumType) (PlumType, error) { return str(a) },
	"prn":     func(a []PlumType) (PlumType, error) { return prn(a) },
	"println": func(a []PlumType) (PlumType, error) { return println(a) },
	"read-string": func(a []PlumType) (PlumType, error) {
		return reader.Read_str(a[0].(string))
	},
	"slurp": slurp,
	"readline": func(a []PlumType) (PlumType, error) {
		return readline.Readline(a[0].(string))
	},

	"<": func(a []PlumType) (PlumType, error) {
		return a[0].(int) < a[1].(int), nil
	},
	"<=": func(a []PlumType) (PlumType, error) {
		return a[0].(int) <= a[1].(int), nil
	},
	">": func(a []PlumType) (PlumType, error) {
		return a[0].(int) > a[1].(int), nil
	},
	">=": func(a []PlumType) (PlumType, error) {
		return a[0].(int) >= a[1].(int), nil
	},
	"+": func(a []PlumType) (PlumType, error) {
		var value float64 = 0
		for _, v := range a {
			switch v.(type) {
			case float64:
				value += v.(float64)
			case int:
				value += float64(v.(int))
			}
		}
		return value, nil
	},
	"-": func(a []PlumType) (PlumType, error) {
		var value int
		if len(a) <= 1 {
			value = 0
		} else {
			value = a[0].(int)
		}
		for i := 1; i < len(a); i++ {
			value -= a[i].(int)
		}
		return value, nil
	},
	"*": func(a []PlumType) (PlumType, error) {
		var value int = 1
		for _, v := range a {
			value *= v.(int)
		}
		return value, nil
	},
	"/": func(a []PlumType) (PlumType, error) {
		var value int
		if len(a) < 1 {
			return nil, errors.New("the expected number of arguments does not match the given number")
		} else if len(a) == 1 {
			return 1 / (a[0].(int)), nil
		} else {
			value = a[0].(int)
		}
		for i := 1; i < len(a); i++ {
			value /= a[i].(int)
		}
		return value, nil
	},
	"time-ms": time_ms,

	"list": func(a []PlumType) (PlumType, error) {
		return List{a, nil}, nil
	},
	"list?": func(a []PlumType) (PlumType, error) {
		return List_Q(a[0]), nil
	},
	"vector": func(a []PlumType) (PlumType, error) {
		return Vector{a, nil}, nil
	},
	"vector?": func(a []PlumType) (PlumType, error) {
		return Vector_Q(a[0]), nil
	},
	"hash-map": func(a []PlumType) (PlumType, error) {
		return NewHashMap(List{a, nil})
	},
	"map?": func(a []PlumType) (PlumType, error) {
		return HashMap_Q(a[0]), nil
	},
	"assoc":  assoc,
	"dissoc": dissoc,
	"get":    get,
	"contains?": func(a []PlumType) (PlumType, error) {
		return contains_Q(a[0], a[1])
	},
	"keys": keys,
	"vals": vals,

	"sequential?": func(a []PlumType) (PlumType, error) {
		return Sequential_Q(a[0]), nil
	},
	"cons":    cons,
	"concat":  concat,
	"nth":     nth,
	"first":   first,
	"rest":    rest,
	"empty?":  empty_Q,
	"count":   count,
	"apply":   apply,
	"funcall": funcall,
	"map":     do_map,
	"conj":    conj,

	"with-meta": with_meta,
	"meta":      meta,
	"atom": func(a []PlumType) (PlumType, error) {
		return &Atom{a[0], nil}, nil
	},
	"atom?": func(a []PlumType) (PlumType, error) {
		return Atom_Q(a[0]), nil
	},
	"deref":  deref,
	"reset!": reset_BANG,
	"swap!":  swap_BANG,
	"exec":   Exec,
	"exit":   Exit,
}
