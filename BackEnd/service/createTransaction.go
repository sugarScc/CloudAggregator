package service

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func CreateOneTransaction() {
	client, err := ethclient.Dial("https://kovan.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")

	_ = client // we'll use this in the upcoming sections
}