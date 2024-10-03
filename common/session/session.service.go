package session

import (
	"errors"
	common "form_management/common/logger"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type SessionService struct {
	logger *common.MyLogger
}

func NewService(logger *common.MyLogger) *SessionService {
	return &SessionService{
		logger: logger,
	}
}

func (a *SessionService) ExtractUserAuth(c echo.Context) (*uint, error) {
	sess, err := session.Get("quiz_app_session", c)
	if err != nil {
		return nil, err
	}

	id := sess.Values["id"]
	if id == nil {
		return nil, errors.New("user data not in session")
	}

	return id.(*uint), nil
}

func (a *SessionService) UpdateSession(c echo.Context, id uint) error {
	sess, err := session.Get("quiz_app_session", c)
	if err != nil {
		return err
	}
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["id"] = id
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	return nil
}

func (a *SessionService) DeleteSession(c echo.Context) error {
	sess, err := session.Get("quiz_app_session", c)
	if err != nil {
		return err
	}

	sess.Options.MaxAge = -1
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	return nil
}
