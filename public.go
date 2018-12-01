package main

import (
	"github.com/goadesign/goa"
)

// PublicController implements the Public resource.
type PublicController struct {
	*goa.Controller
}

// NewPublicController creates a Public controller.
func NewPublicController(service *goa.Service) *PublicController {
	return &PublicController{Controller: service.NewController("PublicController")}
}
