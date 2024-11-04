// models/recovery.go
package models

import (
	"time"
)

// RecoveryRequest представляет запрос на восстановление пароля (фаза 1)
type RecoveryRequest struct {
	Email string `json:"email" binding:"required,email,max=256" example:"v.pupkin@g.nsu.ru"`
}

// RecoveryResponse представляет ответ на восстановление пароля (фаза 1)
type RecoveryResponse struct {
	Error Error `json:"error"`
}

// RecoveryCompleteRequest представляет запрос на восстановление пароля (фаза 2)
type RecoveryCompleteRequest struct {
	Password string `json:"password" binding:"required,len=64" example:"6a4a61f57bccf059abb82fc95589ebc428629326ab965390f25224e262455beb"`
}

// RecoveryCompleteResponse представляет ответ на восстановление пароля (фаза 2)
type RecoveryCompleteResponse struct {
	Error Error `json:"error"`
}

// PasswordRecovery представляет токен восстановления пароля
type PasswordRecovery struct {
	ID             string    `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID         string    `json:"user_id" gorm:"not null"`
	RecoverySuffix string    `json:"recoverySuffix" gorm:"uniqueIndex;not null;size:128"`
	ExpiresAt      time.Time `json:"expires_at" gorm:"not null"`
}

// RecoveryToken представляет токен для восстановления пароля
type RecoveryToken struct {
	ID             string    `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID         string    `json:"user_id" gorm:"type:uuid;not null"`
	RecoverySuffix string    `json:"recovery_suffix" gorm:"uniqueIndex;not null"`
	ExpiresAt      time.Time `json:"expires_at" gorm:"not null"`
}
