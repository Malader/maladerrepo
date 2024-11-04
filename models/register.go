package models

// RegisterRequest представляет запрос на регистрацию пользователя
type RegisterRequest struct {
	Username string `json:"username" binding:"required,max=25" example:"vasyaPupkin"`
	Password string `json:"password" binding:"required,len=64" example:"6a4a61f57bccf059abb82fc95589ebc428629326ab965390f25224e262455beb"`
	Email    string `json:"email" binding:"required,email,max=256" example:"v.pupkin@g.nsu.ru"`
}

// RegisterResponse представляет ответ на регистрацию пользователя
type RegisterResponse struct {
	Error Error `json:"error"`
}
