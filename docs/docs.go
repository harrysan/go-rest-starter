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
        "/api/v1/login": {
            "post": {
                "tags": [
                    "LoginAPI"
                ],
                "summary": "Login system with username and password",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.LoginForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ResponseResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/schema.LoginToken"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.ResponseResult"
                        }
                    }
                }
            }
        },
        "/api/v1/logout": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "LoginAPI"
                ],
                "summary": "Logout system",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.ResponseResult"
                        }
                    }
                }
            }
        },
        "/api/v1/users": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "UserAPI"
                ],
                "summary": "Query user list",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username for login",
                        "name": "username",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Name of user",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Status of user (activated, freezed)",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ResponseResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/schema.User"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/util.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.ResponseResult"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "UserAPI"
                ],
                "summary": "Create user record",
                "parameters": [
                    {
                        "description": "Request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.UserForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ResponseResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/schema.User"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.ResponseResult"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/util.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.ResponseResult"
                        }
                    }
                }
            }
        },
        "/api/v1/users/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "UserAPI"
                ],
                "summary": "Get user record by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "unique id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/util.ResponseResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/schema.User"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/util.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.ResponseResult"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "UserAPI"
                ],
                "summary": "Update user record by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "unique id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.UserForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.ResponseResult"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.ResponseResult"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/util.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.ResponseResult"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "UserAPI"
                ],
                "summary": "Delete user record by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "unique id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.ResponseResult"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/util.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.ResponseResult"
                        }
                    }
                }
            }
        },
        "/api/v1/users/{id}/reset-pwd": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "LoginAPI"
                ],
                "summary": "Change current user password",
                "parameters": [
                    {
                        "type": "string",
                        "description": "unique id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Request body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.UpdateLoginPassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/util.ResponseResult"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/util.ResponseResult"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/util.ResponseResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/util.ResponseResult"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "errors.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "detail": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "schema.LoginForm": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "description": "Login password (md5 hash)",
                    "type": "string"
                },
                "username": {
                    "description": "Login name",
                    "type": "string"
                }
            }
        },
        "schema.LoginToken": {
            "type": "object",
            "properties": {
                "access_token": {
                    "description": "Access token (JWT)",
                    "type": "string"
                },
                "expires_at": {
                    "description": "Expired time (Unit: second)",
                    "type": "integer"
                },
                "token_type": {
                    "description": "Token type (Usage: Authorization=${token_type} ${access_token})",
                    "type": "string"
                }
            }
        },
        "schema.UpdateLoginPassword": {
            "type": "object",
            "required": [
                "new_password",
                "old_password"
            ],
            "properties": {
                "new_password": {
                    "description": "New password (md5 hash)",
                    "type": "string"
                },
                "old_password": {
                    "description": "Old password (md5 hash)",
                    "type": "string"
                }
            }
        },
        "schema.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "Create time",
                    "type": "string"
                },
                "email": {
                    "description": "Email of user",
                    "type": "string"
                },
                "id": {
                    "description": "Unique ID",
                    "type": "integer"
                },
                "name": {
                    "description": "Name of user",
                    "type": "string"
                },
                "phone": {
                    "description": "Phone number of user",
                    "type": "string"
                },
                "remark": {
                    "description": "Remark of user",
                    "type": "string"
                },
                "status": {
                    "description": "Status of user (activated, freezed)",
                    "type": "string"
                },
                "updated_at": {
                    "description": "Update time",
                    "type": "string"
                },
                "username": {
                    "description": "Username for login",
                    "type": "string"
                }
            }
        },
        "schema.UserForm": {
            "type": "object",
            "required": [
                "name",
                "status",
                "username"
            ],
            "properties": {
                "email": {
                    "description": "Email of user",
                    "type": "string",
                    "maxLength": 128
                },
                "name": {
                    "description": "Name of user",
                    "type": "string",
                    "maxLength": 64
                },
                "password": {
                    "description": "Password for login (md5 hash)",
                    "type": "string",
                    "maxLength": 64
                },
                "phone": {
                    "description": "Phone number of user",
                    "type": "string",
                    "maxLength": 32
                },
                "remark": {
                    "description": "Remark of user",
                    "type": "string",
                    "maxLength": 1024
                },
                "status": {
                    "description": "Status of user (activated, freezed)",
                    "type": "string",
                    "enum": [
                        "activated",
                        "freezed"
                    ]
                },
                "username": {
                    "description": "Username for login",
                    "type": "string",
                    "maxLength": 64
                }
            }
        },
        "util.ResponseResult": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "$ref": "#/definitions/errors.Error"
                },
                "success": {
                    "type": "boolean"
                },
                "total": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8040",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Finance Tracker API",
	Description:      "A finance tracker API service based on golang.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
