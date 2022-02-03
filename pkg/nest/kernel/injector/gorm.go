package injector

import (
	"github.com/goava/di"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type OrmConfig struct {
	Entities []interface{}
	Gorm *gorm.Config
}

// Orm godoc
func Orm(dsn gorm.Dialector, config *OrmConfig) di.Option {
	if nil == config {
		config = &OrmConfig{}
	}

	if nil == config.Gorm {
		config.Gorm = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
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
