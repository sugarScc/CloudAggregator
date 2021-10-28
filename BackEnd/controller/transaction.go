package controller

import (
	"backend/service"
	"github.com/gin-gonic/gin"
)

func PostOneTransaction(ctx *gin.Context)  {
	//TODO extract information from request from the ctx
	service.CreateOneTransaction()
}