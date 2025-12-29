package modao

import (
	"context"

	"github.com/spf13/cast"

	"github.com/rrzu/modao/common"
)

// SingleDao 单例对象
type SingleDao[T IGormDb] struct {
	obj      common.Single[T] // 默认
	objDebug common.Single[T] // 调试对象
}

type OptDao struct {
	WithDebug bool
}

func (o *SingleDao[T]) Do(ctx context.Context, f func(withDebug bool) T) (ret T) {
	if debugKey == "" {
		return
	}

	var withDebug = cast.ToBool(ctx.Value(debugKey))

	if withDebug {
		return o.objDebug.Do(func() T {
			dao := f(withDebug)
			if _, ok := dao.Db().Logger.(IModaoLogger); ok {
				dao.Db().Logger.(IModaoLogger).SetOption(OptModaoLogger{
					OnDebug: withDebug,
				})
			}
			return dao
		})
	} else {
		return o.obj.Do(func() T { return f(withDebug) })
	}
}
