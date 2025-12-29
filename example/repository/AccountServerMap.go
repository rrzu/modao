package repository

import (
	"context"

	"github.com/rrzu/modao"
)

var accountServerMapDaoSingle modao.SingleDao[*AccountServerMapDao]

type AccountServerMapDao struct {
	*modao.ClickhouseBaseDao
}

func InstanceAccountServerMapDao(ctx context.Context) *AccountServerMapDao {
	return accountServerMapDaoSingle.Do(ctx, func(withDebug bool) *AccountServerMapDao {
		return &AccountServerMapDao{modao.NewClickhouseBaseDao(&AccountServerMapMod{}, withDebug)}
	})
}

// AccountServerMapMod 应用报名表
type AccountServerMapMod struct{}

func (t *AccountServerMapMod) TableName() string {
	return string(t.Table().TableName)
}

func (t *AccountServerMapMod) Table() *modao.ClickhouseTbl {
	return &modao.ClickhouseTbl{
		ConnectInformation: modao.ConnectInfo{
			ConnectName: "clickhouse-sz-dev",
		},
		DatabaseName: "shuzhi_prod",
		TableName:    "account_server_map",
	}
}
