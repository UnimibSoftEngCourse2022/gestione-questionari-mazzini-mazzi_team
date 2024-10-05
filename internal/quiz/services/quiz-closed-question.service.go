package services

import (
	common "form_management/common/logger"
	closedquestion "form_management/internal/question/closed-question"
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

func (a *QuizClosedQuestionService) FindByQuizID(quizID uint) ([]entities.QuizClosedQuestion, error) {
	quiz, err := a.repository.FindByQuizId(quizID)
	for _, q := range quiz {
		a.logger.Info().Msgf("Found quizID: %d", q.ClosedQuestionID)
	}

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	return quiz, nil
}

func (a *QuizClosedQuestionService) Create(
	closedQuestion closedquestion.ClosedQuestion,
	quiz entities.Quiz,
	order int,
) (*entities.QuizClosedQuestion, error) {

	quizClosedQuestion := &entities.QuizClosedQuestion{
		ClosedQuestion: closedQuestion,
		Quiz:           quiz,
		Order:          order,
	}
	quizClosedQuestion, err := a.repository.Create(quizClosedQuestion)

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	return quizClosedQuestion, nil
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
