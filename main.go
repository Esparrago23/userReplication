package main

import (
	"demo/src/infrastructure"
	"github.com/gin-gonic/gin"
)
	

func main() {
	r := gin.Default()
	infrastructure.Init(r)
	r.Run(":8081")
	
}