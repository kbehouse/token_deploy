package main

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kbehouse/token_deploy/token"
)

var (
	cofigPath string
	name      string
	symbol    string
	supply    int64
)

func init() {

	flag.StringVar(&cofigPath, "config", "config.json", "Config Path")
	flag.StringVar(&name, "name", "", "Token Name")
	flag.StringVar(&symbol, "symbol", "", "Token Symbol")
	flag.Int64Var(&supply, "supply", 1000000000, "`Start` Block")

	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: [options] [root]\n")
	fmt.Fprintf(os.Stderr, "  Currently, there are four URI routes could be used:\n")
	flag.PrintDefaults()
}

func Magnitude(amountIn float64, decimal int64) *big.Int {
	// amountIn := 1.0
	wad := new(big.Int).Exp(big.NewInt(10), big.NewInt(decimal), nil)
	wadF := new(big.Float).SetInt(wad)
	amountInWadF := new(big.Float).Mul(big.NewFloat(amountIn), wadF)

	var amountInWad big.Int
	amountInWadF.Int(&amountInWad)
	return &amountInWad
}

func main() {
	flag.Parse()

	fmt.Println("Loading ", cofigPath)
	var config SwapConfig
	JSONFileRead(cofigPath, &config)

	client, err := ethclient.Dial(config.RPCHTTP)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(config.Key)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(1250000) // in units
	auth.GasPrice = gasPrice

	fmt.Println("Deploy from address: ", fromAddress)
	fmt.Printf("Deploy Token: %s (%s) with mint %v\n", name, symbol, supply)
	// input := "1.0"
	supplyBigInt := Magnitude(float64(supply), 18)
	address, tx, instance, err := token.DeployToken(auth, client, name, symbol, supplyBigInt)
	if err != nil {
		log.Fatalf("Deploy fail, err:%v", err)
	}

	fmt.Printf("Token Address: %s\n", address.Hex())            // 0xb9e266ACfD3dB616F241EFB160f36aE081800E2e
	fmt.Printf("TX: %s%s\n", config.ScanTxURL, tx.Hash().Hex()) // 0xf507628cee1e27f5e70e925c671fe4a5da404a0ed6908dca0a2d31c05eee70a8

	_ = instance
}
