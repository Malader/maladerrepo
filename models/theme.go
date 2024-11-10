package models

type ThemeCategory string

const (
	THEME     ThemeCategory = "THEME"
	METATHEME ThemeCategory = "METATHEME"
)

type Theme struct {
	ID           string        `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	Category     ThemeCategory `gorm:"type:varchar(20);not null" json:"category" binding:"required,oneof=THEME METATHEME"`
	Name         string        `gorm:"unique;not null" json:"name" binding:"required,min=3,max=100"`
	Metathemes   []Theme       `gorm:"many2many:theme_metathemes;association_jointable_foreignkey:metatheme_id" json:"metathemes,omitempty"`
	GameSearches []GameSearch  `gorm:"many2many:game_search_metathemes;association_jointable_foreignkey:game_search_id" json:"game_searches,omitempty"`
}

type AddThemeRequest struct {
	Theme Theme `json:"theme" binding:"required"`
}

type AddThemeResponse struct {
	Error Error `json:"error"`
}

type ConfirmThemeRequest struct {
	Theme      Theme    `json:"theme" binding:"required"`
	Metathemes []string `json:"metathemes"`
}

type ConfirmThemeResponse struct {
	Error Error `json:"error"`
}

type GameSearch struct {
	ID         string  `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	UserID     string  `gorm:"type:uuid;not null" json:"user_id"`
	Metathemes []Theme `gorm:"many2many:game_search_metathemes;association_jointable_foreignkey:metatheme_id" json:"metathemes"`
	Status     string  `gorm:"type:varchar(20);not null" json:"status"`      // "searching", "found", "stopped"
	Spectators []User  `gorm:"many2many:game_spectators;" json:"spectators"` // Добавленное поле для зрителей
}

type AddGameSearchRequest struct {
	Metathemes []string `json:"metathemes" binding:"required"`
}

type AddGameSearchResponse struct {
	Error Error `json:"error"`
}

type StopGameSearchResponse struct {
	Error Error `json:"error"`
}

type AddSpectatorResponse struct {
	Error Error `json:"error"`
}
