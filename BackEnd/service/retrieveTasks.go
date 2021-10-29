package service

import (
	store "backend/contract"
	"backend/env"
	"backend/utils"
	"backend/vo"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

func RetrieveAllUnfinishedTasks() {
	client, _ := utils.GetClientAndAuth(env.KovanInfura, env.PrivateKey)

	address := common.HexToAddress(env.ContractAddress)
	instance, err := store.NewCloudAggregatorCaller(address, client)
	if err != nil {
		log.Fatal(err)
	}

	results, err := instance.RetrieveAllUnfinishedTask(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+vo", results)
}

func RetrieveTasksFromUser(userAddress string) (tasks vo.Tasks ) {
	client, _ := utils.GetClientAndAuth(env.KovanInfura, env.PrivateKey)

	address := common.HexToAddress(env.ContractAddress)
	instance, err := store.NewCloudAggregatorCaller(address, client)
	if err != nil {
		log.Fatal(err)
	}

	results, err := instance.RetrieveTasksInfoFromTargetUser(nil, common.HexToAddress(userAddress))
	if err != nil {
		log.Fatal(err)
	}

	length := len(results.CreationTimeStamps)

	for index := 0; index < length; index++ {
		task := vo.Task{}
		task.TransactionId = results.TransactionIds[index].Int64()
		task.State = Reflect[results.States[index]]
		task.DockerImage = results.DockerImages[index]
		task.Port = results.Ports[index]
		task.FlagMessage = common.Bytes2Hex(results.FlagMessages[index][:])
		task.Url = results.Urls[index]
		task.CreationTimeStamps = results.CreationTimeStamps[index].Int64()
		tasks = append(tasks, task)
	}

	return tasks
}
