package validate

import "github.com/FourSigma/validate/str"

type String string

func (s String) Validate(fn ...str.Handler) *str.Str {
	return str.NewStr(string(s), fn)
}
