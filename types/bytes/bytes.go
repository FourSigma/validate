package integer

import (
	"context"
	"errors"

	"github.com/FourSigma/validate/lib"
)

type HandlerFunc func(context.Context, []byte) error

func (s HandlerFunc) Handle(ctx context.Context, i interface{}) error {
	bytes, ok := i.([]byte)
	if !ok {
		return errors.New("::HandlerFunc::Cannot type assert, must be of type []bytes.")
	}
	return s(ctx, bytes)
}

func NewBytesValidator(s []bytes) *bytes {
	return &bytes{
		s:      s,
		Helper: lib.NewDefaultHelper(s, "Bytes"),
	}
}

type BytesValidator interface {
	Prepend(...HandlerFunc) BytesValidator
	Append(...HandlerFunc) BytesValidator
	lib.Helper
}

//Implementation of Validator interface for float64 primitives.
type bytes struct {
	s []bytes
	lib.Helper
}

func (s *bytes) IsEmpty() bool {
	if len(s.s) == 0 {
		return true
	}
	return false
}

func (s *bytes) Prepend(hf ...HandlerFunc) BytesValidator {
	hdl := append(s.toHandlers(hf...), s.GetHandlers()...)
	s.SetHandlers(hdl...)
	return s
}

func (s *bytes) Append(hf ...HandlerFunc) BytesValidator {
	hdl := append(s.GetHandlers(), s.toHandlers(hf...)...)
	s.SetHandlers(hdl...)
	return s
}
func (s *bytes) toHandlers(list ...HandlerFunc) []lib.Handler {
	rs := make([]lib.Handler, len(list))
	for i, v := range list {
		rs[i] = lib.Handler(v)
	}
	return rs
}
