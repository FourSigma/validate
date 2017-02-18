package logic

import (
	"context"
	"fmt"

	"github.com/FourSigma/validate/lib"
)

func NewAnd(sh ...lib.Handler) lib.Handler {
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

func NewOr(sh ...lib.Handler) lib.Handler {
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
