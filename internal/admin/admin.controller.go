package admin

import (
	"form_management/internal/admin/handlers"
	"github.com/labstack/echo/v4"
)

func Route(group *echo.Group) {
	group.GET("/login", handlers.Login)
	group.GET("/register", handlers.Register)
}
