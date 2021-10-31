package service

import (
	store "backend/contract"
	"backend/env"
	"backend/utils"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stratumn/go-core/types"
	"log"
	"math/big"
)

func CreateOneTransaction(userAddress string, dockerImage string, port string, flagMessage string, url string, creationTimestamp int64) {
	client, auth := utils.GetClientAndAuth()

	address := common.HexToAddress(env.ContractAddress)
	instance, err := store.NewCloudAggregatorTransactor(address, client)
	if err != nil {
		log.Fatal(err)
	}
	bytes32, err := hex.DecodeString(fmt.Sprintf("%x", flagMessage))
	if err != nil {
		fmt.Println(err.Error())
	}
	val := types.NewBytes32FromBytes(bytes32)
	if err != nil {
		fmt.Println(err.Error())
	}
	if err != nil {
		log.Fatal(err)
	}
	userAddressBytes := common.HexToAddress(userAddress)
	//TODO Check if available
	auth.From = userAddressBytes
	// transfer this 0.1 ETH as the Task Fee
	auth.Value = big.NewInt(10000000000000000)

	result, err := instance.PublishTask(auth, dockerImage, port, *val, url, big.NewInt(creationTimestamp))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("CreateOneTransaction finished, transaction hash: %h", result.Hash())
}
