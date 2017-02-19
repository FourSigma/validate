package logic

import (
	"context"

	liblogic "github.com/FourSigma/validate/lib/logic"
	"github.com/FourSigma/validate/types/str"
)

type stringLogic struct{}

func (n *stringLogic) ALL(list ...str.HandlerFunc) str.HandlerFunc {
	andFn := liblogic.NewAnd(str.HandlerFuncList(list).ToHandlers()...)

	return func(ctx context.Context, i *string) error {
		return andFn.Handle(ctx, i)
	}
}

func (n *stringLogic) OR(one, two str.HandlerFunc) str.HandlerFunc {
	orFn := liblogic.NewOr((str.HandlerFuncList{one, two}).ToHandlers()...)

	return func(ctx context.Context, i *string) error {
		return orFn.Handle(ctx, i)
	}
}

func (n *stringLogic) ANY(list ...str.HandlerFunc) str.HandlerFunc {
	orFn := liblogic.NewOr(str.HandlerFuncList(list).ToHandlers()...)

	return func(ctx context.Context, i *string) error {
		return orFn.Handle(ctx, i)
	}
}
