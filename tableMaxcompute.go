package modao

type (
	MaxComputeProjectName ProjectName // maxcompute 项目名
	MaxComputeSchemaName  SchemaName  // maxcompute schema名
	MaxComputeTableName   TableName   // maxcompute 表名
)

type (
	// MaxCompute maxcompute 表信息
	MaxCompute struct {
		ProjectName MaxComputeProjectName // maxcompute 项目名
		SchemaName  MaxComputeSchemaName  // maxcompute schema名
		TableName   MaxComputeTableName   // maxcompute 表名
	}
)
