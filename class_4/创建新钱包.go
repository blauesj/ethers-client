package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"log"
)

// 首先生成一个新的钱包，我们需要导入 go-ethereum crypto 包，该包提供用于生成随机私钥的 GenerateKey 方法。
// 如果已经有了私钥的 Hex 字符串，也可以使用 HexToECDSA 方法恢复私钥：
func main() {

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}

	privateKeyHex := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyHex)[2:])

	//从私钥生成公钥
	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey) //断言成为ecdsa.PublicKey指针
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:])

	// 它接受一个 ECDSA 公钥，并返回公共地址。
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)

	//公共地址其实就是公钥的 Keccak-256 哈希，然后我们取最后 40 个字符（20 个字节）并用“0x”作为前缀。
	//以下是使用 golang.org/x/crypto/sha3 的 Keccak256 函数手动完成的方法。
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))
}
