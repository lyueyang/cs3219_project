package main

import (
	"fmt"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	corsConf := cors.DefaultConfig()
	corsConf.AddAllowMethods("PUT", "GET", "POST", "DELETE", "OPTIONS")
	corsConf.AllowAllOrigins = true

	r.Use(cors.New(corsConf))
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello! this is service bar")
	})

	r.GET("/bar", func(context *gin.Context) {
		context.String(http.StatusOK, "you accessed /bar in service bar")
	})

	err := r.Run(":6789")
	if err != nil {
		fmt.Println("error starting server")
	}
}
