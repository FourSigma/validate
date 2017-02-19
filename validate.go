package validate

import (
	"context"

	"github.com/FourSigma/validate/lib"
	"github.com/FourSigma/validate/lib/logic"
	"github.com/FourSigma/validate/types/bytes"
	"github.com/FourSigma/validate/types/float"
	"github.com/FourSigma/validate/types/integer"
	"github.com/FourSigma/validate/types/str"
)

type String string

func (s String) Validate(list ...str.HandlerFunc) str.StringValidator {
	return str.NewStringValidator((*string)(&s)).Append(list...)
}

type Bytes []byte

func (s Bytes) Validate(list ...bytes.HandlerFunc) bytes.BytesValidator {
	return bytes.NewBytesValidator(s).Append(list...)
}

type Int8 int8

func (s Int8) Validate(list ...integer.HandlerFunc) integer.Int64Validator {
	f := int64(s)
	return integer.NewInt64Validator(&f).Append(list...)
}

type Int16 int16

func (s Int16) Validate(list ...integer.HandlerFunc) integer.Int64Validator {
	f := int64(s)
	return integer.NewInt64Validator(&f).Append(list...)
}

type Int32 int32

func (s Int32) Validate(list ...integer.HandlerFunc) integer.Int64Validator {
	f := int64(s)
	return integer.NewInt64Validator(&f).Append(list...)
}

type Int64 int64

func (s Int64) Validate(list ...integer.HandlerFunc) integer.Int64Validator {
	return integer.NewInt64Validator((*int64)(&s)).Append(list...)
}

type Float32 float32

func (s Float32) Validate(list ...float.HandlerFunc) float.Float64Validator {
	f := float64(s)
	return float.NewFloat64Validator(&f).Append(list...)
}

type Float64 float64

func (s Float64) Validate(list ...float.HandlerFunc) float.Float64Validator {
	return float.NewFloat64Validator((*float64)(&s)).Append(list...)
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

func AND(sh ...lib.Handler) lib.Handler {
	return logic.NewAnd(sh...)
}

func OR(sh ...lib.Handler) lib.Handler {
	return logic.NewOr(sh...)
}
