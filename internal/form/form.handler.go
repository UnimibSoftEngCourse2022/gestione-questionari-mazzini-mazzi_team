package form

import (
	"form_management/common"
	closedquestion "form_management/internal/form/closed-question"
	openquestion "form_management/internal/form/open-question"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type API struct {
	ClosedQuestionService *closedquestion.Service
	OpenQuestionService   *openquestion.Service
}

func NewFormHanlder(closedQuestionService *closedquestion.Service, openQuestionService *openquestion.Service) *API {
	return &API{
		ClosedQuestionService: closedQuestionService,
		OpenQuestionService:   openQuestionService,
	}
}

type RowData struct {
	Question   string
	AnswerType string
	Category   string
	ImageURL   string
}

func (a *API) ListQuestionsHandler(c echo.Context) error {

	closedQuestions, err := a.ClosedQuestionService.FindAll()
	openQuestions, err := a.OpenQuestionService.FindAll()

	if err != nil {
		return c.Render(http.StatusNotFound, "error.html", map[string]interface{}{"error": err.Error()})
	}

	data := []RowData{}

	for _, question := range closedQuestions {
		data = append(
			data,
			RowData{
				Question:   question.Text,
				AnswerType: "CLOSED_QUESTION", //question.AnswerType,
				Category:   question.Category,
				ImageURL:   question.ImageURL,
			},
		)
	}

	for _, question := range openQuestions {
		data = append(
			data,
			RowData{
				Question:   question.Text,
				AnswerType: "OPEN_QUESTION", //question.AnswerType,
				Category:   question.Category,
				ImageURL:   question.ImageURL,
			},
		)
	}

	return c.Render(http.StatusOK, "Table", map[string]interface{}{"questions": data})
}

func (a *API) CreateClosedQuestionsHandler(c echo.Context) error {

	common.Logger.Debug().Msg(c.FormValue("answer"))
	questionText := c.FormValue("questionText")
	questionImageURL := c.FormValue("questionImageURL")
	questionCategory := c.FormValue("questionCategory")
	questionAnswers := c.Request().Form["answer"]

	question, err := a.ClosedQuestionService.Create(
		questionText,
		questionImageURL,
		questionCategory,
		questionAnswers,
	)

	if err != nil {
		return c.Render(http.StatusNotFound, "error.html", map[string]interface{}{"error": err.Error()})
	}

	data := RowData{
		Question:   question.Text,
		AnswerType: question.AnswerType,
		Category:   question.Category,
		ImageURL:   question.ImageURL,
	}
	return c.Render(http.StatusOK, "Row", data)
}

func (a *API) CreateOpenQuestionsHandler(c echo.Context) error {

	questionText := c.FormValue("questionText")
	questionImageURL := c.FormValue("questionImageURL")
	questionCategory := c.FormValue("questionCategory")
	questionMinChar := c.FormValue("questionMinChar")
	questionMaxChar := c.FormValue("questionMaxChar")

	maxChar, err := strconv.Atoi(questionMaxChar)
	minChar, err := strconv.Atoi(questionMinChar)

	if err != nil {
		return c.Render(http.StatusNotFound, "toggle", map[string]interface{}{"error": err.Error()})
	}

	question, err := a.OpenQuestionService.Create(
		questionText,
		questionImageURL,
		questionCategory,
		minChar,
		maxChar,
	)

	if err != nil {
		return c.Render(http.StatusNotFound, "error.html", map[string]interface{}{"error": err.Error()})
	}

	data := RowData{
		Question:   question.Text,
		AnswerType: question.AnswerType,
		Category:   question.Category,
		ImageURL:   question.ImageURL,
	}
	return c.Render(http.StatusOK, "Row", data)
}

func (a *API) RenderNextStepCreationQuetion(c echo.Context) error {
	questionType := c.QueryParam("question-type")

	switch questionType {
	case "open-question":
		return c.Render(http.StatusOK, "ModalOpenQuestion", nil)
	case "closed-question":
		return c.Render(http.StatusOK, "ModalClosedQuestion", nil)
	default:
		return c.Render(http.StatusNotFound, "error.html", "404 not found")
	}

}
