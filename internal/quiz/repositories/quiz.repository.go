package repositories

import (
	"form_management/internal/quiz/entities"

	"gorm.io/gorm"
)

type QuizRepository struct {
	db *gorm.DB
}

func NewQuizRepository(db *gorm.DB) *QuizRepository {
	return &QuizRepository{
		db: db,
	}
}

func (r *QuizRepository) FindAll(userID uint) ([]entities.Quiz, error) {
	var questions []entities.Quiz

	if err := r.db.Where("user_id = ? ", userID).Find(&questions).Error; err != nil {
		return nil, err
	}

	return questions, nil
}

func (r *QuizRepository) Find(question *entities.Quiz) (*entities.Quiz, error) {
	questions := &entities.Quiz{}
	if err := r.db.Where(question).Find(&questions).Error; err != nil {
		return nil, err
	}

	return questions, nil
}

func (r *QuizRepository) Create(question *entities.Quiz) (*entities.Quiz, error) {
	if err := r.db.Create(question).Error; err != nil {
		return nil, err
	}

	return question, nil
}

func (r *QuizRepository) Update(updatedQuestion entities.Quiz) (*entities.Quiz, error) {
	question := &entities.Quiz{}

	if err := r.db.Model(question).Where("id = ?", updatedQuestion.ID).Updates(updatedQuestion).Error; err != nil {
		return nil, err
	}

	return question, nil
}

func (r *QuizRepository) Delete(question *entities.Quiz) error {
	if err := r.db.Delete(question).Error; err != nil {
		return err
	}

	return nil
}
