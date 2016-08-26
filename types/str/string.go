package str

import (
	"errors"
	"fmt"

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

func String(s *string, handlers ...HandlerFunc) *str {
	return &str{s: s, h: handlers}
}

//Implementation of Validator interface for string primitives.
type str struct {
	s        *string
	h        []HandlerFunc
	required bool
	meta     string
}

func (s *str) IsEmpty() bool {
	if *s.s == "" {
		return true
	}
	return false
}

func (s *str) Check() error {
	return lib.DefaultCheck(s)
}

func (s *str) Value() interface{} {
	return s.s
}
func (s *str) Prepend(b ...HandlerFunc) *str {
	s.h = append(b, s.h...)
	return s
}

func (s *str) Append(a ...HandlerFunc) *str {
	s.h = append(s.h, a...)
	return s
}

func (s *str) Name(name string) lib.Validator {
	s.meta = name
	return s
}
func (s *str) String() string {
	if s.meta == "" {
		s.meta = "NO_NAME_ASSIGNED - Use Name()"
	}
	return fmt.Sprintf("%s::%s", "String", s.meta)
}
func (s *str) GetHandlers() []lib.Handler {
	rs := make([]lib.Handler, len(s.h))
	for i, v := range s.h {
		rs[i] = lib.Handler(v)
	}
	return rs
}
func (s *str) Required() lib.Validator {
	s.required = true
	return s
}
func (s *str) IsRequired() bool {
	return s.required
}
