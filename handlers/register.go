package handlers

import (
	"net/http"

	"github.com/Malader/maladerrepo/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUser обрабатывает запросы на регистрацию нового пользователя
// @Summary Регистрация нового пользователя
// @Description Позволяет зарегистрировать пользователя в системе
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.RegisterRequest true "Данные пользователя для регистрации"
// @Success 201 {object} models.RegisterResponse
// @Failure 400 {object} models.RegisterResponse
// @Failure 422 {object} models.RegisterResponse
// @Failure 500 {object} models.RegisterResponse
// @Router /user/register [post]
func RegisterUser(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.RegisterResponse{
			Error: models.Error{
				ErrorCode:        5,
				ErrorDescription: "Некорректный формат email",
			},
		})
		return
	}

	if len(req.Username) > 25 || len(req.Username) == 0 {
		c.JSON(http.StatusUnprocessableEntity, models.RegisterResponse{
			Error: models.Error{
				ErrorCode:        3,
				ErrorDescription: "Некорректный формат username",
			},
		})
		return
	}

	if len(req.Password) != 64 {
		c.JSON(http.StatusUnprocessableEntity, models.RegisterResponse{
			Error: models.Error{
				ErrorCode:        4,
				ErrorDescription: "Некорректный формат password",
			},
		})
		return
	}

	var existingUser models.User
	if err := DB.Where("username = ?", req.Username).Or("email = ?", req.Email).First(&existingUser).Error; err == nil {
		if existingUser.Username == req.Username {
			c.JSON(http.StatusBadRequest, models.RegisterResponse{
				Error: models.Error{
					ErrorCode:        1,
					ErrorDescription: "Указанный username уже зарегистрирован",
				},
			})
			return
		}
		if existingUser.Email == req.Email {
			c.JSON(http.StatusBadRequest, models.RegisterResponse{
				Error: models.Error{
					ErrorCode:        2,
					ErrorDescription: "Указанный email уже зарегистрирован",
				},
			})
			return
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.RegisterResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Некорректное поведение системы",
			},
		})
		return
	}

	user := models.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Team:         "Default", // Или другое значение по умолчанию
	}

	if err := DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.RegisterResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Некорректное поведение системы",
			},
		})
		return
	}

	c.JSON(http.StatusCreated, models.RegisterResponse{
		Error: models.Error{
			ErrorCode:        0,
			ErrorDescription: "",
		},
	})
}
