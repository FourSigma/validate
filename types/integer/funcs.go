package integer

import (
	"context"
	"fmt"
)

func Max(max int64) HandlerFunc {
	return func(ctx context.Context, i *int64) error {
		if *i > max {

			return fmt.Errorf("Max(%d)::Integer(%d) -- integer is greater than max", max, *i)
		}

		return nil
	}
}

func Min(min int64) HandlerFunc {
	return func(ctx context.Context, i *int64) error {
		if *i < min {

			return fmt.Errorf("Min(%d)::Integer(%d) -- integer is less than min", min, *i)
		}

		return nil
	}
}
