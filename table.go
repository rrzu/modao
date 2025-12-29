package modao

import "gorm.io/gorm/schema"

type (
	DatabaseName string // 数据库名
	TableName    string // 表名

	PatternName string // 模式名

	ProjectName string // 项目名
	SchemaName  string // schema名
)

type IMod[T ITbl] interface {
	schema.Tabler // 表名
	Table() T     // 表信息
}

type ITbl interface {
	ConnectInfo() ConnectInfo // 连接信息
	FullTableName() string    // 全表名
	QueryTableName() string   // 查询表名
}
