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

// @Summary Create Account
// @Security ApiKeyAuth
// @Tags accounts
// @Description creates new data about the bank account
// @ID create-bank-account
// @Accept json
// @Produce json
// @Param organization_id path int true "organization id"
// @Param input body entity.Account true "account info"
// @Success 200 {object} createAccountResponse "created account"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /{organization_id}/accounts [post]
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

// @Summary Get Accounts
// @Security ApiKeyAuth
// @Tags accounts
// @Description gets data about all bank accounts
// @ID get-accounts
// @Accept json
// @Produce json
// @Param organization_id path int true "organization id"
// @Success 200 {object} getAccountsResponse "get accounts"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /{organization_id}/accounts [get]
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

// @Summary Update Account
// @Security ApiKeyAuth
// @Tags accounts
// @Description updates data about the bank account
// @ID update-account
// @Accept json
// @Produce json
// @Param organization_id path int true "organization id"
// @Param account_id path int true "account id"
// @Param updateData body entity.UpdateAccountInput true "update account info"
// @Success 200 {object} updateAccountResponse "updated account"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /{organization_id}/accounts/{account_id} [put]
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

// @Summary Delete Account
// @Security ApiKeyAuth
// @Tags accounts
// @Description  deletes bank account
// @ID delete-account
// @Accept json
// @Produce json
// @Param organization_id path int true "organization id"
// @Param account_id path int true "account id"
// @Success 200 {object} statusResponse "ok"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /{organization_id}/accounts/{account_id} [delete]
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
