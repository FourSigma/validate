package funcs

import (
	"context"
	"fmt"

	"github.com/FourSigma/validate/types/integer"
)

type int64Funcs struct{}

func (n *int64Funcs) Max(max int64) integer.HandlerFunc {
	return func(ctx context.Context, i *int64) error {
		if *i > max {

			return fmt.Errorf("Max(%d)::Integer(%d) -- integer is greater than max", max, *i)
		}

		return nil
	}
}

func (n *int64Funcs) Min(min int64) integer.HandlerFunc {
	return func(ctx context.Context, i *int64) error {
		if *i < min {

			return fmt.Errorf("Min(%d)::Integer(%d) -- integer is less than min", min, *i)
		}

		return nil
	}
}
