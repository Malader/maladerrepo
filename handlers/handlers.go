package handlers

import "gorm.io/gorm"

var DB *gorm.DB

// InitDB инициализирует подключение к базе данных для обработчиков
func InitDB(db *gorm.DB) {
	DB = db
}
