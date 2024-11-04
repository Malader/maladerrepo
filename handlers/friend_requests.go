// handlers/friend_requests.go
package handlers

import (
	"net/http"
	"strings"

	"github.com/Malader/maladerrepo/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SendFriendRequestHandler обрабатывает запросы на отправку запроса дружбы
// @Summary Отправление запроса дружбы
// @Description Позволяет отправить запрос дружбы от одного пользователя другому
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "Идентификатор пользователя, отправляющего запрос дружбы" example:"1234567890123456789"
// @Param username path string true "Имя пользователя, получающего запрос дружбы" example:"ExampleUser"
// @Success 200 {object} models.SendFriendRequestResponse
// @Failure 400 {object} models.SendFriendRequestResponse
// @Failure 401 {object} models.SendFriendRequestResponse
// @Failure 500 {object} models.SendFriendRequestResponse
// @Router /user/{id}/friends/{username} [post]
func SendFriendRequestHandler(c *gin.Context) {
	senderID := c.Param("id")
	targetUsername := c.Param("username")

	// Найти отправителя по ID
	var sender models.User
	if err := DB.Where("id = ?", senderID).First(&sender).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, models.SendFriendRequestResponse{
				Error: models.Error{
					ErrorCode:        999,
					ErrorDescription: "Некорректное поведение системы",
				},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.SendFriendRequestResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Некорректное поведение системы",
			},
		})
		return
	}

	// Найти целевого пользователя по username
	var target models.User
	if err := DB.Where("username = ?", targetUsername).First(&target).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, models.SendFriendRequestResponse{
				Error: models.Error{
					ErrorCode:        400,
					ErrorDescription: "Пользователь с указанным username не зарегистрирован в системе",
				},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.SendFriendRequestResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Некорректное поведение системы",
			},
		})
		return
	}

	// Проверить, не отправляет ли пользователь запрос самому себе
	if sender.ID == target.ID {
		c.JSON(http.StatusBadRequest, models.SendFriendRequestResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Некорректное поведение системы",
			},
		})
		return
	}

	// Проверить, не находится ли целевой пользователь в черном списке отправителя или наоборот
	// Предполагается, что есть функция IsUserBlocked(senderID, targetID)
	// Если черный список не реализован, эту проверку можно пропустить или реализовать
	/*
		if IsUserBlocked(sender.ID, target.ID) || IsUserBlocked(target.ID, sender.ID) {
			c.JSON(http.StatusUnauthorized, models.SendFriendRequestResponse{
				Error: models.Error{
					ErrorCode:        401,
					ErrorDescription: "Пользователь с указанным username добавил вас в черный список",
				},
			})
			return
		}
	*/

	// Проверить, являются ли пользователи уже друзьями
	var count int64
	DB.Model(&sender).Where("friends.id = ?", target.ID).Count(&count)
	if count > 0 {
		c.JSON(http.StatusBadRequest, models.SendFriendRequestResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Пользователь уже является вашим другом",
			},
		})
		return
	}

	// Проверить, не отправлен ли уже запрос дружбы
	var existingRequest models.FriendRequest
	if err := DB.Where("from_user_id = ? AND to_user_id = ?", sender.ID, target.ID).First(&existingRequest).Error; err == nil {
		// Запрос уже существует
		c.JSON(http.StatusBadRequest, models.SendFriendRequestResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Запрос дружбы уже отправлен",
			},
		})
		return
	}

	// Создать новый запрос дружбы
	friendRequest := models.FriendRequest{
		FromUserID: sender.ID,
		ToUserID:   target.ID,
		Status:     models.Pending,
	}

	if err := DB.Create(&friendRequest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.SendFriendRequestResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	// Успешно отправлено
	c.JSON(http.StatusOK, models.SendFriendRequestResponse{
		Error: models.Error{
			ErrorCode:        0,
			ErrorDescription: "",
		},
	})
}

