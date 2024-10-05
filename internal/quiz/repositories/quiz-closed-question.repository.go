package repositories

import (
	"form_management/internal/quiz/entities"

	"gorm.io/gorm"
)

type QuizClosedQuestionRepository struct {
	db *gorm.DB
}

func NewQuizClosedQuestionRepository(db *gorm.DB) *QuizClosedQuestionRepository {
	return &QuizClosedQuestionRepository{
		db: db,
	}
}

func (r *QuizClosedQuestionRepository) FindByQuizId(quizID uint) ([]entities.QuizClosedQuestion, error) {
	questions := []entities.QuizClosedQuestion{}
	if err := r.db.Where("quiz_id = ?", quizID).Find(&questions).Error; err != nil {
		return nil, err
	}

	return questions, nil
}

func (r *QuizClosedQuestionRepository) Create(question *entities.QuizClosedQuestion) (*entities.QuizClosedQuestion, error) {
	if err := r.db.Create(question).Error; err != nil {
		return nil, err
	}

	return question, nil
}

func (r *QuizClosedQuestionRepository) Delete(question *entities.QuizClosedQuestion) error {
	if err := r.db.Delete(question).Error; err != nil {
		return err
	}

	return nil
}
