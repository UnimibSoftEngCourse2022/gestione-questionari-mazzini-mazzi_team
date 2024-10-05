package quiz

import (
	session "form_management/common/session"
	closedquestion "form_management/internal/question/closed-question"
	openquestion "form_management/internal/question/open-question"
	service "form_management/internal/quiz/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type API struct {
	quizService               *service.QuizService
	sessionService            *session.SessionService
	quizOpenQuestionService   *service.QuizOpenQuestionService
	quizClosedQuestionService *service.QuizClosedQuestionService
	openQuestionService       *openquestion.Service
	closedQuestionService     *closedquestion.Service
}

func NewQuizHandler(
	quizService *service.QuizService,
	sessionService *session.SessionService,
	quizOpenQuestionService *service.QuizOpenQuestionService,
	quizClosedQuestionService *service.QuizClosedQuestionService,
	openQuestionService *openquestion.Service,
	closedQuestionService *closedquestion.Service,
) *API {
	return &API{
		quizService:               quizService,
		sessionService:            sessionService,
		quizOpenQuestionService:   quizOpenQuestionService,
		quizClosedQuestionService: quizClosedQuestionService,
		openQuestionService:       openQuestionService,
		closedQuestionService:     closedQuestionService,
	}
}

const ErrorPageHandler = "error.html"

type RowData struct {
	QuizID    uint
	QuizTitle string
	Length    int
}

func (a *API) ListQuiz(c echo.Context) error {
	userID, err := a.sessionService.ExtractUserAuth(c)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	quizs, err := a.quizService.FindAll(*userID)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	data := []RowData{}
	for _, quiz := range quizs {
		data = append(
			data,
			RowData{
				QuizID:    quiz.ID,
				QuizTitle: quiz.Title,
				Length:    len(quiz.ClosedQuestions) + len(quiz.OpenQuestions),
			},
		)
	}

	return c.Render(http.StatusOK, "TableQuiz", map[string]interface{}{"quizs": data})
}

func (a *API) FindQuiz(c echo.Context) error {
	idString := c.QueryParam("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	userID, err := a.sessionService.ExtractUserAuth(c)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	quiz, err := a.quizService.FindById(uint(id), *userID)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	quizOpenQuestions, err := a.quizOpenQuestionService.FindByQuizID(uint(id))
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}
	quizClosedQuestions, err := a.quizClosedQuestionService.FindByQuizID(uint(id))
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	idsClosedQuestions := []uint{}
	for _, question := range quizClosedQuestions {
		idsClosedQuestions = append(idsClosedQuestions, question.ClosedQuestionID)
	}
	closedQuestions, err := a.closedQuestionService.FindAllByIds(idsClosedQuestions)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	idsOpenQuestions := []uint{}
	for _, question := range quizOpenQuestions {
		idsOpenQuestions = append(idsOpenQuestions, question.OpenQuestionID)
	}
	openQuestions, err := a.openQuestionService.FindAllByIds(idsOpenQuestions)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	data := map[string]interface{}{
		"QuizTitle":       quiz.Title,
		"OpenQuestions":   openQuestions,
		"ClosedQuestions": closedQuestions,
	}

	return c.Render(http.StatusOK, "CardQuiz", data)
}

func (a *API) CreateQuiz(c echo.Context) error {
	title := c.FormValue("QuizName")
	userID, err := a.sessionService.ExtractUserAuth(c)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	quiz, err := a.quizService.Create(title, *userID)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	data := RowData{
		QuizID:    quiz.ID,
		QuizTitle: quiz.Title,
		Length:    len(quiz.ClosedQuestions) + len(quiz.OpenQuestions),
	}

	return c.Render(http.StatusOK, "QuizRow", data)

}

func (a *API) DeleteQuiz(c echo.Context) error {
	stringId := c.QueryParam("id")
	id, err := strconv.Atoi(stringId)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	userID, err := a.sessionService.ExtractUserAuth(c)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	err = a.quizService.Delete(uint(id), *userID)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	return c.Render(http.StatusOK, "QuizRow", nil)
}

func (a *API) AddOpenQuestionQuiz(c echo.Context) error {
	stringQuizId := c.QueryParam("quizID")
	stringQuestionId := c.QueryParam("questionID")

	quizID, err := strconv.Atoi(stringQuizId)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}
	questionID, err := strconv.Atoi(stringQuestionId)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}
	userID, err := a.sessionService.ExtractUserAuth(c)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	openQuestion, err := a.openQuestionService.FindById(uint(questionID))
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}
	quiz, err := a.quizService.FindById(uint(quizID), *userID)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}
	question, err := a.quizOpenQuestionService.Create(*openQuestion, *quiz, 0)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	data := map[string]interface{}{
		"Text": question.OpenQuestion.Text,
	}
	return c.Render(http.StatusOK, "OpenQuestionSection", data)
}

func (a *API) AddClosedQuestionQuiz(c echo.Context) error {
	stringQuizId := c.QueryParam("quizID")
	stringQuestionId := c.QueryParam("questionID")

	quizID, err := strconv.Atoi(stringQuizId)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}
	questionID, err := strconv.Atoi(stringQuestionId)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}
	userID, err := a.sessionService.ExtractUserAuth(c)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	closedQuestion, err := a.closedQuestionService.FindById(uint(questionID))
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}
	quiz, err := a.quizService.FindById(uint(quizID), *userID)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}
	question, err := a.quizClosedQuestionService.Create(*closedQuestion, *quiz, 0)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	data := map[string]interface{}{
		"Text":    question.ClosedQuestion.Text,
		"Answers": question.ClosedQuestion.Answers,
	}
	return c.Render(http.StatusOK, "ClosedQuestionSection", data)
}
