package account

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		fmt.Println(err.Error())
	}

	// TODO 查询余额
	addr := common.HexToAddress("0x8aBfB96fe5b3E107e6757886b8A0f0a0AfbADEC2")
	balance, _ := client.BalanceAt(context.Background(), addr, nil)
	fmt.Println(balance)

	// TODO 查询在途余额
	pendingBalance, _ := client.PendingBalanceAt(context.Background(), addr)
	fmt.Println(pendingBalance)

	// TODO 生成新钱包
	// 1.生成私钥
	privKey, _ := crypto.GenerateKey()
	privKeyBytes := crypto.FromECDSA(privKey)

	newPrivKey := hexutil.Encode(privKeyBytes)[2:]
	fmt.Println(newPrivKey)

	// 2.由私钥产生公钥
	publicKey := privKey.Public()

	// 3.判断公钥是否为ecdsa类型
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if ok {
		fmt.Println("ecdsa公钥")
	}

	// 4.公钥生成地址
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[2:])

	new_addr := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(new_addr)

	// TODO 账户生成与保存
	importKs()
	createKs()

	// TODO 地址查询
	address := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
	bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		fmt.Println(err)
	}

	isContract := len(bytecode) > 0

	fmt.Printf("is contract: %v\n", isContract)
}

func createKs() {
	ks := keystore.NewKeyStore("./account/tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(account.Address.Hex())
}

func importKs() {
	file := "./tmp/UTC--2022-12-01T07-42-17.161374000Z--5b3c8abe4ffeb46cfbc075ab4ac76ccfcba6d554"
	ks := keystore.NewKeyStore("./account/tmp1", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	password := "secret"
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		fmt.Println("+++++", err)
	}

	fmt.Println(account.Address.Hex())

	if err := os.Remove(file); err != nil {
		fmt.Println(err)
	}
}
