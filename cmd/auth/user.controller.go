package authcmd

import (
	"dotdev/nest/pkg/nest"
	"dotdev/nest/pkg/nest/kernel"
	"dotdev/nest/pkg/user"
	"net/http"
)

const (
	RouteUser = "/api/user/:id"
)

type UserController struct {
	kernel.Controller
	nest.Config
	Crud *user.UserCrud
}

func (c UserController) Router(w *nest.Kernel) {
	w.GET(RouteUser, c.User)
}

// User godoc
func (c UserController) User(ctx nest.Context) error {
	var u user.User
	if err := c.Crud.Find(&u, ctx.Param("id")); err != nil {
		return nest.NewHTTPError(http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK, u)
}
