package form

import (
	"form_management/common"
	"form_management/db"
	closedquestion "form_management/internal/form/closed-question"
	openquestion "form_management/internal/form/open-question"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Route(e *echo.Group) {
	db := db.Init()
	db.AutoMigrate(&closedquestion.ClosedQuestion{})

	logger := common.Logger
	closedQuestionService := closedquestion.NewService(&logger, db)
	openQuestionService := openquestion.NewService(&logger, db)
	handler := NewFormHanlder(closedQuestionService, openQuestionService)

	e.GET("/findAll", handler.ListQuestionsHandler)
	e.POST("/create/open-question", handler.CreateOpenQuestionsHandler)
	e.POST("/create/closed-question", handler.CreateClosedQuestionsHandler)

	e.GET("/create/nextStep", handler.RenderNextStepCreationQuetion)
	e.GET("/create/prevStep", func(c echo.Context) error { return c.Render(http.StatusOK, "ModalBody", nil) })
	// e.POST("/closed-question/create", handler.CreateCloseQuestionsHandler)
}
