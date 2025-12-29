package main

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"github.com/rrzu/modao"
	"github.com/rrzu/modao/example/repository"
)

func main() {
	(&modao.Config{}).
		SetDebugKey("onDebug").
		SetGormDb(modao.ConnectInfo{ConnectName: "mysql-sz-dev"}, &gorm.DB{}).
		SetGormDb(modao.ConnectInfo{ConnectName: "hologres-sz-dev"}, &gorm.DB{}).
		SetGormDb(modao.ConnectInfo{ConnectName: "clickhouse-sz-dev"}, &gorm.DB{}).
		SetGormDb(modao.ConnectInfo{ConnectName: "maxcompute-sz-dev"}, &gorm.DB{}).
		Init()

	ctx := context.Background()

	dao := repository.InstanceAccountServerMapDao(ctx)
	mod := dao.Mod()
	mod.Table()

	dao2 := repository.InstanceAccountDao(ctx)
	dao2.Db()
	mod2 := dao2.Mod()
	mod2.Table()

	dao3 := repository.InstanceDimOrderTypeDao(ctx)
	dao3.Qry()
	mod3 := dao3.Mod()
	mod3.Table()

	sql := dao3.Db().ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Select("*").Find(&map[string]interface{}{})
	})

	fmt.Println(sql)
}
