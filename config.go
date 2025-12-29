package modao

import "gorm.io/gorm"

type Config struct {
	debugKey DebugKey
	gormDbs  map[ConnectInfo]*gorm.DB
}

func (c *Config) Init() {
	RegisterDebugKey(debugKey)
	for name, db := range c.gormDbs {
		RegisterGormDb(name, db)
	}
}

func (c *Config) SetDebugKey(key DebugKey) *Config {
	c.debugKey = key
	return c
}

func (c *Config) SetGormDb(connectionInfo ConnectInfo, db *gorm.DB) *Config {
	if c.gormDbs == nil {
		c.gormDbs = make(map[ConnectInfo]*gorm.DB)
	}

	c.gormDbs[connectionInfo] = db
	return c
}
