package logic

import (
	"context"

	liblogic "github.com/FourSigma/validate/lib/logic"
	"github.com/FourSigma/validate/types/integer"
)

type int64Logic struct{}

func (n *int64Logic) ALL(list ...integer.HandlerFunc) integer.HandlerFunc {
	andFn := liblogic.NewAnd(integer.HandlerFuncList(list).ToHandlers()...)

	return func(ctx context.Context, i *int64) error {
		return andFn.Handle(ctx, i)
	}
}

func (n *int64Logic) OR(one, two integer.HandlerFunc) integer.HandlerFunc {
	orFn := liblogic.NewOr((integer.HandlerFuncList{one, two}).ToHandlers()...)

	return func(ctx context.Context, i *int64) error {
		return orFn.Handle(ctx, i)
	}
}

func (n *int64Logic) ANY(list ...integer.HandlerFunc) integer.HandlerFunc {
	orFn := liblogic.NewOr(integer.HandlerFuncList(list).ToHandlers()...)

	return func(ctx context.Context, i *int64) error {
		return orFn.Handle(ctx, i)
	}
}
