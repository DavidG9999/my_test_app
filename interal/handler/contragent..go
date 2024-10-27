package handler

import (
	"net/http"
	"strconv"

	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/gin-gonic/gin"
)

type createContragentResponse struct {
	Contragent entity.Contragent `json:"contragent"`
}

func (h *Handler) createContragent(c *gin.Context) {
	var input entity.Contragent
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	contragent, err := h.services.Contragent.CreateContragent(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, createContragentResponse{
		Contragent: contragent,
	})
}

type getContragentsResponse struct {
	Contragents []entity.Contragent `json:"contragents"`
}

func (h *Handler) getContragents(c *gin.Context) {
	contragents, err := h.services.Contragent.GetContragents()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getContragentsResponse{
		Contragents: contragents,
	})
}

type updateContragentResponse struct {
	Contragent entity.Contragent `json:"update_contragent"`
}

func (h *Handler) updateContragent(c *gin.Context) {
	contragentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var updateData entity.UpdateContragentInput

	if err := c.BindJSON(&updateData); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	updateContragent, err := h.services.Contragent.UpdateContragent(contragentId, updateData)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, updateContragentResponse{
		Contragent: updateContragent,
	})
}

func (h *Handler) deleteContragent(c *gin.Context) {
	contragentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	err = h.services.Contragent.DeleteContragent(contragentId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
