// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-01-12 22:55:15.49435 +0800 CST m=+0.034748418

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/article": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文章"
                ],
                "summary": "新建文章",
                "parameters": [
                    {
                        "description": "文章",
                        "name": "article",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/service.ArticleCreateService"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.ArticleResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/article/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文章"
                ],
                "summary": "文章详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.ArticleResponse"
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
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文章"
                ],
                "summary": "更新文章",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "更新",
                        "name": "update",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/service.ArticleUpdateService"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.ArticleResponse"
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
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文章"
                ],
                "summary": "删除文章",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/articles": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "文章"
                ],
                "summary": "文章列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "当前页数",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.ArticleListResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/code": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "通用"
                ],
                "summary": "发送短信验证码",
                "parameters": [
                    {
                        "description": "验证码",
                        "name": "code",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/service.SendCodeService"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/sentence": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "通用"
                ],
                "summary": "随机获取定场诗",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.SentenceResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/topic": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "话题"
                ],
                "summary": "新建话题",
                "parameters": [
                    {
                        "description": "话题",
                        "name": "topic",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/service.TopicCreateService"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.TopicResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/topic/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "话题"
                ],
                "summary": "话题详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.TopicResponse"
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
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "话题"
                ],
                "summary": "更新话题",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "更新",
                        "name": "update",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/service.TopicUpdateService"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.TopicResponse"
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
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "话题"
                ],
                "summary": "删除话题",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/topics": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "话题"
                ],
                "summary": "话题列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "当前页数",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.TopicListResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/user": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "当前用户详情",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.UserResponse"
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
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "更新用户",
                "parameters": [
                    {
                        "description": "更新",
                        "name": "update",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/service.UserUpdateService"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.UserResponse"
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
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "删除用户",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/user/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "登录",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/service.UserLoginService"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/user/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "注册",
                        "name": "register",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/service.UserRegisterService"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.UserResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/user/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "指定用户详情",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.UserResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/users": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "每页数量",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "当前页数",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/serializer.UserListResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "serializer.Article": {
            "type": "object",
            "required": [
                "content",
                "profile",
                "tags",
                "title"
            ],
            "properties": {
                "content": {
                    "description": "内容",
                    "type": "string"
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "creator": {
                    "description": "作者",
                    "type": "object",
                    "$ref": "#/definitions/serializer.User"
                },
                "image": {
                    "description": "banner图",
                    "type": "string"
                },
                "profile": {
                    "description": "简介",
                    "type": "string"
                },
                "tags": {
                    "description": "标签",
                    "type": "string"
                },
                "title": {
                    "description": "标题",
                    "type": "string"
                },
                "topic": {
                    "description": "话题",
                    "type": "object",
                    "$ref": "#/definitions/serializer.Topic"
                }
            }
        },
        "serializer.ArticleItem": {
            "type": "object",
            "required": [
                "profile",
                "tags",
                "title"
            ],
            "properties": {
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "creatorAvatar": {
                    "description": "作者头像",
                    "type": "string"
                },
                "creatorName": {
                    "description": "作者名字",
                    "type": "string"
                },
                "image": {
                    "description": "banner图",
                    "type": "string"
                },
                "profile": {
                    "description": "简介",
                    "type": "string"
                },
                "tags": {
                    "description": "标签",
                    "type": "string"
                },
                "title": {
                    "description": "标题",
                    "type": "string"
                },
                "topicName": {
                    "description": "话题名称",
                    "type": "string"
                }
            }
        },
        "serializer.ArticleListResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "状态码",
                    "type": "integer",
                    "example": 2000
                },
                "count": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/serializer.ArticleItem"
                    }
                },
                "message": {
                    "description": "自定义消息",
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "serializer.ArticleResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "状态码",
                    "type": "integer",
                    "example": 2000
                },
                "data": {
                    "type": "object",
                    "$ref": "#/definitions/serializer.Article"
                },
                "message": {
                    "description": "自定义消息",
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "serializer.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "状态码",
                    "type": "integer",
                    "example": 2000
                },
                "data": {
                    "description": "响应数据",
                    "type": "object"
                },
                "message": {
                    "description": "自定义消息",
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "serializer.Sentence": {
            "type": "object",
            "properties": {
                "lines": {
                    "description": "句子",
                    "type": "string"
                }
            }
        },
        "serializer.SentenceResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "状态码",
                    "type": "integer",
                    "example": 2000
                },
                "data": {
                    "type": "object",
                    "$ref": "#/definitions/serializer.Sentence"
                },
                "message": {
                    "description": "自定义消息",
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "serializer.Topic": {
            "type": "object",
            "properties": {
                "icon": {
                    "description": "图标",
                    "type": "string"
                },
                "image": {
                    "description": "背景",
                    "type": "string"
                },
                "info": {
                    "description": "详情",
                    "type": "string"
                },
                "title": {
                    "description": "标题",
                    "type": "string"
                }
            }
        },
        "serializer.TopicListResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "状态码",
                    "type": "integer",
                    "example": 2000
                },
                "count": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/serializer.Topic"
                    }
                },
                "message": {
                    "description": "自定义消息",
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "serializer.TopicResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "状态码",
                    "type": "integer",
                    "example": 2000
                },
                "data": {
                    "type": "object",
                    "$ref": "#/definitions/serializer.Topic"
                },
                "message": {
                    "description": "自定义消息",
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "serializer.User": {
            "type": "object",
            "properties": {
                "avatar": {
                    "description": "用户头像",
                    "type": "string"
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "email": {
                    "description": "用户邮箱",
                    "type": "string"
                },
                "id": {
                    "description": "用户ID",
                    "type": "integer"
                },
                "status": {
                    "description": "用户状态",
                    "type": "integer"
                },
                "username": {
                    "description": "用户名称",
                    "type": "string"
                }
            }
        },
        "serializer.UserListResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "状态码",
                    "type": "integer",
                    "example": 2000
                },
                "count": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/serializer.User"
                    }
                },
                "message": {
                    "description": "自定义消息",
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "serializer.UserResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "状态码",
                    "type": "integer",
                    "example": 2000
                },
                "data": {
                    "description": "用户信息",
                    "type": "object",
                    "$ref": "#/definitions/serializer.User"
                },
                "message": {
                    "description": "自定义消息",
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "service.ArticleCreateService": {
            "type": "object",
            "required": [
                "content",
                "profile",
                "tags",
                "title",
                "topicId"
            ],
            "properties": {
                "content": {
                    "description": "内容",
                    "type": "string"
                },
                "image": {
                    "description": "banner图",
                    "type": "string"
                },
                "profile": {
                    "description": "简介",
                    "type": "string"
                },
                "tags": {
                    "description": "标签",
                    "type": "string"
                },
                "title": {
                    "description": "标题",
                    "type": "string"
                },
                "topicId": {
                    "description": "话题ID",
                    "type": "integer"
                }
            }
        },
        "service.ArticleUpdateService": {
            "type": "object",
            "required": [
                "content",
                "profile",
                "tags",
                "title"
            ],
            "properties": {
                "content": {
                    "description": "内容",
                    "type": "string"
                },
                "image": {
                    "description": "banner图",
                    "type": "string"
                },
                "profile": {
                    "description": "简介",
                    "type": "string"
                },
                "tags": {
                    "description": "标签",
                    "type": "string"
                },
                "title": {
                    "description": "标题",
                    "type": "string"
                },
                "topicId": {
                    "description": "话题ID",
                    "type": "integer"
                }
            }
        },
        "service.SendCodeService": {
            "type": "object",
            "required": [
                "mobile",
                "type"
            ],
            "properties": {
                "mobile": {
                    "description": "手机号码",
                    "type": "string"
                },
                "type": {
                    "description": "类型",
                    "type": "integer"
                }
            }
        },
        "service.TopicCreateService": {
            "type": "object",
            "required": [
                "icon",
                "title"
            ],
            "properties": {
                "icon": {
                    "description": "图标",
                    "type": "string"
                },
                "image": {
                    "description": "背景",
                    "type": "string"
                },
                "info": {
                    "description": "详情",
                    "type": "string"
                },
                "title": {
                    "description": "标题",
                    "type": "string"
                }
            }
        },
        "service.TopicUpdateService": {
            "type": "object",
            "required": [
                "icon",
                "title"
            ],
            "properties": {
                "icon": {
                    "description": "图标",
                    "type": "string"
                },
                "image": {
                    "description": "背景",
                    "type": "string"
                },
                "info": {
                    "description": "详情",
                    "type": "string"
                },
                "title": {
                    "description": "标题",
                    "type": "string"
                }
            }
        },
        "service.UserLoginService": {
            "type": "object",
            "required": [
                "mobile"
            ],
            "properties": {
                "code": {
                    "description": "验证码",
                    "type": "string"
                },
                "mobile": {
                    "description": "手机号",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                }
            }
        },
        "service.UserRegisterService": {
            "type": "object",
            "required": [
                "code",
                "mobile",
                "password"
            ],
            "properties": {
                "code": {
                    "description": "验证码",
                    "type": "string"
                },
                "mobile": {
                    "description": "手机号码",
                    "type": "string"
                },
                "password": {
                    "description": "登录密码",
                    "type": "string"
                }
            }
        },
        "service.UserUpdateService": {
            "type": "object",
            "properties": {
                "avatar": {
                    "description": "头像",
                    "type": "string"
                },
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "token",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "127.0.0.1:3000",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Island Go API",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
