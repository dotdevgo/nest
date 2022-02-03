package hltv

import (
	"dotdev.io/examples/app/hltv/handler/controller"
	"dotdev.io/examples/app/hltv/orm/repository"
)

// NewController creates controller.
func NewTeamCtrl() *controller.TeamController {
	return &controller.TeamController{}
}

// NewTeamRepo godoc
func NewTeamRepo() *repository.TeamRepo {
	return &repository.TeamRepo{}
}
