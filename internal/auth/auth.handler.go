package auth

import (
	session "form_management/common/session"
	"form_management/internal/auth/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type API struct {
	userService    *user.Service
	sessionService *session.SessionService
}

func NewAuthHandler(service *user.Service, sessionService *session.SessionService) *API {
	return &API{
		userService:    service,
		sessionService: sessionService,
	}
}

const ErrorPageHandler = "error.html"

func (a *API) LoginGuest(c echo.Context) error {
	code := c.FormValue("code")
	user, err := a.userService.LoginGuest(code)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	err = a.sessionService.UpdateSession(c, user.ID)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

func (a *API) LoginUser(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	user, err := a.userService.LoginUser(email, password)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	err = a.sessionService.UpdateSession(c, user.ID)
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

	err = a.sessionService.UpdateSession(c, user.ID)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	return c.Redirect(http.StatusSeeOther, "/")
}

func (a *API) Logout(c echo.Context) error {
	err := a.sessionService.DeleteSession(c)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	return c.Redirect(http.StatusSeeOther, "/login")
}

func (a *API) RegisterGuest(c echo.Context) error {
	return c.JSON(http.StatusOK, "register")
}
