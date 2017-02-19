package bytes

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

type HandlerFuncList []HandlerFunc

func (s HandlerFuncList) ToHandlers() []lib.Handler {
	rs := make([]lib.Handler, len(s))
	for i, v := range s {
		rs[i] = lib.Handler(v)
	}
	return rs
}

func NewBytesValidator(s []byte) *bytes {
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
	s []byte
	lib.Helper
}

func (s *bytes) IsEmpty() bool {
	if len(s.s) == 0 {
		return true
	}
	return false
}

func (s *bytes) Prepend(hf ...HandlerFunc) BytesValidator {
	hdl := append(HandlerFuncList(hf).ToHandlers(), s.GetHandlers()...)
	s.SetHandlers(hdl...)
	return s
}

func (s *bytes) Append(hf ...HandlerFunc) BytesValidator {
	hdl := append(s.GetHandlers(), HandlerFuncList(hf).ToHandlers()...)
	s.SetHandlers(hdl...)
	return s
}
