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
	organiationId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid organization_id param")
	}

	var input entity.Account
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	account, err := h.services.Account.CreateAccount(organiationId, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, createAccountResponse{
		Account: account,
	})
}

type getAccountsResponse struct {
	Organization entity.Organization `json:"organization"`
	Accounts     []entity.Account    `json:"accounts"`
}

func (h *Handler) getAccounts(c *gin.Context) {
	organizationId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid organization_id param")
	}
	organization, accounts, err := h.services.Account.GetAccounts(organizationId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAccountsResponse{
		Organization: organization,
		Accounts:     accounts,
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
