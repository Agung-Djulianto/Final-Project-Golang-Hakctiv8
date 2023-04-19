package models

import "time"

type Comment struct {
	ID        string    `gorm:"primaryKey" json:"id_comment"`
	UserID    string    `json:"id_user"`
	PhotoID   string    `json:"id_photo"`
	Massage   string    `gorm:"not null" json:"massage"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommentCreateRequest struct {
	Massage string `json:"message" valid:"required~Message is required"`
}

type CommentUpdateRequest struct {
	Message string `json:"message" valid:"required~Message is required"`
}

type CommentCreateResponse struct {
	ID        string    `json:"id_comment"`
	UserID    string    `json:"id_user"`
	PhotoID   string    `json:"id_photo"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentResponse struct {
	ID        string    `json:"id_comment"`
	UserID    string    `json:"id_user"`
	PhotoID   string    `json:"id_photo"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommentUpdateResponse struct {
	ID        string    `json:"id_comment"`
	UserID    string    `json:"id_user"`
	PhotoID   string    `json:"id_photo"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteCommentResponse struct {
	Message string `json:"message"`
}
