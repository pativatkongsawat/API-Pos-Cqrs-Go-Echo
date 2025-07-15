package user

import (
	"echo_cqrs/internal/src/user"

	"gorm.io/gorm"
)

type DatabaseRequest struct {
	DB *gorm.DB
}

func (u *DatabaseRequest) GetAllUser() ([]user.UserModel, error) {

	user := []user.UserModel{}

	if err := u.DB.Debug().Find(&user).Error; err != nil {

		return nil, err

	}

	return user, nil
}
