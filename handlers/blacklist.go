// handlers/blacklist.go
package handlers

import (
	"net/http"

	"github.com/Malader/maladerrepo/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AddToBlacklistHandler обрабатывает запросы на добавление пользователя в черный список
// @Summary Добавление в черный список
// @Description Позволяет добавить пользователя в черный список
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "Идентификатор пользователя, добавляющего в черный список" example:"1234567890123456789"
// @Param username path string true "Имя пользователя, добавляемого в черный список" example:"ExampleUser"
// @Success 200 {object} models.AddBlacklistResponse
// @Failure 400 {object} models.AddBlacklistResponse
// @Failure 500 {object} models.AddBlacklistResponse
// @Router /user/{id}/blackList/{username} [put]
func AddToBlacklistHandler(c *gin.Context) {
	userID := c.Param("id")
	targetUsername := c.Param("username")

	// Найти пользователя, который добавляет в черный список
	var user models.User
	if err := DB.Where("id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, models.AddBlacklistResponse{
				Error: models.Error{
					ErrorCode:        999,
					ErrorDescription: "Некорректное поведение системы",
				},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.AddBlacklistResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	// Найти целевого пользователя по username
	var target models.User
	if err := DB.Where("username = ?", targetUsername).First(&target).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, models.AddBlacklistResponse{
				Error: models.Error{
					ErrorCode:        400,
					ErrorDescription: "Пользователь с указанным username не зарегистрирован в системе",
				},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.AddBlacklistResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	// Проверить, не пытается ли пользователь добавить самого себя
	if user.ID == target.ID {
		c.JSON(http.StatusBadRequest, models.AddBlacklistResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Некорректное поведение системы",
			},
		})
		return
	}

	// Проверить, не находится ли целевой пользователь уже в черном списке
	has, err := DB.Model(&user).Association("BlacklistedUsers").Has(&target)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.AddBlacklistResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	if has {
		c.JSON(http.StatusBadRequest, models.AddBlacklistResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Пользователь уже находится в черном списке",
			},
		})
		return
	}

	// Добавить пользователя в черный список
	if err := DB.Model(&user).Association("BlacklistedUsers").Append(&target); err != nil {
		c.JSON(http.StatusInternalServerError, models.AddBlacklistResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	// Успешно добавлено
	c.JSON(http.StatusOK, models.AddBlacklistResponse{
		Error: models.Error{
			ErrorCode:        0,
			ErrorDescription: "",
		},
	})
}
