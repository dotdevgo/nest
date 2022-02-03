package hltv

import (
	"github.com/dotdevgo/gosymfony/examples/app/hltv/handler/controller"
	"github.com/dotdevgo/gosymfony/examples/app/hltv/orm/repository"
)

// NewController creates controller.
func NewTeamCtrl() *controller.TeamController {
	return &controller.TeamController{}
}

// NewTeamRepo godoc
func NewTeamRepo() *repository.TeamRepo {
	return &repository.TeamRepo{}
}
