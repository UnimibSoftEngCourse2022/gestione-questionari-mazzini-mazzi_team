package openquestion

import "gorm.io/gorm"

type OpenQuestion struct {
	gorm.Model
	ID         uint `gorm:"primaryKey"`
	Text       string
	ImageURL   string
	AnswerType string
	MinChar    int
	MaxChar    int
	Category   string
}
