package orm

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"

	"github.com/defval/di"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type OrmConfig struct {
	Entities []interface{}
	Gorm     *gorm.Config
}

// NewDsn Orm godoc
func NewDsn(dsn gorm.Dialector, config *OrmConfig) di.Option {
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

// New godoc
func New() di.Option {
	dsn := os.Getenv("DATABASE")
	if len(dsn) == 0 {
		return NewDsn(sqlite.Open("file::memory:?cache=shared"), nil)
	}

	return NewDsn(mysql.Open(os.Getenv("DATABASE")), nil)
}
