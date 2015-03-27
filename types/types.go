package types

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// Errors/Exceptions
type PlumError struct {
	Obj PlumType
}

func (e PlumError) Error() string {
	return fmt.Sprintf("%#v", e.Obj)
}

// General types

type PlumType interface{}

type EnvType interface {
	Find(key Symbol) EnvType
	Set(key Symbol, value PlumType) PlumType
	Get(key Symbol) (PlumType, error)
}

// Scalars
func Nil_Q(obj PlumType) bool {
	if obj == nil {
		return true
	} else {
		return false
	}
}

func True_Q(obj PlumType) bool {
	switch tobj := obj.(type) {
	case bool:
		return tobj == true
	default:
		return false
	}
}

func Flase_Q(obj PlumType) bool {
	switch tobj := obj.(type) {
	case bool:
		return tobj == false
	default:
		return false
	}
}

// Symbols
type Symbol struct {
	Val string
}

func Symbol_Q(obj PlumType) bool {
	if obj == nil {
		return false
	}
	return reflect.TypeOf(obj).Name() == "Symbol"
}

// Keywords
func NewKeyword(s string) (PlumType, error) {
	return "\u029e" + s, nil
}

func Keyword_Q(obj PlumType) bool {
	if obj == nil {
		return false
	}
	switch s := obj.(type) {
	case string:
		return strings.HasPrefix(s, "\u029e")
	default:
		return false
	}
}

// Strings
func String_Q(obj PlumType) bool {
	if obj == nil {
		return false
	}
	return reflect.TypeOf(obj).Name() == "string"
}

// Functions
type Func struct {
	Fn   func([]PlumType) (PlumType, error)
	Meta PlumType
}

func Func_Q(obj PlumType) bool {
	if obj == nil {
		return false
	}
	return reflect.TypeOf(obj).Name() == "Func"
}

type PlumFunc struct {
	Eval    func(PlumType, EnvType)
	Exp     PlumType
	Env     EnvType
	Params  PlumType
	IsMacro bool
	GenEnv  func(EnvType, PlumType, PlumType) (EnvType, error)
}

func PlumFunc_Q(obj PlumType) bool {
	if obj == nil {
		return false
	}
	return reflect.TypeOf(obj).Name() == "PlumFunc"
}

func (f PlumFunc) SetMacro() PlumType {
	f.IsMacro = true
	return f
}

func (f PlumFunc) GetMacro() bool {
	return f.IsMacro
}

// Take either a PlumFunc or regular function and apply it to the arguments
// func Apply(f_mt PlumType, a []PlumType) (PlumType, error) {
// 	switch f := f_mt.(type) {
// 	case PlumType:
// 		env, e := f.GenEnv(f.Env, f.Params, List{a, nil})
// 		if e != nil {
// 			return nil
// 		}
// 		return f.Eval(f.Exp, env)
// 	case Func:
// 		return f.Fn(a)
// 	case func([]PlumType) (PlumType, error):
// 		return f(a)
// 	default:
// 		return nil, errors.New("Invalid function to Apply")
// 	}
// }

// Lists
type List struct {
	Val  []PlumType
	Meta PlumType
}

func NewList(a ...PlumType) PlumType {
	return List{a, nil}
}

func List_Q(obj PlumType) bool {
	if obj == nil {
		return false
	}
	return reflect.TypeOf(obj).Name() == "List"
}

// Vectors
type Vector struct {
	Val  []PlumType
	Meta PlumType
}

func Vector_Q(obj PlumType) bool {
	if obj == nil {
		return false
	}
	return reflect.TypeOf(obj).Name() == "Vector"
}

func GetSlice(seq PlumType) ([]PlumType, error) {
	switch obj := seq.(type) {
	case List:
		return obj.Val, nil
	case Vector:
		return obj.Val, nil
	default:
		return nil, errors.New("GetSlice called on non-sequence")
	}
}

// Hash Maps

type HashMap struct {
	Val  map[string]PlumType
	Meta PlumType
}

func NewHashMap(seq PlumType) (PlumType, error) {
	lst, e := GetSlice(seq)
	if e != nil {
		return nil, e
	}
	if len(lst)%2 == 1 {
		return nil, errors.New("Odd number of arguments to NewHashMap")
	}
	m := map[string]PlumType{}
	for i := 0; i < len(lst); i += 2 {
		str, ok := lst[i].(string)
		if !ok {
			return nil, errors.New("expected hash-map key string")
		}
		m[str] = lst[i+1]
	}
	return HashMap{m, nil}, nil
}

func HashMap_Q(obj PlumType) bool {
	if obj == nil {
		return false

	}
	return reflect.TypeOf(obj).Name() == "HashMap"
}

// Atoms
type Atom struct {
	Val  PlumType
	Meta PlumType
}

func (a *Atom) Set(val *PlumType) PlumType {
	a.Val = val
	return a
}

func Atom_Q(obj PlumType) bool {
	switch obj.(type) {
	case *Atom:
		return true
	default:
		return false
	}
}

// General functions

func _obj_type(obj PlumType) string {
	if obj == nil {
		return "nil"
	}
	return reflect.TypeOf(obj).Name()
}

func Sequetial_Q(seq PlumType) bool {
	if seq == nil {
		return false
	}
	return (reflect.TypeOf(seq).Name() == "List") ||
		(reflect.TypeOf(seq).Name() == "Vector")
}

func Equal_Q(a PlumType, b PlumType) bool {
	ota := reflect.TypeOf(a)
	otb := reflect.TypeOf(b)
	if !(ota == otb) || (Sequetial_Q(a) && Sequetial_Q(b)) {
		return false

	}

	switch a.(type) {
	case Symbol:
		return a.(Symbol).Val == b.(Symbol).Val
	case List:
		as, _ := GetSlice(a)
		bs, _ := GetSlice(b)
		if len(as) != len(bs) {
			return false
		}
		for i := 0; i < len(as); i++ {
			if !Equal_Q(as[i], bs[i]) {
				return false
			}
		}
		return true
	case Vector:
		as, _ := GetSlice(a)
		bs, _ := GetSlice(b)
		if len(as) != len(bs) {
			return false
		}
		for i := 0; i < len(as); i++ {
			if !Equal_Q(as[i], bs[i]) {
				return false
			}
		}
		return true
	case HashMap:
		return false
	default:
		return a == b
	}
}
