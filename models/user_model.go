package models

import "time"

type User struct {
	ID          string `gorm:"primaryKey" json:"user_id"`
	UserName    string `gorm:"not null;unique;type:varchar(30)" json:"user_name"`
	Email       string `gorm:"not null;unique;type:varchar(255)" json:"email"`
	Password    string `gorm:"not null;type:varchar(255)" json:"password"`
	Age         int    `gorm:"not null;size:2" json:"age"`
	Photos      []Photo
	Comment     []Comment
	SocialMedia []SocialMedia
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserRegisterRequest struct {
	Username string `json:"username" binding:"required" validate:"required"`
	Email    string `json:"email" binding:"required" validate:"required,email"`
	Password string `json:"password" binding:"required" validate:"required,min=6"`
	Age      int    `json:"age" binding:"required" validate:"required,gte=8"`
}

type UserRegisterResponse struct {
	ID        string    `json:"user_id"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
}

type UserLoginRequest struct {
	Username string `json:"username" valid:"required~Username is required"`
	Password string `json:"password" valid:"required~Password is required"`
}
type UserLoginResponse struct {
	Token string `json:"token"`
}
