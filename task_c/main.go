package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "users"
)

func main() {
	r := gin.Default()

	r.GET("/", basicAuth, func(context *gin.Context) {
		context.String(http.StatusOK, "hello world")
	})

	err := r.Run(":6789")
	if err != nil {
		fmt.Println("error starting server")
	}
}

func basicAuth(c *gin.Context) {
	user, pw, hasAuth := c.Request.BasicAuth()
	if hasAuth {
		if user == "user1" {
			if pw == "user1pw" {
				c.Status(http.StatusOK)
			} else {
				c.Abort()
				c.Status(http.StatusForbidden)
			}
		} else if user == "admin" {
			if pw == "adminuserpw" {
				c.Status(http.StatusOK)
			} else {
				c.Abort()
				c.Status(http.StatusForbidden)
			}
		} else {
			c.Abort()
			c.Status(http.StatusUnauthorized)
		}
	} else {
		c.Abort()
		c.Status(http.StatusUnauthorized)
	}
}
