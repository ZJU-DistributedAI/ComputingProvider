package main

import (
	"github.com/ZJU-DistributedAI/ComputingProvider/app"
	"github.com/goadesign/goa"
)

// StorageController implements the Storage resource.
type StorageController struct {
	*goa.Controller
}

// NewStorageController creates a Storage controller.
func NewStorageController(service *goa.Service) *StorageController {
	return &StorageController{Controller: service.NewController("StorageController")}
}

// Add runs the add action.
func (c *StorageController) Add(ctx *app.AddStorageContext) error {
	// StorageController_Add: start_implement

	// Put your logic here

	return nil
	// StorageController_Add: end_implement
}

// Cat runs the cat action.
func (c *StorageController) Cat(ctx *app.CatStorageContext) error {
	// StorageController_Cat: start_implement

	// Put your logic here

	return nil
	// StorageController_Cat: end_implement
}
