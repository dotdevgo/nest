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
		nest.NewExtension(func() *countryExt {
			return &countryExt{}
		}),
	)
}

// countryExt godoc
type countryExt struct {
	nest.Extension
}

// Boot godoc
func (p countryExt) OnStart(w *nest.Kernel) error {
	return w.Invoke(p.load)
}

// load Load countries
func (countryExt) load(w *nest.Kernel, db *gorm.DB) error {
	var countries []Country
	if err := db.Model(&Country{}).Find(&countries).Error; err != nil {
		return err
	}

	if len(countries) == 0 {
		for _, c := range cn.All() {
			var cntry Country
			cntry.Code = c.Info().Alpha2
			cntry.Name = c.Info().Name
			countries = append(countries, cntry)

			w.Logger.Printf("Add Country ==> %s %s", cntry.Code, cntry.Name)
		}

		if err := db.CreateInBatches(&countries, 64).Error; err != nil {
			return err
		}

		w.Logger.Infof("==> Loaded %v countries", len(cn.All()))
	} else {
		w.Logger.Infof("==> Loaded %v countries", len(countries))
	}

	return nil
}
