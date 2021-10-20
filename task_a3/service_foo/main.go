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
		context.String(http.StatusOK, "hello! this is service foo")
	})

	r.GET("/foo", func(context *gin.Context) {
		context.String(http.StatusOK, "you accessed /foo in service foo")
	})

	err := r.Run(":6789")
	if err != nil {
		fmt.Println("error starting server")
	}
}
