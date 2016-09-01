package validate

import (
	"context"
	"fmt"
	"testing"

	"github.com/FourSigma/validate/types/str"
)

var list = []str.HandlerFunc{}

func TestCheckStringAPI(t *testing.T) {
	list = append(list, func(ctx context.Context, s *string) error {
		if *s == "HELLO" {
			return nil
		}

		fmt.Println("HERE")
		return nil
	})

	st := "hello"
	ctx := context.Background()
	ss := str.String(&st, list...).Append(list...)
	fmt.Println(ss.Check(ctx))

}
