package str

import (
	"context"
	"errors"

	"github.com/FourSigma/validate/lib"
)

type HandlerFunc func(context.Context, *string) error

func (s HandlerFunc) Handle(ctx context.Context, i interface{}) error {
	str, ok := i.(*string)
	if !ok {
		return errors.New("String::HandlerFunc::Cannot type assert, must be of type *string.")
	}
	return s(ctx, str)
}

type HandlerFuncList []HandlerFunc

func (s HandlerFuncList) ToHandlers() []lib.Handler {
	rs := make([]lib.Handler, len(s))
	for i, v := range s {
		rs[i] = lib.Handler(v)
	}
	return rs
}

func NewStringValidator(s *string) *str {
	return &str{
		s:      s,
		Helper: lib.NewDefaultHelper(s, "String"),
	}
}

type StringValidator interface {
	Prepend(...HandlerFunc) StringValidator
	Append(...HandlerFunc) StringValidator
	lib.Helper
}

//Implementation of Validator interface for string primitives.
type str struct {
	s *string
	lib.Helper
}

func (s *str) IsEmpty() bool {
	if *s.s == "" {
		return true
	}
	return false
}

func (s *str) Prepend(hf ...HandlerFunc) StringValidator {
	hdl := append(HandlerFuncList(hf).ToHandlers(), s.GetHandlers()...)
	s.SetHandlers(hdl...)
	return s
}

func (s *str) Append(hf ...HandlerFunc) StringValidator {
	hdl := append(s.GetHandlers(), HandlerFuncList(hf).ToHandlers()...)
	s.SetHandlers(hdl...)
	return s
}
