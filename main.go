// main.go
package main

import (
	"log"
	"os"

	"github.com/Malader/maladerrepo/docs"
	"github.com/Malader/maladerrepo/handlers"
	"github.com/Malader/maladerrepo/models"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title           User Service API
// @version         1.0
// @description     API сервис пользователей для взаимодействия с сервисом авторизации и сервисом игровых комнат
// @host            localhost:8080
// @BasePath        /api
func main() {
	// Установка режима работы Gin
	gin.SetMode(gin.ReleaseMode)

	// Инициализация маршрутизатора
	router := gin.Default()

	// Отключение доверенных прокси
	router.SetTrustedProxies([]string{})

	// Подключение к базе данных
	dsn := "host=localhost user=postgres password=yourpassword dbname=userdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	// Автоматическая миграция моделей
	err = db.AutoMigrate(&models.User{}, &models.Room{}, &models.FriendRequest{})
	if err != nil {
		log.Fatalf("Не удалось мигрировать базу данных: %v", err)
	}

	// Инициализация базы данных в обработчиках
	handlers.InitDB(db)

	// Swagger документация
	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Группы маршрутов
	api := router.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("/register", handlers.RegisterUser)
			user.POST("/authorize", handlers.AuthorizeUser)
			user.POST("/recovery", handlers.RecoveryPhase1)
			user.PATCH("/recovery/:recoverySuffix", handlers.RecoveryPhase2)
			user.PATCH("/:id/profile", handlers.UpdateProfileHandler)
			user.GET("/:id/profile", handlers.GetProfileHandler)
			user.GET("/:id/friends", handlers.GetFriendsHandler)
			user.POST("/:id/friends/:username", handlers.SendFriendRequestHandler)
			user.PUT("/:id/friends/:username/:confirmation", handlers.ConfirmFriendRequestHandler)
			user.DELETE("/:id/friends/:username", handlers.DeleteFriendHandler)
			user.PUT("/:id/blackList/:username", handlers.AddToBlacklistHandler)
		}

		rooms := api.Group("/rooms")
		{
			rooms.GET("/:room_id/players", handlers.GetPlayersInRoom)
			rooms.POST("/create", handlers.CreateRoomHandler)
			rooms.POST("/:room_id/players", handlers.AddPlayerToRoomHandler)
		}
	}

	// Запуск сервера
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
