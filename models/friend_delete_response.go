// models/friend_delete_response.go
package models

// DeleteFriendResponse представляет ответ на удаление из друзей
type DeleteFriendResponse struct {
	Error Error `json:"error"`
}
