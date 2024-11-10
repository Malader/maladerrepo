package handlers

import (
	"net/http"
	"strings"

	"github.com/Malader/maladerrepo/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AddThemeHandler обрабатывает запросы на предложение новой темы/метатемы
// @Summary Предложение новой темы/метатемы
// @Description Данный метод позволяет предложить новую тему/метатему в системе
// @Tags theme
// @Accept json
// @Produce json
// @Param theme body models.AddThemeRequest true "Предлагаемая тема"
// @Success 200 {object} models.AddThemeResponse "Тема отправлена на валидацию"
// @Failure 400 {object} models.AddThemeResponse "Тема уже существует или некорректные данные"
// @Failure 500 {object} models.AddThemeResponse "Внутренняя ошибка"
// @Router /theme [post]
func AddThemeHandler(c *gin.Context) {
	var req models.AddThemeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.AddThemeResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Некорректные данные запроса",
			},
		})
		return
	}

	theme := req.Theme

	var existingTheme models.Theme
	if err := DB.Where("name = ?", theme.Name).First(&existingTheme).Error; err == nil {
		c.JSON(http.StatusBadRequest, models.AddThemeResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Тема уже существует",
			},
		})
		return
	} else if err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, models.AddThemeResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	if err := DB.Create(&theme).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.AddThemeResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.AddThemeResponse{
		Error: models.Error{
			ErrorCode:        0,
			ErrorDescription: "",
		},
	})
}

// ConfirmThemeHandler обрабатывает запросы на подтверждение или отклонение новой темы/метатемы
// @Summary Подтверждение/отклонение новой темы/метатемы
// @Description Данный метод позволяет зарегистрировать/отклонить новую тему/метатему
// @Tags theme
// @Accept json
// @Produce json
// @Param confirmation path string true "Подтверждение темы/метатемы (YES или NO)" example:"YES"
// @Param theme body models.ConfirmThemeRequest true "Предлагаемая тема"
// @Success 201 {object} models.ConfirmThemeResponse "Тема успешно зарегистрирована"
// @Failure 400 {object} models.ConfirmThemeResponse "Некорректные данные или тема уже существует"
// @Failure 422 {object} models.ConfirmThemeResponse "Некорректные данные для подтверждения темы"
// @Failure 500 {object} models.ConfirmThemeResponse "Внутренняя ошибка"
// @Router /theme/{confirmation} [put]
func ConfirmThemeHandler(c *gin.Context) {
	confirmation := strings.ToUpper(c.Param("confirmation"))

	if confirmation != "YES" && confirmation != "NO" {
		c.JSON(http.StatusUnprocessableEntity, models.ConfirmThemeResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Некорректное поведение системы",
			},
		})
		return
	}

	var req models.ConfirmThemeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ConfirmThemeResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Некорректные данные запроса",
			},
		})
		return
	}

	theme := req.Theme

	var existingTheme models.Theme
	if err := DB.Where("name = ?", theme.Name).First(&existingTheme).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			if confirmation == "YES" {
				if err := DB.Create(&theme).Error; err != nil {
					c.JSON(http.StatusInternalServerError, models.ConfirmThemeResponse{
						Error: models.Error{
							ErrorCode:        999,
							ErrorDescription: "Внутренняя ошибка",
						},
					})
					return
				}

				if theme.Category == models.THEME && len(req.Metathemes) > 0 {
					var metathemes []models.Theme
					if err := DB.Where("id IN (?) AND category = ?", req.Metathemes, models.METATHEME).Find(&metathemes).Error; err != nil {
						c.JSON(http.StatusUnprocessableEntity, models.ConfirmThemeResponse{
							Error: models.Error{
								ErrorCode:        2,
								ErrorDescription: "Некорректный формат метатем в запросе",
							},
						})
						return
					}

					if len(metathemes) != len(req.Metathemes) {
						c.JSON(http.StatusUnprocessableEntity, models.ConfirmThemeResponse{
							Error: models.Error{
								ErrorCode:        1,
								ErrorDescription: "Для метатемы нельзя указывать в запросе другие метатемы",
							},
						})
						return
					}

					if err := DB.Model(&theme).Association("Metathemes").Append(metathemes).Error; err != nil {
						c.JSON(http.StatusInternalServerError, models.ConfirmThemeResponse{
							Error: models.Error{
								ErrorCode:        999,
								ErrorDescription: "Внутренняя ошибка",
							},
						})
						return
					}
				}

				c.JSON(http.StatusCreated, models.ConfirmThemeResponse{
					Error: models.Error{
						ErrorCode:        0,
						ErrorDescription: "",
					},
				})
				return
			} else { // confirmation == "NO"
				// Логика отклонения темы (можно добавить, если необходимо)
				c.JSON(http.StatusOK, models.ConfirmThemeResponse{
					Error: models.Error{
						ErrorCode:        0,
						ErrorDescription: "",
					},
				})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, models.ConfirmThemeResponse{
				Error: models.Error{
					ErrorCode:        999,
					ErrorDescription: "Внутренняя ошибка",
				},
			})
			return
		}
	}

	if confirmation == "YES" {
		c.JSON(http.StatusBadRequest, models.ConfirmThemeResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Тема уже существует",
			},
		})
		return
	} else { // confirmation == "NO"
		// Логика отклонения темы (можно добавить, если необходимо)
		c.JSON(http.StatusOK, models.ConfirmThemeResponse{
			Error: models.Error{
				ErrorCode:        0,
				ErrorDescription: "",
			},
		})
		return
	}
}
