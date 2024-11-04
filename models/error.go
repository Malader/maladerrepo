package models

type Error struct {
	ErrorCode        int    `json:"errorCode"`
	ErrorDescription string `json:"errorDescription,omitempty"`
}
