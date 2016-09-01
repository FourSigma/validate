package str

import (
	"context"
	"fmt"
)

type StringError struct {
	input    interface{}
	funcName string
}

func (s StringError) Error() string {
	return fmt.Sprintf("StringError::%s %s", s.funcName, s.input)
}

func NewStringError(funcName string, input interface{}, list ...interface{}) {
	return
}

func MinLen(min int) HandlerFunc {
	return func(ctx context.Context, s *string) error {
		if len(*s) < min {
			return fmt.Errorf("MinLen: string %s does meet min length %d", *s, min)
		}

		return nil
	}
}
func MaxLen(max int) HandlerFunc {
	return func(ctx context.Context, s *string) error {
		if len(*s) > max {
			return fmt.Errorf("MaxLen: string %s has exceeded max len %d", *s, max)
		}
		return nil

	}
}
