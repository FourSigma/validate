package float

import (
	"context"
	"errors"

	"github.com/FourSigma/validate/lib"
)

type HandlerFunc func(context.Context, *float64) error

func (s HandlerFunc) Handle(ctx context.Context, i interface{}) error {
	flt64, ok := i.(*float64)
	if !ok {
		return errors.New("::HandlerFunc::Cannot type assert, must be of type *float64.")
	}
	return s(ctx, flt64)
}

type HandlerFuncList []HandlerFunc

func (s HandlerFuncList) ToHandlers() []lib.Handler {
	rs := make([]lib.Handler, len(s))
	for i, v := range s {
		rs[i] = lib.Handler(v)
	}
	return rs
}
func NewFloat64Validator(s *float64) *float64Number {
	return &float64Number{
		s:      s,
		Helper: lib.NewDefaultHelper(s, "Float64"),
	}
}

type Float64Validator interface {
	Prepend(...HandlerFunc) Float64Validator
	Append(...HandlerFunc) Float64Validator
	lib.Helper
}

//Implementation of Validator interface for float64 primitives.
type float64Number struct {
	s *float64
	lib.Helper
}

func (s *float64Number) IsEmpty() bool {
	if *s.s == 0 {
		return true
	}
	return false
}

func (s *float64Number) Prepend(hf ...HandlerFunc) Float64Validator {
	hdl := append(HandlerFuncList(hf).ToHandlers(), s.GetHandlers()...)
	s.SetHandlers(hdl...)
	return s
}

func (s *float64Number) Append(hf ...HandlerFunc) Float64Validator {
	hdl := append(s.GetHandlers(), HandlerFuncList(hf).ToHandlers()...)
	s.SetHandlers(hdl...)
	return s
}
