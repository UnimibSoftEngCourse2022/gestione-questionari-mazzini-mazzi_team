package quiz

import (
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

func (a *API) FindAllQuiz(c echo.Context) error {
	return nil
}

func (a *API) FindQuiz(c echo.Context) error {
	return nil
}

func (a *API) CreateQuiz(c echo.Context) error {
	return nil
}

func (a *API) UpdateQuiz(c echo.Context) error {
	return nil
}

func (a *API) DeleteQuiz(c echo.Context) error {
	return nil
}
