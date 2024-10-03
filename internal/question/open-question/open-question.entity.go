package openquestion

import "gorm.io/gorm"

// TODO: "default:18" e "not null"

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
