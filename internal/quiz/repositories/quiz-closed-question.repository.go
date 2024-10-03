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

func (r *QuizClosedQuestionRepository) FindAll(userID uint) ([]entities.QuizClosedQuestion, error) {
	var questions []entities.QuizClosedQuestion

	if err := r.db.Where("user_id = ? ", userID).Find(&questions).Error; err != nil {
		return nil, err
	}

	return questions, nil
}

func (r *QuizClosedQuestionRepository) Find(question *entities.QuizClosedQuestion) (*entities.QuizClosedQuestion, error) {
	questions := &entities.QuizClosedQuestion{}
	if err := r.db.Where(question).Find(&questions).Error; err != nil {
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

// func (r *QuizClosedQuestionRepository) Update(updatedQuestion entities.QuizClosedQuestion) (*entities.QuizClosedQuestion, error) {
// 	question := &entities.QuizClosedQuestion{}

// 	if err := r.db.Model(question).Where("id = ?", updatedQuestion.ID).Updates(updatedQuestion).Error; err != nil {
// 		return nil, err
// 	}

// 	return question, nil
// }

func (r *QuizClosedQuestionRepository) Delete(question *entities.QuizClosedQuestion) error {
	if err := r.db.Delete(question).Error; err != nil {
		return err
	}

	return nil
}
