package closedquestion

import (
	"gorm.io/gorm"
)

type ClosedQuestion struct {
	gorm.Model
	ID         uint `gorm:"primaryKey"`
	Text       string
	ImageURL   string
	Category   string
	AnswerType string
	Answers    []MultipleChoice `gorm:"type:json"`
}

type MultipleChoice struct {
	Text       string `json:"text"`
	IsSelected bool   `json:"isSelected"`
}

// func (q *ClosedQuestion) ToModel() *ClosedQuestion {

// 	return &ClosedQuestion{
// 		Text:       q.Text,
// 		ImageURL:   q.ImageURL,
// 		Category:   q.Category,
// 		AnswerType: q.AnswerType,
// 		Answers:    q.Answers,
// 	}
// }
