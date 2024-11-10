package models

import (
	"time"
)

type Friend struct {
	Username string `json:"username" example:"VasyaPupkin"`
	Image    string `json:"image,omitempty" example:"4AAQSkZJRgABAQEAAAAAAADfjadFHADHda"`
	Status   string `json:"status" example:"active"`
}

type FriendRequestStatus string

const (
	Pending  FriendRequestStatus = "pending"
	Accepted FriendRequestStatus = "accepted"
	Rejected FriendRequestStatus = "rejected"
)

type FriendRequest struct {
	ID         string              `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	FromUserID string              `json:"from_user_id" gorm:"type:uuid;not null"`
	ToUserID   string              `json:"to_user_id" gorm:"type:uuid;not null"`
	Status     FriendRequestStatus `json:"status" gorm:"type:varchar(10);default:'pending'"`
	CreatedAt  time.Time           `json:"created_at" gorm:"autoCreateTime"`

	FromUser User `json:"from_user" gorm:"foreignKey:FromUserID"`
	ToUser   User `json:"to_user" gorm:"foreignKey:ToUserID"`
}

type SendFriendRequestResponse struct {
	Error Error `json:"error"`
}

type ConfirmFriendRequestResponse struct {
	Error Error `json:"error"`
}

type DeleteFriendResponse struct {
	Error Error `json:"error"`
}

type GetFriendsResponse struct {
	Friends []Friend `json:"friends"`
	Error   Error    `json:"error"`
}
