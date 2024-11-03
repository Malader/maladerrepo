package main

import (
	"github.com/Malader/maladerrepo/docs"
	"github.com/Malader/maladerrepo/handlers"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           User Service API
// @version         1.0
// @description     API сервис пользователей для взаимодействия с сервисом авторизации и сервисом игровых комнат
// @host            localhost:8080
// @BasePath        /api
func main() {
	router := gin.Default()

	// Swagger документация
	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/verify", handlers.VerifyCredentials)
		}

		rooms := api.Group("/rooms")
		{
			rooms.GET("/:room_id/players", handlers.GetPlayersInRoom)
		}
	}

	router.Run(":8080")
}
