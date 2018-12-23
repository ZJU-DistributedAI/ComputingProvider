package transaction

import (
	"encoding/hex"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

//var (
//	file     = "/Users/liulifeng/Documents/privatechain/data0/keystore/UTC--2018-11-15T12-49-43.010863282Z--f448d0ae08287173002d06093abdab2ac1d7ce9a"
//	password = "123456"
//)

func getPrivateKeyAndAdress(file string, password string) (string, string) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		os.Exit(1)
	}

	keyjson, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	//解密keystore
	key, err := keystore.DecryptKey(keyjson, password)
	if err != nil {
		panic(err)
	}

	//获得以太坊地址
	address := key.Address.Hex()
	//获得以太坊私钥
	privateKey := hex.EncodeToString(crypto.FromECDSA(key.PrivateKey))

	return privateKey, address
}
