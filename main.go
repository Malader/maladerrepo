package main

import (
	"github.com/Malader/maladerrepo/models"
	"log"
	"os"

	"github.com/Malader/maladerrepo/docs"
	"github.com/Malader/maladerrepo/handlers"
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
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Автоматическая миграция моделей
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Инициализация базы данных в обработчиках
	handlers.InitDB(db)

	// Swagger документация
	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Группы маршрутов
	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/verify", handlers.VerifyCredentials)
			auth.POST("/create", handlers.CreateUserHandler) // Временный эндпоинт для создания пользователей
		}

		rooms := api.Group("/rooms")
		{
			rooms.GET("/:room_id/players", handlers.GetPlayersInRoom)
		}
	}

	// Запуск сервера
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
