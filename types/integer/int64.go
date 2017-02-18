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
		return errors.New("int64::HandlerFunc::Cannot type assert, must be of type *int64.")
	}
	return s(ctx, integer64)
}

func ToHandlers(list ...HandlerFunc) []lib.Handler {
	rs := make([]lib.Handler, len(list))
	for i, v := range list {
		rs[i] = lib.Handler(v)
	}
	return rs
}

func Int64(i *int64, hf ...HandlerFunc) *integer64 {
	ss := &integer64{i: i}
	ss.Helper = lib.NewDefaultHelper(i, "Int64", ToHandlers(hf...)...)
	return ss
}

//Implementation of Validator interface for int64 primitives.
type integer64 struct {
	i *int64
	lib.Helper
}

func (s *integer64) IsEmpty() bool {
	if *s.i == 0 {
		return true
	}
	return false
}

func (s *integer64) Prepend(hf ...HandlerFunc) *integer64 {
	hdl := append(ToHandlers(hf...), s.GetHandlers()...)
	s.SetHandlers(hdl...)
	return s
}

func (s *integer64) Append(hf ...HandlerFunc) *integer64 {
	hdl := append(s.GetHandlers(), ToHandlers(hf...)...)
	s.SetHandlers(hdl...)
	return s
}
