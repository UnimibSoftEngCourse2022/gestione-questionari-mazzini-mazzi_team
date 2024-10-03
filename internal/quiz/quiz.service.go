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

func (a *Service) FindAll(userID uint) ([]Quiz, error) {
	quizs, err := a.repository.FindAll(userID)

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	return quizs, nil
}

func (a *Service) FindById(id uint, userID uint) (*Quiz, error) {
	partialQuiz := Quiz{ID: id, UserID: userID}
	quiz, err := a.repository.Find(&partialQuiz)

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	return quiz, nil
}

func (a *Service) Create(title string, userID uint) (*Quiz, error) {

	quiz := &Quiz{
		Title:  title,
		UserID: userID,
	}
	quiz, err := a.repository.Create(quiz)

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	return quiz, nil
}

func (a *Service) Delete(id uint, UserID uint) error {
	err := a.repository.Delete(&Quiz{ID: id, UserID: UserID})
	if err != nil {
		a.logger.Error().Msg(err.Error())
		return err
	}

	return nil
}

func (a *Service) UpdateClosedQuestion() error {

	return nil
}

func (a *Service) UpdateOpenQuestion() error {
	return nil
}
