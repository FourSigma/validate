package validate

import (
	"context"

	"github.com/FourSigma/validate/lib"
	"github.com/FourSigma/validate/lib/logic"
	"github.com/FourSigma/validate/types/str"
)

type ValidationTypeId string

var valdationRegistry = map[ValidationTypeId]func(context.Context) error{}

func RegisterValidator(id ValidationTypeId, fn lib.Checker) {

}

type String string

func (s String) Validate(list ...str.HandlerFunc) str.StringValidator {
	return str.NewStringValidator((*string)(&s)).Append(list...)
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
	return logic.NewAnd(sh...)
}

func Or(sh ...lib.Handler) lib.Handler {
	return logic.NewOr(sh...)
}
