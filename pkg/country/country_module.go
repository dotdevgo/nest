package country

import (
	"dotdev/nest/pkg/nest"

	cn "github.com/biter777/countries"
	"github.com/goava/di"
	"gorm.io/gorm"
)

// New godoc
func New() di.Option {
	return di.Options(
		di.Invoke(func(db *gorm.DB) error {
			return db.AutoMigrate(&Country{})
		}),
		di.Provide(func() *CountryModule {
			return &CountryModule{}
		}, di.As(new(nest.ContainerModule))),
	)
}

// CountryModule godoc
type CountryModule struct {
	nest.ContainerModule
}

// Boot godoc
func (CountryModule) Boot(w *nest.Kernel) error {
	var db *gorm.DB
	w.ResolveFn(&db)

	var countries []Country
	if err := db.Model(&Country{}).Find(&countries).Error; err != nil {
		return err
	}

	w.Logger.Infof("[App] Countries: %v", len(countries))

	if 0 == len(countries) {
		for _, c := range cn.All() {
			var cntry Country
			cntry.Code = c.Info().Alpha2
			cntry.Name = c.Info().Name
			countries = append(countries, cntry)

			w.Logger.Printf("Insert Country: %s %s", cntry.Code, cntry.Name)
		}

		if err := db.CreateInBatches(&countries, 64).Error; err != nil {
			return err
		}
	}

	return nil
}
