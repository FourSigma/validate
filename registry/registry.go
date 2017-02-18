package registry

import (
	"context"

	"github.com/FourSigma/validate/lib"
)

type ValidationTypeId string

var valdationRegistry = map[ValidationTypeId]func(context.Context) error{}

func RegisterValidator(id ValidationTypeId, fn lib.Checker) {

}
