package models

type UserModel struct {
	ID       *int64 `gorm:"column:id;primary_key"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}

func (UserModel) TableName() string {
	return "users"
}
