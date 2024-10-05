package quiz

import (
	common "form_management/common/logger"
	"form_management/common/session"
	"form_management/db"
	closedquestion "form_management/internal/question/closed-question"
	openquestion "form_management/internal/question/open-question"
	"form_management/internal/quiz/entities"
	services "form_management/internal/quiz/services"

	"github.com/labstack/echo/v4"
)

func Route(e *echo.Group) {
	logger := common.Logger

	db := db.Init()
	db.AutoMigrate(
		&entities.Quiz{},
		&entities.QuizClosedQuestion{},
		&entities.QuizOpenQuestion{},
	)

	sessionService := session.NewService(&logger)
	quizService := services.NewQuizService(&logger, db)
	quizClosedQuestionService := services.NewQuizClosedQuestionService(&logger, db)
	quizOpenQuestionService := services.NewQuizOpenQuestionService(&logger, db)
	openQuestionsService := openquestion.NewService(&logger, db)
	closedQuestionsService := closedquestion.NewService(&logger, db)

	handler := NewQuizHandler(
		quizService,
		sessionService,
		quizOpenQuestionService,
		quizClosedQuestionService,
		openQuestionsService,
		closedQuestionsService,
	)

	e.GET("/findAll", handler.ListQuiz)
	e.GET("/find", handler.FindQuiz)
	e.POST("/create", handler.CreateQuiz)
	e.DELETE("/delete", handler.DeleteQuiz)

	e.PUT("/update/add-closed-question", handler.AddClosedQuestionQuiz)
	e.PUT("/update/add-open-question", handler.AddOpenQuestionQuiz)
}
