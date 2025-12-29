package modao

import (
	"sync"

	"gorm.io/gorm"
)

// gorm连接对象容器
var gormConnectContainer sync.Map

// RegisterGormDb 注册gorm连接对象
func RegisterGormDb(connectionInfo ConnectInfo, db *gorm.DB) {
	if GetGormDb(connectionInfo, false) != nil {
		return
	}

	gormConnectContainer.Store(key(connectionInfo), db)
}

// ModifyGormDb 修改gorm连接对象
func ModifyGormDb(connectionInfo ConnectInfo, db *gorm.DB) {
	gormConnectContainer.Store(key(connectionInfo), db)
}

// GetGormDb 获取gorm连接对象
func GetGormDb(connectionInfo ConnectInfo, withDebug bool) (gormDB *gorm.DB) {
	defer func() {
		if withDebug && gormDB != nil {
			gormDB = gormDB.Debug()
		}
	}()

	if db, ok := gormConnectContainer.Load(key(connectionInfo)); ok {
		return db.(*gorm.DB)
	}
	return nil
}

func key(connectionInfo ConnectInfo) string {
	return string(connectionInfo.ConnectName)
}
