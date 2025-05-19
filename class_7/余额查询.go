package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math"
	"math/big"
)

func main() {

	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/tzMLGExhnP3XLwe1z8dJAOASCK3S-fW6")
	if err != nil {
		panic(err)
	}
	account := common.HexToAddress("0x25836239F7b632635F815689389C537133248edb")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(balance)

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue) // 25.729324269165216041
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println(pendingBalance)

}
