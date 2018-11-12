package main

import (
    "context"
    "log"

    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    client, _ := ethclient.Dial("https://example-rpc-server.com:port")
    defer client.Close()

    contractAddr := common.HexToAddress("0x_contract_address")
    callMsg := ethereum.CallMsg{
        To:   &contractAddr,
        Data: common.FromHex("0x8da5cb5b"),
    }

    res, err := client.CallContract(context.Background(), callMsg, nil)
    if err != nil {
        log.Fatalf("Error calling contract: %v", err)
    }
    log.Printf("Owner: %v", common.BytesToAddress(res).Hex())
}
