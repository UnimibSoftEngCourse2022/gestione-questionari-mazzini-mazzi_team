package service

import (
	common "form_management/common/logger"
	"form_management/internal/quiz/entities"
	"form_management/internal/quiz/repositories"

	"gorm.io/gorm"
)

type QuizService struct {
	logger     *common.MyLogger
	repository *repositories.QuizRepository
}

func NewQuizService(logger *common.MyLogger, db *gorm.DB) *QuizService {
	return &QuizService{
		logger:     logger,
		repository: repositories.NewQuizRepository(db),
	}
}

func (a *QuizService) FindAll(userID uint) ([]entities.Quiz, error) {
	quizs, err := a.repository.FindAll(userID)

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	return quizs, nil
}

func (a *QuizService) FindById(id uint, userID uint) (*entities.Quiz, error) {
	partialQuiz := entities.Quiz{ID: id, UserID: userID}
	quiz, err := a.repository.Find(&partialQuiz)

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	return quiz, nil
}

func (a *QuizService) Create(title string, userID uint) (*entities.Quiz, error) {

	newQuiz := &entities.Quiz{
		Title:  title,
		UserID: userID,
	}
	quiz, err := a.repository.Create(newQuiz)

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	return quiz, nil
}

func (a *QuizService) Delete(id uint, UserID uint) error {
	err := a.repository.Delete(&entities.Quiz{ID: id, UserID: UserID})
	if err != nil {
		a.logger.Error().Msg(err.Error())
		return err
	}

	return nil
}
