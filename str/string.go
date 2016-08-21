package str

import (
	"errors"
	"log"

	"github.com/FourSigma/validate"
	. "github.com/FourSigma/validate/misc/err"
)

type Handler func(*string) error

func (s Handler) Handle(i interface{}) error {
	str, ok := i.(*string)
	if !ok {
		return errors.New("Something went wrong")
	}

	return s(str)
}

//Implementation of Validator interface for string primitives.
type str struct {
	s        *string
	h        []StringHandler
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

func (s *str) Add(b ...validate.Handler) *str {
	s.h = append(getStringHandler(b...), s.h...)
	return s
}

func (s *str) Finally(a ...Handler) *str {
	s.h = append(s.h, getStringHandler(a...)...)
	return s
}

func (s *str) Name(name string) *str {
	s.meta = name
	return s
}

func (s *str) Required() *str {
	s.required = true
	return s
}

//Ignores non-string Handlers in the slice.
func getStringHandler(list ...Handler) (rs []StringHandler) {
	for _, v := range list {
		switch x := v.(type) {
		case StringHandler:
			rs = append(rs, x)
		default:
			continue
		}
	}

	return
}
