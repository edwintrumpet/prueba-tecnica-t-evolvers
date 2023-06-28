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
        "/customers": {
            "get": {
                "description": "List customers with is_active in true",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customers"
                ],
                "summary": "List active customers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Customer"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customers"
                ],
                "summary": "Create",
                "parameters": [
                    {
                        "description": "initial data to create a customer",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateCustomer"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Customer"
                        }
                    }
                }
            }
        },
        "/workorders": {
            "get": {
                "description": "List all work orders with filters by status and date",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Work orders"
                ],
                "summary": "List all work orders",
                "parameters": [
                    {
                        "type": "string",
                        "description": "iso date",
                        "name": "until",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "iso date",
                        "name": "since",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "new, cancelled, done",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.WorkOrder"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new work order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Work orders"
                ],
                "summary": "Create",
                "parameters": [
                    {
                        "description": "initial data to create a work order",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateWorkOrder"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.WorkOrder"
                        }
                    }
                }
            }
        },
        "/workorders/customer/{customerId}": {
            "get": {
                "description": "List all work orders by a customer",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Work orders"
                ],
                "summary": "List work orders by customer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "uuid",
                        "name": "customerId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.WorkOrder"
                            }
                        }
                    }
                }
            }
        },
        "/workorders/finish": {
            "post": {
                "description": "Finishes a work order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Work orders"
                ],
                "summary": "Finish",
                "parameters": [
                    {
                        "description": "ids to find the work order to finish",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.FinishWorkOrder"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.WorkOrder"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CreateCustomer": {
            "type": "object",
            "required": [
                "address",
                "firstName",
                "lastName"
            ],
            "properties": {
                "address": {
                    "type": "string",
                    "example": "742 Evergreen Terrace"
                },
                "firstName": {
                    "type": "string",
                    "example": "Homer"
                },
                "lastName": {
                    "type": "string",
                    "example": "Simpson"
                }
            }
        },
        "models.CreateWorkOrder": {
            "type": "object",
            "required": [
                "customerId",
                "plannedTimeBegin",
                "plannedTimeEnd",
                "title"
            ],
            "properties": {
                "customerId": {
                    "type": "string",
                    "format": "uuid",
                    "example": "256c1214-3385-4235-9cfe-1dc85a5f2a46"
                },
                "plannedTimeBegin": {
                    "type": "string",
                    "example": "2023-06-27T17:45:00.408032Z"
                },
                "plannedTimeEnd": {
                    "type": "string",
                    "example": "2023-06-27T17:45:00.408032Z"
                },
                "title": {
                    "type": "string",
                    "example": "something"
                }
            }
        },
        "models.Customer": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "example": "742 Evergreen Terrace"
                },
                "createdAt": {
                    "type": "string",
                    "format": "iso",
                    "example": "2023-06-26T17:45:00.408032Z"
                },
                "endDate": {
                    "type": "string",
                    "format": "iso",
                    "example": "0001-01-01T00:00:00Z"
                },
                "firstName": {
                    "type": "string",
                    "example": "Homer"
                },
                "id": {
                    "type": "string",
                    "format": "uuid",
                    "example": "256c1214-3385-4235-9cfe-1dc85a5f2a46"
                },
                "isActive": {
                    "type": "boolean",
                    "example": false
                },
                "lastName": {
                    "type": "string",
                    "example": "Simpson"
                },
                "startDate": {
                    "type": "string",
                    "format": "iso",
                    "example": "0001-01-01T00:00:00Z"
                }
            }
        },
        "models.FinishWorkOrder": {
            "type": "object",
            "required": [
                "customerId",
                "workOrderId"
            ],
            "properties": {
                "customerId": {
                    "type": "string",
                    "format": "uuid",
                    "example": "256c1214-3385-4235-9cfe-1dc85a5f2a46"
                },
                "workOrderId": {
                    "type": "string",
                    "format": "uuid",
                    "example": "cd9bde09-5374-4749-86a6-34866c100e6e"
                }
            }
        },
        "models.WorkOrder": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string",
                    "format": "iso",
                    "example": "2023-06-26T17:45:00.408032Z"
                },
                "customer": {
                    "$ref": "#/definitions/models.Customer"
                },
                "customerId": {
                    "type": "string",
                    "format": "uuid",
                    "example": "256c1214-3385-4235-9cfe-1dc85a5f2a46"
                },
                "id": {
                    "type": "string",
                    "format": "uuid",
                    "example": "cd9bde09-5374-4749-86a6-34866c100e6e"
                },
                "plannedTimeBegin": {
                    "type": "string",
                    "example": "2023-06-27T17:45:00.408032Z"
                },
                "plannedTimeEnd": {
                    "type": "string",
                    "example": "2023-06-27T17:45:00.408032Z"
                },
                "status": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.workOrderStatus"
                        }
                    ],
                    "example": "ok"
                },
                "title": {
                    "type": "string",
                    "example": "something"
                }
            }
        },
        "models.workOrderStatus": {
            "type": "string",
            "enum": [
                "new"
            ],
            "x-enum-varnames": [
                "New"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "http://35.175.235.82",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "enerBit API",
	Description:      "API to manage customers and work orders",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}