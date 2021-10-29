package service

import (
	store "backend/contract"
	"backend/env"
	"backend/utils"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
)

func WithdrawDeposit(transactionId int64)  {
	client, auth := utils.GetClientAndAuth(env.KovanInfura, env.PrivateKey)

	address := common.HexToAddress(env.ContractAddress)
	instance, err := store.NewCloudAggregatorTransactor(address, client)
	if err != nil {
		log.Fatal(err)
	}
	result,err := instance.ReturnMoneyBack(auth, big.NewInt(transactionId))
	if err != nil{
		log.Fatal(err)
	}
	log.Printf("withdraw transactionHash : %h",result.Hash())
}
