package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

type SwapInterface struct {
	Factory common.Address `json:"factory"`
	Router  common.Address `json:"router"`
	Fee     int64          `json:"fee"`
}

type SwapConfig struct {
	Network   string         `json:"network"`
	RPCHTTP   string         `json:"rpc_http"`
	Address   common.Address `json:"address"`
	ScanTxURL string         `json:"scan_tx_url"`
	Key       string         `json:"privkey"`
}

// JSONFileRead read JSON to obj
func JSONFileRead(filePath string, obj interface{}) {
	jsonFile, err := os.Open(filePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("Successfully Opened ", filePath)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// read JSON
	err = json.Unmarshal(byteValue, obj)
	if err != nil {
		fmt.Println("Parse ", filePath, "fail, err:", err)
	}
}
