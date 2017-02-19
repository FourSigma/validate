package logic

import (
	"context"

	liblogic "github.com/FourSigma/validate/lib/logic"

	"github.com/FourSigma/validate/types/bytes"
)

type bytesLogic struct{}

func (n *bytesLogic) ALL(list ...bytes.HandlerFunc) bytes.HandlerFunc {
	andFn := liblogic.NewAnd(bytes.HandlerFuncList(list).ToHandlers()...)
	return func(ctx context.Context, i []byte) error {
		return andFn.Handle(ctx, i)
	}
}

func (n *bytesLogic) OR(one, two bytes.HandlerFunc) bytes.HandlerFunc {
	orFn := liblogic.NewOr((bytes.HandlerFuncList{one, two}).ToHandlers()...)

	return func(ctx context.Context, i []byte) error {
		return orFn.Handle(ctx, i)
	}
}

func (n *bytesLogic) ANY(list ...bytes.HandlerFunc) bytes.HandlerFunc {
	orFn := liblogic.NewOr(bytes.HandlerFuncList(list).ToHandlers()...)

	return func(ctx context.Context, i []byte) error {
		return orFn.Handle(ctx, i)
	}
}
