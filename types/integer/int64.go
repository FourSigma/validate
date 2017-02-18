package integer

import (
	"context"
	"errors"

	"github.com/FourSigma/validate/lib"
)

type HandlerFunc func(context.Context, *int64) error

func (s HandlerFunc) Handle(ctx context.Context, i interface{}) error {
	integer64, ok := i.(*int64)
	if !ok {
		return errors.New("::HandlerFunc::Cannot type assert, must be of type *int64.")
	}
	return s(ctx, integer64)
}

func NewInt64Validator(s *int64) *int64Integer {
	return &int64Integer{
		s:      s,
		Helper: lib.NewDefaultHelper(s, "Unsigned Number"),
	}
}

type Int64Validator interface {
	Prepend(...HandlerFunc) Int64Validator
	Append(...HandlerFunc) Int64Validator
	lib.Helper
}

//Implementation of Validator interface for int64 primitives.
type int64Integer struct {
	s *int64
	lib.Helper
}

func (s *int64Integer) IsEmpty() bool {
	if *s.s == 0 {
		return true
	}
	return false
}

func (s *int64Integer) Prepend(hf ...HandlerFunc) Int64Validator {
	hdl := append(s.toHandlers(hf...), s.GetHandlers()...)
	s.SetHandlers(hdl...)
	return s
}

func (s *int64Integer) Append(hf ...HandlerFunc) Int64Validator {
	hdl := append(s.GetHandlers(), s.toHandlers(hf...)...)
	s.SetHandlers(hdl...)
	return s
}
func (s *int64Integer) toHandlers(list ...HandlerFunc) []lib.Handler {
	rs := make([]lib.Handler, len(list))
	for i, v := range list {
		rs[i] = lib.Handler(v)
	}
	return rs
}
