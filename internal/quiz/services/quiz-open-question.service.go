package services

import (
	common "form_management/common/logger"
	openquestion "form_management/internal/question/open-question"
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

func (a *QuizOpenQuestionService) FindByQuizID(quizID uint) ([]entities.QuizOpenQuestion, error) {
	partialQuiz := entities.QuizOpenQuestion{QuizID: quizID}
	quiz, err := a.repository.Find(&partialQuiz)

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	return quiz, nil
}

func (a *QuizOpenQuestionService) Create(
	openQuestion openquestion.OpenQuestion,
	quiz entities.Quiz,
	order int,
) (*entities.QuizOpenQuestion, error) {

	quizOpenQuestion := &entities.QuizOpenQuestion{
		OpenQuestion: openQuestion,
		Quiz:         quiz,
		Order:        order,
	}

	createdQuizOpenQuestion, err := a.repository.Create(quizOpenQuestion)
	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	return createdQuizOpenQuestion, nil
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
