package models

// User представляет пользователя системы
type User struct {
	ID       string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Username string `json:"username" example:"john_doe"`
	Email    string `json:"email" example:"john.doe@example.com"`
	Team     string `json:"team" example:"Red"`
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
	Error string `json:"error,omitempty"`
}

// PlayersInRoomResponse представляет ответ с информацией об игроках в комнате
type PlayersInRoomResponse struct {
	RoomID  string `json:"room_id" example:"room123"`
	Players []User `json:"players"`
	Error   string `json:"error,omitempty"`
}
