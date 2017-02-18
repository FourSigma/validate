package validate

import (
	"context"
	"testing"

	"github.com/FourSigma/validate/types/str"
)

var myst = "hello"

var list = []str.HandlerFunc{str.MaxLen(10), str.Contains(myst), str.Contains(myst), str.Contains(myst), str.Contains(myst)}

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
