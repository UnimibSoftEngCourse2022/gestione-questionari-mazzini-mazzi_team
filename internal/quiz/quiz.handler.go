package quiz

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type API struct {
	QuizService *Service
}

func NewQuizHanlder(service *Service) *API {
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

func (a *API) ListQuiz(c echo.Context) error {
	quizs, err := a.QuizService.FindAll()

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

	quiz, err := a.QuizService.FindById(uint(id))
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
	quiz, err := a.QuizService.Create(title)

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

	err = a.QuizService.Delete(uint(id))
	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	return c.Render(http.StatusOK, "QuizRow", nil)
}
