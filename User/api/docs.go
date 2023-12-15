package api

import (
	"github.com/gofrs/uuid"
	"github.com/swaggo/swag"
	"Projet_Middleware/User/internal/models" 
)

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
        "/users": {
            "get": {
                "description": "Obtenir des utilisateurs.",
                "tags": ["users"],
                "summary": "Obtenir des utilisateurs.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    },
                    "500": {
                        "description": "Quelque chose s'est mal passé"
                    }
                }
            },
            "post": {
                "description": "Créer un nouvel utilisateur",
                "tags": ["users"],
                "summary": "Créer un nouvel utilisateur",
                "parameters": [
                    {
                        "name": "user",
                        "in": "body",
                        "description": "L'objet utilisateur à créer",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "500": {
                        "description": "Quelque chose s'est mal passé"
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "Obtenir un utilisateur.",
                "tags": ["users"],
                "summary": "Obtenir un utilisateur.",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "ID de l'utilisateur au format UUID",
                        "required": true,
                        "type": "string"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "422": {
                        "description": "Impossible d'analyser l'ID"
                    },
                    "500": {
                        "description": "Quelque chose s'est mal passé"
                    }
                }
            },
            "put": {
                "description": "Mettre à jour un utilisateur spécifique par ID.",
                "tags": ["users"],
                "summary": "Mettre à jour un utilisateur spécifique par ID.",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "ID de l'utilisateur au format UUID",
                        "required": true,
                        "type": "string"
                    },
                    {
                        "name": "user",
                        "in": "body",
                        "description": "Données utilisateur mises à jour",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "422": {
                        "description": "Impossible d'analyser l'ID"
                    },
                    "500": {
                        "description": "Quelque chose s'est mal passé"
                    }
                }
            },
            "delete": {
                "description": "Supprimer un utilisateur spécifique par ID.",
                "tags": ["users"],
                "summary": "Supprimer un utilisateur spécifique par ID.",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "ID de l'utilisateur au format UUID",
                        "required": true,
                        "type": "string"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "422": {
                        "description": "Impossible d'analyser l'ID"
                    },
                    "500": {
                        "description": "Quelque chose s'est mal passé"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.User": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo détient les informations Swagger exportées pour que les clients puissent les modifier
var SwaggerInfo = &swag.Spec{
    Version:          "1.0.0",
    Host:             "",
    BasePath:         "/",
    Schemes:          []string{"http"},
    Title:            "Projet_Middleware/User",
    Description:      "API pour gérer des utilisateurs. ",
    InfoInstanceName: "swagger",
    SwaggerTemplate:  docTemplate,
    LeftDelim:        "{{",
    RightDelim:       "}}",
}

func init() {
    swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
