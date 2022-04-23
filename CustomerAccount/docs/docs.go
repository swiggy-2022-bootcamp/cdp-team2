// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2022-04-23 18:39:36.8993584 +0530 IST m=+0.123177001

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
        "termsOfService": "demo.com",
        "contact": {
            "name": "API Support"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/account/{id}": {
            "get": {
                "description": "API to Get Customers-Frontstore's Customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer Account Get"
                ],
                "summary": "Customer Account Get Route",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Customer Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Account2"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/string"
                        }
                    }
                }
            },
            "put": {
                "description": "API to Update Customer-Frontstore's Customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer Account Update"
                ],
                "summary": "Customer Account Create Route",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Customer Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Account Update",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/%7B%7D"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Account"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/string"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "API to check Customer-Frontstore's health",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "Health Check Route",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/string"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "API to Create Customers-Frontstore Customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Create"
                ],
                "summary": "Account Create Route",
                "parameters": [
                    {
                        "description": "Account info",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Account"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Account"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Account": {
            "type": "object",
            "properties": {
                "agree": {
                    "type": "string"
                },
                "confirm": {
                    "type": "string"
                },
                "customer_group_id": {
                    "type": "integer"
                },
                "customer_id": {
                    "type": "string"
                },
                "date_added": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "telephone": {
                    "type": "string"
                },
                "user_balance": {
                    "type": "number"
                }
            }
        },
        "models.Account2": {
            "type": "object",
            "properties": {
                "cart": {
                    "type": "object",
                    "$ref": "#/definitions/models.Cart"
                },
                "customer_group_id": {
                    "type": "integer"
                },
                "customer_id": {
                    "type": "string"
                },
                "date_added": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "reward": {
                    "type": "object",
                    "$ref": "#/definitions/models.Reward"
                },
                "telephone": {
                    "type": "string"
                },
                "user_balance": {
                    "description": "RewardTotal     Reward    ` + "`" + `json:\"reward_total\"` + "`" + `",
                    "type": "number"
                }
            }
        },
        "models.Cart": {
            "type": "object",
            "properties": {
                "product": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Product"
                    }
                }
            }
        },
        "models.Product": {
            "type": "object",
            "properties": {
                "product_id": {
                    "type": "string"
                },
                "quality": {
                    "type": "integer"
                }
            }
        },
        "models.Reward": {
            "type": "object",
            "properties": {
                "customer_id": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "points": {
                    "type": "integer"
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
	Host:        "localhost:8092",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Account API",
	Description: "This is Account CRUD service.",
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
