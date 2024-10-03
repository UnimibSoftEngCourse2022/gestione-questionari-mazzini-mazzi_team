package closedquestion

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

func (r *Repository) FindAll() ([]ClosedQuestion, error) {
	var questions []ClosedQuestion

	if err := r.db.Find(&questions).Error; err != nil {
		return nil, err
	}

	return questions, nil
}

func (r *Repository) Find(question *ClosedQuestion) (*ClosedQuestion, error) {
	questions := &ClosedQuestion{}
	if err := r.db.Where(question).Find(&questions).Error; err != nil {
		return nil, err
	}

	return questions, nil
}

func (r *Repository) Create(question *ClosedQuestion) (*ClosedQuestion, error) {
	if err := r.db.Create(question).Error; err != nil {
		return nil, err
	}

	return question, nil
}

func (r *Repository) Update(updatedQuestion ClosedQuestion) (*ClosedQuestion, error) {
	question := &ClosedQuestion{}
	if err := r.db.Model(question).Where("id = ?", updatedQuestion.ID).Updates(updatedQuestion).Error; err != nil {
		return nil, err
	}

	return question, nil
}

func (r *Repository) Delete(question *ClosedQuestion) error {
	if err := r.db.Delete(question).Error; err != nil {
		return err
	}

	return nil
}
