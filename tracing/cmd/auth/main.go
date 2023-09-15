package main

import (
	"github.com/gin-gonic/gin"
	"tracing/cmd/auth/internal"
)

func main() {
	r := gin.Default()

	r.POST("/genToken", internal.GenTokenHandler())
	r.POST("/verifyToken", internal.VerifyTokenHandler())
	r.POST("/deleteToken", internal.DeleteTokenHandler())

	r.Run(":8003")
}
