package quiz

import (
	common "form_management/common/logger"
	"form_management/db"

	"github.com/labstack/echo/v4"
)

func Route(e *echo.Group) {
	logger := common.Logger

	db := db.Init()
	db.AutoMigrate(&Quiz{}, &QuizClosedQuestion{}, &QuizOpenQuestion{})

	service := NewService(&logger, db)
	handler := NewQuizHanlder(service)

	e.GET("/findAll", handler.ListQuiz)
	e.GET("/find", handler.FindQuiz)

	e.POST("/create", handler.CreateQuiz)

	e.DELETE("/delete", handler.DeleteQuiz)

}
