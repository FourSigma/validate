package funcs

import (
	"context"
	"fmt"

	"github.com/FourSigma/validate/types/float"
)

type float64Funcs struct{}

func (n *float64Funcs) Max(max float64) float.HandlerFunc {
	return func(ctx context.Context, i *float64) error {
		if *i > max {

			return fmt.Errorf("Max(%d)::Integer(%d) -- integer is greater than max", max, *i)
		}

		return nil
	}
}

func (n *float64Funcs) Min(min float64) float.HandlerFunc {
	return func(ctx context.Context, i *float64) error {
		if *i < min {

			return fmt.Errorf("Min(%d)::Integer(%d) -- integer is less than min", min, *i)
		}

		return nil
	}
}
