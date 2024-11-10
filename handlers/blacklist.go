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

	if user.ID == target.ID {
		c.JSON(http.StatusBadRequest, models.AddBlacklistResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Некорректное поведение системы",
			},
		})
		return
	}

	var count int64
	DB.Table("user_blacklists").Where("user_id = ? AND blacklisted_user_id = ?", user.ID, target.ID).Count(&count)
	if count > 0 {
		c.JSON(http.StatusBadRequest, models.AddBlacklistResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Пользователь уже находится в черном списке",
			},
		})
		return
	}

	if err := DB.Model(&user).Association("BlacklistedUsers").Append(&target); err != nil {
		c.JSON(http.StatusInternalServerError, models.AddBlacklistResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.AddBlacklistResponse{
		Error: models.Error{
			ErrorCode:        0,
			ErrorDescription: "",
		},
	})
}
