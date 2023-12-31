// Package api Code generated by swaggo/swag. DO NOT EDIT
package api

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Ibrahima Coulibaly.",
            "email": "ibrahima.coulibaly@etu.uca.fr"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/songs": {
            "get": {
                "description": "Get songs.",
                "tags": ["songs"],
                "summary": "Get songs.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Song"
                            }
                        }
                    },
                    "500": {
                        "description": "Something went wrong"
                    }
                }
            },
            "post": {
                "description": "Create a new song",
                "tags": ["songs"],
                "summary": "Create a new song",
                "parameters": [
                    {
                        "name": "song",
                        "in": "body",
                        "description": "The song object to be created",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Song"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Song"
                        }
                    },
                    "500": {
                        "description": "Something went wrong"
                    }
                }
            }
        },
        "/songs/{id}": {
            "get": {
                "description": "Get a song.",
                "tags": ["songs"],
                "summary": "Get a song.",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "Song UUID formatted ID",
                        "required": true,
                        "type": "string"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Song"
                        }
                    },
                    "422": {
                        "description": "Cannot parse id"
                    },
                    "500": {
                        "description": "Something went wrong"
                    }
                }
            },
            "put": {
                "description": "Update a specific song by ID.",
                "tags": ["songs"],
                "summary": "Update a specific song by ID.",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "Song UUID formatted ID",
                        "required": true,
                        "type": "string"
                    },
                    {
                        "name": "song",
                        "in": "body",
                        "description": "Updated song data",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Song"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Song"
                        }
                    },
                    "422": {
                        "description": "Cannot parse id"
                    },
                    "500": {
                        "description": "Something went wrong"
                    }
                }
            },
            "delete": {
                "description": "Delete a specific song by ID.",
                "tags": ["songs"],
                "summary": "Delete a specific song by ID.",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "Song UUID formatted ID",
                        "required": true,
                        "type": "string"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "422": {
                        "description": "Cannot parse id"
                    },
                    "500": {
                        "description": "Something went wrong"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Song": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "artist": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
    Version:          "1.0.0",
    Host:             "",
    BasePath:         "/",
    Schemes:          []string{"http"},
    Title:            "Projet_Middleware/Song",
    Description:      "API to manage songs. ",
    InfoInstanceName: "swagger",
    SwaggerTemplate:  docTemplate,
    LeftDelim:        "{{",
    RightDelim:       "}}",
}

func init() {
    swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}