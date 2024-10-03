package user

import (
	"encoding/base64"
	"errors"
	common "form_management/common/logger"
	t "form_management/common/type"
	"math/rand/v2"
	"time"

	"golang.org/x/crypto/bcrypt"
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

func generateCode() string {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	randomComponent := rand.Uint32()

	combined := (uint64(timestamp) << 32) | uint64(randomComponent)
	bytes := make([]byte, 12)
	for i := 11; i >= 0; i-- {
		bytes[i] = byte(combined & 0xff)
		combined >>= 8
	}
	customID := base64.URLEncoding.EncodeToString(bytes)
	customID = customID[:len(customID)-1]

	return customID
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func checkPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (a *Service) IsLogged(id uint) (*UserInfo, error) {
	partialUser := User{ID: id}
	user, err := a.repository.Find(&partialUser)

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	userInfo := UserInfo{
		ID:   user.ID,
		Code: user.Code,
		Role: t.GUEST,
	}

	return &userInfo, nil
}

func (a *Service) LoginGuest(code string) (*UserInfo, error) {
	partialUser := User{Code: code}
	user, err := a.repository.Find(&partialUser)

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	if user.Email != "" {
		return nil, errors.New("This user is not a guest")
	}

	userInfo := UserInfo{
		ID:   user.ID,
		Code: user.Code,
		Role: t.GUEST,
	}

	return &userInfo, nil
}

func (a *Service) LoginUser(email string, password string) (*UserInfo, error) {
	partialUser := User{Email: email}
	user, err := a.repository.Find(&partialUser)

	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	if checkPasswordHash(password, user.Password) {
		userInfo := UserInfo{
			ID:    user.ID,
			Code:  user.Code,
			Email: user.Email,
			Role:  t.USER,
		}
		return &userInfo, nil
	}

	return nil, errors.New("password not correct")
}

func (a *Service) RegisterUser(email string, password string) (*UserInfo, error) {
	code := generateCode()
	hashedPassword, err := hashPassword(password)
	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	existingUser, err := a.repository.Find(&User{Email: email})
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	if existingUser.Email != "" {
		return nil, errors.New("user with this email already exists")
	}

	newUser := User{
		Email:    email,
		Password: hashedPassword,
		Role:     t.USER,
		Code:     code,
	}

	user, err := a.repository.Create(&newUser)
	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	return user, nil
}

func (a *Service) RegisterGuest() (*UserInfo, error) {
	code := generateCode()
	newUser := User{
		Role: t.GUEST,
		Code: code,
	}

	user, err := a.repository.Create(&newUser)
	if err != nil {
		a.logger.Error().Msg(err.Error())
		return nil, err
	}

	return user, nil
}
