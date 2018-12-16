package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"mime/multipart"
	"net/http"
	"os"
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

var IPFS_HOST = os.Getenv("IPFS_HOST")

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
	// upload file start
	file, err := ctx.Payload.File.Open()
	fmt.Println("ctx file===========>", err)
	if err != nil {
		return ctx.BadRequest(
			goa.ErrBadRequest("Could not open file", "API", "add", "Err", err.Error()))
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	localFileName := ctx.Payload.File.Filename
	part, err := writer.CreateFormFile("uploadfile", localFileName)
	fmt.Println("ctx part===========>", err)
	if err != nil {
		return ctx.InternalServerError(
			goa.ErrInternal("Could not create form file", "API", "add", "Err", err.Error()))
	}
	_, err = io.Copy(part, file)
	err = writer.Close()
	fmt.Println("ctx err===========>", err)
	if err != nil {
		return ctx.InternalServerError(
			goa.ErrInternal("Could not close form writter", "API", "add", "Err", err.Error()))
	}

	// url := fmt.Sprintf("http://%s:5001/api/v0/add", IPFS_HOST)
	url := fmt.Sprintf("http://47.52.231.230:8899/storage")
	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	fmt.Println("ctx body===========>", body)
	if err != nil {
		return ctx.InternalServerError(
			goa.ErrInternal("Error creating post request", "API", "add", "Err", err.Error()))
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	fmt.Println("ctx resp===========>", resp)
	if err != nil {
		return ctx.InternalServerError(
			goa.ErrInternal("Error posting request to IPFS", "API", "add", "Err", err.Error()))
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	fmt.Println("ctx responseBody===========>", responseBody)
	if err != nil {
		return ctx.InternalServerError(
			goa.ErrInternal("Error closing IPFS response body", "API", "add", "Err", err.Error()))
	}
	// get hash to signature offline, then send raw transaction to ethereum
	// type ResponseStruct struct {
	// 	Name string `json:"name"`
	// 	Hash string `json:"card_balance"`
	// }
	// var responseStruct ResponseStruct
	// err = json.Unmarshal(responseBody, &responseStruct) //json = > struct
	fmt.Println("responseBody===========>", responseBody)
	// fmt.Println("ctx json.Unmarshal===========>", err)
	if err != nil {
		return ctx.BadRequest(
			goa.ErrBadRequest("json parse failure"))
	}
	// hash := responseStruct.Hash
	hash := string(responseBody[:])

	ctx.Hash = hash
	// return ctx.OK([]byte(responseBody))
	// uplaod file end
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

	// generate transaction
	tx, err := generateTransaction("add", ctx.Hash, ctx.PrivateKey, config)
	if err != nil {
		fmt.Println("generateTransaction tx===========>", tx)
		return ctx.InternalServerError(
			goa.ErrInternal("Generate transaction failed!"))
	}

	// sign transaction
	signedTx, err := signTransaction(tx, ctx.PrivateKey)
	if err != nil {
		fmt.Println("signTransaction signedTx===========>", signedTx)
		return ctx.InternalServerError(
			goa.ErrInternal("Fail to sign transaction"))
	}

	// send transaction
	transactionHash, err := sendTransaction(signedTx, config.ETH_HOST)
	if err != nil {
		fmt.Println("sendTransaction transactionHash===========>", transactionHash)
		return ctx.InternalServerError(
			goa.ErrInternal("Fail to send transaction"))
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
	// ComputingProviderController_Del: start_implement

	// Put your logic here

	return nil
	// ComputingProviderController_Del: end_implement
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

func generateTransaction(op string, hash string, privateKeyStr string, config *TransactionConfig) (*types.Transaction, error) {

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
