package services

import (
	common "form_management/common/logger"
	"form_management/internal/quiz/entities"
	"form_management/internal/quiz/repositories"

	"gorm.io/gorm"
)

type QuizOpenQuestionService struct {
	logger     *common.MyLogger
	repository *repositories.QuizOpenQuestionRepository
}

func NewQuizOpenQuestionService(logger *common.MyLogger, db *gorm.DB) *QuizOpenQuestionService {
	return &QuizOpenQuestionService{
		logger:     logger,
		repository: repositories.NewQuizOpenQuestionRepository(db),
	}
}

func (a *QuizOpenQuestionService) FindByQuizID(quizID uint) (*entities.QuizOpenQuestion, error) {
	partialQuiz := entities.QuizOpenQuestion{QuizID: quizID}
	quiz, err := a.repository.Find(&partialQuiz)

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	return quiz, nil
}

func (a *QuizOpenQuestionService) Create(
	openQuestionID uint,
	quizID uint,
	order int,
) (*entities.QuizOpenQuestion, error) {

	quizClosedQuestion := &entities.QuizOpenQuestion{
		OpenQuestionID: openQuestionID,
		QuizID:         quizID,
		Order:          order,
	}
	quiz, err := a.repository.Create(quizClosedQuestion)

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	return quiz, nil
}

func (a *QuizOpenQuestionService) Delete(
	openQuestionID uint,
	quizID uint,
) error {
	err := a.repository.Delete(
		&entities.QuizOpenQuestion{
			OpenQuestionID: openQuestionID,
			QuizID:         quizID,
		},
	)
	if err != nil {
		a.logger.Error().Msg(err.Error())
		return err
	}

	return nil
}
