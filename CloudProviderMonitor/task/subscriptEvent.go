package task

import (
	CloudAggregator "CloudProviderMonitor/contract"
	"CloudProviderMonitor/dto"
	"CloudProviderMonitor/env"
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
)

func SubscriptEvent(channel chan dto.Task) {
	log.Println("start Subscript")
	client, err := ethclient.Dial(env.KovanInfura)
	if err != nil {
		log.Println(err.Error())
		return
	}

	contractAddress := common.HexToAddress(env.ContractAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Println(err.Error())
		return
	}
	contractAbi, err := abi.JSON(strings.NewReader(CloudAggregator.CloudAggregatorABI))

	for {
		select {
		case err := <-sub.Err():
			log.Println(err.Error())
			continue
		case vLog := <-logs:
			result, err := contractAbi.Unpack("taskCommit", vLog.Data)
			if err != nil {
				log.Println(err.Error())
				continue
			}
			dockerImage := result[0].(string)
			port := result[1].(string)
			transactionId := result[2].(*big.Int)
			channel <- dto.Task{
				DockerImage:   dockerImage,
				Port:          port,
				TransactionId: transactionId.Int64(),
			}
		}
	}
}
