package handlers

import (
	"net/http"

	"github.com/Malader/maladerrepo/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Room struct {
	ID      string        `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name    string        `json:"name" gorm:"uniqueIndex;not null"`
	Players []models.User `json:"players" gorm:"many2many:room_players;"`
}

type PlayersInRoomResponse struct {
	RoomID  string        `json:"room_id" example:"room123"`
	Players []models.User `json:"players"`
	Error   models.Error  `json:"error,omitempty"`
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
// @Failure 500 {object} models.PlayersInRoomResponse
// @Router /rooms/{room_id}/players [get]
func GetPlayersInRoom(c *gin.Context) {
	roomID := c.Param("room_id")

	var room Room
	if err := DB.Preload("Players").Where("id = ?", roomID).First(&room).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, models.PlayersInRoomResponse{
				RoomID:  roomID,
				Players: nil,
				Error: models.Error{
					ErrorCode:        1,
					ErrorDescription: "Комната не найдена",
				},
			})
			return
		}

		c.JSON(http.StatusInternalServerError, models.PlayersInRoomResponse{
			RoomID:  roomID,
			Players: nil,
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.PlayersInRoomResponse{
		RoomID:  room.ID,
		Players: room.Players,
		Error: models.Error{
			ErrorCode:        0,
			ErrorDescription: "",
		},
	})
}

type CreateRoomRequest struct {
	Name string `json:"name" binding:"required,max=100" example:"Room One"`
}

type CreateRoomResponse struct {
	Room  Room         `json:"room,omitempty"`
	Error models.Error `json:"error"`
}

// CreateRoomHandler создает новую игровую комнату
// @Summary Создание новой игровой комнаты
// @Description Создает новую игровую комнату с заданным именем
// @Tags room
// @Accept json
// @Produce json
// @Param room body CreateRoomRequest true "Информация о комнате"
// @Success 201 {object} CreateRoomResponse
// @Failure 400 {object} CreateRoomResponse
// @Failure 500 {object} CreateRoomResponse
// @Router /rooms/create [post]
func CreateRoomHandler(c *gin.Context) {
	var req CreateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, CreateRoomResponse{
			Room: Room{},
			Error: models.Error{
				ErrorCode:        3,
				ErrorDescription: "Некорректный формат имени комнаты",
			},
		})
		return
	}

	var existingRoom Room
	if err := DB.Where("name = ?", req.Name).First(&existingRoom).Error; err == nil {
		c.JSON(http.StatusBadRequest, CreateRoomResponse{
			Room: Room{},
			Error: models.Error{
				ErrorCode:        1,
				ErrorDescription: "Комната с таким именем уже существует",
			},
		})
		return
	} else if err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, CreateRoomResponse{
			Room: Room{},
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	room := Room{
		Name: req.Name,
	}

	if err := DB.Create(&room).Error; err != nil {
		c.JSON(http.StatusInternalServerError, CreateRoomResponse{
			Room: Room{},
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	c.JSON(http.StatusCreated, CreateRoomResponse{
		Room: room,
		Error: models.Error{
			ErrorCode:        0,
			ErrorDescription: "",
		},
	})
}

type AddPlayerToRoomRequest struct {
	Username string `json:"username" binding:"required,max=25" example:"john_doe"`
}

type AddPlayerToRoomResponse struct {
	Room  Room         `json:"room,omitempty"`
	Error models.Error `json:"error"`
}

// AddPlayerToRoomHandler добавляет игрока в игровую комнату
// @Summary Добавление игрока в комнату
// @Description Добавляет пользователя в указанную игровую комнату
// @Tags room
// @Accept json
// @Produce json
// @Param room_id path string true "Идентификатор комнаты" example:"room123"
// @Param player body AddPlayerToRoomRequest true "Информация об игроке"
// @Success 200 {object} AddPlayerToRoomResponse
// @Failure 400 {object} AddPlayerToRoomResponse
// @Failure 404 {object} AddPlayerToRoomResponse
// @Failure 500 {object} AddPlayerToRoomResponse
// @Router /rooms/{room_id}/players [post]
func AddPlayerToRoomHandler(c *gin.Context) {
	roomID := c.Param("room_id")

	var req AddPlayerToRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, AddPlayerToRoomResponse{
			Room: Room{},
			Error: models.Error{
				ErrorCode:        3,
				ErrorDescription: "Некорректный формат запроса",
			},
		})
		return
	}

	var room Room
	if err := DB.Preload("Players").Where("id = ?", roomID).First(&room).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, AddPlayerToRoomResponse{
				Room: Room{},
				Error: models.Error{
					ErrorCode:        1,
					ErrorDescription: "Комната не найдена",
				},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, AddPlayerToRoomResponse{
			Room: Room{},
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	var user models.User
	if err := DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, AddPlayerToRoomResponse{
				Room: Room{},
				Error: models.Error{
					ErrorCode:        2,
					ErrorDescription: "Пользователь не найден",
				},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, AddPlayerToRoomResponse{
			Room: Room{},
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	for _, p := range room.Players {
		if p.ID == user.ID {
			c.JSON(http.StatusBadRequest, AddPlayerToRoomResponse{
				Room: Room{},
				Error: models.Error{
					ErrorCode:        4,
					ErrorDescription: "Пользователь уже находится в комнате",
				},
			})
			return
		}
	}

	if err := DB.Model(&room).Association("Players").Append(&user); err != nil {
		c.JSON(http.StatusInternalServerError, AddPlayerToRoomResponse{
			Room: Room{},
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Не удалось добавить игрока в комнату",
			},
		})
		return
	}

	room.Players = append(room.Players, user)

	c.JSON(http.StatusOK, AddPlayerToRoomResponse{
		Room: room,
		Error: models.Error{
			ErrorCode:        0,
			ErrorDescription: "",
		},
	})
}
