package handlers

import (
	"net/http"

	"github.com/Malader/maladerrepo/models"
	"github.com/gin-gonic/gin"
)

// Mocked user data for demonstration purposes
var users = []models.User{
	{ID: "1", Username: "john_doe", Email: "john.doe@example.com", Team: "Red"},
	{ID: "2", Username: "jane_smith", Email: "jane.smith@example.com", Team: "Blue"},
}

// VerifyCredentials обрабатывает запросы на проверку учетных данных пользователей
// @Summary Проверка учетных данных пользователя
// @Description Проверяет правильность введенных данных пользователя
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body models.VerifyCredentialsRequest true "Учетные данные"
// @Success 200 {object} models.VerifyCredentialsResponse
// @Failure 400 {object} models.VerifyCredentialsResponse
// @Router /auth/verify [post]
func VerifyCredentials(c *gin.Context) {
	var req models.VerifyCredentialsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.VerifyCredentialsResponse{
			Valid: false,
			Error: "Неверный формат запроса",
		})
		return
	}

	// Здесь должна быть логика проверки пароля, например, через хэширование
	// Для демонстрации проверяем только наличие пользователя
	for _, user := range users {
		if user.Username == req.Username {
			// Предположим, что пароль всегда верный для упрощения
			c.JSON(http.StatusOK, models.VerifyCredentialsResponse{
				Valid: true,
				User:  user,
			})
			return
		}
	}

	c.JSON(http.StatusOK, models.VerifyCredentialsResponse{
		Valid: false,
		Error: "Пользователь не найден или неверный пароль",
	})
}
