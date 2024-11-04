package models

import (
	"time"
)

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
