package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"

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
	// TODO 上传  格式 2:运算资源描述hash:上传运算资源地址hash？
	transactionHash, err := transaction.OperateTransaction(transaction.ADD, ctx.Hash, ctx.ETHKey, config)
	if err != nil {
		fmt.Println("err===========>", err)
		return ctx.InternalServerError(
			goa.ErrInternal("operateTrasaction failure"))
	}
	// save privatekey
	os.Setenv("Agree_ETHKey", ctx.ETHKey)
	return ctx.OK([]byte(transactionHash))
}

// Agree runs the agree action.
func (c *ComputingProviderController) Agree(ctx *app.AgreeComputingProviderContext) error {
	// 获取判断swaggerUI上的参数参数；TODO同意并发送到以太坊 => 离线签名？
	if checkAgreeArgments(ctx.ETHKey, ctx.ComputingHash, ctx.ContractHash, ctx.PublicKey) == false {
		fmt.Println("ctx.Hash===========>", ctx)
		return ctx.BadRequest(
			goa.ErrBadRequest("Agree action Invalid arguments!"))
	}

	// read config
	config := readConfig()
	if config == nil {
		fmt.Println("readConfig config===========>", config)
		goa.LogInfo(context.Background(), "Config of computing provider error")
		return ctx.InternalServerError(
			goa.ErrInternal("Config of computing provider error"))
	}

	// computingHashAdress := os.Getenv("Del_to_address")
	// send2Ethereum TODO 内容的定义
	// agreeHash, err := send2Ethereum(transaction.AGREE, computingHashAdress)
	content := ctx.ComputingHash + ":" + ctx.ContractHash + ":" + ctx.PublicKey
	agreeHash, err := transaction.OperateTransaction(transaction.AGREE, content, ctx.ETHKey, config)
	if err != nil {
		fmt.Println("send2Ethereum err===========>", err)
		return ctx.BadRequest(
			goa.ErrBadRequest("Agree send2Ethereum error!"))
	}
	return ctx.OK([]byte(agreeHash))
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
	// check arguments
	if checkArguments(ctx.AesHash, ctx.ETHKey) == false {
		fmt.Println("ctx.Hash===========>", ctx.AesHash)
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
	content := ctx.AesHash + ":" + ctx.ResHash
	uploadHash, err := transaction.OperateTransaction(transaction.AGREE, content, ctx.ETHKey, config)
	if err != nil {
		fmt.Println("send2Ethereum err===========>", err)
		return ctx.BadRequest(
			goa.ErrBadRequest("upload result error!"))
	}
	return ctx.OK([]byte(uploadHash))
}

// Train runs the train action.
func (c *ComputingProviderController) Train(ctx *app.TrainComputingProviderContext) error {
	// TODO 检查参数，本地训练[完成代码链接，执行训练]
	if checkTrainArgments(ctx.ETHKey, ctx.ComputingHash, ctx.ContractHash, ctx.PublicKey) == false {
		fmt.Println("ctx.Hash===========>", ctx)
		return ctx.BadRequest(
			goa.ErrBadRequest("Train action Invalid arguments!"))
	}

	// training
	path := "_test.sh"
	if trainModelByPython(path) == false {
		fmt.Println("trainModelByPython error")
		return ctx.BadRequest(
			goa.ErrBadRequest("Train action trainModelByPython error!"))
	}
	return ctx.OK([]byte(os.Getenv("Del_to_address")))
}

// ----- Add and Del methods start -----
// check arguments
func checkArguments(hash string, privateKey string) bool {
	// easy check
	// if len(hash) != 46 || len(privateKey) != 64 {
	// 	return false
	// }

	return true
}

// set transaction argments  c 代表 computing provider； m -> model;  d->data
func setTransactionArgments() error {
	err := os.Setenv("Add_to_address", "0af5013bb6f5c65d04abc69c9843697d708d3b5d")
	if err != nil {
		return err
	}
	err = os.Setenv("Add_data_prefix", "cadd ")
	if err != nil {
		return err
	}
	err = os.Setenv("Del_to_address", "7aa5414d58026ed3e3d3d87c97698c33e3f1602d")
	if err != nil {
		return err
	}
	err = os.Setenv("Del_data_prefix", "cdel ")
	if err != nil {
		return err
	}
	err = os.Setenv("Agree_data_prefix", "cagree ")
	if err != nil {
		return err
	}
	err = os.Setenv("Upload_data_prefix", "cupload ")
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

// readConfig
func readConfig() *transaction.TransactionConfig {

	config := &transaction.TransactionConfig{
		Add_to_address:  os.Getenv("Add_to_address"),
		Add_data_prefix: os.Getenv("Add_data_prefix"),

		Del_to_address:  os.Getenv("Del_to_address"),
		Del_data_prefix: os.Getenv("Del_data_prefix"),

		Agree_data_prefix:  os.Getenv("Agree_data_prefix"),
		Upload_data_prefix: os.Getenv("Upload_data_prefix"),

		ETH_HOST:  os.Getenv("ETH_HOST"),
		Value:     os.Getenv("Value"),
		Gas_price: os.Getenv("Gas_price"),
		Gas_limit: os.Getenv("Gas_limit"),
	}

	return config
}

// ----- Agree methods start -----
func checkAgreeArgments(agreeETHKey string, agreeComputingHash string, agreeContractHash string, agreePublicKey string) bool {
	// if len(agreeComputingHash) != 46 || len(agreeContractHash) != 46 || len(agreeETHKey) != 64 || len(agreePublicKey) != 64 {
	// 	fmt.Println("argments is not valid")
	// 	return false
	// }

	// judge ETHKey
	AgreeETHKey := os.Getenv("Agree_ETHKey")
	if AgreeETHKey != agreeETHKey {
		fmt.Println("ETHKey is not matching")
		return false
	}
	return true
}

// TODO send2Ethereum send agree to ethereum
func send2Ethereum(op transaction.OpType, computingAddressHash string) (string, error) {
	//send hash to ethereum

	return "TODO", nil
}

// ----- Train methods start -----
func checkTrainArgments(trainETHKey string, trainComputingHash string, trainContractHash string, trainPublicKey string) bool {
	// if len(trainComputingHash) != 46 || len(trainContractHash) != 46 || len(trainETHKey) != 64 || len(trainPublicKey) != 64 {
	// 	fmt.Println("argments is not valid")
	// 	return false
	// }
	// judge ETHKey
	AgreeETHKey := os.Getenv("Agree_ETHKey")
	if AgreeETHKey != trainETHKey {
		fmt.Println("ETHKey is not matching")
		return false
	}
	return true
}

//TODO trainModelByPython
func trainModelByPython(commandFilePath string) bool {
	// command := `./_test.sh`
	command := "./" + commandFilePath + "`"
	cmd := exec.Command("/bin/bash", "-c", command)

	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		return false
	}
	fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
	return true
}
