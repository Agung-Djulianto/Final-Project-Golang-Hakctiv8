package models

import "time"

type Photo struct {
	ID        string `gorm:"primaryKey" json:"id_photo"`
	Title     string `gorm:"not null;type:varchar(100)" json:"tittle"`
	Caption   string `gorm:"not null;type:varchar(255)" json:"caption"`
	PhotoURL  string `gorm:"not null;type:varchar(255)" json:"photo_url"`
	UserID    string
	Comment   []Comment
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PhotoCreateRequest struct {
	Title    string `json:"title" valid:"required~Title is required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" valid:"required~Photo URL is required"`
}

type PhotoUpdateRequest struct {
	Title    string `json:"title" valid:"required~Title is required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" valid:"required~Photo URL is required"`
}

type PhotoCreateResponse struct {
	ID        string    `json:"id_photo"`
	UserID    string    `json:"id_user"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	CreatedAt time.Time `json:"created_at"`
}

type PhotoUpdateResponse struct {
	ID        string    `json:"id_photo"`
	UserID    string    `json:"id_user"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PhotoResponse struct {
	ID        string    `json:"id_photo"`
	UserID    string    `json:"id_user"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeletePhotoResponse struct {
	Message string `json:"message"`
}
