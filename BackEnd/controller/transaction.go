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

func PostOneTransaction(ctx *gin.Context) {
	//extract information from request from the ctx
	var commitRequest dto.CommitRequest
	bodyBytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(500, vo.Response{
			Code: 500,
			Message: err.Error(),
		})
		return
	}
	err = json.Unmarshal(bodyBytes, &commitRequest)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(500, vo.Response{
			Code: 500,
			Message: err.Error(),
		})
		return
	}
	result, err := service.CreateOneTransaction(commitRequest.UserAddress, commitRequest.DockerImage, commitRequest.Port, commitRequest.FlagMessage, commitRequest.Url, time.Now().Unix())
	if err != nil {
		log.Printf("error occur in createTransaction, %s",err.Error())
		ctx.JSON(500, vo.Response{
			Code: 500,
			Message: err.Error(),
		})
		return
	} else {
		ctx.JSON(201, vo.Response{
			Code: 201,
			Data: result,
		})
		return
	}
}

func RetrieveAllUnfinishedTasks(ctx *gin.Context) {
	service.RetrieveAllUnfinishedTasks()
}

func RetrieveTasksFromUser(ctx *gin.Context) {
	userAddress := ctx.Param("userAddress")
	tasks,err := service.RetrieveTasksFromUser(userAddress)

	if err != nil {
		log.Printf("error occur in retrieve task, %s",err.Error())
		ctx.JSON(500, vo.Response{
			Code: 500,
			Message: err.Error(),
		})
	} else {
		ctx.JSON(200, vo.Response{
			Code: 200,
			Data: tasks,
		})
	}
}

func WithdrawDeposit(ctx *gin.Context) {
	transactionId, _ := strconv.Atoi(ctx.Param("transactionId"))
	result, err := service.WithdrawDeposit(int64(transactionId), ctx.Param("userAddress"))
	if err != nil {
		log.Printf("error occur in withdraw, %s",err.Error())
		ctx.JSON(500, vo.Response{
			Code: 500,
			Message: err.Error(),
		})
	} else {
		ctx.JSON(201, vo.Response{
			Code: 201,
			Data: result,
		})
	}
}
