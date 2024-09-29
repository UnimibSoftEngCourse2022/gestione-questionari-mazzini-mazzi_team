package closedquestion

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type MultipleChoice struct {
	Text       string `json:"text"`
	IsSelected bool   `json:"is_selected"`
}

type ClosedQuestion struct {
	gorm.Model
	ID         uint `gorm:"primaryKey"`
	Text       string
	ImageURL   string
	Category   string
	AnswerType string
	Answers    MultipleChoiceArray `gorm:"type:jsonb[]" json:"answers"`
}

type MultipleChoiceArray []MultipleChoice

func (a MultipleChoiceArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *MultipleChoiceArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan value, not []byte")
	}
	return json.Unmarshal(bytes, a)
}
