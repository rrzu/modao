package modao

import "gorm.io/gorm"

// IGormDb 获取数据库查询对象
type IGormDb interface {
	Db(obj ...*gorm.DB) *gorm.DB
}

// ---------------------------
// ----      BaseDao      ----
// ---------------------------

type BaseDao[T ITbl] struct {
	db  *gorm.DB
	mod IMod[T]
}

func NewBaseDao[T ITbl](m IMod[T], withDebug bool) *BaseDao[T] {
	d := new(BaseDao[T])
	d.db = GetGormDb(m.Table().ConnectInfo(), withDebug)
	d.mod = m
	return d
}

func (d *BaseDao[T]) Db(obj ...*gorm.DB) *gorm.DB {
	return d.db
}

func (d *BaseDao[T]) Qry() *gorm.DB {
	return d.Db().Table(d.Mod().Table().QueryTableName())
}

func (d *BaseDao[T]) Mod() IMod[T] {
	return d.mod
}
