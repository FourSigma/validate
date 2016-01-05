package validate

import (
	"bytes"
	"fmt"
	"testing"
)

import "github.com/FourSigma/validate/str"

var APITests = []struct {
	id         string
	s          string
	h          []str.Handler
	didpass    bool
	shouldpass bool
}{

	{
		id:         "Empty Handler Test -- Pass",
		s:          "Hello",
		h:          []str.Handler{},
		shouldpass: true,
	},
	{
		id:         "Single Handler Test -- Pass",
		s:          "Hello",
		h:          []str.Handler{str.MaxLen(5)},
		shouldpass: true,
	},
	{
		id:         "Single Handler Test -- Fail",
		s:          "Hello",
		h:          []str.Handler{str.MaxLen(2)},
		shouldpass: false,
	},
	{
		id:         "Multiple Handlers Test -- Pass",
		s:          "Hello",
		h:          []str.Handler{str.MaxLen(5), str.MaxLen(5)},
		shouldpass: true,
	},
	{
		id:         "Multiple Handlers Test -- Fail",
		s:          "Hello",
		h:          []str.Handler{str.MaxLen(2), str.MaxLen(5)},
		shouldpass: false,
	},
	{
		id:         "Multiple Handlers Test -- Fail",
		s:          "Hello",
		h:          []str.Handler{str.MaxLen(5), str.MaxLen(2)},
		shouldpass: false,
	},
}

func TestAPI(t *testing.T) {
	for _, v := range APITests {
		err := Check(
			String(&v.s).Validate(v.h...),
		)
		if err == nil {
			v.didpass = true
		}

		if v.shouldpass != v.didpass {
			t.Errorf(v.id, "Failed")
		}
	}

}

func ToUpper(b []byte) ([]byte, error) {
	return bytes.ToUpper(b), nil
}

func ToLower(b []byte) ([]byte, error) {
	return bytes.ToLower(b), nil
}
func TestTransformAPI(t *testing.T) {
	s := "Hello"
	String(&s).Transform(ToLower, ToUpper, ToLower).Transform()
	fmt.Println("NOW", s)

}
