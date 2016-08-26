package validate

import (
	"errors"
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
		return errors.New("Incorrect greeting!")
	})
	list = append(list, list...)
	st := "Hello"
	err := str.String(&st, list...).Append().Check()
	fmt.Println(err)
}
