// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2022-04-19 20:46:59.9392674 +0530 IST m=+0.072217401

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
            "name": "API Support",
            "url": "http://demo.com/support"
        },
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/customers": {
            "post": {
                "description": "API to Create Customer-Admin's Customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Create"
                ],
                "summary": "Customer Create Route",
                "parameters": [
                    {
                        "description": "Customer info",
                        "name": "Customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Customer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Customer"
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
        "/customers/email/{email}": {
            "get": {
                "description": "API to Get Customer-Admin's Customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer Get"
                ],
                "summary": "Customer Get By Email Route",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Customer email",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Customer"
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
        "/customers/{id}": {
            "get": {
                "description": "API to Get Customer-Admin's Customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer Get"
                ],
                "summary": "Customer Get Route",
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
                            "$ref": "#/definitions/models.Customer"
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
                "description": "API to Update Customer-Admin's Customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer Update"
                ],
                "summary": "Customer Update Route",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Customer Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Customer Update",
                        "name": "customer",
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
                            "$ref": "#/definitions/models.Customer"
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
            "delete": {
                "description": "API to Delete Customer-Admin's Customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Delete"
                ],
                "summary": "Customer Delete Route",
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
                            "$ref": "#/definitions/models.Customer"
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
        "models.Address": {
            "type": "object",
            "properties": {
                "address_1": {
                    "type": "string"
                },
                "address_2": {
                    "type": "string"
                },
                "address_id": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "company": {
                    "type": "string"
                },
                "country_id": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "postcode": {
                    "type": "integer"
                }
            }
        },
        "models.Customer": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Address"
                    }
                },
                "approved": {
                    "type": "string"
                },
                "cart": {
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
                "fax": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "newsletter": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "reward_total": {
                    "type": "integer"
                },
                "safe": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "telephone": {
                    "type": "string"
                },
                "user_balance": {
                    "type": "number"
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
	Host:        "localhost:8080",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Customers API",
	Description: "This is customers CRUD service.",
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
