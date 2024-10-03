package services

import (
	common "form_management/common/logger"
	"form_management/internal/quiz/entities"
	"form_management/internal/quiz/repositories"

	"gorm.io/gorm"
)

type QuizClosedQuestionService struct {
	logger     *common.MyLogger
	repository *repositories.QuizClosedQuestionRepository
}

func NewQuizClosedQuestionService(logger *common.MyLogger, db *gorm.DB) *QuizClosedQuestionService {
	return &QuizClosedQuestionService{
		logger:     logger,
		repository: repositories.NewQuizClosedQuestionRepository(db),
	}
}

func (a *QuizClosedQuestionService) FindByQuizID(quizID uint) (*entities.QuizClosedQuestion, error) {
	partialQuiz := entities.QuizClosedQuestion{QuizID: quizID}
	quiz, err := a.repository.Find(&partialQuiz)

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	return quiz, nil
}

func (a *QuizClosedQuestionService) Create(
	closedQuestionID uint,
	quizID uint,
	order int,
) (*entities.QuizClosedQuestion, error) {

	quizClosedQuestion := &entities.QuizClosedQuestion{
		ClosedQuestionID: closedQuestionID,
		QuizID:           quizID,
		Order:            order,
	}
	quiz, err := a.repository.Create(quizClosedQuestion)

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	return quiz, nil
}

func (a *QuizClosedQuestionService) Delete(
	closedQuestionID uint,
	quizID uint,
) error {
	err := a.repository.Delete(
		&entities.QuizClosedQuestion{
			ClosedQuestionID: closedQuestionID,
			QuizID:           quizID,
		},
	)
	if err != nil {
		a.logger.Error().Msg(err.Error())
		return err
	}

	return nil
}
