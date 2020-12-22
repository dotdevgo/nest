package provider

import (
	"github.com/goava/di"
	"gorm.io/gorm"
)

type OrmConfig struct {
	Entities []interface{}
	Gorm *gorm.Config
}

func Orm(dsn gorm.Dialector, config *OrmConfig) di.Option {
	if nil == config.Gorm {
		config.Gorm = &gorm.Config{}
	}

	db, err := gorm.Open(dsn, config.Gorm)

	if err != nil {
		panic(err)
	}

	for _, entity := range config.Entities {
		if err := db.AutoMigrate(entity); err != nil {
			panic(err)
		}
	}

	return di.Options(
		di.Provide(func() *gorm.DB {
			return db
		}),
	)
}
