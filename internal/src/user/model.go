package user

import "time"

type UserModel struct {
	ID       string    `gorm:"primaryKey" json:"id"`
	UserId   string    `gorm:"not null;unique" json:"user_id"`
	UserName string    `gorm:"not null" json:"user_name"`
	Email    string    `gorm:"unique" json:"email"`
	Password string    `gorm:"not null" json:"password"`
	RoleId   int       `gorm:"not null" json:"role_id"`
	StoreId  int       `gorm:"not null" json:"store_id"`
	IsActive bool      `json:"is_active"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

func (UserModel) TableName() string {

	return "users"

}
