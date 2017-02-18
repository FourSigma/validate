package float

import (
	"context"
	"errors"

	"github.com/FourSigma/validate/lib"
)

type HandlerFunc func(context.Context, *float64) error

func (s HandlerFunc) Handle(ctx context.Context, i interface{}) error {
	integer64, ok := i.(*float64)
	if !ok {
		return errors.New("::HandlerFunc::Cannot type assert, must be of type *float64.")
	}
	return s(ctx, integer64)
}

func NewFloat64Validator(s *float64) *float64Integer {
	return &float64Integer{
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
type float64Integer struct {
	s *float64
	lib.Helper
}

func (s *float64Integer) IsEmpty() bool {
	if *s.s == 0 {
		return true
	}
	return false
}

func (s *float64Integer) Prepend(hf ...HandlerFunc) Float64Validator {
	hdl := append(s.toHandlers(hf...), s.GetHandlers()...)
	s.SetHandlers(hdl...)
	return s
}

func (s *float64Integer) Append(hf ...HandlerFunc) Float64Validator {
	hdl := append(s.GetHandlers(), s.toHandlers(hf...)...)
	s.SetHandlers(hdl...)
	return s
}
func (s *float64Integer) toHandlers(list ...HandlerFunc) []lib.Handler {
	rs := make([]lib.Handler, len(list))
	for i, v := range list {
		rs[i] = lib.Handler(v)
	}
	return rs
}
