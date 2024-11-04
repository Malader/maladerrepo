package models

import "time"

type UpdateProfileRequest struct {
	Username string `json:"username" binding:"omitempty,max=25" example:"vasyaPupkin"`
	Password string `json:"password" binding:"omitempty,len=64" example:"6a4a61f57bccf059abb82fc95589ebc428629326ab965390f25224e262455beb"`
	Image    string `json:"image" binding:"omitempty,base64" example:"TWFuIGlzIGRpc3Rpbmd1aXNoZWQsIG5vdCBvbmx5IGJ5IGhpcyByZWFzb24sIGJ1dCAuLi4="`
}

type UpdateProfileResponse struct {
	Error Error `json:"error"`
}

type GetProfileResponse struct {
	Email            string    `json:"email" example:"example@example.com"`
	Username         string    `json:"username" example:"ExampleUser"`
	Password         string    `json:"password" example:"hashedPassword123"`
	RegistrationDate time.Time `json:"registrationDate" example:"2024-01-01T00:00:00Z"`
	LastActivityDate time.Time `json:"lastActivityDate" example:"2024-05-22T12:34:56Z"`
	Image            string    `json:"image" example:"4AAQSkZJRgABAQEAAAAAAADfjadFHADHda"`
	Error            Error     `json:"error"`
}
