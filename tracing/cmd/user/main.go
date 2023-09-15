package main

import (
	"github.com/gin-gonic/gin"
	"tracing/cmd/user/internal"
)

func main() {
	r := gin.Default()

	r.POST("/login", internal.LoginHandler())
	r.POST("/logout", internal.LogoutHandler())

	r.Run(":8004")
}
