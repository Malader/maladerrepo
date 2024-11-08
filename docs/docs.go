// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/rooms/create": {
            "post": {
                "description": "Создает новую игровую комнату с заданным именем",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "room"
                ],
                "summary": "Создание новой игровой комнаты",
                "parameters": [
                    {
                        "description": "Информация о комнате",
                        "name": "room",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.CreateRoomRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/handlers.CreateRoomResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.CreateRoomResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.CreateRoomResponse"
                        }
                    }
                }
            }
        },
        "/rooms/{room_id}/players": {
            "get": {
                "description": "Возвращает список игроков в указанной игровой комнате",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "room"
                ],
                "summary": "Получение информации об игроках в комнате",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор комнаты",
                        "name": "room_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.PlayersInRoomResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.PlayersInRoomResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.PlayersInRoomResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Добавляет пользователя в указанную игровую комнату",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "room"
                ],
                "summary": "Добавление игрока в комнату",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор комнаты",
                        "name": "room_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Информация об игроке",
                        "name": "player",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.AddPlayerToRoomRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.AddPlayerToRoomResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.AddPlayerToRoomResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handlers.AddPlayerToRoomResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.AddPlayerToRoomResponse"
                        }
                    }
                }
            }
        },
        "/user/authorize": {
            "post": {
                "description": "Позволяет авторизовать пользователя в системе",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Авторизация пользователя",
                "parameters": [
                    {
                        "description": "Учетные данные пользователя",
                        "name": "credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AuthorizeRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.AuthorizeResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.AuthorizeResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.AuthorizeResponse"
                        }
                    }
                }
            }
        },
        "/user/recovery": {
            "post": {
                "description": "Позволяет начать процесс восстановления пароля через email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recovery"
                ],
                "summary": "Инициация восстановления пароля",
                "parameters": [
                    {
                        "description": "Email пользователя",
                        "name": "recovery",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RecoveryRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.RecoveryResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.RecoveryResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.RecoveryResponse"
                        }
                    }
                }
            }
        },
        "/user/recovery/{recoverySuffix}": {
            "patch": {
                "description": "Позволяет завершить процесс восстановления пароля по ссылке",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recovery"
                ],
                "summary": "Завершение восстановления пароля",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Уникальный суффикс восстановления",
                        "name": "recoverySuffix",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Новый пароль",
                        "name": "recovery",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ChangePasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.ChangePasswordResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ChangePasswordResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ChangePasswordResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/models.ChangePasswordResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ChangePasswordResponse"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "description": "Позволяет зарегистрировать пользователя в системе",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Регистрация нового пользователя",
                "parameters": [
                    {
                        "description": "Данные пользователя для регистрации",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.RegisterResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.RegisterResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/models.RegisterResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.RegisterResponse"
                        }
                    }
                }
            }
        },
        "/user/{id}/blackList/{username}": {
            "put": {
                "description": "Позволяет добавить пользователя в черный список",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Добавление в черный список",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор пользователя, добавляющего в черный список",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Имя пользователя, добавляемого в черный список",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AddBlacklistResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.AddBlacklistResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.AddBlacklistResponse"
                        }
                    }
                }
            }
        },
        "/user/{id}/friends": {
            "get": {
                "description": "Позволяет получить список друзей пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Получение списка друзей",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор пользователя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.GetFriendsResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.GetFriendsResponse"
                        }
                    },
                    "504": {
                        "description": "Gateway Timeout",
                        "schema": {
                            "$ref": "#/definitions/models.GetFriendsResponse"
                        }
                    }
                }
            }
        },
        "/user/{id}/friends/{username}": {
            "post": {
                "description": "Позволяет отправить запрос дружбы от одного пользователя другому",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Отправление запроса дружбы",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор пользователя, отправляющего запрос дружбы",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Имя пользователя, получающего запрос дружбы",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.SendFriendRequestResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.SendFriendRequestResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.SendFriendRequestResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.SendFriendRequestResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Позволяет удалить из друзей пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Удаление из друзей",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор пользователя, удаляющего друга",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Имя пользователя, удаляемого из списка друзей",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.DeleteFriendResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.DeleteFriendResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.DeleteFriendResponse"
                        }
                    }
                }
            }
        },
        "/user/{id}/friends/{username}/{confirmation}": {
            "put": {
                "description": "Позволяет ответить на запрос дружбы от одного пользователя другому",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Подтверждение/отклонение запроса дружбы",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор пользователя, принимающего запрос дружбы",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Имя пользователя, отправившего запрос дружбы",
                        "name": "username",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Подтверждение запроса дружбы (YES или NO)",
                        "name": "confirmation",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ConfirmFriendRequestResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ConfirmFriendRequestResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.ConfirmFriendRequestResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ConfirmFriendRequestResponse"
                        }
                    }
                }
            }
        },
        "/user/{id}/profile": {
            "get": {
                "description": "Позволяет получить информацию о пользователе",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Получение информации профиля",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор профиля",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.GetProfileResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.GetProfileResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.GetProfileResponse"
                        }
                    },
                    "504": {
                        "description": "Gateway Timeout",
                        "schema": {
                            "$ref": "#/definitions/models.GetProfileResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Позволяет изменить информацию о пользователе",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Изменение информации профиля",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор профиля",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Информация профиля",
                        "name": "profile",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateProfileRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.UpdateProfileResponse"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/models.UpdateProfileResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.UpdateProfileResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.AddPlayerToRoomRequest": {
            "type": "object",
            "required": [
                "username"
            ],
            "properties": {
                "username": {
                    "type": "string",
                    "maxLength": 25,
                    "example": "john_doe"
                }
            }
        },
        "handlers.AddPlayerToRoomResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/models.Error"
                },
                "room": {
                    "$ref": "#/definitions/handlers.Room"
                }
            }
        },
        "handlers.CreateRoomRequest": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "maxLength": 100,
                    "example": "Room One"
                }
            }
        },
        "handlers.CreateRoomResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/models.Error"
                },
                "room": {
                    "$ref": "#/definitions/handlers.Room"
                }
            }
        },
        "handlers.Room": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "players": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.User"
                    }
                }
            }
        },
        "models.AddBlacklistResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/models.Error"
                }
            }
        },
        "models.AuthorizeRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "example": "6a4a61f57bccf059abb82fc95589ebc428629326ab965390f25224e262455beb"
                },
                "username": {
                    "type": "string",
                    "maxLength": 25,
                    "example": "vasyaPupkin"
                }
            }
        },
        "models.AuthorizeResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/models.Error"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "models.ChangePasswordRequest": {
            "type": "object",
            "required": [
                "newPassword"
            ],
            "properties": {
                "newPassword": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 8,
                    "example": "newSecurePassword123"
                }
            }
        },
        "models.ChangePasswordResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/models.Error"
                }
            }
        },
        "models.ConfirmFriendRequestResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/models.Error"
                }
            }
        },
        "models.DeleteFriendResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/models.Error"
                }
            }
        },
        "models.Error": {
            "type": "object",
            "properties": {
                "errorCode": {
                    "type": "integer"
                },
                "errorDescription": {
                    "type": "string"
                }
            }
        },
        "models.Friend": {
            "type": "object",
            "properties": {
                "image": {
                    "type": "string",
                    "example": "4AAQSkZJRgABAQEAAAAAAADfjadFHADHda"
                },
                "status": {
                    "type": "string",
                    "example": "active"
                },
                "username": {
                    "type": "string",
                    "example": "VasyaPupkin"
                }
            }
        },
        "models.GetFriendsResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/models.Error"
                },
                "friends": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Friend"
                    }
                }
            }
        },
        "models.GetProfileResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "example@example.com"
                },
                "error": {
                    "$ref": "#/definitions/models.Error"
                },
                "image": {
                    "type": "string",
                    "example": "4AAQSkZJRgABAQEAAAAAAADfjadFHADHda"
                },
                "lastActivityDate": {
                    "type": "string",
                    "example": "2024-05-22T12:34:56Z"
                },
                "password": {
                    "type": "string",
                    "example": "hashedPassword123"
                },
                "registrationDate": {
                    "type": "string",
                    "example": "2024-01-01T00:00:00Z"
                },
                "username": {
                    "type": "string",
                    "example": "ExampleUser"
                }
            }
        },
        "models.PlayersInRoomResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/models.Error"
                },
                "players": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.User"
                    }
                },
                "room_id": {
                    "type": "string",
                    "example": "room123"
                }
            }
        },
        "models.RecoveryRequest": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 256,
                    "example": "v.pupkin@g.nsu.ru"
                }
            }
        },
        "models.RecoveryResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/models.Error"
                }
            }
        },
        "models.RegisterRequest": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "maxLength": 256,
                    "example": "v.pupkin@g.nsu.ru"
                },
                "password": {
                    "type": "string",
                    "example": "6a4a61f57bccf059abb82fc95589ebc428629326ab965390f25224e262455beb"
                },
                "username": {
                    "type": "string",
                    "maxLength": 25,
                    "example": "vasyaPupkin"
                }
            }
        },
        "models.RegisterResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/models.Error"
                }
            }
        },
        "models.SendFriendRequestResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/models.Error"
                }
            }
        },
        "models.UpdateProfileRequest": {
            "type": "object",
            "properties": {
                "image": {
                    "type": "string",
                    "example": "TWFuIGlzIGRpc3Rpbmd1aXNoZWQsIG5vdCBvbmx5IGJ5IGhpcyByZWFzb24sIGJ1dCAuLi4="
                },
                "password": {
                    "type": "string",
                    "example": "6a4a61f57bccf059abb82fc95589ebc428629326ab965390f25224e262455beb"
                },
                "username": {
                    "type": "string",
                    "maxLength": 25,
                    "example": "vasyaPupkin"
                }
            }
        },
        "models.UpdateProfileResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/models.Error"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "lastActivityDate": {
                    "type": "string"
                },
                "registrationDate": {
                    "type": "string"
                },
                "team": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "User Service API",
	Description:      "API сервис пользователей для взаимодействия с сервисом авторизации и сервисом игровых комнат",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
