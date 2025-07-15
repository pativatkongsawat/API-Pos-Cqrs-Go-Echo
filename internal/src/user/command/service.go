package user

import "gorm.io/gorm"

type DatabaseRequest struct {
	DB *gorm.DB
}
