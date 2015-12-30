package validate

import "testing"
import "github.com/FourSigma/validate/str"

var StringTests = []struct {
	id         string
	s          string
	h          []str.Handler
	didpass    bool
	shouldpass bool
}{
	{
		id:         "API_Test",
		s:          "Hello",
		h:          []str.Handler{},
		shouldpass: true,
	},
}

func TestString(t *testing.T) {
	for _, v := range StringTests {
		err := Check(
			String(v.s).Validate(v.h...),
		)
		if err != nil {
			v.didpass = false
			t.Log(v.id, "Failed")
		}
	}

}
