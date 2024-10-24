package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getIdUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": userId,
	})
}

func (h *Handler) getUser(c *gin.Context) {

}

func (h *Handler) updateName(c *gin.Context) {

}

func (h *Handler) updatePassword(c *gin.Context) {

}

func (h *Handler) deleteUser(c *gin.Context) {

}
