// handlers/recovery_phase2.go
package handlers

import (
	"net/http"
	"time"

	"github.com/Malader/maladerrepo/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// RecoveryPhase2 обрабатывает запросы на смену пароля по ссылке восстановления
// @Summary Завершение восстановления пароля
// @Description Позволяет завершить процесс восстановления пароля по ссылке
// @Tags recovery
// @Accept json
// @Produce json
// @Param recoverySuffix path string true "Уникальный суффикс восстановления" example:"some-suffix-here"
// @Param recovery body models.ChangePasswordRequest true "Новый пароль"
// @Success 201 {object} models.ChangePasswordResponse
// @Failure 400 {object} models.ChangePasswordResponse
// @Failure 404 {object} models.ChangePasswordResponse
// @Failure 422 {object} models.ChangePasswordResponse
// @Failure 500 {object} models.ChangePasswordResponse
// @Router /user/recovery/{recoverySuffix} [patch]
func RecoveryPhase2(c *gin.Context) {
	recoverySuffix := c.Param("recoverySuffix")

	var req models.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ChangePasswordResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Некорректное поведение системы",
			},
		})
		return
	}

	var token models.RecoveryToken
	if err := DB.Where("recovery_suffix = ?", recoverySuffix).First(&token).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ChangePasswordResponse{
			Error: models.Error{
				ErrorCode:        1,
				ErrorDescription: "Уникальный суффикс восстановления не найден",
			},
		})
		return
	}

	if time.Now().After(token.ExpiresAt) {
		c.JSON(http.StatusUnprocessableEntity, models.ChangePasswordResponse{
			Error: models.Error{
				ErrorCode:        2,
				ErrorDescription: "Время действия ссылки истекло, повторите процедуру",
			},
		})
		return
	}

	var user models.User
	if err := DB.Where("id = ?", token.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ChangePasswordResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ChangePasswordResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	user.PasswordHash = string(hashedPassword)

	if err := DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ChangePasswordResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	if err := DB.Delete(&token).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ChangePasswordResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	c.JSON(http.StatusCreated, models.ChangePasswordResponse{
		Error: models.Error{
			ErrorCode:        0,
			ErrorDescription: "",
		},
	})
}
