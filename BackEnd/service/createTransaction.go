package service

import (
	store "backend/contract"
	"backend/env"
	"backend/utils"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stratumn/go-core/types"
	"math/big"
)

func CreateOneTransaction(userAddress string, dockerImage string, port string, flagMessage string, url string, creationTimestamp int64) (interface{}, error) {
	client, auth := utils.GetClientAndAuth()

	address := common.HexToAddress(env.ContractAddress)
	instance, err := store.NewCloudAggregatorTransactor(address, client)
	if err != nil {
		return nil,err
	}
	bytes32, err := hex.DecodeString(fmt.Sprintf("%x", flagMessage))
	if err != nil {
		return nil,err
	}
	val := types.NewBytes32FromBytes(bytes32)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	userAddressBytes := common.HexToAddress(userAddress)
	//TODO Check if available
	auth.From = userAddressBytes
	// transfer this 0.1 ETH as the Task Fee
	auth.Value = big.NewInt(10000000000000000)

	result, err := instance.PublishTask(auth, dockerImage, port, *val, url, big.NewInt(creationTimestamp))
	if err != nil {
		return nil, err
	}
	fmt.Printf("CreateOneTransaction finished, transaction hash: %h", result.Hash())
	return result.Hash().String(), nil
}
