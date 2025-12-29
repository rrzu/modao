package common

import (
	"sync"
)

// Single 单例对象
type Single[T any] struct {
	isset     bool // 是否已设置
	obj       T    // 实例对象
	sync.Once      // 确保单例对象只被创建一次
}

func (o *Single[T]) Do(f func() T) T {
	if o.isset {
		return o.obj
	}

	o.Once.Do(func() {
		o.isset = true
		o.obj = f()
	})
	return o.obj
}
