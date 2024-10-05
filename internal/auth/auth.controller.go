package auth

import (
	common "form_management/common/logger"
	session "form_management/common/session"
	"form_management/db"
	"form_management/internal/auth/user"

	"github.com/labstack/echo/v4"
)

func Route(e *echo.Group) {
	logger := common.Logger

	db := db.Init()
	db.AutoMigrate(&user.User{})

	userService := user.NewService(&logger, db)
	sessionService := session.NewService(&logger)
	handler := NewAuthHandler(userService, sessionService)

	e.POST("/login/user", handler.LoginUser)
	e.POST("/login/guest", nil)

	e.POST("/register/guest", nil)
	e.POST("/register/user", handler.RegisterUser)

	e.GET("/logout", handler.Logout)
}
