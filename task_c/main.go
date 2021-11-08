package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"io/ioutil"
	"net/http"

	_ "github.com/lib/pq"
)

var secret = []byte("shhsecret")

func main() {
	r := gin.Default()

	r.POST("/gettoken", issueToken)

	r.GET("/", advancedAuth, func(context *gin.Context) {
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

func issueToken(c *gin.Context) {
	type acctDetails struct {
		Name     string
		Password string
	}

	var receivedAcct acctDetails

	reqBody, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(reqBody, &receivedAcct)

	if err != nil {
		// unable to retrieve user info
		c.Status(http.StatusForbidden) // no authentication info found
		return
	}

	var token *jwt.Token
	if receivedAcct.Name == "adminuser" && receivedAcct.Password == "adminpw" {
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"name": receivedAcct.Name,
			"role": "admin",
		})
	} else if receivedAcct.Name == "normaluser" && receivedAcct.Password == "userpw" {
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"name": receivedAcct.Name,
			"role": "user",
		})
	} else {
		c.String(http.StatusForbidden, "your credentials don't check out :(")
		return
	}

	if token != nil {
		out, _ := token.SignedString(secret)
		c.String(http.StatusOK, out)
		return
	}
}

func advancedAuth(c *gin.Context) {
	tokenString := c.GetHeader("x-auth-token")

	if tokenString == "" {
		c.Abort()
		c.String(http.StatusUnauthorized, "I don't know who you are") // no identification == no authentication provided
		return
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check authorisation
		name := claims["name"]
		role := claims["role"]

		if (name == "adminuser" && role == "admin") || (name == "normaluser" && role == "user") {
			// credentials check out, ok :)
			c.Status(http.StatusOK)
		} else {
			// doesn't seem like the right person, sorry
			c.String(http.StatusForbidden, "your credentials don't check out :(")
		}
	} else {
		c.Abort()
		c.String(http.StatusUnauthorized, "your credentials don't check out :(") // token invalid == inauthentic
	}
}
