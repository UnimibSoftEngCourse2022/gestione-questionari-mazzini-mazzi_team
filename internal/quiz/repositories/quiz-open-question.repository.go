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

func (r *QuizOpenQuestionRepository) FindAll(userID uint) ([]entities.QuizOpenQuestion, error) {
	var questions []entities.QuizOpenQuestion

	if err := r.db.Where("user_id = ? ", userID).Find(&questions).Error; err != nil {
		return nil, err
	}

	return questions, nil
}

func (r *QuizOpenQuestionRepository) Find(question *entities.QuizOpenQuestion) (*entities.QuizOpenQuestion, error) {
	questions := &entities.QuizOpenQuestion{}
	if err := r.db.Where(question).Find(&questions).Error; err != nil {
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

// func (r *QuizOpenQuestionRepository) Update(updatedQuestion entities.QuizOpenQuestion) (*entities.QuizOpenQuestion, error) {
// 	question := &entities.QuizOpenQuestion{}

// 	if err := r.db.Model(question).Where("id = ?", updatedQuestion.ID).Updates(updatedQuestion).Error; err != nil {
// 		return nil, err
// 	}

// 	return question, nil
// }

func (r *QuizOpenQuestionRepository) Delete(question *entities.QuizOpenQuestion) error {
	if err := r.db.Delete(question).Error; err != nil {
		return err
	}

	return nil
}
