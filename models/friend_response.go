// models/friend_response.go
package models

// ConfirmFriendRequestResponse представляет ответ на подтверждение/отклонение запроса дружбы
type ConfirmFriendRequestResponse struct {
	Error Error `json:"error"`
}
