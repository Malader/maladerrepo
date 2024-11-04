// handlers/recovery_phase1.go
package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/Malader/maladerrepo/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RecoveryPhase1 обрабатывает запросы на инициацию восстановления пароля
// @Summary Инициация восстановления пароля
// @Description Позволяет начать процесс восстановления пароля через email
// @Tags recovery
// @Accept json
// @Produce json
// @Param recovery body models.RecoveryRequest true "Email пользователя"
// @Success 201 {object} models.RecoveryResponse
// @Failure 400 {object} models.RecoveryResponse
// @Failure 500 {object} models.RecoveryResponse
// @Router /user/recovery [post]
func RecoveryPhase1(c *gin.Context) {
	var req models.RecoveryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.RecoveryResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Некорректное поведение системы",
			},
		})
		return
	}

	var user models.User
	if err := DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, models.RecoveryResponse{
			Error: models.Error{
				ErrorCode:        1,
				ErrorDescription: "Аккаунт с указанным Email не существует",
			},
		})
		return
	}

	recoverySuffix := uuid.New().String()

	recoveryToken := models.RecoveryToken{
		UserID:         user.ID,
		RecoverySuffix: recoverySuffix,
		ExpiresAt:      time.Now().Add(24 * time.Hour),
	}

	if err := DB.Create(&recoveryToken).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.RecoveryResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	recoveryLink := "http://localhost:8080/api/user/recovery/" + recoverySuffix
	log.Printf("Recovery link for user %s: %s", user.Email, recoveryLink)

	c.JSON(http.StatusCreated, models.RecoveryResponse{
		Error: models.Error{
			ErrorCode:        0,
			ErrorDescription: "",
		},
	})
}
