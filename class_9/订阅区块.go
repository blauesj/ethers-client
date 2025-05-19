package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	//  创建一个ethclient
	//  需要一个websocket链接
	client, err := ethclient.Dial("wss://eth-sepolia.g.alchemy.com/v2/tzMLGExhnP3XLwe1z8dJAOASCK3S-fW6")
	if err != nil {
		panic(err)
	}
	// 订阅区块头
	headers := make(chan *types.Header)
	//  创建订阅
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		panic(err)
	}
	for {
		// 管道处理
		select {
		case err := <-sub.Err():
			// 处理区块头
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			// 获取区块头对应的区块信息
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Hash().Hex())        // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f
			fmt.Println(block.Number().Uint64())   // 3477413
			fmt.Println(block.Time())              // 1529525947
			fmt.Println(block.Nonce())             // 130524141876765836
			fmt.Println(len(block.Transactions())) // 7

		}
	}

}
