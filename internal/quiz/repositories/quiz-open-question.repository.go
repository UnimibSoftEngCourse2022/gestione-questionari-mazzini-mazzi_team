package repositories

import (
	"form_management/internal/quiz/entities"

	"gorm.io/gorm"
)

type QuizOpenQuestionRepository struct {
	db *gorm.DB
}

func NewQuizOpenQuestionRepository(db *gorm.DB) *QuizOpenQuestionRepository {
	return &QuizOpenQuestionRepository{
		db: db,
	}
}

func (r *QuizOpenQuestionRepository) FindByQuizID(quizID uint) ([]entities.QuizOpenQuestion, error) {
	questions := []entities.QuizOpenQuestion{}
	if err := r.db.Where("quiz_id = ?", quizID).Find(&questions).Error; err != nil {
		return nil, err
	}

	return questions, nil
}

func (r *QuizOpenQuestionRepository) Create(question *entities.QuizOpenQuestion) (*entities.QuizOpenQuestion, error) {
	if err := r.db.Create(question).Error; err != nil {
		return nil, err
	}

	return question, nil
}

func (r *QuizOpenQuestionRepository) Delete(question *entities.QuizOpenQuestion) error {
	if err := r.db.Delete(question).Error; err != nil {
		return err
	}

	return nil
}
