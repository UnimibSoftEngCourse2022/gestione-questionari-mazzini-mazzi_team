package openquestion

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

func (r *Repository) FindAll() ([]OpenQuestion, error) {
	var questions []OpenQuestion

	if err := r.db.Find(&questions).Error; err != nil {
		return nil, err
	}

	return questions, nil
}

func (r *Repository) Find(id int) (*OpenQuestion, error) {
	questions := &OpenQuestion{}
	if err := r.db.Where("id = ?", id).First(&questions).Error; err != nil {
		return nil, err
	}

	return questions, nil
}

func (r *Repository) Create(question *OpenQuestion) (*OpenQuestion, error) {
	if err := r.db.Create(question).Error; err != nil {
		return nil, err
	}

	return question, nil
}