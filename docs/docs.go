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
        "contact": {
            "name": "API Support (My GitHub)",
            "url": "https://github.com/saber-khakbiz",
            "email": "saberkhakbiz73@g"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/nids/{id}": {
            "get": {
                "description": "Get City Information by National Identification",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/html"
                ],
                "tags": [
                    "ID"
                ],
                "summary": "Get City name  Information by nID",
                "operationId": "get-city-by-id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "National ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Project Golang Maktabkhoneh API",
	Description:      "This is a web application that with the help of gin library and swagger document can verify the national number of people and also return the information of that person in response.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}