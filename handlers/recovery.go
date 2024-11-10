package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/Malader/maladerrepo/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// RecoveryHandler обрабатывает запросы на восстановление пароля
// @Summary Восстановление пароля
// @Description Позволяет начать и завершить процесс восстановления пароля через email
// @Tags recovery
// @Accept json
// @Produce json
// @Param recoverySuffix path string false "Уникальный суффикс восстановления" example:"some-suffix-here"
// @Param recovery body models.RecoveryRequest true "Email пользователя или новый пароль"
// @Success 201 {object} models.RecoverySuccessResponse "Успешное восстановление пароля"
// @Failure 400 {object} models.RecoveryBadRequestResponse "Некорректные данные запроса"
// @Failure 404 {object} models.RecoveryNotFoundResponse "Суффикс восстановления не найден"
// @Failure 422 {object} models.RecoveryUnprocessableEntityResponse "Время действия ссылки истекло"
// @Failure 500 {object} models.RecoveryInternalServerErrorResponse "Внутренняя ошибка"
// @Router /user/recovery [post]
// @Router /user/recovery/{recoverySuffix} [patch]
func RecoveryHandler(c *gin.Context) {
	recoverySuffix := c.Param("recoverySuffix")

	if recoverySuffix == "" {
		// Фаза 1: Инициация восстановления пароля
		var req models.RecoveryRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, models.RecoveryBadRequestResponse{
				Error: models.RecoveryBadRequestError{
					ErrorCode:        999,
					ErrorDescription: "Некорректные данные запроса",
				},
			})
			return
		}

		var user models.User
		if err := DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, models.RecoveryBadRequestResponse{
				Error: models.RecoveryBadRequestError{
					ErrorCode:        1,
					ErrorDescription: "Аккаунт с указанным Email не существует",
				},
			})
			return
		}

		recoverySuffix = uuid.New().String()

		recoveryToken := models.RecoveryToken{
			UserID:         user.ID,
			RecoverySuffix: recoverySuffix,
			ExpiresAt:      time.Now().Add(24 * time.Hour),
		}

		if err := DB.Create(&recoveryToken).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.RecoveryInternalServerErrorResponse{
				Error: models.RecoveryInternalServerError{
					ErrorCode:        999,
					ErrorDescription: "Внутренняя ошибка",
				},
			})
			return
		}

		recoveryLink := "http://localhost:8080/api/user/recovery/" + recoverySuffix
		log.Printf("Recovery link for user %s: %s", user.Email, recoveryLink)

		c.JSON(http.StatusCreated, models.RecoverySuccessResponse{
			Error: models.RecoverySuccessError{
				ErrorCode:        0,
				ErrorDescription: "",
			},
		})
	} else {
		// Фаза 2: Завершение восстановления пароля
		var req models.ChangePasswordRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, models.RecoveryBadRequestResponse{
				Error: models.RecoveryBadRequestError{
					ErrorCode:        999,
					ErrorDescription: "Некорректные данные запроса",
				},
			})
			return
		}

		var token models.RecoveryToken
		if err := DB.Where("recovery_suffix = ?", recoverySuffix).First(&token).Error; err != nil {
			c.JSON(http.StatusNotFound, models.RecoveryNotFoundResponse{
				Error: models.RecoveryNotFoundError{
					ErrorCode:        1,
					ErrorDescription: "Уникальный суффикс восстановления не найден",
				},
			})
			return
		}

		if time.Now().After(token.ExpiresAt) {
			c.JSON(http.StatusUnprocessableEntity, models.RecoveryUnprocessableEntityResponse{
				Error: models.RecoveryUnprocessableEntityError{
					ErrorCode:        2,
					ErrorDescription: "Время действия ссылки истекло, повторите процедуру",
				},
			})
			return
		}

		var user models.User
		if err := DB.Where("id = ?", token.UserID).First(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.RecoveryInternalServerErrorResponse{
				Error: models.RecoveryInternalServerError{
					ErrorCode:        999,
					ErrorDescription: "Внутренняя ошибка",
				},
			})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.RecoveryInternalServerErrorResponse{
				Error: models.RecoveryInternalServerError{
					ErrorCode:        999,
					ErrorDescription: "Внутренняя ошибка",
				},
			})
			return
		}

		user.PasswordHash = string(hashedPassword)

		if err := DB.Save(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.RecoveryInternalServerErrorResponse{
				Error: models.RecoveryInternalServerError{
					ErrorCode:        999,
					ErrorDescription: "Внутренняя ошибка",
				},
			})
			return
		}

		if err := DB.Delete(&token).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.RecoveryInternalServerErrorResponse{
				Error: models.RecoveryInternalServerError{
					ErrorCode:        999,
					ErrorDescription: "Внутренняя ошибка",
				},
			})
			return
		}

		c.JSON(http.StatusCreated, models.RecoverySuccessResponse{
			Error: models.RecoverySuccessError{
				ErrorCode:        0,
				ErrorDescription: "",
			},
		})
	}
}
