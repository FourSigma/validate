package validate

import "testing"

import "github.com/FourSigma/validate/str"

var APITests = []struct {
	id         string
	s          string
	h          []str.Handler
	didpass    bool
	shouldpass bool
}{

	{
		id:         "Empty Hanlder Test -- Pass",
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
}

func TestAPI(t *testing.T) {
	for _, v := range APITests {
		err := Check(
			String(v.s).Validate(v.h...),
		)
		if err == nil {
			v.didpass = true
		}

		if v.shouldpass != v.didpass {
			t.Errorf(v.id, "Failed")
		}
	}

}
