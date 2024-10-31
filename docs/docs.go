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
        "/medfri-securiry/user/{id}": {
            "delete": {
                "description": "Este endpoint permite eliminar un usuario específico dado su ID.",
                "tags": [
                    "usuarios"
                ],
                "summary": "Eliminar un usuario por ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del usuario",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Usuario eliminado con éxito"
                    }
                }
            }
        },
        "/medfri-security/entity": {
            "put": {
                "description": "Este endpoint permite actualizar la información de una entidad existente.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "entidades"
                ],
                "summary": "Actualizar una entidad",
                "parameters": [
                    {
                        "description": "Información de la entidad actualizada",
                        "name": "entity",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Entity"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Entidad actualizada con éxito",
                        "schema": {
                            "$ref": "#/definitions/entity.Entity"
                        }
                    },
                    "400": {
                        "description": "Error en el cuerpo de la solicitud",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Este endpoint permite crear una nueva entidad en el sistema.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "entidades"
                ],
                "summary": "Crear una entidad",
                "parameters": [
                    {
                        "description": "Información de la entidad",
                        "name": "entity",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Entity"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Entidad creada con éxito",
                        "schema": {
                            "$ref": "#/definitions/entity.Entity"
                        }
                    },
                    "400": {
                        "description": "Error en el cuerpo de la solicitud",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/medfri-security/entity/{id}": {
            "get": {
                "description": "Este endpoint permite obtener la información de una entidad específica usando su ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "entidades"
                ],
                "summary": "Obtener una entidad por ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID de la entidad",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Entidad encontrada",
                        "schema": {
                            "$ref": "#/definitions/entity.Entity"
                        }
                    },
                    "404": {
                        "description": "Entidad no encontrada",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Este endpoint permite eliminar una entidad específica usando su ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "entidades"
                ],
                "summary": "Eliminar una entidad",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID de la entidad",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Entidad eliminada con éxito"
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/medfri-security/menu": {
            "put": {
                "description": "Este endpoint permite actualizar la información de un menu existente.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menus"
                ],
                "summary": "Actualizar un menu",
                "parameters": [
                    {
                        "description": "Información del menu actualizada",
                        "name": "resource",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Menu"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "menu actualizado con éxito",
                        "schema": {
                            "$ref": "#/definitions/entity.Menu"
                        }
                    },
                    "400": {
                        "description": "Error en el cuerpo de la solicitud",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Este endpoint permite crear un nuevo menu en el sistema.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menus"
                ],
                "summary": "Crear un menu",
                "parameters": [
                    {
                        "description": "Información del menu",
                        "name": "resource",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Menu"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Menu creado con éxito",
                        "schema": {
                            "$ref": "#/definitions/entity.Menu"
                        }
                    },
                    "400": {
                        "description": "Error en el cuerpo de la solicitud",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/medfri-security/menu/{id}": {
            "get": {
                "description": "Este endpoint devuelve la información de un menu específico dado su ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menus"
                ],
                "summary": "Obtener un menu por ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del menu",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "menu encontrado",
                        "schema": {
                            "$ref": "#/definitions/entity.Menu"
                        }
                    }
                }
            },
            "delete": {
                "description": "Este endpoint permite eliminar un menu específico usando su ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "menus"
                ],
                "summary": "Eliminar un menu",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del menu",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "menu eliminado con éxito"
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/medfri-security/resources": {
            "put": {
                "description": "Este endpoint permite actualizar la información de un recurso existente.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recursos"
                ],
                "summary": "Actualizar un recurso",
                "parameters": [
                    {
                        "description": "Información del recurso actualizada",
                        "name": "resource",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Resource"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Recurso actualizado con éxito",
                        "schema": {
                            "$ref": "#/definitions/entity.Resource"
                        }
                    },
                    "400": {
                        "description": "Error en el cuerpo de la solicitud",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Este endpoint permite crear un nuevo recurso en el sistema.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recursos"
                ],
                "summary": "Crear un recurso",
                "parameters": [
                    {
                        "description": "Información del recurso",
                        "name": "resource",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Resource"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Recurso creado con éxito",
                        "schema": {
                            "$ref": "#/definitions/entity.Resource"
                        }
                    },
                    "400": {
                        "description": "Error en el cuerpo de la solicitud",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/medfri-security/resources/{id}": {
            "get": {
                "description": "Este endpoint permite obtener la información de un recurso específico usando su ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recursos"
                ],
                "summary": "Obtener un recurso por ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del recurso",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Recurso encontrado",
                        "schema": {
                            "$ref": "#/definitions/entity.Resource"
                        }
                    },
                    "404": {
                        "description": "Recurso no encontrado",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "delete": {
                "description": "Este endpoint permite eliminar un recurso específico usando su ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recursos"
                ],
                "summary": "Eliminar un recurso",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del recurso",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Recurso eliminado con éxito"
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/medfri-security/rol": {
            "put": {
                "description": "Este endpoint permite actualizar la información de un Rol existente.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rols"
                ],
                "summary": "Actualizar un Rol",
                "parameters": [
                    {
                        "description": "Información del Rol actualizada",
                        "name": "resource",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Rol"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Rol actualizado con éxito",
                        "schema": {
                            "$ref": "#/definitions/entity.Rol"
                        }
                    },
                    "400": {
                        "description": "Error en el cuerpo de la solicitud",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Este endpoint permite crear un nuevo Rol en el sistema.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rols"
                ],
                "summary": "Crear un Rol",
                "parameters": [
                    {
                        "description": "Información del Rol",
                        "name": "resource",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Rol"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Rol creado con éxito",
                        "schema": {
                            "$ref": "#/definitions/entity.Rol"
                        }
                    },
                    "400": {
                        "description": "Error en el cuerpo de la solicitud",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/medfri-security/rol/{id}": {
            "get": {
                "description": "Este endpoint devuelve la información de un Rol específico dado su ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rols"
                ],
                "summary": "Obtener un Rol por ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del Rol",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Rol encontrado",
                        "schema": {
                            "$ref": "#/definitions/entity.Rol"
                        }
                    }
                }
            },
            "delete": {
                "description": "Este endpoint permite eliminar un Rol específico usando su ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "rols"
                ],
                "summary": "Eliminar un Rol",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del Rol",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Rol eliminado con éxito"
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/medfri-security/user": {
            "put": {
                "description": "Este endpoint permite actualizar la información de un usuario existente.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "usuarios"
                ],
                "summary": "Actualizar un usuario",
                "parameters": [
                    {
                        "description": "Información del usuario",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Usuario actualizado",
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    }
                }
            },
            "post": {
                "description": "Este endpoint permite crear un nuevo usuario en el sistema",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "usuarios"
                ],
                "parameters": [
                    {
                        "description": "Información del usuario",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    }
                }
            }
        },
        "/medfri-security/user-rol/asignar": {
            "post": {
                "description": "Este endpoint permite crear un nueva relacion usuarioRol en el sistema.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "re-usuario-rol"
                ],
                "summary": "Crear un recurso",
                "parameters": [
                    {
                        "description": "Información del recurso",
                        "name": "UserRol",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.UserRol"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Recurso creado con éxito",
                        "schema": {
                            "$ref": "#/definitions/entity.UserRol"
                        }
                    },
                    "400": {
                        "description": "Error en el cuerpo de la solicitud",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/medfri-security/user-rol/desasignar/{id}": {
            "delete": {
                "description": "Este endpoint permite eliminar una relacion usuario rol específico usando su ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "re-usuario-rol"
                ],
                "summary": "Eliminar una relacion usuario rol",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del usuario rol",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "usuario rol eliminado con éxito"
                    },
                    "500": {
                        "description": "Error interno del servidor",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/medfri-security/user/{id}": {
            "get": {
                "description": "Este endpoint devuelve la información de un usuario específico dado su ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "usuarios"
                ],
                "summary": "Obtener un usuario por ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del usuario",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Usuario encontrado",
                        "schema": {
                            "$ref": "#/definitions/entity.User"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.Entity": {
            "type": "object",
            "properties": {
                "activo": {
                    "type": "boolean"
                },
                "email": {
                    "type": "string"
                },
                "entidad_id": {
                    "type": "integer"
                },
                "nit": {
                    "type": "integer"
                },
                "nit_alterno": {
                    "type": "string"
                },
                "razon_social": {
                    "type": "string"
                }
            }
        },
        "entity.Menu": {
            "type": "object",
            "properties": {
                "descripcion": {
                    "type": "string"
                },
                "entidad": {
                    "$ref": "#/definitions/entity.Entity"
                },
                "entidad_id": {
                    "type": "integer"
                },
                "menuID": {
                    "type": "integer"
                },
                "menu_padre": {
                    "$ref": "#/definitions/entity.Menu"
                },
                "menu_padre_id": {
                    "type": "integer"
                },
                "nombre": {
                    "type": "string"
                },
                "recurso": {
                    "$ref": "#/definitions/entity.Resource"
                },
                "recurso_id": {
                    "type": "integer"
                }
            }
        },
        "entity.Resource": {
            "type": "object",
            "properties": {
                "acceso": {
                    "type": "string"
                },
                "descripcion": {
                    "type": "string"
                },
                "recurso_id": {
                    "type": "integer"
                }
            }
        },
        "entity.Rol": {
            "type": "object",
            "properties": {
                "activo": {
                    "type": "boolean"
                },
                "descripcion": {
                    "type": "string"
                },
                "entidad_id": {
                    "type": "integer"
                },
                "nombre": {
                    "type": "string"
                },
                "rolID": {
                    "type": "integer"
                }
            }
        },
        "entity.User": {
            "type": "object",
            "properties": {
                "activo": {
                    "type": "boolean"
                },
                "apellido_materno": {
                    "type": "string"
                },
                "apellido_paterno": {
                    "type": "string"
                },
                "cambiar_clave": {
                    "type": "boolean"
                },
                "clave": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "fecha_cambio_clave": {
                    "type": "string"
                },
                "fecha_creacion": {
                    "type": "string"
                },
                "fecha_retiro": {
                    "type": "string"
                },
                "nombre_1": {
                    "type": "string"
                },
                "nombre_2": {
                    "type": "string"
                },
                "tiempo_valides_token": {
                    "type": "string"
                },
                "usuario": {
                    "type": "integer"
                },
                "usuario_id": {
                    "type": "integer"
                }
            }
        },
        "entity.UserRol": {
            "type": "object",
            "properties": {
                "entidad_id": {
                    "type": "integer"
                },
                "fecha_fin": {
                    "type": "string"
                },
                "fecha_inicio": {
                    "type": "string"
                },
                "indefinido": {
                    "type": "string"
                },
                "rol_id": {
                    "type": "integer"
                },
                "usuarioRolID": {
                    "type": "integer"
                },
                "usuario_id": {
                    "type": "integer"
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
