package hltv

import (
	"dotdev.io/internal/app/hltv/handler/controller"
	"dotdev.io/internal/app/hltv/orm/entity"
	"dotdev.io/pkg/nest"
	"github.com/goava/di"
	"gorm.io/gorm"
)

// Router godoc
func Router(api nest.ApiGroup, team *controller.TeamController) {
	e := api.(*nest.Group)
	e.GET("/hltv/teams", team.List)
	e.POST("/hltv/teams", team.Save)
	e.PUT("/hltv/teams/:id", team.Save)
}

// Provider godoc
func Provider() di.Option {
	return di.Options(
		di.Provide(NewTeamCtrl),
		di.Provide(NewTeamRepo),
		di.Invoke(func(db *gorm.DB) error {
			return db.AutoMigrate(&entity.Team{})
		}),
	)
}
