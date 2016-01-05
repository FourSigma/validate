package validate

import "github.com/FourSigma/validate/str"

type stng struct {
	t *str.TransStr //TransStr
	s *str.Str      //Checker
}

func (s stng) Validate(fn ...str.Handler) *str.Str {
	s.s.Add(fn...)
	return s.s
}

func (s *stng) Transform(fn ...str.TransHandler) *str.TransStr {
	s.t.Add(fn...)
	return s.t
}

func String(s *string) *stng {
	return &stng{
		s: str.NewChkStr(*s, nil),
		t: str.NewTransStr(s, nil),
	}
}
