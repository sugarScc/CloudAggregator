package main

import (
	"backend/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.V1Router(r)
	_ = r.Run(":4128")
	return
}
