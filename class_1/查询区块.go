package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {

	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/tzMLGExhnP3XLwe1z8dJAOASCK3S-fW6")

	if err != nil {
		log.Fatal(err)
	}
	blockNumber := big.NewInt(5671744)
	header, err := client.HeaderByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("header.Number:", header.Number.Uint64())
	fmt.Println("header.Time：", header.Time)
	fmt.Println("header.Difficulty：", header.Difficulty.Uint64())
	fmt.Println("header.Hash：", header.Hash().Hex())

	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("block.Number:", block.Number().Uint64())
	fmt.Println("block.Time：", block.Time())
	fmt.Println("block.Difficulty：", block.Difficulty().Uint64())
	fmt.Println("block.Hash：", block.Hash().Hex())
	fmt.Println("block.Transactions：", len(block.Transactions()))

}
