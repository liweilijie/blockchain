package main

import (
	"context"
	"fmt"
	"github.com/bnb-chain/go-sdk/client"
	"github.com/tendermint/tendermint/libs/pubsub/query"
	"github.com/tendermint/tendermint/types"
)

func main() {
	// wss://bsc.getblock.io/api_key/mainnet/
	// wss://bsc.getblock.io/api_key/testnet/
	client := client.NewHTTP("tcp://bsc.getblock.io/api_key/testnet/", "/websocket")
	err := client.Start()
	if err != nil {

		// handle error

	}
	defer client.Stop()
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	//query := query.MustParse("tm.event = 'Tx' AND tx.height = 3")
	query := query.MustParse("tm.event = 'Tx'")
	txs := make(chan interface{})
	err = client.Subscribe(ctx, "test-client", query, txs)

	go func() {

		for e := range txs {
			fmt.Println("got ", e.(types.EventDataTx))
		}

	}()

}
