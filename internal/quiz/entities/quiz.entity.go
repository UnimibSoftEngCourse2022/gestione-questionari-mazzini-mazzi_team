package entities

import (
	"form_management/internal/auth/user"

	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	ID              uint                 `gorm:"primaryKey"`
	ClosedQuestions []QuizClosedQuestion `gorm:"foreignKey:QuizID"`
	OpenQuestions   []QuizOpenQuestion   `gorm:"foreignKey:QuizID"`
	Title           string               `gorm:"not null"`
	UserID          uint                 //`gorm:"not null"`
	User            user.User            `gorm:"foreignKey:UserID;references:ID"`
}
