//go:generate goagen bootstrap -d github.com/ZJU-DistributedAI/ComputingProvider/design

package main

import (
	"github.com/ZJU-DistributedAI/ComputingProvider/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("computingProvider service APIs")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "ComputingProvider" controller
	c := NewComputingProviderController(service)
	app.MountComputingProviderController(service, c)
	// Mount "swagger" controller
	c2 := NewSwaggerController(service)
	app.MountSwaggerController(service, c2)
	// Mount "swagger-ui-dist" controller
	c3 := NewSwaggerUIDistController(service)
	app.MountSwaggerUIDistController(service, c3)

	// Start service
	if err := service.ListenAndServe(":8899"); err != nil {
		service.LogError("startup", "err", err)
	}

}
