package services

import (
	common "form_management/common/logger"
	"form_management/internal/quiz/entities"
	"form_management/internal/quiz/repositories"

	"gorm.io/gorm"
)

type QuizService struct {
	logger                    *common.MyLogger
	repository                *repositories.QuizRepository
	quizOpenQuestionService   *QuizOpenQuestionService
	quizClosedQuestionService *QuizClosedQuestionService
}

func NewQuizService(logger *common.MyLogger, db *gorm.DB) *QuizService {
	return &QuizService{
		logger:                    logger,
		repository:                repositories.NewQuizRepository(db),
		quizOpenQuestionService:   NewQuizOpenQuestionService(logger, db),
		quizClosedQuestionService: NewQuizClosedQuestionService(logger, db),
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

func (a *QuizService) AddOpenQuestion(openQuestionID uint, quizID uint, userID uint) (*entities.Quiz, error) {
	quiz, err := a.FindById(quizID, userID)
	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	order := 0
	if len(quiz.OpenQuestions) > order {
		order = len(quiz.OpenQuestions) + 1
	}

	if len(quiz.ClosedQuestions) > order {
		order = len(quiz.ClosedQuestions) + 1
	}

	quizOpenQuestion, err := a.quizOpenQuestionService.Create(openQuestionID, quizID, order)
	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	quiz.OpenQuestions = append(quiz.OpenQuestions, *quizOpenQuestion)
	updatedQuiz, err := a.repository.Update(*quiz)
	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	return updatedQuiz, nil
}

func (a *QuizService) AddClosedQuestion(closedQuestionID uint, quizID uint, userID uint) (*entities.Quiz, error) {
	quiz, err := a.FindById(quizID, userID)
	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	order := 0
	if len(quiz.OpenQuestions) > order {
		order = len(quiz.OpenQuestions) + 1
	}

	if len(quiz.ClosedQuestions) > order {
		order = len(quiz.ClosedQuestions) + 1
	}

	quizOpenQuestion, err := a.quizOpenQuestionService.Create(closedQuestionID, quizID, order)
	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	quiz.OpenQuestions = append(quiz.OpenQuestions, *quizOpenQuestion)
	updatedQuiz, err := a.repository.Update(*quiz)
	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	return updatedQuiz, nil
}
