package main

import (
	"context"
	"ethers-client/task_1/count"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {

	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/tzMLGExhnP3XLwe1z8dJAOASCK3S-fW6")
	if err != nil {
		panic(err)
	}
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	account := common.HexToAddress("0xcontractAdress")

	countContract, err := count.NewCount(account, client)

	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	tx, err := countContract.Increment(opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx hash:", tx.Hash().Hex())

	callOpt := &bind.CallOpts{Context: context.Background()}
	countResult, err := countContract.Count(callOpt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("count:", countResult)
}
