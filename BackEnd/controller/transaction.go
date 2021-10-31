package controller

import (
	"backend/dto"
	"backend/service"
	"backend/vo"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"strconv"
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
	service.CreateOneTransaction(commitRequest.UserAddress, commitRequest.DockerImage, commitRequest.Port, commitRequest.FlagMessage, commitRequest.Url, time.Now().Unix())
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
	transactionId,_ := strconv.Atoi(ctx.Param("transactionId"))
	service.WithdrawDeposit(int64(transactionId), ctx.Param("userAddress"))
	ctx.JSON(204, vo.Response{
		Code: 204,
	})
}