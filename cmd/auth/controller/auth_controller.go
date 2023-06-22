package controller

import (
	"fmt"
	"net"
	"net/http"
	"net/url"

	"dotdev/nest/pkg/auth"
	"dotdev/nest/pkg/nest"
	"dotdev/nest/pkg/user"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

const (
	RouteAuthSignup         = "/auth/signup"
	RouteAuthSignin         = "/auth/signin"
	RouteAuthConfirm        = "/auth/confirm/:token"
	RouteAuthRestore        = "/auth/restore"
	RouteAuthResetToken     = "/auth/reset/:user/:token"
	RouteOauth              = "/auth/oauth/:provider"
	RouteOauthCallback      = "/auth/callback/:provider"
	RouteAuthChangePassword = "/api/auth/password"
	RouteAuthUpdate         = "/api/auth/update"
	RouteAuthMe             = "/api/auth/me"
)

type AuthController struct {
	nest.Controller
	nest.Config
	Crud *user.UserCrud
	Auth *auth.AuthManager
}

func (c AuthController) Router(w *nest.Kernel) {
	w.POST(RouteAuthSignup, c.SignUp)
	w.POST(RouteAuthSignin, c.SignIn)
	w.GET(RouteAuthConfirm, c.Confirm)
	w.POST(RouteAuthRestore, c.Restore)
	w.GET(RouteAuthResetToken, c.ResetToken)
	w.GET(RouteOauth, c.OAuth)
	w.GET(RouteOauthCallback, c.OAuth)

	api := w.Api()
	api.GET(RouteAuthMe, c.Me)
	api.POST(RouteAuthChangePassword, c.ChangePassword)
	api.POST(RouteAuthUpdate, c.Update)
}

// OAuth godoc
func (c AuthController) OAuth(ctx nest.Context) error {
	req, res := ctx.Request(), ctx.Response()

	// https://groups.google.com/g/golang-nuts/c/Dur6uGUEKKk
	values := req.URL.Query()
	values.Add("provider", ctx.Param("provider"))
	req.URL.RawQuery = values.Encode()

	if oauthUser, err := gothic.CompleteUserAuth(res, req); err == nil {
		oauth, err := c.Auth.OAuth(oauthUser)
		if err != nil {
			return nest.NewHTTPError(http.StatusBadRequest, err)
		}

		token, err := c.Auth.NewToken(oauth.User)
		if err != nil {
			return nest.NewHTTPError(http.StatusBadRequest, err)
		}

		u, err := url.Parse(c.Config.HTTP.Origin)
		if err != nil {
			return nest.NewHTTPError(http.StatusBadRequest, err)
		}

		domain, _, _ := net.SplitHostPort(u.Host)
		if len(domain) == 0 {
			domain = u.Host
		}

		// TODO: config cookie name and other values.
		cookie := &http.Cookie{
			Name:     "APP_SSO_TOKEN",
			Value:    token,
			Path:     "/",
			Domain:   domain,
			HttpOnly: false,
			Secure:   ctx.IsTLS(),
			SameSite: http.SameSiteLaxMode,
		}

		ctx.SetCookie(cookie)

		redirectUrl := fmt.Sprintf("%s/auth/oauth", c.Config.HTTP.Origin)

		return ctx.Redirect(http.StatusMovedPermanently, redirectUrl)
	}

	gothic.BeginAuthHandler(res, req)
	return nil
}

// SignUp godoc
func (c AuthController) SignUp(ctx nest.Context) error {
	var input auth.SignUpDto
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
func (c AuthController) SignIn(ctx nest.Context) error {
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
func (c AuthController) Confirm(ctx nest.Context) error {
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
func (c AuthController) Restore(ctx nest.Context) error {
	var input auth.IdentityDto
	if err := c.Crud.IsValid(ctx, &input); err != nil {
		return nest.NewValidatorError(ctx, err)
	}

	if err := c.Auth.Restore(input); err != nil {
		return nest.NewHTTPError(http.StatusBadRequest, err)
	}

	return ctx.NoContent(http.StatusOK)
}

// ResetToken godoc
func (c AuthController) ResetToken(ctx nest.Context) error {
	var u user.User
	if err := c.Crud.Find(&u, ctx.Param("user")); err != nil {
		return nest.NewHTTPError(http.StatusNotFound, err)
	}

	if err := c.Auth.ResetToken(&u, ctx.Param("token")); err != nil {
		return nest.NewHTTPError(http.StatusBadRequest, err)
	}

	return ctx.Redirect(http.StatusMovedPermanently, c.Config.CORS.Origin)
}

// Me godoc
func (c AuthController) Me(ctx nest.Context) error {
	cc := auth.Context(ctx)
	u := cc.User()
	if u == nil {
		return nest.NewHTTPError(http.StatusBadRequest)
	}

	return ctx.JSON(http.StatusOK, u)
}

// Update godoc
func (c AuthController) Update(ctx nest.Context) error {
	cc := auth.Context(ctx)
	u := cc.User()
	if u == nil {
		return nest.NewHTTPError(http.StatusBadRequest)
	}

	var input user.UserDto
	// input.Pk = u.Pk
	// input.ID = u.ID

	if err := ctx.Bind(&input); err != nil {
		return err
	}

	if err := c.Auth.Save(u, input); err != nil {
		return nest.NewValidatorError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, u)
}

// ChangePassword godoc
func (c AuthController) ChangePassword(ctx nest.Context) error {
	cc := auth.Context(ctx)
	u := cc.User()
	if u == nil {
		return nest.NewHTTPError(http.StatusBadRequest)
	}

	var input auth.ChangePasswordDto
	if err := c.Crud.IsValid(ctx, &input); err != nil {
		return nest.NewValidatorError(ctx, err)
	}

	err := c.Auth.ChangePassword(u, input)
	if err != nil {
		return nest.NewHTTPError(http.StatusBadRequest, err)
	}

	return ctx.NoContent(http.StatusNoContent)
}
