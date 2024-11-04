package handlers

import (
	"net/http"
	"time"

	"github.com/Malader/maladerrepo/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("$JWT_SECRET")

// AuthorizeUser обрабатывает запросы на авторизацию пользователя
// @Summary Авторизация пользователя
// @Description Позволяет авторизовать пользователя в системе
// @Tags user
// @Accept json
// @Produce json
// @Param credentials body models.AuthorizeRequest true "Учетные данные пользователя"
// @Success 201 {object} models.AuthorizeResponse
// @Failure 400 {object} models.AuthorizeResponse
// @Failure 500 {object} models.AuthorizeResponse
// @Router /user/authorize [post]
func AuthorizeUser(c *gin.Context) {
	var req models.AuthorizeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.AuthorizeResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Некорректное поведение системы",
			},
		})
		return
	}

	var user models.User
	if err := DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, models.AuthorizeResponse{
			Error: models.Error{
				ErrorCode:        1,
				ErrorDescription: "Неуспешная авторизация",
			},
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		c.JSON(http.StatusBadRequest, models.AuthorizeResponse{
			Error: models.Error{
				ErrorCode:        1,
				ErrorDescription: "Неуспешная авторизация",
			},
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.AuthorizeResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Некорректное поведение системы",
			},
		})
		return
	}

	c.JSON(http.StatusCreated, models.AuthorizeResponse{
		Error: models.Error{
			ErrorCode:        0,
			ErrorDescription: "",
		},
		Token: tokenString,
	})
}
