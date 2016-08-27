package validate

import (
	"fmt"
	"testing"

	"github.com/FourSigma/validate/types/str"
)

var list = []str.HandlerFunc{}

func TestCheckStringAPI(t *testing.T) {
	list = append(list, func(s *string) error {
		if *s == "HELLO" {
			return nil
		}

		fmt.Println("HERE")
		return nil
	})

	st := "hello"
	ss := str.String(&st, list...).Append(list...)
	fmt.Println(ss.Check())

}
