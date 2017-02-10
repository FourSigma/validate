package validate

import (
	"context"
	"fmt"

	"github.com/FourSigma/validate/lib"
	"github.com/FourSigma/validate/types/str"
)

type String string

func (s String) Validate(list ...str.HandlerFunc) str.StringValidator {
	return str.NewStringValidator((*string)(&s), list...)
}

func Check(ctx context.Context, c ...lib.Checker) error {
	for _, v := range c {
		err := v.Check(ctx)
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

func (a and) Handle(ctx context.Context, i interface{}) error {
	for _, v := range a {
		err := v.Handle(ctx, i)
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

func (a or) Handle(ctx context.Context, i interface{}) error {
	for _, v := range a {
		err := v.Handle(ctx, i)
		if err == nil {
			return nil
		}
	}
	return fmt.Errorf("OR error for type %v", i)
}
