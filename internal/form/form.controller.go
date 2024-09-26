package form

import (
	"form_management/common"
	"form_management/db"
	closedquestion "form_management/internal/form/closed-question"

	"github.com/labstack/echo/v4"
)

func Route(e *echo.Group) {
	db := db.Init()
	db.AutoMigrate(&closedquestion.ClosedQuestion{})

	logger := common.Logger
	service := closedquestion.NewService(&logger, db)
	handler := NewFormHanlder(service)

	e.GET("/findAll", handler.ListQuestionsHandler)
	e.GET("/find", handler.FindQuestionsHandler)
	e.GET("/create", handler.CreateQuestionsHandler)
}
