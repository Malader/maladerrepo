// models/change_password.go
package models

// ChangePasswordRequest представляет запрос на смену пароля
type ChangePasswordRequest struct {
	NewPassword string `json:"newPassword" binding:"required,min=8,max=64" example:"newSecurePassword123"`
}

// ChangePasswordResponse представляет ответ на смену пароля
type ChangePasswordResponse struct {
	Error Error `json:"error"`
}
