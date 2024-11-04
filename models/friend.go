package models

type Friend struct {
	Username string `json:"username" example:"VasyaPupkin"`
	Image    string `json:"image" example:"4AAQSkZJRgABAQEAAAAAAADfjadFHADHda"`
	Status   string `json:"status" example:"active"`
}
