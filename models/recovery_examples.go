package models

type RecoverySuccessResponse struct {
	Error RecoverySuccessError `json:"error"`
}

type RecoverySuccessError struct {
	ErrorCode        int    `json:"errorCode" example:"0"`
	ErrorDescription string `json:"errorDescription,omitempty" example:""`
}

type RecoveryBadRequestResponse struct {
	Error RecoveryBadRequestError `json:"error"`
}

type RecoveryBadRequestError struct {
	ErrorCode        int    `json:"errorCode" example:"1"`
	ErrorDescription string `json:"errorDescription,omitempty" example:"Аккаунт с указанным Email не существует"`
}

type RecoveryNotFoundResponse struct {
	Error RecoveryNotFoundError `json:"error"`
}

type RecoveryNotFoundError struct {
	ErrorCode        int    `json:"errorCode" example:"1"`
	ErrorDescription string `json:"errorDescription,omitempty" example:"Уникальный суффикс восстановления не найден"`
}

type RecoveryUnprocessableEntityResponse struct {
	Error RecoveryUnprocessableEntityError `json:"error"`
}

type RecoveryUnprocessableEntityError struct {
	ErrorCode        int    `json:"errorCode" example:"2"`
	ErrorDescription string `json:"errorDescription,omitempty" example:"Время действия ссылки истекло, повторите процедуру"`
}

type RecoveryInternalServerErrorResponse struct {
	Error RecoveryInternalServerError `json:"error"`
}

type RecoveryInternalServerError struct {
	ErrorCode        int    `json:"errorCode" example:"999"`
	ErrorDescription string `json:"errorDescription,omitempty" example:"Внутренняя ошибка"`
}
