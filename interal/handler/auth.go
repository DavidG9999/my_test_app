package handler

import (
	"net/http"

	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/gin-gonic/gin"
)

// @Summary SignUp
// @Tags auth
// @Description creates new account
// @ID create-user-account
// @Accept json
// @Produce json
// @Param input body entity.User true "account information"
// @Success 200 {integer} integer "id"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input entity.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary SignIn
// @Tags auth
// @Description authentication and JWT-token issuance
// @ID login
// @Accept json
// @Produce json
// @Param input body signInInput true "credentials (email and pasword)"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

// @Summary Get User
// @Security ApiKeyAuth
// @Tags user
// @Description gives information about the user
// @ID get-user-info
// @Accept json
// @Produce json
// @Success 200 {object} entity.User "user"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user [get]
func (h *Handler) getUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	user, err := h.services.Authorization.GetUserById(userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, user)

}

// @Summary Update Name
// @Security ApiKeyAuth
// @Tags user
// @Description updates username
// @ID update-name
// @Accept json
// @Produce json
// @Param updateData body entity.UpdateNameUserInput true "update password"
// @Success 200 {object} entity.User "updated user info"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/name [put]
func (h *Handler) updateName(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var updateData entity.UpdateNameUserInput
	if err := c.BindJSON(&updateData); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	updateUser, err := h.services.Authorization.UpdateName(userId, updateData)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, updateUser)
}

// @Summary Update Password
// @Security ApiKeyAuth
// @Tags user
// @Description updates user password
// @ID update-password
// @Accept json
// @Produce json
// @Param updateData body entity.UpdatePasswordUserInput true "update password"
// @Success 200 {object} entity.User "updated user info"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/password [put]
func (h *Handler) updatePassword(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var updateData entity.UpdatePasswordUserInput
	if err := c.BindJSON(&updateData); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	updateUser, err := h.services.Authorization.UpdatePassword(userId, updateData)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, updateUser)

}

// @Summary Delete Account
// @Security ApiKeyAuth
// @Tags user
// @Description deletes user account
// @ID delete-user-account
// @Accept json
// @Produce json
// @Success 200 {object} statusResponse "ok"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user [delete]
func (h *Handler) deleteUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	err = h.services.Authorization.DeleteUser(userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}
