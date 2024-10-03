package quiz

import (
	"errors"
	service "form_management/internal/quiz/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type API struct {
	QuizService *service.QuizService
}

func NewQuizHanlder(service *service.QuizService) *API {
	return &API{
		QuizService: service,
	}
}

const ErrorPageHandler = "error.html"

type RowData struct {
	QuizID    uint
	QuizTitle string
	Length    int
}

func extractUserAuth(c echo.Context) (*uint, error) {
	sess, err := session.Get("quiz_app_session", c)
	if err != nil {
		return nil, err
	}

	id := sess.Values["id"]
	if id == nil {
		return nil, errors.New("user data not in session")
	}

	parsedId := id.(uint)
	return &parsedId, nil
}

func (a *API) ListQuiz(c echo.Context) error {
	userID, err := extractUserAuth(c)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	quizs, err := a.QuizService.FindAll(*userID)
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

	userID, err := extractUserAuth(c)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	quiz, err := a.QuizService.FindById(uint(id), *userID)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	data := map[string]interface{}{
		"QuizTitle": quiz.Title,
	}

	return c.Render(http.StatusOK, "CardQuiz", data)
}

func (a *API) CreateQuiz(c echo.Context) error {
	title := c.FormValue("QuizName")
	userID, err := extractUserAuth(c)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	quiz, err := a.QuizService.Create(title, *userID)
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

func (a *API) UpdateQuiz(c echo.Context) error {
	return nil
}

func (a *API) DeleteQuiz(c echo.Context) error {
	stringId := c.QueryParam("id")
	id, err := strconv.Atoi(stringId)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	userID, err := extractUserAuth(c)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	err = a.QuizService.Delete(uint(id), *userID)
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	return c.Render(http.StatusOK, "QuizRow", nil)
}
