package validate

import (
	"errors"
	"fmt"
	"testing"

	"github.com/FourSigma/validate/str"
)

var list []str.HandlerFunc

func TestCheckStringAPI(t *testing.T) {
	list = append(list, func(s *string) error {
		if *s == "HELLO" {
			return nil
		}
		return errors.New("Incorrect greeting!")
	})
	err := String("HELLO").Add(list...).Check()
	fmt.Println(err)
}
