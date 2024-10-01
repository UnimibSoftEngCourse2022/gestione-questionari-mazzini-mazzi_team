package quiz

import (
	"form_management/common"
	"form_management/db"

	"github.com/labstack/echo/v4"
)

func Route(e *echo.Group) {
	logger := common.Logger

	db := db.Init()
	db.AutoMigrate(&Quiz{})

	service := NewService(&logger, db)
	handler := NewQuizHanlder(service)

	e.GET("/findAll", handler.FindAllQuiz)

}