// ConfirmFriendRequestHandler обрабатывает запросы на подтверждение или отклонение запроса дружбы
// @Summary Подтверждение/отклонение запроса дружбы
// @Description Позволяет ответить на запрос дружбы от одного пользователя другому
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "Идентификатор пользователя, принимающего запрос дружбы" example:"1234567890123456789"
// @Param username path string true "Имя пользователя, отправившего запрос дружбы" example:"ExampleUser"
// @Param confirmation path string true "Подтверждение запроса дружбы (YES или NO)" example:"YES"
// @Success 200 {object} models.ConfirmFriendRequestResponse
// @Failure 400 {object} models.ConfirmFriendRequestResponse
// @Failure 401 {object} models.ConfirmFriendRequestResponse
// @Failure 500 {object} models.ConfirmFriendRequestResponse
// @Router /user/{id}/friends/{username}/{confirmation} [put]
func ConfirmFriendRequestHandler(c *gin.Context) {
	receiverID := c.Param("id")
	senderUsername := c.Param("username")
	confirmation := strings.ToUpper(c.Param("confirmation"))

	// Валидация параметра confirmation
	if confirmation != "YES" && confirmation != "NO" {
		c.JSON(http.StatusBadRequest, models.ConfirmFriendRequestResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Некорректное поведение системы",
			},
		})
		return
	}

	// Найти получателя по ID
	var receiver models.User
	if err := DB.Where("id = ?", receiverID).First(&receiver).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, models.ConfirmFriendRequestResponse{
				Error: models.Error{
					ErrorCode:        999,
					ErrorDescription: "Некорректное поведение системы",
				},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ConfirmFriendRequestResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Некорректное поведение системы",
			},
		})
		return
	}

	// Найти отправителя по username
	var sender models.User
	if err := DB.Where("username = ?", senderUsername).First(&sender).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, models.ConfirmFriendRequestResponse{
				Error: models.Error{
					ErrorCode:        400,
					ErrorDescription: "Пользователь с указанным username не зарегистрирован в системе",
				},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ConfirmFriendRequestResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Некорректное поведение системы",
			},
		})
		return
	}

	// Найти запрос дружбы от отправителя к получателю
	var friendRequest models.FriendRequest
	if err := DB.Where("from_user_id = ? AND to_user_id = ?", sender.ID, receiver.ID).First(&friendRequest).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, models.ConfirmFriendRequestResponse{
				Error: models.Error{
					ErrorCode:        999,
					ErrorDescription: "Некорректное поведение системы",
				},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ConfirmFriendRequestResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Некорректное поведение системы",
			},
		})
		return
	}

	// Проверить текущий статус запроса
	if friendRequest.Status != models.Pending {
		c.JSON(http.StatusBadRequest, models.ConfirmFriendRequestResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Некорректное поведение системы",
			},
		})
		return
	}

	if confirmation == "YES" {
		// Подтверждение запроса дружбы
		friendRequest.Status = models.Accepted

		// Добавление в друзья обоих пользователей
		if err := DB.Model(&receiver).Association("Friends").Append(&sender); err != nil {
			c.JSON(http.StatusInternalServerError, models.ConfirmFriendRequestResponse{
				Error: models.Error{
					ErrorCode:        999,
					ErrorDescription: "Некорректное поведение системы",
				},
			})
			return
		}
		if err := DB.Model(&sender).Association("Friends").Append(&receiver); err != nil {
			c.JSON(http.StatusInternalServerError, models.ConfirmFriendRequestResponse{
				Error: models.Error{
					ErrorCode:        999,
					ErrorDescription: "Некорректное поведение системы",
				},
			})
			return
		}
	} else {
		// Отклонение запроса дружбы
		friendRequest.Status = models.Rejected
	}

	// Сохранение обновленного статуса запроса дружбы
	if err := DB.Save(&friendRequest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ConfirmFriendRequestResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	// Успешно обработано
	c.JSON(http.StatusOK, models.ConfirmFriendRequestResponse{
		Error: models.Error{
			ErrorCode:        0,
			ErrorDescription: "",
		},
	})
}

// DeleteFriendHandler обрабатывает запросы на удаление друга
// @Summary Удаление из друзей
// @Description Позволяет удалить из друзей пользователя
// @Tags user
// @Accept json
// @Produce json
// @Param id path string true "Идентификатор пользователя, удаляющего друга" example:"1234567890123456789"
// @Param username path string true "Имя пользователя, удаляемого из списка друзей" example:"ExampleUser"
// @Success 200 {object} models.DeleteFriendResponse
// @Failure 400 {object} models.DeleteFriendResponse
// @Failure 500 {object} models.DeleteFriendResponse
// @Router /user/{id}/friends/{username} [delete]
func DeleteFriendHandler(c *gin.Context) {
	userID := c.Param("id")
	friendUsername := c.Param("username")

	// Найти пользователя, который удаляет друга
	var user models.User
	if err := DB.Where("id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, models.DeleteFriendResponse{
				Error: models.Error{
					ErrorCode:        999,
					ErrorDescription: "Некорректное поведение системы",
				},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.DeleteFriendResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	// Найти друга по username
	var friend models.User
	if err := DB.Where("username = ?", friendUsername).First(&friend).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, models.DeleteFriendResponse{
				Error: models.Error{
					ErrorCode:        400,
					ErrorDescription: "Пользователь с указанным username не зарегистрирован в системе",
				},
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.DeleteFriendResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	// Проверить, являются ли пользователи друзьями
	has, err := DB.Model(&user).Association("Friends").Has(&friend)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DeleteFriendResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	if !has {
		c.JSON(http.StatusBadRequest, models.DeleteFriendResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Пользователь не является вашим другом",
			},
		})
		return
	}

	// Удалить дружбу в обоих направлениях
	if err := DB.Model(&user).Association("Friends").Delete(&friend); err != nil {
		c.JSON(http.StatusInternalServerError, models.DeleteFriendResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	if err := DB.Model(&friend).Association("Friends").Delete(&user); err != nil {
		c.JSON(http.StatusInternalServerError, models.DeleteFriendResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	// Успешно удалено
	c.JSON(http.StatusOK, models.DeleteFriendResponse{
		Error: models.Error{
			ErrorCode:        0,
			ErrorDescription: "",
		},
	})
}
