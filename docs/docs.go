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
        "/orders": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get a order board",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "Get a order board",
                "parameters": [
                    {
                        "enum": [
                            "live",
                            "history",
                            "removed"
                        ],
                        "type": "string",
                        "x-enum-varnames": [
                            "Live",
                            "History",
                            "Removed"
                        ],
                        "description": "FIXME: add pagination with var below\npageReq\nif true, return orders created by the user",
                        "name": "board_type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/api.pageResp"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.Order"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.errorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.errorResp"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Make a order",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "Make a order",
                "parameters": [
                    {
                        "description": "order id to attend and user's email",
                        "name": "jsonBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.makeOrderBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.errorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.errorResp"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Delete a order",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "Delete a order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of order",
                        "name": "order_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.errorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.errorResp"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Take a order",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "order"
                ],
                "summary": "Take a order",
                "parameters": [
                    {
                        "description": "order id to attend and user's email",
                        "name": "jsonBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.takeOrderBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.errorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.errorResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.errorResp": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "request_id": {
                    "type": "string"
                }
            }
        },
        "api.makeOrderBody": {
            "type": "object",
            "required": [
                "action",
                "amount",
                "price"
            ],
            "properties": {
                "action": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.OrderAction"
                        }
                    ],
                    "example": "buy"
                },
                "amount": {
                    "type": "integer",
                    "minimum": 1,
                    "example": 100
                },
                "price": {
                    "type": "integer",
                    "minimum": 1,
                    "example": 10
                }
            }
        },
        "api.pageResp": {
            "type": "object",
            "properties": {
                "data": {},
                "next": {
                    "type": "string",
                    "example": "next cursor value"
                }
            }
        },
        "api.takeOrderBody": {
            "type": "object",
            "required": [
                "action",
                "amount"
            ],
            "properties": {
                "action": {
                    "description": "FIXME:user_id should be retrieved from the user's jwt token",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.OrderAction"
                        }
                    ],
                    "example": "buy"
                },
                "amount": {
                    "type": "integer",
                    "minimum": 1,
                    "example": 100
                }
            }
        },
        "models.Order": {
            "type": "object",
            "properties": {
                "action": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.OrderAction"
                        }
                    ],
                    "example": "buy"
                },
                "amount": {
                    "type": "integer",
                    "example": 100
                },
                "created_at": {
                    "type": "string",
                    "example": "2021-01-01T00:00:00Z"
                },
                "id": {
                    "type": "string",
                    "example": "uuid"
                },
                "price": {
                    "description": "using int instead of float64 to avoid floating point precision issue",
                    "type": "integer",
                    "example": 10
                }
            }
        },
        "models.OrderAction": {
            "type": "string",
            "enum": [
                "buy",
                "sell"
            ],
            "x-enum-varnames": [
                "Buy",
                "Sell"
            ]
        },
        "models.OrderBoardType": {
            "type": "string",
            "enum": [
                "live",
                "history",
                "removed"
            ],
            "x-enum-varnames": [
                "Live",
                "History",
                "Removed"
            ]
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Type \"Bearer\" followed by a space and JWT.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v1",
	Host:             "kickstart.gcp_project_name.tw",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Order (aka Broadcast Service) API",
	Description:      "This service provides message broadcasting service for official accounts,\nit also supports audience filtering and performance reporting.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
