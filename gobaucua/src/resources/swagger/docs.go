// Package swagger Code generated by swaggo/swag. DO NOT EDIT
package swagger

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://github.com/BevisDev",
        "contact": {
            "name": "Truong Thanh Binh",
            "url": "https://github.com/BevisDev",
            "email": "dev.binhtt@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/db": {
            "get": {
                "description": "Check health DB System",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Check Health"
                ],
                "summary": "Ping DB API",
                "responses": {
                    "200": {
                        "description": "Successful",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "Check health system",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Check Health"
                ],
                "summary": "Ping System API",
                "responses": {
                    "200": {
                        "description": "Successful",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/redis": {
            "get": {
                "description": "Check health Redis System",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Check Health"
                ],
                "summary": "Ping Redis API",
                "responses": {
                    "200": {
                        "description": "Successful",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/signin": {
            "post": {
                "security": [
                    {
                        "AccessTokenAuth": []
                    }
                ],
                "description": "sign in web app",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Sign In API",
                "responses": {
                    "200": {
                        "description": "Successful",
                        "schema": {
                            "$ref": "#/definitions/response.Data"
                        }
                    },
                    "400": {
                        "description": "Client Error",
                        "schema": {
                            "$ref": "#/definitions/response.DataError"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.DataError"
                        }
                    }
                }
            }
        },
        "/signup": {
            "post": {
                "security": [
                    {
                        "AccessTokenAuth": []
                    }
                ],
                "description": "sign up web app",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Sign Up API",
                "responses": {
                    "200": {
                        "description": "Successful",
                        "schema": {
                            "$ref": "#/definitions/response.Data"
                        }
                    },
                    "400": {
                        "description": "Client Error",
                        "schema": {
                            "$ref": "#/definitions/response.DataError"
                        }
                    },
                    "500": {
                        "description": "Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.DataError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "response.Data": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 2000
                },
                "data": {},
                "is_success": {
                    "type": "boolean",
                    "example": true
                },
                "message": {
                    "type": "string",
                    "example": "Success"
                },
                "response_at": {
                    "type": "string",
                    "example": "2025-01-14 16:44:47.510"
                },
                "state": {
                    "type": "string",
                    "example": "8137ce10-305b-42f5-8f14-9c48dd6f23f0"
                }
            }
        },
        "response.DataError": {
            "type": "object",
            "properties": {
                "error": {
                    "$ref": "#/definitions/response.Error"
                },
                "is_success": {
                    "type": "boolean",
                    "example": false
                },
                "response_at": {
                    "type": "string",
                    "example": "2025-01-14 16:44:47.510"
                },
                "state": {
                    "type": "string",
                    "example": "8137ce10-305b-42f5-8f14-9c48dd6f23f0"
                }
            }
        },
        "response.Error": {
            "type": "object",
            "properties": {
                "error_code": {
                    "type": "integer",
                    "example": 3000
                },
                "message": {
                    "type": "string",
                    "example": "Invalid RequestLogger"
                }
            }
        }
    },
    "securityDefinitions": {
        "AccessTokenAuth": {
            "description": "Description for what is this security definition being used",
            "type": "apiKey",
            "name": "AccessToken",
            "in": "header"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8089",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "API Specification",
	Description:      "There are APIs in project",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
