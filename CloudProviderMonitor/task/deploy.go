package task

import (
	CloudAggregator "CloudProviderMonitor/contract"
	"CloudProviderMonitor/dto"
	"CloudProviderMonitor/env"
	"CloudProviderMonitor/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
	"os/exec"
)

func Deploy(taskChan chan dto.Task) {
	log.Println("start deploy")
	for {
		select {
		case task := <-taskChan:
			log.Printf("retrieve task:%v\n", task)
			cmd := exec.Command("docker", "run", "-d", "-p", fmt.Sprintf("%s:%s", task.Port, task.Port), task.DockerImage)
			fmt.Printf("cmd command is : %s\n", cmd.String())
			err := cmd.Run()
			if err != nil {
				log.Printf("deploy failed:%s \n", err.Error())
			} else {
				log.Println("success deploy the image")
				commitTask(task.TransactionId)
			}
		}
	}
}

func commitTask(transactionId int64) {
	client, auth := utils.GetClientAndAuth()

	address := common.HexToAddress(env.ContractAddress)
	instance, err := CloudAggregator.NewCloudAggregatorTransactor(address, client)
	if err != nil {
		log.Println(err.Error())
	}
	result, err := exec.Command("curl", "ifconfig.me").Output()
	if err != nil {
		log.Fatal(err)
	}
	ip := string(result)
	auth.GasLimit = 2999999
	log.Println("the ip address is:", ip, "the transactionId is :", transactionId)
	transaction, err := instance.CommitTask(auth, ip, big.NewInt(transactionId))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("withdraw transactionHash : %h", transaction.Hash())
}
