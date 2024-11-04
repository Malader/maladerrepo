// handlers/friends.go
package handlers

import (
	"net/http"

	"github.com/Malader/maladerrepo/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetFriendsHandler обрабатывает запросы на получение списка друзей пользователя
// @Summary Получение списка друзей
// @Description Позволяет получить список друзей пользователя
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "Идентификатор пользователя" example:"1234567890123456789"
// @Success 200 {object} models.GetFriendsResponse
// @Failure 500 {object} models.GetFriendsResponse
// @Failure 504 {object} models.GetFriendsResponse
// @Router /user/{id}/friends [get]
func GetFriendsHandler(c *gin.Context) {
	userID := c.Param("id")

	var user models.User
	if err := DB.Preload("Friends").Where("id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusInternalServerError, models.GetFriendsResponse{
				Error: models.Error{
					ErrorCode:        999,
					ErrorDescription: "Некорректное поведение системы",
				},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.GetFriendsResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Некорректное поведение системы",
			},
		})
		return
	}

	friends := make([]models.Friend, 0)
	for _, f := range user.Friends {
		friends = append(friends, models.Friend{
			Username: f.Username,
			Image:    f.Image,
			Status:   "active", // Здесь может быть логика для определения статуса
		})
	}

	c.JSON(http.StatusOK, models.GetFriendsResponse{
		Friends: friends,
		Error: models.Error{
			ErrorCode:        0,
			ErrorDescription: "",
		},
	})
}
