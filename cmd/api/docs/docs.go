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
        "/cards": {
          "post": {
            "tags": [
              "cards"
            ],
            "summary": "Inserta la información de una nueva carta a la BD",
            "parameters": [
              {
                "in": "body",
                "name": "card data",
                "description": "Carta a crear",
                "required": true,
                "schema": {
                  "$ref": "#/definitions/BodyCardsPost"
                }
              }
            ],
            "responses": {
              "201": {
                "description": "(Created) La información de la carta se guardó correctamente",
                "schema": {
                  "$ref": "#/definitions/CreatedCardsPost"
                }
              },
              "400": {
                "$ref": "#/definitions/BadRequestBody"
              },
              "500": {
                "$ref": "#/definitions/CreatingInternalServerError"
              }
            }
          },
          "get": {
            "tags": [
              "cards"
            ],
            "summary": "Obtiene la información de todas las cartas guardadas en la base de datos",
            "responses": {
              "200": {
                "description": "(OK) La información de todas las cartas se obtuvo correctamente",
                "schema": {
                  "$ref": "#/definitions/AllCardsResponse"
                }
              },
              "500": {
                "$ref": "#/definitions/AllCardsInternalServerError"
              }
            }
          }
        },
        "/cards/{card_id}": {
          "get": {
            "tags": [
              "cards"
            ],
            "summary": "Obtiene la información de una carta especifica de la base de datos",
            "parameters": [
              {
                "name": "card_id",
                "in": "path",
                "description": "Identificador de la carta a obtener",
                "required": true,
                "type": "string"
              }
            ],
            "responses": {
              "200": {
                "description": "(OK) La información de la carta se obtuvo correctamente",
                "schema": {
                  "$ref": "#/definitions/CompleteCardsBody"
                }
              },
              "404": {
                "$ref": "#/definitions/NotFound"
              }
            }
          },
          "put": {
            "tags": [
              "cards"
            ],
            "summary": "Actualiza la información de una carta.",
            "parameters": [
              {
                "in": "body",
                "name": "card data",
                "description": "Carta a actualizar",
                "required": true,
                "schema": {
                  "$ref": "#/definitions/BodyCardsPost"
                }
              },
              {
                "name": "card_id",
                "in": "path",
                "description": "Identificador de la carta a obtener",
                "required": true,
                "type": "string"
              }
            ],
            "responses": {
              "202": {
                "description": "(Accepted) La información de la carta se guardó correctamente",
                "schema": {
                  "$ref": "#/definitions/UpdatedCardPut"
                }
              },
              "500": {
                "$ref": "#/definitions/CreatingInternalServerError"
              }
            }
          },
          "delete": {
            "tags": [
              "cards"
            ],
            "summary": "Elimina la información de una carta.",
            "parameters": [
              {
                "name": "card_id",
                "in": "path",
                "description": "Identificador de la carta a obtener",
                "required": true,
                "type": "string"
              }
            ],
            "responses": {
              "202": {
                "description": "(Accepted) La información de la carta se guardó correctamente",
                "schema": {
                  "$ref": "#/definitions/DeletedCard"
                }
              },
              "500": {
                "$ref": "#/definitions/CreatingInternalServerError"
              }
            }
          }
        }
      },
      "definitions": {
        "BodyCardsPost": {
          "type": "object",
          "properties": {
            "name": {
              "type": "string",
              "description": "Nombre de la carta"
            },
            "first_edition": {
              "type": "boolean",
              "description": "Indica si la carta es de la primera edición"
            },
            "serie_code": {
              "type": "string",
              "description": "Código de serie de la carta"
            },
            "subtype_id": {
              "type": "integer",
              "description": "Indica el subtipo de la carta (este luego se relaciona con el tipo internamente)"
            },
            "atk": {
              "type": "integer",
              "description": "Puntos de ataque (si corresponde)"
            },
            "def": {
              "type": "integer",
              "description": "Puntos de defensa (si corresponde)"
            },
            "stars": {
              "type": "integer",
              "description": "Cantidad de estrellas (si corresponde)"
            },
            "description": {
              "type": "string",
              "description": "Descripción de la carta"
            },
            "price": {
              "type": "integer",
              "description": "Precio de la carta"
            },
            "image_id": {
              "type": "integer",
              "description": "Indica la imagen de la carta"
            }
          }
        },
        "CreatedCardsPost": {
          "type": "object",
          "properties": {
            "message": {
              "type": "string",
              "description": "Mensaje que indica que se ha creado la carta correctamente."
            },
            "status": {
              "type": "integer",
              "description": "Código de respuesta"
            }
          }
        },
        "UpdatedCardPut": {
          "type": "object",
          "properties": {
            "message": {
              "type": "string",
              "description": "Mensaje que indica que se ha actualizado la carta correctamente."
            },
            "status": {
              "type": "integer",
              "description": "Código de respuesta"
            }
          }
        },
        "DeletedCard": {
          "type": "object",
          "properties": {
            "message": {
              "type": "string",
              "description": "Mensaje que indica que se ha eliminado la carta correctamente."
            },
            "status": {
              "type": "integer",
              "description": "Código de respuesta"
            }
          }
        },
        "AllCardsResponse": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/CompleteCardsBody"
          }
        },
        "CompleteCardsBody": {
          "type": "object",
          "properties": {
            "id": {
              "type": "integer",
              "description": "Id de la carta"
            },
            "name": {
              "type": "string",
              "description": "Nombre de la carta"
            },
            "first_edition": {
              "type": "boolean",
              "description": "Indica si la carta es de la primera edición"
            },
            "serie_code": {
              "type": "string",
              "description": "Código de serie de la carta"
            },
            "subtype_id": {
              "type": "integer",
              "description": "Indica el subtipo de la carta (este luego se relaciona con el tipo internamente)"
            },
            "subtype_description": {
              "type": "string",
              "description": "Indica el subtipo de la carta"
            },
            "type_id": {
              "type": "integer",
              "description": "Indica el tipo de la carta (esta relacionado con el subtipo)"
            },
            "type_description": {
              "type": "string",
              "description": "Indica el tipo de la carta"
            },
            "atk": {
              "type": "integer",
              "description": "Puntos de ataque (si corresponde)"
            },
            "def": {
              "type": "integer",
              "description": "Puntos de defensa (si corresponde)"
            },
            "stars": {
              "type": "integer",
              "description": "Cantidad de estrellas (si corresponde)"
            },
            "description": {
              "type": "string",
              "description": "Descripción de la carta"
            },
            "price": {
              "type": "integer",
              "description": "Precio de la carta"
            },
            "image_id": {
              "type": "integer",
              "description": "Indica la imagen de la carta"
            },
            "image_description": {
              "type": "string",
              "description": "Indica la imagen de la carta"
            }
          }
        },
        "BadRequestBody": {
          "description": "(BadRequest) Los datos enviados son incorrectos o hay datos que no corresponden con la estructura de la carta"
        },
        "CreatingInternalServerError": {
          "description": "(InternalServerError) Error al crear el registro en la base de datos."
        },
        "AllCardsInternalServerError": {
          "description": "(InternalServerError) Error obteniendo todas las cartas de la base de datos."
        },
        "NotFound": {
          "description": "(InternalServerError) Carta no encontrada en la base de datos"
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