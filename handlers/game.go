package handlers

import (
	"net/http"
	"strings"

	"github.com/Malader/maladerrepo/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AddGameSearchHandler обрабатывает запросы на начало поиска игры
// @Summary Начать поиск игры
// @Description Данный метод позволяет пользователю начать поиск игры с заданным списком метатем.
// @Tags game
// @Accept json
// @Produce json
// @Param id path string true "Идентификатор пользователя" example:"1234567890123456789"
// @Param metathemes body models.AddGameSearchRequest true "Список метатем"
// @Success 200 {object} models.AddGameSearchResponse "Успешное начало поиска игры"
// @Failure 400 {object} models.AddGameSearchResponse "Некорректные данные запроса или пользователь не найден"
// @Failure 500 {object} models.AddGameSearchResponse "Внутренняя ошибка"
// @Router /game/{id} [post]
func AddGameSearchHandler(c *gin.Context) {
	userID := c.Param("id")

	var req models.AddGameSearchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.AddGameSearchResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Некорректные данные запроса",
			},
		})
		return
	}

	var user models.User
	if err := DB.Where("id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, models.AddGameSearchResponse{
				Error: models.Error{
					ErrorCode:        999,
					ErrorDescription: "Некорректное поведение системы",
				},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.AddGameSearchResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	var metathemes []models.Theme
	if err := DB.Where("id IN (?) AND category = ?", req.Metathemes, models.METATHEME).Find(&metathemes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.AddGameSearchResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	if len(metathemes) != len(req.Metathemes) {
		c.JSON(http.StatusBadRequest, models.AddGameSearchResponse{
			Error: models.Error{
				ErrorCode:        2,
				ErrorDescription: "Некорректный формат метатем в запросе",
			},
		})
		return
	}

	gameSearch := models.GameSearch{
		UserID:     user.ID,
		Metathemes: metathemes,
		Status:     "searching",
	}

	if err := DB.Create(&gameSearch).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.AddGameSearchResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.AddGameSearchResponse{
		Error: models.Error{
			ErrorCode:        0,
			ErrorDescription: "",
		},
	})
}

// StopGameSearchHandler обрабатывает запросы на остановку поиска игры
// @Summary Остановить поиск игры
// @Description Данный метод позволяет пользователю остановить поиск игры.
// @Tags game
// @Accept json
// @Produce json
// @Param id path string true "Идентификатор пользователя" example:"1234567890123456789"
// @Success 200 {object} models.StopGameSearchResponse "Успешная остановка поиска игры"
// @Failure 500 {object} models.StopGameSearchResponse "Внутренняя ошибка"
// @Router /game/{id} [delete]
func StopGameSearchHandler(c *gin.Context) {
	userID := c.Param("id")

	var user models.User
	if err := DB.Where("id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, models.StopGameSearchResponse{
				Error: models.Error{
					ErrorCode:        999,
					ErrorDescription: "Некорректное поведение системы",
				},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.StopGameSearchResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	var gameSearches []models.GameSearch
	if err := DB.Where("user_id = ? AND status = ?", user.ID, "searching").Find(&gameSearches).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.StopGameSearchResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	if len(gameSearches) == 0 {
		c.JSON(http.StatusOK, models.StopGameSearchResponse{
			Error: models.Error{
				ErrorCode:        0,
				ErrorDescription: "",
			},
		})
		return
	}

	if err := DB.Model(&gameSearches).Update("status", "stopped").Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.StopGameSearchResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.StopGameSearchResponse{
		Error: models.Error{
			ErrorCode:        0,
			ErrorDescription: "",
		},
	})
}

// AddSpectatorHandler обрабатывает запросы на начало просмотра игры как зритель
// @Summary Начать просмотр текущей игры как зритель
// @Description Данный метод позволяет пользователю начать просмотр игры с заданным идентификатором игры.
// @Tags game
// @Accept json
// @Produce json
// @Param user_id path string true "Идентификатор пользователя" example:"1234567890123456789"
// @Param game_id path string true "Идентификатор игры" example:"9876543210987654321"
// @Success 200 {object} models.AddSpectatorResponse "Успешное подключение к игре"
// @Failure 400 {object} models.AddSpectatorResponse "Игра больше не существует (завершена) или некорректные данные запроса"
// @Failure 500 {object} models.AddSpectatorResponse "Внутренняя ошибка"
// @Router /game/spectator/{user_id}/{game_id} [post]
func AddSpectatorHandler(c *gin.Context) {
	userID := c.Param("user_id")
	gameID := c.Param("game_id")

	var user models.User
	if err := DB.Where("id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, models.AddSpectatorResponse{
				Error: models.Error{
					ErrorCode:        999,
					ErrorDescription: "Некорректное поведение системы",
				},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.AddSpectatorResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	var game models.GameSearch
	if err := DB.Where("id = ?", gameID).First(&game).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, models.AddSpectatorResponse{
				Error: models.Error{
					ErrorCode:        400,
					ErrorDescription: "Игра больше не существует (завершена)",
				},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.AddSpectatorResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	if strings.ToLower(game.Status) != "searching" && strings.ToLower(game.Status) != "found" {
		c.JSON(http.StatusBadRequest, models.AddSpectatorResponse{
			Error: models.Error{
				ErrorCode:        400,
				ErrorDescription: "Игра больше не существует (завершена)",
			},
		})
		return
	}

	var count int64
	DB.Table("game_spectators").Where("game_search_id = ? AND user_id = ?", game.ID, user.ID).Count(&count)
	if count > 0 {
		c.JSON(http.StatusOK, models.AddSpectatorResponse{
			Error: models.Error{
				ErrorCode:        0,
				ErrorDescription: "",
			},
		})
		return
	}

	if err := DB.Model(&game).Association("Spectators").Append(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.AddSpectatorResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.AddSpectatorResponse{
		Error: models.Error{
			ErrorCode:        0,
			ErrorDescription: "",
		},
	})
}
