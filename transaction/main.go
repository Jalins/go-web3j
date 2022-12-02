package transaction

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func queryBlockAndTrans(client *ethclient.Client) {
	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(block.Number().Uint64())

	for index, tx := range block.Transactions() {
		fmt.Println("=================================== 第 " + strconv.Itoa(index+1) + "笔交易 ===================================")
		fmt.Println("交易hash值为: " + tx.Hash().Hex())           // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
		fmt.Println("转账交易额度为: " + tx.Value().String())        // 10000000000000000
		fmt.Printf("话费gas为: %v\n", tx.Gas())                  // 105000
		fmt.Printf("gasPrice为: %v\n", tx.GasPrice().Uint64()) // 102000000000
		fmt.Printf("nonce值为: %v\n", tx.Nonce())               // 110644
		fmt.Println("data为: " + fmt.Sprintf("%x", tx.Data())) // []
		fmt.Println("接收方地址为: " + tx.To().Hex())               // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e

		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			fmt.Println(err)
		}

		if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID), nil); err == nil {
			fmt.Println("发送方地址为：" + msg.From().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
		}

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("调用合约地址为：%v\n", receipt.ContractAddress) // 1
	}
}

func Main() {
	client, err := ethclient.Dial("https://goerli.infura.io/v3/82a8978ba71f4a578f39ccb2b808e527")
	if err != nil {
		fmt.Println(err)
	}
	queryBlockAndTrans(client)

}
