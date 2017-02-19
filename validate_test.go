package validate

import (
	"context"
	"testing"

	fn "github.com/FourSigma/validate/funcs"
	"github.com/FourSigma/validate/types/str"
)

var myst = "hello"

var list = []str.HandlerFunc{
	fn.String.MaxLen(10),
	OR(
		fn.String.Contains(myst),
		fn.String.Contains(myst),
	),
	fn.String.Contains(myst),
	fn.String.Contains(myst),
}

func TestString(tst *testing.T) {

	st := "hello12789fffffffff"
	ctx := context.Background()

	var ss str.StringValidator
	ss = String(st).Validate(list...).Prepend(list...)
	err := ss.Check(ctx)
	if err != nil {
		tst.Error(err)
	}

}
