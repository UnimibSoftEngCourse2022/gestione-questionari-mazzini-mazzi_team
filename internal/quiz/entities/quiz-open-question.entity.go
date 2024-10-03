package entities

import (
	openquestion "form_management/internal/question/open-question"
)

type QuizOpenQuestion struct {
	OpenQuestionID uint                      `gorm:"primaryKey"`
	QuizID         uint                      `gorm:"primaryKey"`
	OpenQuestion   openquestion.OpenQuestion `gorm:"foreignKey:OpenQuestionID;references:ID"`
	Quiz           Quiz                      `gorm:"foreignKey:QuizID;references:ID"`
	Order          int                       `gorm:"not null"`
}
