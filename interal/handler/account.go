package handler

import (
	"net/http"
	"strconv"

	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/gin-gonic/gin"
)

type createAccountResponse struct {
	Account entity.Account `json:"account"`
}

func (h *Handler) createAccount(c *gin.Context) {
	var input entity.Account
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	account, err := h.services.Account.CreateAccount(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, createAccountResponse{
		Account: account,
	})
}

type getAccountsResponse struct {
	Accounts []entity.Account `json:"accounts"`
}

func (h *Handler) getAccounts(c *gin.Context) {
	accounts, err := h.services.Account.GetAccounts()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAccountsResponse{
		Accounts: accounts,
	})
}

type getAccountByAccountNumberResponse struct {
	Account entity.Account
}

func (h *Handler) getAccountByAccountNumber(c *gin.Context) {
	account_number := c.Param("account_number")

	if len(account_number) != 20 {
		NewErrorResponse(c, http.StatusBadRequest, "invalid account_number param")
		return
	}

	account, err := h.services.Account.GetAccountByAccountNumber(account_number)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAccountByAccountNumberResponse{
		Account: account,
	})
}

type updateAccountResponse struct {
	Account entity.Account
}

func (h *Handler) updateAccount(c *gin.Context) {
	accountId, err := strconv.Atoi(c.Param("account_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid account_id param")
		return
	}
	var updateData entity.UpdateAccountInput

	if err := c.BindJSON(&updateData); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	updateAccount, err := h.services.Account.UpdateAccount(accountId, updateData)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, updateAccountResponse{
		Account: updateAccount,
	})

}

func (h *Handler) deleteAccount(c *gin.Context) {
	accountId, err := strconv.Atoi(c.Param("account_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid account_id param")
		return
	}

	if err := h.services.Account.DeleteAccount(accountId); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}
