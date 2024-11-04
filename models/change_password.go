package models

type ChangePasswordRequest struct {
	NewPassword string `json:"newPassword" binding:"required,min=8,max=64" example:"newSecurePassword123"`
}

type ChangePasswordResponse struct {
	Error Error `json:"error"`
}
