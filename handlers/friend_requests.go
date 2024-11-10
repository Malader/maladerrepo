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

	var existingRequest models.FriendRequest
	if err := DB.Where("from_user_id = ? AND to_user_id = ?", sender.ID, target.ID).First(&existingRequest).Error; err == nil {
		c.JSON(http.StatusBadRequest, models.SendFriendRequestResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Запрос дружбы уже отправлен",
			},
		})
		return
	}

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

	if confirmation != "YES" && confirmation != "NO" {
		c.JSON(http.StatusBadRequest, models.ConfirmFriendRequestResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Некорректное поведение системы",
			},
		})
		return
	}

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
		friendRequest.Status = models.Accepted

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
		friendRequest.Status = models.Rejected
	}

	if err := DB.Save(&friendRequest).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ConfirmFriendRequestResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

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

	var count int64
	DB.Table("user_friends").Where("user_id = ? AND friend_id = ?", user.ID, friend.ID).Count(&count)
	if count == 0 {
		c.JSON(http.StatusBadRequest, models.DeleteFriendResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Пользователь не является вашим другом",
			},
		})
		return
	}

	DB.Model(&user).Association("Friends").Delete(&friend)
	if DB.Error != nil {
		c.JSON(http.StatusInternalServerError, models.DeleteFriendResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	DB.Model(&friend).Association("Friends").Delete(&user)
	if DB.Error != nil {
		c.JSON(http.StatusInternalServerError, models.DeleteFriendResponse{
			Error: models.Error{
				ErrorCode:        999,
				ErrorDescription: "Внутренняя ошибка",
			},
		})
		return
	}

	c.JSON(http.StatusOK, models.DeleteFriendResponse{
		Error: models.Error{
			ErrorCode:        0,
			ErrorDescription: "",
		},
	})
}
