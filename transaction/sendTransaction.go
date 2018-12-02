package transaction
import (
	"fmt"
	"math/big"
	"context"
	"io/ioutil"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

)

var ETH_HOST = "http://localhost:8545"

func SendTransaction(
	KEYJSON_FILEDIR string,
	SIGN_PASSPHRASE string,
	KEYSTORE_DIR string,
	COINBASE_ADDR_HEX string,
	ALTER_ADDR_HEX string,
	amount *big.Int,
	gasLimit uint64,
	gasPirce *big.Int,
	data string) (string ,error) {

	// 初始化keystore
	ks := keystore.NewKeyStore(
		KEYSTORE_DIR,
		keystore.LightScryptN,
		keystore.LightScryptP)
	fmt.Println(ks)

	// 创建账户 COINBASE_ADDR_HEX
	fromAccDef := accounts.Account{
		Address: common.HexToAddress(COINBASE_ADDR_HEX),
	}

	toAccDef := accounts.Account{
		Address: common.HexToAddress(ALTER_ADDR_HEX),
	}

	// 查找将给定的帐户解析为密钥库中的唯一条目:找到签名的账户
	signAcc, err := ks.Find(fromAccDef)
	if err != nil {
		fmt.Println("account keystore find error:")
		panic(err)
	}
	fmt.Printf("account found: signAcc.addr=%s; signAcc.url=%s\n", signAcc.Address.String(), signAcc.URL)
	fmt.Println()

	// 解锁签名的账户
	errUnlock := ks.Unlock(signAcc, SIGN_PASSPHRASE)
	if errUnlock != nil {
		fmt.Println("account unlock error:")
		panic(err)
	}
	fmt.Printf("account unlocked: signAcc.addr=%s; signAcc.url=%s\n", signAcc.Address.String(), signAcc.URL)
	fmt.Println()

	// 建立交易
	tx := types.NewTransaction(
		0x0,
		toAccDef.Address,
		amount,
		gasLimit,
		gasPirce,
		[]byte(data))

	// 打开账户私钥文件
	keyJson, readErr := ioutil.ReadFile(KEYJSON_FILEDIR)
	if readErr != nil {
		fmt.Println("key json read error:")
		panic(readErr)
	}

	// 解析私钥文件
	keyWrapper, keyErr := keystore.DecryptKey(keyJson, SIGN_PASSPHRASE)
	if keyErr != nil {
		fmt.Println("key decrypt error:")
		panic(keyErr)
	}
	fmt.Printf("key extracted: addr=%s", keyWrapper.Address.String())

	// Define signer and chain id
	// chainID := big.NewInt(CHAIN_ID)
	// signer := types.NewEIP155Signer(chainID)
	signer := types.HomesteadSigner{}

	//用私钥签署交易签名
	signature, signatureErr := crypto.Sign(tx.Hash().Bytes(), keyWrapper.PrivateKey)
	if signatureErr != nil {
		fmt.Println("signature create error:")
		panic(signatureErr)
	}

	signedTx, signErr := tx.WithSignature(signer, signature)
	if signErr != nil {
		fmt.Println("signer with signature error:")
		panic(signErr)
	}

	//连接客户端
	client, err := ethclient.Dial(ETH_HOST)
	if err != nil {
		fmt.Println("client connection error:")
		panic(err)
	}
	fmt.Println("client connected")
	fmt.Println()

	//发送交易到网络
	txErr := client.SendTransaction(context.Background(), signedTx)
	if txErr != nil {
		fmt.Println("send tx error:")
		panic(txErr)
	}
	fmt.Printf("send success tx.hash=%s\n", signedTx.Hash().String())

	return signedTx.Hash().String(), nil
}