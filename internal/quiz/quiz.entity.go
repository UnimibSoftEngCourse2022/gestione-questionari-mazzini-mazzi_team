package quiz

import (
	"form_management/internal/auth/user"
	closedquestion "form_management/internal/form/closed-question"
	openquestion "form_management/internal/form/open-question"

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

type QuizClosedQuestion struct {
	ClosedQuestionID uint                          `gorm:"primaryKey"`
	QuizID           uint                          `gorm:"primaryKey"`
	ClosedQuestion   closedquestion.ClosedQuestion `gorm:"foreignKey:ClosedQuestionID;references:ID"`
	Quiz             Quiz                          `gorm:"foreignKey:QuizID;references:ID"`
	Order            int                           `gorm:"not null"`
}

type QuizOpenQuestion struct {
	OpenQuestionID uint                      `gorm:"primaryKey"`
	QuizID         uint                      `gorm:"primaryKey"`
	OpenQuestion   openquestion.OpenQuestion `gorm:"foreignKey:OpenQuestionID;references:ID"`
	Quiz           Quiz                      `gorm:"foreignKey:QuizID;references:ID"`
	Order          int                       `gorm:"not null"`
}
