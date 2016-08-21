package str

import (
	"errors"
	"log"

	"github.com/FourSigma/validate/lib"
	. "github.com/FourSigma/validate/misc/err"
)

type HandlerFunc func(*string) error

func (s HandlerFunc) Handle(i interface{}) error {
	str, ok := i.(*string)
	if !ok {
		return errors.New("Something went wrong")
	}

	return s(str)
}

func NewStringValidator(s *string) *str {
	return &str{s: s}
}

//Implementation of Validator interface for string primitives.
type str struct {
	s        *string
	h        []HandlerFunc
	required bool

	meta string
	log  *log.Logger
}

func (s *str) Empty() bool {
	if *s.s == "" {
		return true
	}

	return false
}

func (s *str) Check() error {
	if s.Empty() && s.required == true {
		return errors.New("String value required")
	}

	if s.Empty() {
		return nil
	}

	for _, v := range s.h {
		err := v(s.s)
		if err != nil {
			if _, ok := err.(TerminateLoop); ok {
				break
			}
			return err
		}
	}
	return nil
}

func (s *str) Add(b ...lib.Handler) lib.Validator {
	s.h = append(getHandlerFunc(b...), s.h...)
	return s
}

func (s *str) Finally(a ...lib.Handler) lib.Validator {
	s.h = append(s.h, getHandlerFunc(a...)...)
	return s
}

func (s *str) Name(name string) lib.Validator {
	s.meta = name
	return s
}

func (s *str) Required() lib.Validator {
	s.required = true
	return s
}

//Ignores non-string Handlers in the slice.
func getHandlerFunc(list ...lib.Handler) (rs []HandlerFunc) {
	for _, v := range list {
		switch x := v.(type) {
		case HandlerFunc:
			rs = append(rs, x)
		default:
			continue
		}
	}

	return
}
