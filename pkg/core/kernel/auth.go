package kernel

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	nest "github.com/dotdevgo/nest/pkg/core"
	"github.com/labstack/echo/v4"
)

func AuthController(ctx nest.Context) error {
	username := ctx.FormValue("username")

	claims := &JwtClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
