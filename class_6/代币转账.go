package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
	"log"
	"math/big"
)

func main() {

	// 创建一个 Ethereum 客户端
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/tzMLGExhnP3XLwe1z8dJAOASCK3S-fW6")
	if err != nil {
		log.Fatal(err)
	}
	// 代币发送的eth为0
	value := big.NewInt(0)

	//目标地址
	toAddress := common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4")
	//代币地址
	tokenAddress := common.HexToAddress("0xfadea654ea83c00e5003d2ea15c59830b65471c0")
	//from address
	privateKey, err := crypto.HexToECDSA("9af8075a23db01e091288cb0cbf4eb3b7140571a7398973e018605868cdaee39")
	// 获取公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	// 获取公钥对应的地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//transfer函数切片
	transferFnSignature := []byte("transfer(address,uint256)")
	//生成函数选择器
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID))

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)

	// 设置代币数量
	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10) // 1000 tokens  实际数字后多出18个0
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	//获取nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	gasLimit := uint64(21000)

	//  获取当前网络的 gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//  创建一个交易
	//tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	tx := types.NewTx(&types.LegacyTx{
		Nonce: nonce, GasPrice: gasPrice, Gas: gasLimit, To: &tokenAddress, Value: value, Data: data,
	})

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
