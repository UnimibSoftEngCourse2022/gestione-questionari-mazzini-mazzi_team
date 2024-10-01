package quiz

import (
	closedquestion "form_management/internal/form/closed-question"
	openquestion "form_management/internal/form/open-question"

	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	ID              uint                 `gorm:"primaryKey"`
	ClosedQuestions []QuizClosedQuestion `gorm:"many2many:quiz_closed_questions;"`
	OpenQuestions   []QuizOpenQuestion   `gorm:"many2many:quiz_open_questions;"`
	Title           string
}

type QuizClosedQuestion struct {
	QuizID           Quiz                          `gorm:"primaryKey"`
	ClosedQuestionID closedquestion.ClosedQuestion `gorm:"primaryKey"`
	Order            int
}

type QuizOpenQuestion struct {
	QuizID         Quiz                      `gorm:"primaryKey"`
	OpenQuestionID openquestion.OpenQuestion `gorm:"primaryKey"`
	Order          int
}
