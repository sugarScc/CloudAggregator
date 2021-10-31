package router

import (
	"backend/controller"
	"github.com/gin-gonic/gin"
)

func V1Router(router *gin.Engine)  {
	v1 := router.Group("/v1")
	{
		v1.POST("/transaction",controller.PostOneTransaction)
		v1.GET("/transactions/user/:userAddress", controller.RetrieveTasksFromUser)
		//v1.GET("/unfinished/transactions")
		v1.PUT("/withdraw/transactionId/:transactionId/user/:userAddress", controller.WithdrawDeposit)
	}
}
