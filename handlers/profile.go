package handlers

import (
	"encoding/base64"
	"github.com/Malader/maladerrepo/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

// UpdateProfileHandler обрабатывает запросы на изменение информации профиля пользователя
// @Summary Изменение информации профиля
// @Description Позволяет изменить информацию о пользователе
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "Идентификатор профиля" example:"1234567890123456789"
// @Param profile body models.UpdateProfileRequest true "Информация профиля"
// @Success 201 {object} models.UpdateProfileResponse
// @Failure 422 {object} models.UpdateProfileResponse
// @Failure 500 {object} models.UpdateProfileResponse
// @Router /user/{id}/profile [patch]
func UpdateProfileHandler(c *gin.Context) {
	userID := c.Param("id")

	var req models.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.UpdateProfileResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Некорректное поведение системы",
			},
		})
		return
	}

	var user models.User
	if err := DB.Where("id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, models.UpdateProfileResponse{
				Error: models.Error{
					ErrorCode:        999,
					ErrorDescription: "Некорректное поведение системы",
				},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.UpdateProfileResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Некорректное поведение системы",
			},
		})
		return
	}

	if req.Username != "" && req.Username != user.Username {
		var existingUser models.User
		if err := DB.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
			c.JSON(http.StatusUnprocessableEntity, models.UpdateProfileResponse{
				Error: models.Error{
					ErrorCode:        1,
					ErrorDescription: "Указанный username уже зарегистрирован",
				},
			})
			return
		} else if err != gorm.ErrRecordNotFound {
			c.JSON(http.StatusInternalServerError, models.UpdateProfileResponse{
				Error: models.Error{
					ErrorCode:        999,
					ErrorDescription: "Некорректное поведение системы",
				},
			})
			return
		}
		user.Username = req.Username
	}

	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.UpdateProfileResponse{
				Error: models.Error{
					ErrorCode:        999,
					ErrorDescription: "Некорректное поведение системы",
				},
			})
			return
		}
		user.PasswordHash = string(hashedPassword)
	}

	if req.Image != "" {
		decodedImage, err := base64.StdEncoding.DecodeString(req.Image)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, models.UpdateProfileResponse{
				Error: models.Error{
					ErrorCode:        3,
					ErrorDescription: "Некорректный формат изображения",
				},
			})
			return
		}
		if len(decodedImage) > 10*1024*1024 { // лимит на 10MB
			c.JSON(http.StatusUnprocessableEntity, models.UpdateProfileResponse{
				Error: models.Error{
					ErrorCode:        4,
					ErrorDescription: "Размер изображения больше предела допустимого",
				},
			})
			return
		}
		user.Image = req.Image
	}

	if err := DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.UpdateProfileResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Некорректное поведение системы",
			},
		})
		return
	}

	c.JSON(http.StatusCreated, models.UpdateProfileResponse{
		Error: models.Error{
			ErrorCode:        0,
			ErrorDescription: "",
		},
	})
}

// GetProfileHandler обрабатывает запросы на получение информации о профиле пользователя
// @Summary Получение информации профиля
// @Description Позволяет получить информацию о пользователе
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "Идентификатор профиля" example:"1234567890123456789"
// @Success 201 {object} models.GetProfileResponse
// @Failure 404 {object} models.GetProfileResponse
// @Failure 500 {object} models.GetProfileResponse
// @Failure 504 {object} models.GetProfileResponse
// @Router /user/{id}/profile [get]
func GetProfileHandler(c *gin.Context) {
	userID := c.Param("id")

	var user models.User
	if err := DB.Where("id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, models.GetProfileResponse{
				Error: models.Error{
					ErrorCode:        999,
					ErrorDescription: "Некорректное поведение системы",
				},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.GetProfileResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Некорректное поведение системы",
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.GetProfileResponse{
		Email:            user.Email,
		Username:         user.Username,
		Password:         user.PasswordHash,
		RegistrationDate: user.CreatedAt,
		LastActivityDate: user.LastActivityDate,
		Image:            user.Image,
		Error: models.Error{
			ErrorCode:        0,
			ErrorDescription: "",
		},
	})
}
