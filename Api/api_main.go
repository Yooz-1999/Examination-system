package main

import (
	"Api/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.LoadUserRouter(r)
	r.Run(":9999")
}
