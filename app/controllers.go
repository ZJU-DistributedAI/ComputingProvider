// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// unnamed API: Application Controllers
//
// Command:
// $ goagen
// --design=ComputingProvider/design
// --out=$(GOPATH)\src\ComputingProvider
// --version=v1.3.1

package app

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// ComputingProviderController is the controller interface for the ComputingProvider actions.
type ComputingProviderController interface {
	goa.Muxer
	Add(*AddComputingProviderContext) error
	Agree(*AgreeComputingProviderContext) error
	Del(*DelComputingProviderContext) error
	UploadRes(*UploadResComputingProviderContext) error
}

// MountComputingProviderController "mounts" a ComputingProvider resource controller on the given service.
func MountComputingProviderController(service *goa.Service, ctrl ComputingProviderController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewAddComputingProviderContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Add(rctx)
	}
	service.Mux.Handle("POST", "/computing/add/:hash/:private_key", ctrl.MuxHandler("add", h, nil))
	service.LogInfo("mount", "ctrl", "ComputingProvider", "action", "Add", "route", "POST /computing/add/:hash/:private_key")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewAgreeComputingProviderContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Agree(rctx)
	}
	service.Mux.Handle("POST", "/computing/agree/:hash/:ETH_key/:request_id", ctrl.MuxHandler("agree", h, nil))
	service.LogInfo("mount", "ctrl", "ComputingProvider", "action", "Agree", "route", "POST /computing/agree/:hash/:ETH_key/:request_id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDelComputingProviderContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Del(rctx)
	}
	service.Mux.Handle("POST", "/computing/del/:hash/:private_key", ctrl.MuxHandler("del", h, nil))
	service.LogInfo("mount", "ctrl", "ComputingProvider", "action", "Del", "route", "POST /computing/del/:hash/:private_key")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUploadResComputingProviderContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.UploadRes(rctx)
	}
	service.Mux.Handle("POST", "/computing/upload/:res_hash/:aes_hash/:ETH_key/:request_id", ctrl.MuxHandler("uploadRes", h, nil))
	service.LogInfo("mount", "ctrl", "ComputingProvider", "action", "UploadRes", "route", "POST /computing/upload/:res_hash/:aes_hash/:ETH_key/:request_id")
}

// SwaggerController is the controller interface for the Swagger actions.
type SwaggerController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl SwaggerController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/swagger.json", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))

	h = ctrl.FileHandler("/swagger.json", "swagger/swagger.json")
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/swagger.json", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "files", "swagger/swagger.json", "route", "GET /swagger.json")
}

// handleSwaggerOrigin applies the CORS response headers corresponding to the origin.
func handleSwaggerOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// SwaggerUIDistController is the controller interface for the SwaggerUIDist actions.
type SwaggerUIDistController interface {
	goa.Muxer
	goa.FileServer
}

// MountSwaggerUIDistController "mounts" a SwaggerUIDist resource controller on the given service.
func MountSwaggerUIDistController(service *goa.Service, ctrl SwaggerUIDistController) {
	initService(service)
	var h goa.Handler

	h = ctrl.FileHandler("/swagger-ui-dist/*filepath", "swagger-ui-dist/")
	service.Mux.Handle("GET", "/swagger-ui-dist/*filepath", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "SwaggerUIDist", "files", "swagger-ui-dist/", "route", "GET /swagger-ui-dist/*filepath")

	h = ctrl.FileHandler("/swagger-ui-dist/", "swagger-ui-dist\\index.html")
	service.Mux.Handle("GET", "/swagger-ui-dist/", ctrl.MuxHandler("serve", h, nil))
	service.LogInfo("mount", "ctrl", "SwaggerUIDist", "files", "swagger-ui-dist\\index.html", "route", "GET /swagger-ui-dist/")
}
