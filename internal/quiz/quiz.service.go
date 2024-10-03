package quiz

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

func (a *Service) FindAll() ([]Quiz, error) {
	quizs, err := a.repository.FindAll()

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	return quizs, nil
}

func (a *Service) FindById(id uint) (*Quiz, error) {
	partialQuiz := Quiz{ID: id}
	quiz, err := a.repository.Find(&partialQuiz)

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	return quiz, nil
}

func (a *Service) Create(title string) (*Quiz, error) {

	quiz := &Quiz{
		Title: title,
	}
	quiz, err := a.repository.Create(quiz)

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	return quiz, nil
}

// func (a *Service) Update(id uint, text string, imageURL string, category string, minChar int, maxChar int) (*Quiz, error) {
// 	updatedQuestion := Quiz{
// 		Text:       text,
// 		ImageURL:   imageURL,
// 		Category:   category,
// 		AnswerType: "OPEN_QUESTION",
// 		MinChar:    minChar,
// 		MaxChar:    maxChar,
// 	}
// 	quiz, err := a.repository.Update(updatedQuestion)

// 	if err != nil {
// 		a.logger.Error().Msg(err.Error())
// 		return nil, err
// 	}

// 	return quiz, nil
// }

func (a *Service) Delete(id uint) error {
	err := a.repository.Delete(&Quiz{ID: id})

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return err
	}

	return nil
}
