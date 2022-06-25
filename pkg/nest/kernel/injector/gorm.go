package injector

import (
	"log"
	"os"

	"github.com/goava/di"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type OrmConfig struct {
	Entities []interface{}
	Gorm     *gorm.Config
}

// Orm godoc
func OrmWithDsn(dsn gorm.Dialector, config *OrmConfig) di.Option {
	if nil == config {
		config = &OrmConfig{}
	}

	if nil == config.Gorm {
		config.Gorm = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		}
	}

	db, err := gorm.Open(dsn, config.Gorm)

	if err != nil {
		log.Fatal(err)
	}

	for _, entity := range config.Entities {
		if err := db.AutoMigrate(entity); err != nil {
			log.Fatal(err)
		}
	}

	return di.Options(
		di.Provide(func() *gorm.DB {
			return db
		}),
	)
}

func Orm() di.Option {
	return OrmWithDsn(mysql.Open(os.Getenv("DATABASE")), nil)
}
