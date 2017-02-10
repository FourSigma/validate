package validate

import (
	"context"
	"testing"

	"github.com/FourSigma/validate/types/str"
)

var myst = "hello"

var list = []str.HandlerFunc{str.MaxLen(10), str.Contains(myst), str.Contains(myst), str.Contains(myst), str.Contains(myst)}

func BenchmarkCheckStringAPI(b *testing.B) {

	st := "hello"
	ctx := context.Background()

	var ss str.StringValidator
	for i := 0; i < b.N; i++ {
		ss = String(st).Validate(list...).Prepend(list...)
		//fmt.Println(ss.Check(ctx))
		ss.Check(ctx)
	}

}
