package main

import (
	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	if err != nil {
		log.Println("Не удалось загрузить .env файл")
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.SetTrustedProxies([]string{})

	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=" + os.Getenv("DB_SSLMODE") +
		" TimeZone=" + os.Getenv("DB_TIMEZONE")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Room{}, &models.FriendRequest{}, &models.Theme{}, &models.GameSearch{})
	if err != nil {
		log.Fatalf("Не удалось мигрировать базу данных: %v", err)
	}

	handlers.InitDB(db)

	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		userRoutes := api.Group("/user")
		{
			userRoutes.POST("/register", handlers.RegisterUser)
			userRoutes.POST("/authorize", handlers.AuthorizeUser)
			userRoutes.POST("/recovery", handlers.RecoveryHandler)
			userRoutes.PATCH("/recovery/:recoverySuffix", handlers.RecoveryHandler)
			userRoutes.PATCH("/:id/profile", handlers.UpdateProfileHandler)
			userRoutes.GET("/:id/profile", handlers.GetProfileHandler)
			userRoutes.GET("/:id/friends", handlers.GetFriendsHandler)
			userRoutes.POST("/:id/friends/:username", handlers.SendFriendRequestHandler)
			userRoutes.PUT("/:id/friends/:username/:confirmation", handlers.ConfirmFriendRequestHandler)
			userRoutes.DELETE("/:id/friends/:username", handlers.DeleteFriendHandler)
			userRoutes.PUT("/:id/blacklist/:username", handlers.AddToBlacklistHandler)
		}

		themeRoutes := api.Group("/theme")
		{
			themeRoutes.POST("", handlers.AddThemeHandler)
			themeRoutes.PUT("/:confirmation", handlers.ConfirmThemeHandler)
		}

		gameRoutes := api.Group("/game")
		{
			gameRoutes.POST("/:id", handlers.AddGameSearchHandler)
			gameRoutes.DELETE("/:id", handlers.StopGameSearchHandler)
			gameRoutes.POST("/spectator/:user_id/:game_id", handlers.AddSpectatorHandler)
		}

		roomRoutes := api.Group("/rooms")
		{
			roomRoutes.POST("/create", handlers.CreateRoomHandler)
			roomRoutes.GET("/:room_id/players", handlers.GetPlayersInRoom)
			roomRoutes.POST("/:room_id/players", handlers.AddPlayerToRoomHandler)
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
