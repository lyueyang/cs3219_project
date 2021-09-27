package main

import (
	"cs3219_project/handlers"
	"cs3219_project/storage"
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
		context.String(http.StatusOK, "hello world :)")
	})

	localStorage := storage.Storage{}
	hdlr := handlers.AccountHandler{Service: localStorage}

	r.GET("/accounts", hdlr.HandleGetAccounts)
	r.POST("/accounts", hdlr.HandleCreateAccounts)
	r.DELETE("/accounts", hdlr.HandleDeleteAccount)
	r.PUT("/accounts", hdlr.HandleUpdateAccount)

	err := r.Run(":6789")
	if err != nil {
		fmt.Println("error starting server")
	}
}
