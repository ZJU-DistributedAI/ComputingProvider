package main

import (
	"bytes"
	"fmt"
	"github.com/ZJU-DistributedAI/ComputingProvider/app"
	"github.com/goadesign/goa"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"github.com/ZJU-DistributedAI/ComputingProvider/transaction"
	"math/big"
)
var IPFS_HOST = os.Getenv("IPFS_HOST")
// ComputingInfoController implements the ComputingInfo resource.
type ComputingInfoController struct {
	*goa.Controller
}

// NewComputingInfoController creates a ComputingInfo controller.
func NewComputingInfoController(service *goa.Service) *ComputingInfoController {
	return &ComputingInfoController{Controller: service.NewController("ComputingInfoController")}
}

// Cat runs the cat action.
func (c *ComputingInfoController) Cat(ctx *app.CatComputingInfoContext) error {
	// ComputingInfoController_Cat: start_implement
	var url = fmt.Sprintf("http://%s:5001/api/v0/cat?arg=%s", IPFS_HOST, ctx.Address)
	goa.LogInfo(ctx, fmt.Sprintf("Calling IPFS Cat API, at url: %s", url))
	resp, err := http.Get(url)
	if err != nil {
		return ctx.InternalServerError(
			goa.ErrInternal("Failed calling IPFS", "API", "cat", "Err", err.Error()))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ctx.InternalServerError(
			goa.ErrInternal("Failed reading IPFS response", "API", "cat", "Err", err.Error()))
	}
	return ctx.OK([]byte(body))
	// ComputingInfoController_Cat: end_implement
}

// Upload runs the upload action.
func (c *ComputingInfoController) Upload(ctx *app.UploadComputingInfoContext) error {
	// ComputingInfoController_Upload: start_implement
	file, err := ctx.Payload.File.Open()
	if err != nil {
		return ctx.BadRequest(
			goa.ErrBadRequest("Could not open file", "API", "upload", "Err", err.Error()))
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("", "")
	if err != nil {
		return ctx.InternalServerError(
			goa.ErrInternal("Could not create form file", "API", "upload", "Err", err.Error()))
	}
	_, err = io.Copy(part, file)
	err = writer.Close()
	if err != nil {
		return ctx.InternalServerError(
			goa.ErrInternal("Could not close form writter", "API", "upload", "Err", err.Error()))
	}

	url := fmt.Sprintf("http://%s:5001/api/v0/upload", IPFS_HOST)
	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	if err != nil {
		return ctx.InternalServerError(
			goa.ErrInternal("Error creating post request", "API", "upload", "Err", err.Error()))
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ctx.InternalServerError(
			goa.ErrInternal("Error posting request to IPFS", "API", "upload", "Err", err.Error()))
	}
	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ctx.InternalServerError(
			goa.ErrInternal("Error closing IPFS response body", "API", "upload", "Err", err.Error()))
	}

	// get hash to signature offline, then send raw transaction to ethereum
	//TODO(responseBody content)
	var hash = string(responseBody)

	if len(hash) != 46 {
		return ctx.BadRequest(
			goa.ErrBadRequest("Hash invalid!"))
	}

	// key文件数据； 账户密码； key文件数据路径； 账户； 带转入账户
	var KEYJSON_FILEDIR   = `./UTC--2018-11-15T13-17-11.427354517Z--9893e46b95e70035cf11c103d5ca425166b0532b`
	var SIGN_PASSPHRASE   = `123456`
	var KEYSTORE_DIR      = `.`
	var COINBASE_ADDR_HEX = `0x9893e46b95e70035cf11c103d5ca425166b0532b`
	var ALTER_ADDR_HEX    = `0xa0c34337a7b0ab1de7462899cb037d3588d1db92`
	amount := big.NewInt(0)
	gasPrice := big.NewInt(2000000000)
	gasLimit := uint64(2711301)
	data := "add " + hash

	returnHash, err := transaction.SendTransaction(
		KEYJSON_FILEDIR,
		SIGN_PASSPHRASE,
		KEYSTORE_DIR,
		COINBASE_ADDR_HEX,
		ALTER_ADDR_HEX,
		amount,
		gasLimit,
		gasPrice,
		data);

	if err != nil{
		return ctx.BadRequest(
			goa.ErrBadRequest("upload block failure"))
	}
	fmt.Printf(returnHash);

	return ctx.OK([]byte(responseBody))
	// ComputingInfoController_Upload: end_implement
}

func setTransactionInfo()  {
	
}
