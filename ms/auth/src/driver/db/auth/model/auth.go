package model

type Auth struct {
	UserID   string `gorm:"column:user_id" json:"user_id"`
	Password string `gorm:"column:password" json:"password"`
	Email    string `gorm:"column:email" json:"email"`
	Scope    string `gorm:"column:scope" json:"scope"`
}

func (a *Auth) TableName() string {
	return "auth"
}
