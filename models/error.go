package models

// Error представляет структуру ошибки в ответах API
type Error struct {
	ErrorCode        int    `json:"errorCode"`
	ErrorDescription string `json:"errorDescription,omitempty"`
}
