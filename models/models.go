// models/models.go
package models

import (
	"gorm.io/gorm"
	"time"
)

// User представляет пользователя системы
type User struct {
	ID               string         `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Username         string         `json:"username" gorm:"uniqueIndex;not null"`
	Email            string         `json:"email" gorm:"uniqueIndex;not null"`
	PasswordHash     string         `json:"-" gorm:"not null"`
	Team             string         `json:"team"`
	Image            string         `json:"image" gorm:"type:text"`
	RegistrationDate time.Time      `json:"registrationDate" gorm:"autoCreateTime"`
	LastActivityDate time.Time      `json:"lastActivityDate" gorm:"autoUpdateTime"`
	CreatedAt        gorm.DeletedAt `json:"-"`
	Friends          []User         `json:"-" gorm:"many2many:user_friends;"`
	BlacklistedUsers []User         `json:"-" gorm:"many2many:user_blacklists;"`
}

// Room представляет игровую комнату
type Room struct {
	ID      string `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name    string `json:"name" gorm:"uniqueIndex;not null"`
	Players []User `json:"players" gorm:"many2many:room_players;"`
}

// PlayersInRoomResponse представляет ответ с информацией об игроках в комнате
type PlayersInRoomResponse struct {
	RoomID  string `json:"room_id" example:"room123"`
	Players []User `json:"players"`
	Error   Error  `json:"error,omitempty"`
}
