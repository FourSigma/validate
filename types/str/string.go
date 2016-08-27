package str

import (
	"errors"

	"github.com/FourSigma/validate/lib"
)

type HandlerFunc func(*string) error

func (s HandlerFunc) Handle(i interface{}) error {
	str, ok := i.(*string)
	if !ok {
		return errors.New("String::HandlerFunc::Cannot type assert, must be of type *string.")
	}
	return s(str)
}

func ToHandlers(list ...HandlerFunc) []lib.Handler {
	rs := make([]lib.Handler, len(list))
	for i, v := range list {
		rs[i] = lib.Handler(v)
	}
	return rs

}

func String(s *string, hf ...HandlerFunc) *str {
	ss := &str{s: s, h: hf}
	ss.Helper = lib.NewHelper(s, "String", ToHandlers(hf...)...)
	return ss
}

//Implementation of Validator interface for string primitives.
type str struct {
	s *string
	h []HandlerFunc
	lib.Helper
}

func (s *str) IsEmpty() bool {
	if *s.s == "" {
		return true
	}
	return false
}

func (s *str) Prepend(hf ...HandlerFunc) *str {
	hdl := append(ToHandlers(hf...), s.GetHandlers()...)
	s.SetHandlers(hdl...)
	return s
}

func (s *str) Append(hf ...HandlerFunc) *str {
	hdl := append(s.GetHandlers(), ToHandlers(hf...)...)
	s.SetHandlers(hdl...)
	return s
}
