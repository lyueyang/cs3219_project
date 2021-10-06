//go:generate mockgen -source=handle_accounts.go -destination=mocks/mock_storage.go -package=mock

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
	StoreAccount(a storage.Account)
	DeleteAccount(a string)
	UpdateAccount(a storage.Account)
}

type AccountSchema struct {
	Username    string
	Description string
}

type ErrorResponse struct {
	Message string
}

type AccountHandler struct {
	Service StorageService
}

func (a AccountHandler) HandleGetAccounts(c *gin.Context) {
	accounts := a.Service.GetAccounts()

	var out []AccountSchema

	for _, acct := range accounts {
		out = append(out, AccountSchema{
			Username:    acct.Name,
			Description: acct.Description,
		})
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
	a.Service.StoreAccount(storage.Account{
		Name:        acct.Username,
		Description: acct.Description,
	})
	c.Status(http.StatusOK)
}

func (a AccountHandler) HandleDeleteAccount(c *gin.Context) {
	var acct AccountSchema

	reqBody, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(reqBody, &acct)

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "json parsing error"})
		return
	}

	// intentionally choose not to report on account deletion status for security purpose
	a.Service.DeleteAccount(acct.Username)
	c.Status(http.StatusOK)
}

func (a AccountHandler) HandleUpdateAccount(c *gin.Context) {
	var acct AccountSchema

	reqBody, _ := ioutil.ReadAll(c.Request.Body)
	err := json.Unmarshal(reqBody, &acct)

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "json parsing error"})
		return
	}

	// intentionally choose not to report on account updating status for security purpose
	a.Service.UpdateAccount(storage.Account{
		Name:        acct.Username,
		Description: acct.Description,
	})
}
