package modao

type ClickhouseBaseDao struct {
	*BaseDao[*ClickhouseTbl]
}

func NewClickhouseBaseDao(mod IMod[*ClickhouseTbl], withDebug bool) *ClickhouseBaseDao {
	return &ClickhouseBaseDao{
		NewBaseDao(mod, withDebug),
	}
}

func (ch *ClickhouseBaseDao) Mod() IMod[*ClickhouseTbl] {
	return ch.mod.(IMod[*ClickhouseTbl])
}
