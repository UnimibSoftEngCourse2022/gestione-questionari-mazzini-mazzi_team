package closedquestion

import (
	"form_management/common"

	"gorm.io/gorm"
)

type Service struct {
	logger     *common.MyLogger
	repository *Repository
}

func NewService(logger *common.MyLogger, db *gorm.DB) *Service {
	return &Service{
		logger:     logger,
		repository: NewRepository(db),
	}
}

func (a *Service) FindAll() ([]ClosedQuestion, error) {
	questions, err := a.repository.FindAll()

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	return questions, nil
}

// func (a *Service) Find() (*ClosedQuestion, error) {

// 	question, err := a.repository.Find()

// 	if err != nil {
// 		a.logger.Error().Msg(err.Error())
// 		return nil, err
// 	}

// 	return question, nil
// }

func (a *Service) Create(
	text string,
	imageURL string,
	category string,
	answersText []string,
) (*ClosedQuestion, error) {

	answers := []MultipleChoice{}
	for _, answer := range answersText {
		answers = append(
			answers,
			MultipleChoice{
				Text:       answer,
				IsSelected: false,
			},
		)
	}

	question := &ClosedQuestion{
		Text:       text,
		ImageURL:   imageURL,
		Category:   category,
		AnswerType: "CLOSED_QUESTION",
		Answers:    answers,
	}
	question, err := a.repository.Create(question)

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	return question, nil
}
