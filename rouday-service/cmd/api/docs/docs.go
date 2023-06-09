// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/ping": {
            "get": {
                "description": "Simple ping method for testing purposes",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ping"
                ],
                "summary": "Ping",
                "responses": {
                    "200": {
                        "description": "pong",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/week-configurations": {
            "post": {
                "description": "Creates a new week configuration",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "week-configurations"
                ],
                "summary": "Create Week Configuration",
                "parameters": [
                    {
                        "description": "Create Week Configuration",
                        "name": "weekConfiguration",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.WeekConfiguration"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.WeekConfiguration"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Routine": {
            "type": "object",
            "properties": {
                "comment": {
                    "type": "string"
                },
                "daily_average_hours": {
                    "type": "number"
                },
                "friday": {
                    "type": "number"
                },
                "id": {
                    "type": "string"
                },
                "monday": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                },
                "saturday": {
                    "type": "number"
                },
                "sunday": {
                    "type": "number"
                },
                "thursday": {
                    "type": "number"
                },
                "tuesday": {
                    "type": "number"
                },
                "wednesday": {
                    "type": "number"
                }
            }
        },
        "models.WeekConfiguration": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "routines": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Routine"
                    }
                },
                "util_days_per_week": {
                    "type": "integer"
                },
                "week_start": {
                    "type": "string"
                },
                "work_hours_per_day": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
