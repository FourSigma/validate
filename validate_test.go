package validate

import (
	"context"
	"fmt"
	"testing"

	"github.com/FourSigma/validate/lib"
	"github.com/FourSigma/validate/types/str"
)

var list = []lib.Handler{}

func TestCheckStringAPI(t *testing.T) {
	list = append(list, str.MaxLen(10), str.Contains("h"))

	st := "hello"
	ctx := context.Background()
	ss := str.String(&st, list...).Append(list...)
	fmt.Println(ss.Check(ctx))

}
