package main

import (
	"backend/router"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	router.V1Router(r)
	_ = r.Run(":4128")
	return
}
