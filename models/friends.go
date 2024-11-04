// models/friends.go
package models

// GetFriendsResponse представляет ответ на получение списка друзей пользователя
type GetFriendsResponse struct {
	Friends []Friend `json:"friends"`
	Error   Error    `json:"error"`
}
