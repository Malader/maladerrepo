package handlers

import (
	"net/http"

	"github.com/Malader/maladerrepo/models"
	"github.com/gin-gonic/gin"
)

// Mocked room data for demonstration purposes
var roomPlayers = map[string][]models.User{
	"room123": {
		{ID: "1", Username: "john_doe", Email: "john.doe@example.com", Team: "Red"},
		{ID: "2", Username: "jane_smith", Email: "jane.smith@example.com", Team: "Blue"},
	},
	"room456": {
		{ID: "3", Username: "alice_jones", Email: "alice.jones@example.com", Team: "Green"},
	},
}

// GetPlayersInRoom обрабатывает запросы на получение информации об игроках в комнате
// @Summary Получение информации об игроках в комнате
// @Description Возвращает список игроков в указанной игровой комнате
// @Tags room
// @Accept json
// @Produce json
// @Param room_id path string true "Идентификатор комнаты" example:"room123"
// @Success 200 {object} models.PlayersInRoomResponse
// @Failure 404 {object} models.PlayersInRoomResponse
// @Router /rooms/{room_id}/players [get]
func GetPlayersInRoom(c *gin.Context) {
	roomID := c.Param("room_id")
	players, exists := roomPlayers[roomID]
	if !exists {
		c.JSON(http.StatusNotFound, models.PlayersInRoomResponse{
			RoomID: roomID,
			Error:  "Комната не найдена",
		})
		return
	}

	c.JSON(http.StatusOK, models.PlayersInRoomResponse{
		RoomID:  roomID,
		Players: players,
	})
}
