package openquestion

import (
	"form_management/common"

	"gorm.io/gorm"
)

type Service struct {
	logger     *common.MyLogger
	repository *Repository
}

func NewService(logger *common.MyLogger, db *gorm.DB) *Service {
	return &Service{
		logger:     logger,
		repository: NewRepository(db),
	}
}

func (a *Service) FindAll() ([]OpenQuestion, error) {
	questions, err := a.repository.FindAll()

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	return questions, nil
}

// func (a *Service) Find() (*OpenQuestion, error) {
// 	question, err := a.repository.Find(1)

// 	if err != nil {
// 		a.logger.Error().Msg(err.Error())
// 		return nil, err
// 	}

// 	return question, nil
// }

func (a *Service) Create(text string, image_url string, category string, minChar int, maxChar int) (*OpenQuestion, error) {

	question := &OpenQuestion{
		Text:       text,
		ImageURL:   image_url,
		Category:   category,
		AnswerType: "OPEN_QUESTION",
		MinChar:    maxChar,
		MaxChar:    minChar,
	}
	question, err := a.repository.Create(question)

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	return question, nil
}
