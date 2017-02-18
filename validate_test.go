package validate

import (
	"context"
	"testing"

	"github.com/FourSigma/validate"
	"github.com/FourSigma/validate/types/str"
	. "github.com/FourSigma/validate/types/str/funcs"
)

var myst = "hello"

var list = []str.HandlerFunc{
	String.MaxLen(10),
	validate.OR(
		String.Contains(myst),
		String.Contains(myst),
	),
	validate.XOR(
		String.MinLen(10),
		String.MaxLen(100),
	),
	fn.String.Contains(myst),
	fn.String.Contains(myst),
}

func TestString(tst *testing.T) {

	st := "hello12789"
	ctx := context.Background()

	var ss str.StringValidator
	ss = String(st).Validate(list...).Prepend(list...)
	err := ss.Check(ctx)
	if err != nil {
		tst.Error(err)
	}

}
