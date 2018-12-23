package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ZJU-DistributedAI/ComputingProvider/app"
	"github.com/ZJU-DistributedAI/ComputingProvider/transaction"
	"github.com/goadesign/goa"
)

var IPFS_API = os.Getenv("IPFS_API")

// ComputingProviderController implements the ComputingProvider resource.
type ComputingProviderController struct {
	*goa.Controller
}

// NewComputingProviderController creates a ComputingProvider controller.
func NewComputingProviderController(service *goa.Service) *ComputingProviderController {
	os.Setenv("IPFS_API", "http://47.52.231.230:8899")
	// set transaction argments
	err := setTransactionArgments()
	if err != nil {
		fmt.Println("setTransactionArgments err:", err)
	}
	return &ComputingProviderController{Controller: service.NewController("ComputingProviderController")}
}

// Add runs the add action.
func (c *ComputingProviderController) Add(ctx *app.AddComputingProviderContext) error {

	// check arguments
	if checkArguments(ctx.Hash, ctx.ETHKey) == false {
		fmt.Println("ctx.Hash===========>", ctx.Hash)
		return ctx.BadRequest(
			goa.ErrBadRequest("Invalid arguments!"))
	}

	// read config
	config := readConfig()
	if config == nil {
		fmt.Println("readConfig config===========>", config)
		goa.LogInfo(context.Background(), "Config of computing provider error")
		return ctx.InternalServerError(
			goa.ErrInternal("Config of computing provider error"))
	}

	// operate transaction
	transactionHash, err := transaction.OperateTransaction(transaction.DEL, ctx.Hash, ctx.ETHKey, config)
	if err != nil {
		fmt.Println("err===========>", err)
		return ctx.InternalServerError(
			goa.ErrInternal("operateTrasaction failure"))
	}
	return ctx.OK([]byte(transactionHash))
}

// Agree runs the agree action.
func (c *ComputingProviderController) Agree(ctx *app.AgreeComputingProviderContext) error {
	// ComputingProviderController_Agree: start_implement

	// Put your logic here

	return nil
	// ComputingProviderController_Agree: end_implement
}

// Del runs the del action.
func (c *ComputingProviderController) Del(ctx *app.DelComputingProviderContext) error {
	// check arguments
	if checkArguments(ctx.Hash, ctx.ETHKey) == false {
		fmt.Println("ctx.Hash===========>", ctx.Hash)
		return ctx.BadRequest(
			goa.ErrBadRequest("Invalid arguments!"))
	}

	// read config
	config := readConfig()
	if config == nil {
		fmt.Println("readConfig config===========>", config)
		goa.LogInfo(context.Background(), "Config of computing provider error")
		return ctx.InternalServerError(
			goa.ErrInternal("Config of computing provider error"))
	}

	// operate transaction
	// transactionHash, err := operateTrasaction(DEL, ctx.Hash, ctx.PrivateKey, config)
	transactionHash, err := transaction.OperateTransaction(transaction.DEL, ctx.Hash, ctx.ETHKey, config)
	if err != nil {
		fmt.Println("err===========>", err)
		return ctx.InternalServerError(
			goa.ErrInternal("operateTrasaction failure"))
	}
	return ctx.OK([]byte(transactionHash))
}

// UploadRes runs the uploadRes action.
func (c *ComputingProviderController) UploadRes(ctx *app.UploadResComputingProviderContext) error {
	// ComputingProviderController_UploadRes: start_implement

	// Put your logic here

	return nil
	// ComputingProviderController_UploadRes: end_implement
}

// check arguments
func checkArguments(hash string, privateKey string) bool {
	// easy check
	if len(hash) != 46 || len(privateKey) != 64 {
		return false
	}
	return true
}

// set transaction argments
func setTransactionArgments() error {
	err := os.Setenv("Add_to_address", "0af5013bb6f5c65d04abc69c9843697d708d3b5d")
	if err != nil {
		return err
	}
	err = os.Setenv("Add_data_prefix", "add ")
	if err != nil {
		return err
	}
	err = os.Setenv("Del_to_address", "7aa5414d58026ed3e3d3d87c97698c33e3f1602d")
	if err != nil {
		return err
	}
	err = os.Setenv("Del_data_prefix", "del ")
	if err != nil {
		return err
	}
	err = os.Setenv("ETH_HOST", "http://localhost:8545")
	if err != nil {
		return err
	}
	err = os.Setenv("Value", "0")
	if err != nil {
		return err
	}
	err = os.Setenv("Gas_price", "200")
	if err != nil {
		return err
	}
	err = os.Setenv("Gas_limit", "300000")
	if err != nil {
		return err
	}
	return nil
}

func readConfig() *transaction.TransactionConfig {
	// os.Setenv("TransactionConfig", )
	// config := &transaction.TransactionConfig{
	// 	Add_to_address:  "0af5013bb6f5c65d04abc69c9843697d708d3b5d",
	// 	Add_data_prefix: "add ",

	// 	Del_to_address:  "7aa5414d58026ed3e3d3d87c97698c33e3f1602d",
	// 	Del_data_prefix: "del ",

	// 	ETH_HOST:  "http://localhost:8545",
	// 	Value:     "0",
	// 	Gas_price: "200",
	// 	Gas_limit: "300000",
	// }

	config := &transaction.TransactionConfig{
		Add_to_address:  os.Getenv("Add_to_address"),
		Add_data_prefix: os.Getenv("Add_data_prefix"),

		Del_to_address:  os.Getenv("Del_to_address"),
		Del_data_prefix: os.Getenv("Del_data_prefix"),

		ETH_HOST:  os.Getenv("ETH_HOST"),
		Value:     os.Getenv("Value"),
		Gas_price: os.Getenv("Gas_price"),
		Gas_limit: os.Getenv("Gas_limit"),
	}

	return config
}
