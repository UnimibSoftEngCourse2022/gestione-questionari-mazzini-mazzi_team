package closedquestion

import (
	common "form_management/common/logger"

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

func (a *Service) FindAllByIds(ids []uint) ([]ClosedQuestion, error) {
	questions, err := a.repository.FindAllByIds(ids)
	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}
	return questions, nil
}

func (a *Service) FindById(id uint) (*ClosedQuestion, error) {
	partialQuestion := ClosedQuestion{ID: id}
	question, err := a.repository.Find(&partialQuestion)

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	return question, nil
}

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

func (a *Service) Update(id uint, text string, imageURL string, category string, answersText []string) (*ClosedQuestion, error) {
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

	updatedQuestion := ClosedQuestion{
		Text:       text,
		ImageURL:   imageURL,
		Category:   category,
		AnswerType: "CLOSED_QUESTION",
		Answers:    answers,
	}
	question, err := a.repository.Update(updatedQuestion)

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	return question, nil
}

func (a *Service) Delete(id uint) error {
	err := a.repository.Delete(&ClosedQuestion{ID: id})

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return err
	}

	return nil
}
