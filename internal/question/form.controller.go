package question

import (
	common "form_management/common/logger"
	"form_management/db"
	closedquestion "form_management/internal/question/closed-question"
	openquestion "form_management/internal/question/open-question"

	"github.com/labstack/echo/v4"
)

func Route(e *echo.Group) {
	logger := common.Logger

	db := db.Init()
	db.AutoMigrate(&closedquestion.ClosedQuestion{})
	db.AutoMigrate(&openquestion.OpenQuestion{})

	closedQuestionService := closedquestion.NewService(&logger, db)
	openQuestionService := openquestion.NewService(&logger, db)
	handler := NewFormHanlder(closedQuestionService, openQuestionService)

	e.GET("/findAll", handler.ListQuestions)
	e.GET("/find", handler.FindQuestion)

	e.POST("/create/open-question", handler.CreateOpenQuestions)
	e.POST("/create/closed-question", handler.CreateClosedQuestions)

	e.PUT("/update/closed-question", handler.UpdateClosedQuestions)
	e.PUT("/update/open-question", handler.UpdateOpenQuestions)

	e.DELETE("/delete/closed-question", handler.DeleteClosedQuestions)
	e.DELETE("/delete/open-question", handler.DeleteOpenQuestions)

	e.GET("/closed-question/create-page", handler.RenderCardClosedQuestion)
	e.GET("/open-question/create-page", handler.RenderCardOpenQuestion)

}
