package modao

type (
	MysqlDatabaseName DatabaseName // mysql 数据库名
	MysqlTableName    TableName    // mysql 表名

	// MysqlTbl mysql 表信息
	MysqlTbl struct {
		ConnectInformation ConnectInfo       // mysql 连接信息
		DatabaseName       MysqlDatabaseName // mysql 数据库名
		TableName          MysqlTableName    // mysql 表名
	}
)

func (ch *MysqlTbl) ConnectInfo() ConnectInfo {
	return ch.ConnectInformation
}

func (ch *MysqlTbl) FullTableName() string {
	return string(ch.DatabaseName) + "." + string(ch.TableName)
}

func (ch *MysqlTbl) QueryTableName() string {
	return string(ch.TableName)
}
