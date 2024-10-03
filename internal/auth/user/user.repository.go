package user

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Find(partialUser *User) (*User, error) {
	user := &User{}
	if err := r.db.Where(partialUser).Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Repository) Create(user *User) (*UserInfo, error) {
	newUser := &UserInfo{}
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}

	newUser = &UserInfo{
		Email: user.Email,
		Role:  user.Role,
		ID:    user.ID,
		Code:  user.Code,
	}

	return newUser, nil
}

func (r *Repository) Delete(user *User) error {
	if err := r.db.Delete(user).Error; err != nil {
		return err
	}
	return nil
}
