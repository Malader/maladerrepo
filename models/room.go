package models

type Room struct {
	ID      string `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name    string `json:"name" gorm:"uniqueIndex;not null"`
	Players []User `json:"players" gorm:"many2many:room_players;"`
}

type PlayersInRoomResponse struct {
	RoomID  string `json:"room_id" example:"room123"`
	Players []User `json:"players"`
	Error   Error  `json:"error,omitempty"`
}
