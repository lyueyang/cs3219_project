package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello world :)")
	})

	err := r.Run(":6789")
	if err != nil {
		fmt.Println("error starting server")
	}
}
