package validate

import "github.com/FourSigma/validate/str"

type String []byte

func (s String) Validate(fn ...str.Handler) *str.Str {
	return str.NewChkStr(string(s), fn)
}

func (s String) Transform(fn ...str.TransHandler) *str.TransStr {
	return str.NewTransStr(s, fn)
}
