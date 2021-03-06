// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/checkout/{cart_id}": {
            "post": {
                "tags": [
                    "Checkout"
                ],
                "summary": "Start Checkout.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "cart_id",
                        "name": "cart_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ApiResponseWithErr"
                        }
                    }
                }
            }
        },
        "/order/{order_id}": {
            "get": {
                "tags": [
                    "Checkout"
                ],
                "summary": "Get Order .",
                "parameters": [
                    {
                        "type": "string",
                        "description": "order_id",
                        "name": "order_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ApiResponseWithErr"
                        }
                    }
                }
            }
        },
        "/pay": {
            "post": {
                "tags": [
                    "Checkout"
                ],
                "summary": "Pay Order price.",
                "parameters": [
                    {
                        "description": "pay \u0026 confirm order input",
                        "name": "pay",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.PayRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ApiResponseWithErr"
                        }
                    }
                }
            }
        },
        "/rewards": {
            "put": {
                "tags": [
                    "Checkout"
                ],
                "summary": "Apply Reward on Order.",
                "parameters": [
                    {
                        "description": "apply reward input",
                        "name": "applyReward",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api.ApplyRewardRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ApiResponseWithErr"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.ApiResponseWithErr": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "api.ApplyRewardRequest": {
            "type": "object",
            "properties": {
                "order_id": {
                    "type": "string"
                },
                "rewards": {
                    "type": "integer"
                }
            }
        },
        "api.PayRequest": {
            "type": "object",
            "properties": {
                "order_id": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "35.84.28.237:30206",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Checkout Api",
	Description:      "This is checkout service.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
