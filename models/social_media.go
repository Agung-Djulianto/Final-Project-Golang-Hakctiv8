package models

import "time"

type SocialMedia struct {
	ID             string `gorm:"primaryKey" json:"id_sosmed"`
	Name           string `gorm:"not null;type:varchar(255)" json:"name"`
	SocialMediaURL string `gorm:"not null;type:varchar(255);column:social_media_url" json:"social_media_url"`
	UserID         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type SocialMediaCreateRequest struct {
	Name           string `json:"name" valid:"required~Social Media name is required"`
	SocialMediaURL string `json:"social_media_url" valid:"required~Social Media URL is required"`
}

type SocialMediaUpdateRequest struct {
	Name           string `json:"name" valid:"required~Social Media name is required"`
	SocialMediaURL string `json:"social_media_url" valid:"required~Social Media URL is required"`
}

type SocialMediaCreateResponse struct {
	ID             string    `json:"id_sosmed"`
	UserID         string    `json:"user_id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	CreatedAt      time.Time `json:"created_at"`
}

type SocialMediaUpdateResponse struct {
	ID             string    `json:"id_sosmed"`
	UserID         string    `json:"user_id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type SocialMediaResponse struct {
	ID             string    `json:"id_sosmed"`
	UserID         string    `json:"user_id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type DeleteSocialMediaResponse struct {
	Message string `json:"message"`
}
