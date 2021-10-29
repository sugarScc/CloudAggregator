package service

import (
	store "backend/contract"
	"backend/env"
	"backend/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stratumn/go-core/types"
	"log"
	"math/big"
)

func CreateOneTransaction(dockerImage string, port string, flagMessage string, url string, creationTimestamp int64) {
	client, auth := utils.GetClientAndAuth(env.KovanInfura, env.PrivateKey)

	address := common.HexToAddress(env.ContractAddress)
	instance, err := store.NewCloudAggregatorTransactor(address, client)
	if err != nil {
		log.Fatal(err)
	}
	// transfer the params into using structure
	flagMessageBytes32, err := types.NewBytes32FromString(flagMessage)
	if err != nil {
		log.Fatal(err)
	}

	result, err := instance.PublishTask(auth, dockerImage, port, *flagMessageBytes32, url, big.NewInt(creationTimestamp))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("CreateOneTransaction finished, transaction hash: %h",result.Hash())
}
