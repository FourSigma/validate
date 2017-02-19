package logic

import (
	"context"

	liblogic "github.com/FourSigma/validate/lib/logic"
	"github.com/FourSigma/validate/types/float"
)

type float64Logic struct{}

func (n *float64Logic) ALL(list ...float.HandlerFunc) float.HandlerFunc {
	andFn := liblogic.NewAnd(float.HandlerFuncList(list).ToHandlers()...)

	return func(ctx context.Context, i *float64) error {
		return andFn.Handle(ctx, i)
	}
}

func (n *float64Logic) OR(one, two float.HandlerFunc) float.HandlerFunc {
	orFn := liblogic.NewOr((float.HandlerFuncList{one, two}).ToHandlers()...)

	return func(ctx context.Context, i *float64) error {
		return orFn.Handle(ctx, i)
	}
}

func (n *float64Logic) ANY(list ...float.HandlerFunc) float.HandlerFunc {
	orFn := liblogic.NewOr(float.HandlerFuncList(list).ToHandlers()...)

	return func(ctx context.Context, i *float64) error {
		return orFn.Handle(ctx, i)
	}
}
