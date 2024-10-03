package auth

import (
	common "form_management/common/logger"
	"form_management/db"
	"form_management/internal/auth/user"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(handler echo.HandlerFunc) echo.HandlerFunc {
	mydb := db.Init()
	logger := common.NewLogger()
	userService := user.NewService(&logger, mydb)

	return func(c echo.Context) error {
		sess, err := session.Get("quiz_app_session", c)
		if err != nil {
			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}

		id := sess.Values["id"]
		if id == nil {
			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}
		_, err = userService.IsLogged(id.(uint))
		if err != nil {
			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}

		return handler(c)
	}
}
