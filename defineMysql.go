package modao

type MysqlBaseDao struct {
	*BaseDao[*MysqlTbl]
}

func NewMysqlBaseDao(mod IMod[*MysqlTbl], withDebug bool) *MysqlBaseDao {
	return &MysqlBaseDao{
		NewBaseDao(mod, withDebug),
	}
}

func (ch *MysqlBaseDao) Mod() IMod[*MysqlTbl] {
	return ch.mod.(IMod[*MysqlTbl])
}
