package question

import (
	closedquestion "form_management/internal/question/closed-question"
	openquestion "form_management/internal/question/open-question"
	"net/http"
	"strconv"
	"strings"

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

const ErrorPageHandler = "error.html"

type RowData struct {
	QuestionID    uint
	Question      string
	AnswerType    string
	Category      string
	ImageURL      string
	QuestionIndex int
}

type CardHTMX struct {
	Method string
	Target string
	URL    string
	Swap   string
}

func updateCard(id uint, answerType string) CardHTMX {
	paredAnswerType := strings.ReplaceAll(strings.ToLower(answerType), "_", "-")
	return CardHTMX{
		Method: "put",
		Target: "#question-row-" + strconv.FormatUint(uint64(id), 10) + "-" + paredAnswerType,
		URL:    "/form/update/closed-question?id=" + strconv.FormatUint(uint64(id), 10),
		Swap:   "outerHTML",
	}
}

func createCard(answerType string) CardHTMX {
	paredAnswerType := strings.ReplaceAll(strings.ToLower(answerType), "_", "-")
	return CardHTMX{
		Method: "post",
		URL:    "/form/create/" + paredAnswerType,
		Target: "#questions",
		Swap:   "afterbegin",
	}
}

func (a *API) FindQuestion(c echo.Context) error {
	idString := c.QueryParam("id")
	answerType := strings.ReplaceAll(strings.ToLower(c.QueryParam("answerType")), "_", "-")
	id, err := strconv.Atoi(idString)

	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	if answerType == "closed-question" {
		question, _ := a.ClosedQuestionService.FindById(uint(id))
		data := map[string]interface{}{
			"QuestionID":   question.ID,
			"QuestionText": question.Text,
			"AnswerType":   question.AnswerType,
			"Category":     question.Category,
			"ImageURL":     question.ImageURL,
			"Answers":      question.Answers,
			"htmx":         updateCard(question.ID, question.AnswerType),
		}
		return c.Render(http.StatusOK, "ClosedQuestionCard", data)

	} else if answerType == "open-question" {
		question, _ := a.OpenQuestionService.FindById(uint(id))
		data := map[string]interface{}{
			"QuestionID":   question.ID,
			"QuestionText": question.Text,
			"AnswerType":   question.AnswerType,
			"Category":     question.Category,
			"ImageURL":     question.ImageURL,
			"MinChar":      question.MinChar,
			"MaxChar":      question.MaxChar,
			"htmx":         updateCard(question.ID, question.AnswerType),
		}
		return c.Render(http.StatusOK, "OpenQuestionCard", data)
	}

	return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": "Not Found"})

}

func (a *API) ListQuestions(c echo.Context) error {

	closedQuestions, err := a.ClosedQuestionService.FindAll()
	openQuestions, err := a.OpenQuestionService.FindAll()

	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	data := []RowData{}

	for index, question := range closedQuestions {
		data = append(
			data,
			RowData{
				QuestionID:    question.ID,
				Question:      question.Text,
				AnswerType:    "closed-question", //question.AnswerType,
				Category:      question.Category,
				ImageURL:      question.ImageURL,
				QuestionIndex: index,
			},
		)
	}

	len := len(data)
	for index, question := range openQuestions {
		data = append(
			data,
			RowData{
				QuestionID:    question.ID,
				Question:      question.Text,
				AnswerType:    "open-question", //question.AnswerType,
				Category:      question.Category,
				ImageURL:      question.ImageURL,
				QuestionIndex: index + len,
			},
		)
	}

	return c.Render(http.StatusOK, "TableQuestions", map[string]interface{}{"questions": data})
}

func (a *API) CreateClosedQuestions(c echo.Context) error {

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
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	data := RowData{
		QuestionID: question.ID,
		Question:   question.Text,
		AnswerType: question.AnswerType,
		Category:   question.Category,
		ImageURL:   question.ImageURL,
	}
	return c.Render(http.StatusOK, "QuestionRow", data)
}

func (a *API) CreateOpenQuestions(c echo.Context) error {

	questionText := c.FormValue("questionText")
	questionImageURL := c.FormValue("questionImageURL")
	questionCategory := c.FormValue("questionCategory")
	questionMinChar := c.FormValue("questionMinChar")
	questionMaxChar := c.FormValue("questionMaxChar")

	maxChar, err := strconv.Atoi(questionMaxChar)
	minChar, err := strconv.Atoi(questionMinChar)

	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	question, err := a.OpenQuestionService.Create(
		questionText,
		questionImageURL,
		questionCategory,
		minChar,
		maxChar,
	)

	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	data := RowData{
		QuestionID: question.ID,
		Question:   question.Text,
		AnswerType: question.AnswerType,
		Category:   question.Category,
		ImageURL:   question.ImageURL,
	}
	return c.Render(http.StatusOK, "QuestionRow", data)
}

func (a *API) UpdateClosedQuestions(c echo.Context) error {

	questionID := c.QueryParam("id")
	questionText := c.FormValue("questionText")
	questionImageURL := c.FormValue("questionImageURL")
	questionCategory := c.FormValue("questionCategory")
	questionAnswers := c.Request().Form["answer"]

	id, err := strconv.Atoi(questionID)

	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	question, err := a.ClosedQuestionService.Update(
		uint(id),
		questionText,
		questionImageURL,
		questionCategory,
		questionAnswers,
	)

	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	data := RowData{
		QuestionID: question.ID,
		Question:   question.Text,
		AnswerType: question.AnswerType,
		Category:   question.Category,
		ImageURL:   question.ImageURL,
	}
	return c.Render(http.StatusOK, "QuestionRow", data)
}

func (a *API) UpdateOpenQuestions(c echo.Context) error {

	questionID := c.QueryParam("id")
	questionText := c.FormValue("questionText")
	questionImageURL := c.FormValue("questionImageURL")
	questionCategory := c.FormValue("questionCategory")
	questionMinChar := c.FormValue("questionMinChar")
	questionMaxChar := c.FormValue("questionMaxChar")

	maxChar, err := strconv.Atoi(questionMaxChar)
	minChar, err := strconv.Atoi(questionMinChar)
	id, err := strconv.Atoi(questionID)

	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	question, err := a.OpenQuestionService.Update(
		uint(id),
		questionText,
		questionImageURL,
		questionCategory,
		minChar,
		maxChar,
	)

	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	data := RowData{
		QuestionID: question.ID,
		Question:   question.Text,
		AnswerType: question.AnswerType,
		Category:   question.Category,
		ImageURL:   question.ImageURL,
	}

	return c.Render(http.StatusOK, "QuestionRow", data)

}

func (a *API) DeleteClosedQuestions(c echo.Context) error {
	questionId := c.QueryParam("id")
	id, err := strconv.Atoi(questionId)

	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	err = a.ClosedQuestionService.Delete(uint(id))
	return c.Render(http.StatusOK, "QuestionRow", nil)
}

func (a *API) DeleteOpenQuestions(c echo.Context) error {
	questionId := c.QueryParam("id")
	id, err := strconv.Atoi(questionId)

	if err != nil {
		return c.Render(http.StatusNotFound, ErrorPageHandler, map[string]interface{}{"error": err.Error()})
	}

	err = a.OpenQuestionService.Delete(uint(id))
	return c.Render(http.StatusOK, "QuestionRow", nil)
}

func (a *API) RenderCardClosedQuestion(c echo.Context) error {
	data := map[string]interface{}{
		"htmx": createCard("closed-question"),
	}
	return c.Render(http.StatusOK, "ClosedQuestionCard", data)
}

func (a *API) RenderCardOpenQuestion(c echo.Context) error {
	data := map[string]interface{}{
		"htmx": createCard("open-question"),
	}
	return c.Render(http.StatusOK, "OpenQuestionCard", data)
}
