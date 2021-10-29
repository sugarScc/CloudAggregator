package controller

import (
	"backend/dto"
	"backend/service"
	"backend/vo"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"time"
)

func PostOneTransaction(ctx *gin.Context)  {
	//extract information from request from the ctx
	var commitRequest dto.CommitRequest
	bodyBytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil{
		log.Fatal(err)
	}
	err = json.Unmarshal(bodyBytes, &commitRequest)
	if err != nil{
		log.Fatal(err)
	}
	service.CreateOneTransaction(commitRequest.DockerImage, commitRequest.Port, commitRequest.FlagMessage, commitRequest.Url, time.Now().Unix())
	ctx.JSON(204, vo.Response{
		Code: 204,
	})
}

func RetrieveAllUnfinishedTasks(ctx *gin.Context)  {
	service.RetrieveAllUnfinishedTasks()
}

func RetrieveTasksFromUser(ctx *gin.Context) {
	userAddress := ctx.Param("userAddress")
	tasks := service.RetrieveTasksFromUser(userAddress)
	ctx.JSON(200, vo.Response{
		Code: 200,
		Data: tasks,
	})
}

func WithdrawDeposit(ctx *gin.Context)  {
	var withdrawRequest dto.WithdrawRequest
	bodyBytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil{
		log.Fatal(err)
	}
	err = json.Unmarshal(bodyBytes, &withdrawRequest)
	if err != nil{
		log.Fatal(err)
	}
	service.WithdrawDeposit(withdrawRequest.TransactionId)
	ctx.JSON(204, vo.Response{
		Code: 204,
	})
}