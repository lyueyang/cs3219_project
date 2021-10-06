package handlers

import (
	"bytes"
	mock "cs3219_project/handlers/mocks"
	"cs3219_project/storage"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAccountHandler_HandleCreateAccounts(t *testing.T) {
	t.Run("GIVEN valid account to create THEN return 200", func(t *testing.T) {
		// test setup
		ctrl := gomock.NewController(t)
		mockStorage := mock.NewMockStorageService(ctrl)
		mockStorage.EXPECT().StoreAccount(storage.Account{
			Name:        "user1",
			Description: "this is the first user",
		})
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// add information that would be sent
		inputAccount, _ := json.Marshal(AccountSchema{
			Username:    "user1",
			Description: "this is the first user",
		})
		req := httptest.NewRequest("POST", "/accounts", bytes.NewReader(inputAccount))
		c.Request = req

		a := AccountHandler{Service: mockStorage}
		a.HandleCreateAccounts(c)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("GIVEN invalid account to create THEN return 500", func(t *testing.T) {
		// test setup
		ctrl := gomock.NewController(t)
		mockStorage := mock.NewMockStorageService(ctrl)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// add information that would be sent
		req := httptest.NewRequest("POST", "/accounts", bytes.NewReader([]byte("account information")))
		c.Request = req

		a := AccountHandler{Service: mockStorage}
		a.HandleCreateAccounts(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestAccountHandler_HandleDeleteAccount(t *testing.T) {
	t.Run("GIVEN valid account to delete THEN return 200", func(t *testing.T) {
		// test setup
		ctrl := gomock.NewController(t)
		mockStorage := mock.NewMockStorageService(ctrl)
		mockStorage.EXPECT().DeleteAccount("user1")
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// add information that would be sent
		inputAccount, _ := json.Marshal(AccountSchema{
			Username: "user1",
		})
		req := httptest.NewRequest("DELETE", "/accounts", bytes.NewReader(inputAccount))
		c.Request = req

		a := AccountHandler{Service: mockStorage}
		a.HandleDeleteAccount(c)

		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("GIVEN invalid account to create THEN return 500", func(t *testing.T) {
		// test setup
		ctrl := gomock.NewController(t)
		mockStorage := mock.NewMockStorageService(ctrl)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// add information that would be sent
		req := httptest.NewRequest("DELETE", "/accounts", bytes.NewReader([]byte("account information")))
		c.Request = req

		a := AccountHandler{Service: mockStorage}
		a.HandleDeleteAccount(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestAccountHandler_HandleGetAccounts(t *testing.T) {
	t.Run("GIVEN get accounts THEN return accounts", func(t *testing.T) {
		// test setup
		ctrl := gomock.NewController(t)
		mockStorage := mock.NewMockStorageService(ctrl)
		mockStorage.EXPECT().GetAccounts().Return([]storage.Account{
			{
				Name:        "user1",
				Description: "this is the first user",
			},
			{
				Name:        "user2",
				Description: "this is the second user",
			},
		})

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		a := AccountHandler{Service: mockStorage}

		a.HandleGetAccounts(c)

		wantAccounts, _ := json.Marshal([]AccountSchema{
			{
				Username:    "user1",
				Description: "this is the first user",
			},
			{
				Username:    "user2",
				Description: "this is the second user",
			},
		})

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, wantAccounts, w.Body.Bytes())
	})
}

func TestAccountHandler_HandleUpdateAccount(t *testing.T) {
	t.Run("GIVEN valid account to update THEN return 200", func(t *testing.T) {
		// test setup
		ctrl := gomock.NewController(t)
		mockStorage := mock.NewMockStorageService(ctrl)
		mockStorage.EXPECT().UpdateAccount(storage.Account{
			Name:        "user1",
			Description: "new description",
		})
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// add information that would be sent
		inputAccount, _ := json.Marshal(AccountSchema{
			Username:    "user1",
			Description: "new description",
		})
		req := httptest.NewRequest("PUT", "/accounts", bytes.NewReader(inputAccount))
		c.Request = req

		a := AccountHandler{Service: mockStorage}
		a.HandleUpdateAccount(c)

		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("GIVEN invalid account to update THEN return 500", func(t *testing.T) {
		// test setup
		ctrl := gomock.NewController(t)
		mockStorage := mock.NewMockStorageService(ctrl)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// add information that would be sent
		req := httptest.NewRequest("PUT", "/accounts", bytes.NewReader([]byte("account information")))
		c.Request = req

		a := AccountHandler{Service: mockStorage}
		a.HandleUpdateAccount(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
