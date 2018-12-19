package main

import (
	"ComputingProvider/app"
	"github.com/goadesign/goa"
)

// ComputingProviderController implements the ComputingProvider resource.
type ComputingProviderController struct {
	*goa.Controller
}

// NewComputingProviderController creates a ComputingProvider controller.
func NewComputingProviderController(service *goa.Service) *ComputingProviderController {
	return &ComputingProviderController{Controller: service.NewController("ComputingProviderController")}
}

// Add runs the add action.
func (c *ComputingProviderController) Add(ctx *app.AddComputingProviderContext) error {

	return ctx.NotImplemented(goa.ErrInternal("Not implemented"))
}

// Agree runs the agree action.
func (c *ComputingProviderController) Agree(ctx *app.AgreeComputingProviderContext) error {

	return ctx.NotImplemented(goa.ErrInternal("Not implemented"))
}

// Del runs the del action.
func (c *ComputingProviderController) Del(ctx *app.DelComputingProviderContext) error {

	return ctx.NotImplemented(goa.ErrInternal("Not implemented"))
}

// UploadRes runs the uploadRes action.
func (c *ComputingProviderController) UploadRes(ctx *app.UploadResComputingProviderContext) error {

	return ctx.NotImplemented(goa.ErrInternal("Not implemented"))
}
