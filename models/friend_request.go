// models/friend_request.go
package models

import (
	"time"
)

// FriendRequestStatus представляет статус запроса дружбы
type FriendRequestStatus string

const (
	Pending  FriendRequestStatus = "pending"
	Accepted FriendRequestStatus = "accepted"
	Rejected FriendRequestStatus = "rejected"
)

// FriendRequest представляет запрос дружбы между пользователями
type FriendRequest struct {
	ID         string              `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	FromUser   User                `json:"from_user" gorm:"foreignKey:FromUserID"`
	FromUserID string              `json:"from_user_id" gorm:"type:uuid;not null"`
	ToUser     User                `json:"to_user" gorm:"foreignKey:ToUserID"`
	ToUserID   string              `json:"to_user_id" gorm:"type:uuid;not null"`
	Status     FriendRequestStatus `json:"status" gorm:"type:varchar(10);default:'pending'"`
	CreatedAt  time.Time           `json:"created_at" gorm:"autoCreateTime"`
}

// SendFriendRequestResponse представляет ответ на отправку запроса дружбы
type SendFriendRequestResponse struct {
	Error Error `json:"error"`
}
