// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "A simple Rest-Api to submit Fibonacci Job through Stack.",
    "title": "3 Simple Job",
    "version": "0.1.0"
  },
  "paths": {
    "/": {
      "get": {
        "tags": [
          "jobs"
        ],
        "operationId": "findJobs",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "page",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int64",
            "default": 20,
            "name": "pagesize",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "list all created jobs",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/job"
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "tags": [
          "jobs"
        ],
        "operationId": "addOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/job"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/job"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/{id}": {
      "put": {
        "tags": [
          "jobs"
        ],
        "operationId": "updateOne",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/job"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/job"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "jobs"
        ],
        "operationId": "destroyOne",
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "format": "int64",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "job": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "id": {
          "description": "Job Identifier",
          "type": "integer",
          "readOnly": true
        },
        "name": {
          "description": "Job Description",
          "type": "string"
        },
        "quantifier": {
          "description": "Job Quantifier",
          "type": "integer"
        },
        "status": {
          "description": "Job Status",
          "type": "string",
          "enum": [
            "Completed",
            "Executing",
            "Validated",
            "Started",
            "Queued",
            "Pending"
          ]
        }
      }
    }
  }
}`))
}