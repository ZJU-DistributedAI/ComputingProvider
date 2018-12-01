package main

import (
	"github.com/goadesign/goa"
)

// SwaggerUIDistController implements the swagger-ui-dist resource.
type SwaggerUIDistController struct {
	*goa.Controller
}

// NewSwaggerUIDistController creates a swagger-ui-dist controller.
func NewSwaggerUIDistController(service *goa.Service) *SwaggerUIDistController {
	return &SwaggerUIDistController{Controller: service.NewController("SwaggerUIDistController")}
}
