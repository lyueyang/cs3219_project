package handlers

import (
	"cs3219_project/storage"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type StorageService interface {
	GetAccounts() []storage.Account
	StoreAccount(a *storage.Account)
}

type AccountSchema struct {
	Username    string
	Description string
}

type ErrorResponse struct {
	Message string
}

type AccountHandler struct {
	DBService 		StorageService
	RedisService 	StorageService
}

func (a AccountHandler) HandleGetAccounts(c *gin.Context) {
	var accounts []storage.Account
	isDB := false

	if res := a.RedisService.GetAccounts(); res != nil {
		accounts = res
	} else {
		accounts = a.DBService.GetAccounts()
		isDB = true
	}

	var out []AccountSchema

	for _, acct := range accounts {
		out = append(out, AccountSchema{
			Username:    acct.Name,
			Description: acct.Description,
		})

		if isDB {
			a.RedisService.StoreAccount(&acct)
		}
	}

	c.JSON(http.StatusOK, out)
}

func (a AccountHandler) HandleCreateAccounts(c *gin.Context) {
	var acct AccountSchema

	reqBody, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(reqBody, &acct)

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "json parsing error"})
		return
	}

	// we are intentionally choosing not to report account creation status for security purposes
	a.DBService.StoreAccount(&storage.Account{
		Name:        acct.Username,
		Description: acct.Description,
	})
	c.Status(http.StatusOK)
}
