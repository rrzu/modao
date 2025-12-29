package repository

import (
	"context"

	"github.com/rrzu/modao"
)

var accountDaoSingle modao.SingleDao[*AccountDao]

type AccountDao struct {
	*modao.MysqlBaseDao
}

func InstanceAccountDao(ctx context.Context) *AccountDao {
	return accountDaoSingle.Do(ctx, func(withDebug bool) *AccountDao {
		return &AccountDao{modao.NewMysqlBaseDao(&AccountMod{}, withDebug)}
	})
}

// AccountMod 账号表
type AccountMod struct {
	Id                int    `json:"id" gorm:"column:id"`                                   //
	Phone             string `json:"phone" gorm:"column:phone"`                             // 手机
	Name              string `json:"name" gorm:"column:name"`                               // 账号名称
	Password          string `json:"password" gorm:"column:password"`                       // 密码
	Role              int8   `json:"role" gorm:"column:role"`                               // 角色：-1 : 超管，0 : 普通，1：普通运营
	SysUserId         int    `json:"sys_user_id" gorm:"column:sys_user_id"`                 // 系统user_id（对应zulin.user.id）
	CreatedTime       int    `json:"created_time" gorm:"column:created_time"`               // 创建时间
	UpdatedTime       int    `json:"updated_time" gorm:"column:updated_time"`               // 修改时间
	PasswordUpdatedAt int    `json:"password_updated_at" gorm:"column:password_updated_at"` // 密码更新时间
}

func (t *AccountMod) TableName() string {
	return string(t.Table().TableName)
}

func (t *AccountMod) Table() *modao.MysqlTbl {
	return &modao.MysqlTbl{
		ConnectInformation: modao.ConnectInfo{
			ConnectName: "mysql-sz-dev",
		},
		DatabaseName: "shuzhi",
		TableName:    "account",
	}
}
