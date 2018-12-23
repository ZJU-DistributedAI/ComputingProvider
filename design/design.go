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

/*
********************************************************
(2)  Computing Provider Client
********************************************************
*/

var _ = Resource("ComputingProvider", func() {
	BasePath("/computing")

	Action("add", func() {
		Description("add computing resource")
		Routing(POST("/add/:hash/:ETH_key"))
		Params(func() {
			Param("hash", String, "data IPFS address")                  // 运算资源的ipfs地址
			Param("ETH_key", String, "ETH private key for transaction") // 以太坊交易秘钥，以后会隐藏
		})
		Response(OK, "plain/text")
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(NotImplemented, ErrorMedia)
	})

	Action("del", func() {
		Description("delete computing resource")
		Routing(POST("/del/:hash/:ETH_key"))
		Params(func() {
			Param("hash", String, "data IPFS address")                  // 运算资源的ipfs地址
			Param("ETH_key", String, "ETH private key for transaction") // 以太坊交易秘钥，以后会隐藏
		})
		Response(OK, "plain/text")
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(NotImplemented, ErrorMedia)
	})

	Action("agree", func() {
		Description("agree computing request for request[ID]")
		Routing(POST("/agree/:ETH_key/:computing_hash/:contract_hash/:public_key"))
		Params(func() {
			// 智能合约地址，被请求的运算资源地址，请求运算资源的客户端钱包地址可以成为运算资源请求的唯一标识
			Param("ETH_key", String, "ETH private key for transaction")   // 以太坊交易秘钥，以后会隐藏
			Param("computing_hash", String, "computing resourse hash")    // 被请求的数据的运算资源地址
			Param("contract_hash", String, "smart contract hash")         // 智能合约的地址
			Param("public_key", String, "ETH public key(Wallet address)") // 数据方客户端的公钥，即钱包地址
		})
		Response(OK, "plain/text")
		Response(InternalServerError, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(NotImplemented, ErrorMedia)
	})

	// TODO
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
		Response(NotImplemented, ErrorMedia)
	})
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET") // Allow all origins to retrieve the Swagger JSON (CORS)
	})
	Files("/swagger.json", "swagger/swagger.json")
})

var _ = Resource("swagger-ui", func() {

	Files("/swagger-ui/*filepath", "swagger-ui/")
})
