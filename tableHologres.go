package modao

type (
	HologresDatabaseName DatabaseName // hologres 数据库名
	HologresPatternName  PatternName  // hologres 模式名
	HologresTableName    TableName    // hologres 表名

	// HologresTbl hologres 表信息
	HologresTbl struct {
		ConnectInformation ConnectInfo          // hologres 连接信息
		DatabaseName       HologresDatabaseName // hologres 数据库名
		PatternName        HologresPatternName  // hologres 模式名
		TableName          HologresTableName    // hologres 表名
	}
)

func (ch *HologresTbl) ConnectInfo() ConnectInfo {
	return ch.ConnectInformation
}

func (ch *HologresTbl) FullTableName() string {
	return string(ch.DatabaseName) + "." + string(ch.PatternName) + "." + string(ch.TableName)
}

func (ch *HologresTbl) QueryTableName() string {
	return string(ch.PatternName) + "." + string(ch.TableName)
}
