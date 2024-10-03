package entities

import (
	closedquestion "form_management/internal/question/closed-question"
)

type QuizClosedQuestion struct {
	ClosedQuestionID uint                          `gorm:"primaryKey"`
	QuizID           uint                          `gorm:"primaryKey"`
	ClosedQuestion   closedquestion.ClosedQuestion `gorm:"foreignKey:ClosedQuestionID;references:ID"`
	Quiz             Quiz                          `gorm:"foreignKey:QuizID;references:ID"`
	Order            int                           `gorm:"not null"`
}
