package validate

import (
	"bytes"
	"testing"
)

import "github.com/FourSigma/validate/str"

var CheckStringAPITests = []struct {
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

func TestCheckStringAPI(t *testing.T) {
	for _, v := range CheckStringAPITests {
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

var TransformStringAPITests = []struct {
	id         string
	s          string
	expected   string
	h          []str.TransHandler
	didpass    bool
	shouldpass bool
}{

	{
		id:         "Empty Handler Test -- Pass",
		s:          "hello",
		expected:   "hello",
		h:          []str.TransHandler{},
		shouldpass: true,
	},
	{
		id:         "Single Handler Test -- Pass",
		s:          "hello",
		expected:   "HELLO",
		h:          []str.TransHandler{ToUpper},
		shouldpass: true,
	},
	{
		id:         "Multiple Handlers Test -- Pass",
		s:          "hEllo",
		expected:   "HELLO",
		h:          []str.TransHandler{ToUpper, ToLower, ToUpper},
		shouldpass: true,
	},
	{
		id:         "Multiple Handlers Test -- Pass",
		s:          "hEllo",
		expected:   "HELLO",
		h:          []str.TransHandler{ToLower, ToLower, ToUpper},
		shouldpass: true,
	},
	{
		id:         "Multiple Handlers Test -- Fail",
		s:          "Hello",
		expected:   "Hello",
		h:          []str.TransHandler{ToUpper},
		shouldpass: false,
	},
	{
		id:         "Multiple Handlers Test -- Fail",
		s:          "Hello",
		expected:   "Hello",
		h:          []str.TransHandler{ToUpper, ToLower},
		shouldpass: false,
	},
}

func TestTransformStringAPI(t *testing.T) {
	for _, v := range TransformStringAPITests {
		err := Transform(
			String(&v.s).Transform(v.h...),
		)
		if err == nil {
			if v.expected == v.s {
				v.didpass = true
			}
		}
		if v.shouldpass != v.didpass {
			t.Errorf(v.id, "Failed")
		}
	}
}
