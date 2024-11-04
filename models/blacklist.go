// models/blacklist.go
package models

// AddBlacklistResponse представляет ответ на добавление в черный список
type AddBlacklistResponse struct {
	Error Error `json:"error"`
}
