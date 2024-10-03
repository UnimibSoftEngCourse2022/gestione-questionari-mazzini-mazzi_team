package openquestion

import (
	common "form_management/common/logger"

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

func (a *Service) FindById(id uint) (*OpenQuestion, error) {
	partialQuestion := OpenQuestion{ID: id}
	question, err := a.repository.Find(&partialQuestion)

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	return question, nil
}

func (a *Service) Create(text string, image_url string, category string, minChar int, maxChar int) (*OpenQuestion, error) {

	question := &OpenQuestion{
		Text:       text,
		ImageURL:   image_url,
		Category:   category,
		AnswerType: "OPEN_QUESTION",
		MinChar:    minChar,
		MaxChar:    maxChar,
	}
	question, err := a.repository.Create(question)

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	return question, nil
}

func (a *Service) Update(id uint, text string, imageURL string, category string, minChar int, maxChar int) (*OpenQuestion, error) {
	updatedQuestion := OpenQuestion{
		Text:       text,
		ImageURL:   imageURL,
		Category:   category,
		AnswerType: "OPEN_QUESTION",
		MinChar:    minChar,
		MaxChar:    maxChar,
	}
	question, err := a.repository.Update(updatedQuestion)

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	return question, nil
}

func (a *Service) Delete(id uint) error {
	err := a.repository.Delete(&OpenQuestion{ID: id})

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return err
	}

	return nil
}
