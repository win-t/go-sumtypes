package json

import (
	"sort"
	"strconv"
	"strings"
)

func Stringify(v Value) string {
	var ret string
	v.Case(
		func(v String) {
			ret = `"` + strings.ReplaceAll(string(v), `"`, `\"`) + `"`
		},

		func(v Number) {
			ret = strconv.FormatFloat(float64(v), 'f', -1, 64)
		},

		func(v Bool) {
			if bool(v) {
				ret = "true"
			} else {
				ret = "false"
			}
		},

		func(v Null) {
			ret = "null"
		},

		func(v Array) {
			ret = "["
			for i, v := range v {
				if i != 0 {
					ret += ","
				}
				ret += Stringify(v)
			}
			ret += "]"
		},

		func(v Object) {
			var keys []string
			for k := range v {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			ret = "{"
			for i, k := range keys {
				if i != 0 {
					ret += ","
				}
				ret += Stringify(String(k).Value())
				ret += ":"
				ret += Stringify(v[k])
			}
			ret += "}"
		},
	)
	return ret
}
