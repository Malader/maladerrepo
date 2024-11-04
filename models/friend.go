// models/friend.go
package models

// Friend представляет структуру друга
type Friend struct {
	Username string `json:"username" example:"VasyaPupkin"`
	Image    string `json:"image" example:"4AAQSkZJRgABAQEAAAAAAADfjadFHADHda"`
	Status   string `json:"status" example:"active"`
}
