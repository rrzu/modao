package modao

type (
	ClickhouseDatabaseName DatabaseName // clickhouse 数据库名
	ClickhouseTableName    TableName    // clickhouse 表名

	// ClickhouseTbl clickhouse 表信息
	ClickhouseTbl struct {
		ConnectInformation ConnectInfo            // clickhouse 连接信息
		DatabaseName       ClickhouseDatabaseName // clickhouse 数据库名
		TableName          ClickhouseTableName    // clickhouse 表名
	}
)

func (ch *ClickhouseTbl) ConnectInfo() ConnectInfo {
	return ch.ConnectInformation
}

func (ch *ClickhouseTbl) FullTableName() string {
	return string(ch.DatabaseName) + "." + string(ch.TableName)
}

func (ch *ClickhouseTbl) QueryTableName() string {
	return string(ch.TableName)
}
