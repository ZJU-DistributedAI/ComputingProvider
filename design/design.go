package design

// The goa API design language is a DSL implemented in Go and is not Go.
// The generated code or any of the actual Go code in goa does not make
// use of “dot imports”. Using this technique for the DSL results in far
// cleaner looking code.
//
// For more details refer to https://goa.design/design/overview/

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("computingProvider service APIs", func() {
	Title("ComputingProvider service APIs documentation")
	Description("This API includes a list of computingProvider utilities which can be used by any participants in our system")
	Host("localhost:8899")
	Scheme("http")
})

var _ = Resource("ComputingProvider", func() {
	BasePath("/computing")

	Action("add", func() {
		Description("add computing resource")
		Routing(POST("/add/:hash/:private_key"))
		Params(func() {
			Param("hash", String, "computing resource IPFS address")
			Param("private_key", String, "ETH private key for transaction")
		})
		Response(OK, "plain/text")
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})

	Action("del", func() {
		Description("delete computing resource")
		Routing(POST("/del/:hash/:private_key"))
		Params(func() {
			Param("hash", String, "computing resource IPFS address")
			Param("private_key", String, "ETH private key for transaction")
		})
		Response(OK, "plain/text")
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})

	Action("agree", func() {
		Description("agree computing request for request[ID]")
		Routing(POST("/agree/:hash/:ETH_key/:request_id"))
		Params(func() {
			Param("ETH_key", String, "ETH private key for transaction")
			Param("request_id", Integer, "request[ID]")
		})
		Response(OK, "plain/text")
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})

	Action("uploadRes", func() {
		Description("upload result hash for [request_id]")
		Routing(POST("/upload/:res_hash/:aes_hash/:ETH_key/:request_id"))
		Params(func() {
			Param("res_hash", String, "encrypted result hash")
			Param("aes_hash", String, "encrypted aes key hash")
			Param("ETH_key", String, "ETH private key for transaction")
			Param("request_id", Integer, " [request_id]")

		})
		Response(OK, "plain/text")
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET") // Allow all origins to retrieve the Swagger JSON (CORS)
	})
	Files("/swagger.json", "swagger/swagger.json")
})

var _ = Resource("swagger-ui-dist", func() {

	Files("/swagger-ui-dist/*filepath", "swagger-ui-dist/")
})
