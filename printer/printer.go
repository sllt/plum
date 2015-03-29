package printer

import (
	"fmt"
	"strings"
)

import (
	"github.com/sllt/plum/types"
)

func Pr_list(lst []types.PlumType, pr bool,
	start string, end string, join string) string {
	str_list := make([]string, 0, len(lst))
	for _, e := range lst {
		str_list = append(str_list, Pr_str(e, pr))
	}
	return start + strings.Join(str_list, join) + end
}

func Pr_str(obj types.PlumType, print_readably bool) string {
	switch tobj := obj.(type) {
	case types.List:
		return Pr_list(tobj.Val, print_readably, "(", ")", " ")
	case types.Vector:
		return Pr_list(tobj.Val, print_readably, "[", "]", " ")
	case types.HashMap:
		str_list := make([]string, 0, len(tobj.Val)*2)
		for k, v := range tobj.Val {
			str_list = append(str_list, Pr_str(k, print_readably))
			str_list = append(str_list, Pr_str(v, print_readably))
		}
		return "{" + strings.Join(str_list, " ") + "}"
	case string:
		if strings.HasPrefix(tobj, "\u029e") {
			return ":" + tobj[2:len(tobj)]
		} else if print_readably {
			return `"` + strings.Replace(
				strings.Replace(
					strings.Replace(tobj, `\`, `\\`, -1),
					`"`, `\"`, -1),
				"\n", `\n`, -1) + `"`
		} else {
			return tobj
		}
	case types.Symbol:
		return tobj.Val
	case nil:
		return "nil"
	case types.PlumFunc:
		return "(fn* " +
			Pr_str(tobj.Params, true) + " " +
			Pr_str(tobj.Exp, true) + ")"
	case func([]types.PlumType) (types.PlumType, error):
		return fmt.Sprintf("<function %v>", obj)
	case *types.Atom:
		return "(atom " +
			Pr_str(tobj.Val, true) + ")"
	default:
		return fmt.Sprintf("%v", obj)
	}
}
