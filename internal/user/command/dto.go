package user

type CreateUser struct {
	UserId   string `gorm:"not null;unique" json:"user_id"`
	UserName string `gorm:"not null" json:"user_name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `gorm:"not null" json:"password"`
	RoleId   int    `gorm:"not null" json:"role_id"`
	StoreId  int    `gorm:"not null" json:"store_id"`
}

type UpdateUser struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleId   int    `json:"role_id"`
	StoreId  int    `json:"store_id"`
}
