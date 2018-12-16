package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"strconv"

	"github.com/ZJU-DistributedAI/ComputingProvider/app"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/goadesign/goa"
)

type TransactionConfig struct {
	// add info
	Add_to_address  string
	Add_data_prefix string

	// del info
	Del_to_address  string
	Del_data_prefix string

	// public info
	ETH_HOST  string
	Value     string
	Gas_price string
	Gas_limit string
}

type OpType string

const (
	ADD OpType = "add"
	DEL OpType = "del"
)

var IPFS_API = "http://47.52.231.230:8899"

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
	// check arguments
	if checkArguments(ctx.Hash, ctx.PrivateKey) == false {
		return ctx.BadRequest(
			goa.ErrBadRequest("Invalid arguments!"))
	}

	// read config
	config := readConfig()
	if config == nil {
		goa.LogInfo(context.Background(), "Config of computing provider error")
		return ctx.InternalServerError(
			goa.ErrInternal("Config of computing provider error"))
	}

	// operate transaction
	transactionHash, err := operateTrasaction(ADD, ctx.Hash, ctx.PrivateKey, config)
	if err != nil {
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
	if checkArguments(ctx.Hash, ctx.PrivateKey) == false {
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
	transactionHash, err := operateTrasaction(DEL, ctx.Hash, ctx.PrivateKey, config)
	if err != nil {
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

func readConfig() *TransactionConfig {

	// read file
	configJSON, err := ioutil.ReadFile("transaction_config.json")
	if err != nil {
		return nil
	}

	// parse json string
	config := &TransactionConfig{}
	err = json.Unmarshal([]byte(configJSON), &config)
	if err != nil {
		return nil
	}

	return config
}

// generate sign, sign, send transaction, return hash
func operateTrasaction(op OpType, hash string, privateKey string, config *TransactionConfig) (string, error) {
	// generate transaction
	tx, err := generateTransaction(op, hash, privateKey, config)
	if err != nil {
		fmt.Println("Generate transaction failed!")
		return "", err
	}

	// sign transaction
	signedTx, err := signTransaction(tx, privateKey)
	if err != nil {
		fmt.Println("Fail to sign transaction")
		return "", err
	}

	// send transaction
	transactionHash, err := sendTransaction(signedTx, config.ETH_HOST)
	if err != nil {
		fmt.Println("Fail to send transaction")
		return "", err
	}
	return transactionHash, nil
}

func generateTransaction(op OpType, hash string, privateKeyStr string, config *TransactionConfig) (*types.Transaction, error) {

	// get paraments of  transaction
	value, gasLimite, gasPrice, err := trans_type(config)
	if err != nil {
		return new(types.Transaction), err
	}

	// data
	to := config.Add_to_address
	data := config.Add_data_prefix + hash
	if op != "add" {
		to = config.Del_to_address
		data = config.Del_data_prefix + hash
	}
	fmt.Println(data)

	// get valid nonce
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return new(types.Transaction), err
	}
	client, err := ethclient.Dial(config.ETH_HOST)
	if err != nil {
		return new(types.Transaction), err
	}
	nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(privateKey.PublicKey))
	if err != nil {
		return new(types.Transaction), err
	}

	// a new Transaction
	tx := types.NewTransaction(
		nonce,
		common.HexToAddress(to),
		value,
		gasLimite,
		gasPrice,
		[]byte(data))

	return tx, nil
}

func trans_type(config *TransactionConfig) (*big.Int, uint64, *big.Int, error) {

	// trans value
	value, err := new(big.Int).SetString(config.Value, 10)
	if err == false {
		goa.LogInfo(context.Background(), "Trans value failed")
		return new(big.Int), uint64(0), new(big.Int), fmt.Errorf("Trans value failed")
	}

	// trans gasLimit
	gas_limit, err_gas := strconv.ParseInt(config.Gas_limit, 16, 64)
	if err_gas != nil {
		goa.LogInfo(context.Background(), "Trans value failed")
		return new(big.Int), uint64(0), new(big.Int), fmt.Errorf("Trans value failed")
	}
	gasLimit := uint64(gas_limit)

	// trans gasPrice
	gasPrice, err_price := new(big.Int).SetString(config.Gas_price, 10)
	if err_price == false {
		goa.LogInfo(context.Background(), "Trans gasPrice failed")
		return new(big.Int), uint64(0), new(big.Int), fmt.Errorf("Trans gasPrice failed")
	}

	return value, gasLimit, gasPrice, nil
}

func signTransaction(transaction *types.Transaction, private_key_str string) (*types.Transaction, error) {

	// get private key
	privity_key, err := crypto.HexToECDSA(private_key_str)
	if err != nil {
		return new(types.Transaction), err
	}

	// get auth for sign
	auth := bind.NewKeyedTransactor(privity_key)
	auth.Nonce = big.NewInt(int64(transaction.Nonce()))
	auth.Value = transaction.Value()
	auth.GasLimit = transaction.Gas()
	auth.GasPrice = transaction.GasPrice()
	auth.From = crypto.PubkeyToAddress(privity_key.PublicKey)

	//chainID := big.NewInt(int64(ChainID))
	//signer := types.NewEIP155Signer(chainID)

	// sign
	signer := types.HomesteadSigner{}
	signedTx, err := auth.Signer(signer, auth.From, transaction)
	return signedTx, err
}

func sendTransaction(signedTx *types.Transaction, ETH_HOST string) (string, error) {
	// get client
	client, err := ethclient.Dial(ETH_HOST)
	if err != nil {
		return "", err
	}

	// send
	txErr := client.SendTransaction(context.Background(), signedTx)
	if txErr != nil {
		return "", txErr
	}

	_, bind_err := bind.WaitMined(context.Background(), client, signedTx)
	if bind_err != nil {
		return "", bind_err
	}

	return signedTx.Hash().String(), nil
}
