package validate

import (
	"fmt"

	"github.com/FourSigma/validate/lib"
)

func Check(c ...lib.Checker) error {
	for _, v := range c {
		err := v.Check()
		if err != nil {
			return err
		}
	}
	return nil
}

func And(sh ...lib.Handler) lib.Handler {
	return and(sh)
}

type and []lib.Handler

func (a and) Handle(i interface{}) error {
	for _, v := range a {
		err := v.Handle(i)
		if err != nil {
			return fmt.Errorf("AND error for type %v", i)
		}
	}

	return nil
}

func Or(sh ...lib.Handler) lib.Handler {
	return or(sh)
}

type or []lib.Handler

func (a or) Handle(i interface{}) error {
	for _, v := range a {
		err := v.Handle(i)
		if err == nil {
			return nil
		}
	}

	return fmt.Errorf("OR error for type %v", i)
}
