package quiz

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) FindAll(userID uint) ([]Quiz, error) {
	var questions []Quiz

	if err := r.db.Where("user_id = ? ", userID).Find(&questions).Error; err != nil {
		return nil, err
	}

	return questions, nil
}

func (r *Repository) Find(question *Quiz) (*Quiz, error) {
	questions := &Quiz{}
	if err := r.db.Where(question).Find(&questions).Error; err != nil {
		return nil, err
	}

	return questions, nil
}

func (r *Repository) Create(question *Quiz) (*Quiz, error) {
	if err := r.db.Create(question).Error; err != nil {
		return nil, err
	}

	return question, nil
}

func (r *Repository) Update(updatedQuestion Quiz) (*Quiz, error) {
	question := &Quiz{}

	if err := r.db.Model(question).Where("id = ?", updatedQuestion.ID).Updates(updatedQuestion).Error; err != nil {
		return nil, err
	}

	return question, nil
}

func (r *Repository) Delete(question *Quiz) error {
	if err := r.db.Delete(question).Error; err != nil {
		return err
	}

	return nil
}
