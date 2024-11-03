package handlers

import (
	"net/http"
	"time"

	"github.com/Malader/maladerrepo/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// Secret key для JWT. В продакшен-среде храните его безопасно (например, в переменных окружения)
var jwtSecret = []byte("your_secret_key")

// VerifyCredentials обрабатывает запросы на проверку учетных данных пользователей
// @Summary Проверка учетных данных пользователя
// @Description Проверяет правильность введенных данных пользователя
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body models.VerifyCredentialsRequest true "Учетные данные"
// @Success 200 {object} models.VerifyCredentialsResponse
// @Failure 400 {object} models.VerifyCredentialsResponse
// @Failure 401 {object} models.VerifyCredentialsResponse
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

	var user models.User
	// Поиск пользователя по имени
	if err := DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, models.VerifyCredentialsResponse{
			Valid: false,
			Error: "Пользователь не найден или неверный пароль",
		})
		return
	}

	// Проверка пароля
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, models.VerifyCredentialsResponse{
			Valid: false,
			Error: "Пользователь не найден или неверный пароль",
		})
		return
	}

	// Генерация JWT-токена
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(), // Токен действителен 72 часа
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.VerifyCredentialsResponse{
			Valid: false,
			Error: "Не удалось сгенерировать токен",
		})
		return
	}

	// Возвращение успешного ответа с токеном
	c.JSON(http.StatusOK, models.VerifyCredentialsResponse{
		Valid: true,
		User: models.User{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Team:     user.Team,
		},
		Token: tokenString,
	})
}

// CreateUserHandler создает нового пользователя
// @Summary Создание нового пользователя
// @Description Создает нового пользователя с хэшированным паролем
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "Пользователь"
// @Success 201 {object} models.User
// @Failure 400 {object} gin.H
// @Router /auth/create [post]
func CreateUserHandler(c *gin.Context) {
	var req struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
		Team     string `json:"team" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверка, существует ли пользователь с таким же именем или email
	var existingUser models.User
	if err := DB.Where("username = ?", req.Username).Or("email = ?", req.Email).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Пользователь с таким именем или email уже существует"})
		return
	}

	// Хэширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось хэшировать пароль"})
		return
	}

	user := models.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Team:         req.Team,
	}

	if err := DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать пользователя"})
		return
	}

	c.JSON(http.StatusCreated, models.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Team:     user.Team,
	})
}
