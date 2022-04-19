package authcmd

import (
	"net/http"

	"github.com/dotdevgo/nest/pkg/auth"
	"github.com/dotdevgo/nest/pkg/nest"
	"github.com/dotdevgo/nest/pkg/nest/kernel"
	"github.com/dotdevgo/nest/pkg/user"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

const (
	RouteAuthSignup     = "/auth/signup"
	RouteAuthSignin     = "/auth/signin"
	RouteAuthConfirm    = "/auth/confirm/:token"
	RouteAuthRestore    = "/auth/restore"
	RouteAuthResetToken = "/auth/reset/:user/:token"
	RouteAuthMe         = "/auth/me"

	// Oauth
	RouteAuthOauth = "/auth/oauth"
)

// AuthController godoc
type AuthController struct {
	kernel.Controller
	nest.Config
	Crud *user.UserCrud
	Auth *auth.AuthService
}

// Router godoc
func (c *AuthController) Router(w *nest.Kernel) {
	w.POST(RouteAuthSignup, c.SignUp)
	w.POST(RouteAuthSignin, c.SignIn)
	w.GET(RouteAuthConfirm, c.Confirm)
	w.POST(RouteAuthRestore, c.Restore)
	w.GET(RouteAuthResetToken, c.ResetToken)
	w.GET(RouteAuthOauth, c.OAuth)
	w.GET("/auth/callback/steam", c.OAuth)

	api := w.ApiGroup()
	api.GET(RouteAuthMe, c.Me)
}

// SignUp godoc
func (c *AuthController) OAuth(ctx nest.Context) error {
	req := ctx.Request()
	res := ctx.Response()

	if user, err := gothic.CompleteUserAuth(res, req); err == nil {
		return ctx.JSON(http.StatusOK, user)
	}

	gothic.BeginAuthHandler(res, req)
	return nil
}

// SignUp godoc
func (c *AuthController) SignUp(ctx nest.Context) error {
	var input auth.SignUpDto // = new(auth.SignUpDto)
	if err := c.Crud.IsValid(ctx, &input); err != nil {
		return nest.NewValidatorError(ctx, err)
	}

	u, err := c.Auth.SignUp(input)
	if err != nil {
		return nest.NewHTTPError(http.StatusBadRequest, err)
	}

	token, err := c.Auth.NewToken(u)
	if err != nil {
		return nest.NewHTTPError(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"user":  u,
		"token": token,
	})
}

// SignIn godoc
func (c *AuthController) SignIn(ctx nest.Context) error {
	var input auth.SignInDto
	if err := c.Crud.IsValid(ctx, &input); err != nil {
		return nest.NewValidatorError(ctx, err)
	}

	u, err := c.Auth.Validate(input)
	if err != nil {
		return nest.NewHTTPError(http.StatusBadRequest, err)
	}

	token, err := c.Auth.NewToken(u)
	if err != nil {
		return nest.NewHTTPError(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"user":  u,
		"token": token,
	})
}

// Confirm godoc
func (c *AuthController) Confirm(ctx nest.Context) error {
	token := ctx.Param("token")
	if token == "" {
		return nest.NewHTTPError(http.StatusBadRequest)
	}

	if err := c.Auth.Confirm(token); err != nil {
		return nest.NewHTTPError(http.StatusBadRequest, err)
	}

	return ctx.Redirect(http.StatusMovedPermanently, c.Config.CORS.Origin)
}

// Restore godoc
func (c *AuthController) Restore(ctx nest.Context) error {
	var input auth.RestoreDto
	if err := c.Crud.IsValid(ctx, input); err != nil {
		return nest.NewValidatorError(ctx, err)
	}

	if err := c.Auth.Restore(input); err != nil {
		return nest.NewHTTPError(http.StatusBadRequest, err)
	}

	return ctx.NoContent(http.StatusOK)
}

// ResetToken godoc
func (c *AuthController) ResetToken(ctx nest.Context) error {
	var u user.User
	if err := c.Crud.Find(&u, ctx.Param("user")); err != nil {
		return nest.NewHTTPError(http.StatusNotFound, err)
	}

	if err := c.Auth.ResetToken(u, ctx.Param("token")); err != nil {
		return nest.NewHTTPError(http.StatusBadRequest, err)
	}

	return ctx.Redirect(http.StatusMovedPermanently, c.Config.CORS.Origin)
}

// Me godoc
func (c *AuthController) Me(ctx nest.Context) error {
	cc := auth.NewContext(ctx)

	u := cc.User()
	if u == nil {
		return nest.NewHTTPError(http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK, u)
}
