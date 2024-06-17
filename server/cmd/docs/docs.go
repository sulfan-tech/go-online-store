// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Sulfan Aidid",
            "email": "sulfantech@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {}
}`

var SwaggerInfo = &swag.Spec{
    Version:          "1.0.0",
    Host:             "localhost:1313",
    BasePath:         "/v1",
    Schemes:          []string{"http", "https"},
    Title:            "Go Online Store",
    Description:      "CRUD online store using golang",
    InfoInstanceName: "swagger",
    SwaggerTemplate:  docTemplate,
    LeftDelim:        "{{",
    RightDelim:       "}}",
}

func init() {
    swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}