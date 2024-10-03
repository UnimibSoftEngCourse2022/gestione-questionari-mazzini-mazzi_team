package auth

import (
	"form_management/internal/auth/user"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type API struct {
	userService *user.Service
}

func NewAuthHandler(service *user.Service) *API {
	return &API{
		userService: service,
	}
}

const ErrorPageHandler = "error.html"

func updateSession(c echo.Context, id uint) error {
	sess, err := session.Get("quiz_app_session", c)
	if err != nil {
		return err
	}

	// sess.Values["code"] = code
	sess.Values["id"] = id
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	return nil
}

func (a *API) LoginGuest(c echo.Context) error {
	code := c.FormValue("code")
	user, err := a.userService.LoginGuest(code)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	err = updateSession(c, user.ID)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	return c.Redirect(http.StatusSeeOther, "/login")
}

func (a *API) LoginUser(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	user, err := a.userService.LoginUser(email, password)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	err = updateSession(c, user.ID)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

func (a *API) RegisterUser(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	user, err := a.userService.RegisterUser(email, password)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	err = updateSession(c, user.ID)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

func (a *API) RegisterGuest(c echo.Context) error {
	return c.JSON(http.StatusOK, "register")
}
