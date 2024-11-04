// models/authorize.go
package models

// AuthorizeRequest представляет запрос на авторизацию пользователя
type AuthorizeRequest struct {
	Username string `json:"username" binding:"required,max=25" example:"vasyaPupkin"`
	Password string `json:"password" binding:"required,len=64" example:"6a4a61f57bccf059abb82fc95589ebc428629326ab965390f25224e262455beb"`
}

// AuthorizeResponse представляет ответ на авторизацию пользователя
type AuthorizeResponse struct {
	Error Error  `json:"error"`
	Token string `json:"token,omitempty"`
}
