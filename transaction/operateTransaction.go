package transaction

import (
	"context"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/goadesign/goa"
)

// TransactionConfig all argments
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

// OpType operate Type
type OpType string

const (
	//ADD add
	ADD OpType = "add"
	//DEL del
	DEL OpType = "del"
)

// OperateTrasaction generate generate, sign, send transaction, return hash
func OperateTransaction(op OpType, hash string, privateKey string, config *TransactionConfig) (string, error) {
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
