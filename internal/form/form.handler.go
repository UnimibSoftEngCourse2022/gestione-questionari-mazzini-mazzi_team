package form

import (
	closedquestion "form_management/internal/form/closed-question"
	"net/http"

	"github.com/labstack/echo/v4"
)

type API struct {
	ClosedQuestionService *closedquestion.Service
}

func NewFormHanlder(service *closedquestion.Service) *API {
	return &API{
		ClosedQuestionService: service,
	}
}

func (a *API) ListQuestionsHandler(c echo.Context) error {
	questions, err := a.ClosedQuestionService.FindAll()

	if err != nil {
		return c.Render(http.StatusNotFound, "error.html", map[string]interface{}{"error": err.Error()})
	}

	data := map[string]interface{}{
		"title":     "Home Page",
		"message":   "Hello, World!",
		"questions": questions,
	}
	return c.Render(http.StatusOK, "test.html", data)
}

func (a *API) FindQuestionsHandler(c echo.Context) error {
	question, err := a.ClosedQuestionService.Find()

	if err != nil {
		return c.Render(http.StatusNotFound, "error.html", map[string]interface{}{"error": err.Error()})
	}

	data := map[string]interface{}{
		"title":   "Home Page",
		"message": "Hello, World!",
		"isEmpty": question.Text,
	}
	return c.Render(http.StatusOK, "find.html", data)
}

func (a *API) CreateQuestionsHandler(c echo.Context) error {
	question, err := a.ClosedQuestionService.Create("text", "http://sas.com", "CATEGORY")

	if err != nil {
		return c.Render(http.StatusNotFound, "error.html", map[string]interface{}{"error": err.Error()})
	}

	data := map[string]interface{}{
		"title":   "Home Page",
		"message": "Hello, World!",
		"isEmpty": question.Text,
	}
	return c.Render(http.StatusOK, "find.html", data)
}
