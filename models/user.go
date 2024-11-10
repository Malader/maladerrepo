package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID               string         `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Username         string         `json:"username" gorm:"uniqueIndex;not null" binding:"required,max=25"`
	Email            string         `json:"email" gorm:"uniqueIndex;not null" binding:"required,email,max=256"`
	PasswordHash     string         `json:"-" gorm:"not null"`
	Team             string         `json:"team,omitempty"`
	Image            string         `json:"image,omitempty" gorm:"type:text" binding:"omitempty,base64"`
	RegistrationDate time.Time      `json:"registrationDate" gorm:"autoCreateTime"`
	LastActivityDate time.Time      `json:"lastActivityDate" gorm:"autoUpdateTime"`
	CreatedAt        time.Time      `json:"-" gorm:"autoCreateTime"`
	DeletedAt        gorm.DeletedAt `json:"-"`

	Friends          []User `json:"-" gorm:"many2many:user_friends;"`
	BlacklistedUsers []User `json:"-" gorm:"many2many:user_blacklists;"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,max=25" example:"vasyaPupkin"`
	Password string `json:"password" binding:"required,min=8,max=64" example:"securePassword123"`
	Email    string `json:"email" binding:"required,email,max=256" example:"v.pupkin@g.nsu.ru"`
}

type RegisterResponse struct {
	Error Error `json:"error"`
}

type AuthorizeRequest struct {
	Username string `json:"username" binding:"required,max=25" example:"vasyaPupkin"`
	Password string `json:"password" binding:"required,min=8,max=64" example:"securePassword123"`
}

type AuthorizeResponse struct {
	Error Error  `json:"error"`
	Token string `json:"token,omitempty"`
}

type UpdateProfileRequest struct {
	Username string `json:"username,omitempty" binding:"omitempty,max=25" example:"vasyaPupkin"`
	Password string `json:"password,omitempty" binding:"omitempty,min=8,max=64" example:"newSecurePassword123"`
	Image    string `json:"image,omitempty" binding:"omitempty,base64" example:"TWFuIGlzIGRpc3Rpbmd1aXNoZWQsIG5vdCBvbmx5IGJ5IGhpcyByZWFzb24sIGJ1dCAuLi4="`
}

type UpdateProfileResponse struct {
	Error Error `json:"error"`
}

type GetProfileResponse struct {
	Email            string    `json:"email" example:"example@example.com"`
	Username         string    `json:"username" example:"ExampleUser"`
	Password         string    `json:"password" example:"hashedPassword123"`
	RegistrationDate time.Time `json:"registrationDate" example:"2024-01-01T00:00:00Z"`
	LastActivityDate time.Time `json:"lastActivityDate" example:"2024-05-22T12:34:56Z"`
	Image            string    `json:"image" example:"4AAQSkZJRgABAQEAAAAAAADfjadFHADHda"`
	Error            Error     `json:"error"`
}

type RecoveryRequest struct {
	Email string `json:"email" binding:"required,email,max=256" example:"v.pupkin@g.nsu.ru"`
}

type RecoveryResponse struct {
	Error Error `json:"error"`
}

type RecoveryToken struct {
	ID             string    `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID         string    `json:"user_id" gorm:"type:uuid;not null"`
	RecoverySuffix string    `json:"recovery_suffix" gorm:"uniqueIndex;not null"`
	ExpiresAt      time.Time `json:"expires_at" gorm:"not null"`
}

type ChangePasswordRequest struct {
	NewPassword string `json:"newPassword" binding:"required,min=8,max=64" example:"newSecurePassword123"`
}

type ChangePasswordResponse struct {
	Error Error `json:"error"`
}
