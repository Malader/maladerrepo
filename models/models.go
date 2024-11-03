package models

import "gorm.io/gorm"

// User представляет пользователя системы
type User struct {
	ID           string         `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Username     string         `json:"username" gorm:"uniqueIndex;not null"`
	Email        string         `json:"email" gorm:"uniqueIndex;not null"`
	PasswordHash string         `json:"-" gorm:"not null"`
	Team         string         `json:"team"`
	CreatedAt    gorm.DeletedAt `json:"-"`
}

// VerifyCredentialsRequest представляет запрос на проверку учетных данных
type VerifyCredentialsRequest struct {
	Username string `json:"username" binding:"required" example:"john_doe"`
	Password string `json:"password" binding:"required" example:"securepassword"`
}

// VerifyCredentialsResponse представляет ответ на проверку учетных данных
type VerifyCredentialsResponse struct {
	Valid bool   `json:"valid" example:"true"`
	User  User   `json:"user,omitempty"`
	Token string `json:"token,omitempty"`
	Error string `json:"error,omitempty"`
}

// PlayersInRoomResponse представляет ответ с информацией об игроках в комнате
type PlayersInRoomResponse struct {
	RoomID  string `json:"room_id" example:"room123"`
	Players []User `json:"players"`
	Error   string `json:"error,omitempty"`
}
