package handler

import (
	"net/http"
	"strconv"

	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/gin-gonic/gin"
)

type createDispetcherResponse struct {
	Dispetcher entity.Dispetcher `json:"dispetcher"`
}

func (h *Handler) createDispetcher(c *gin.Context) {
	var input entity.Dispetcher
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	dispetcher, err := h.services.Dispetcher.CreateDispetcher(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, createDispetcherResponse{
		Dispetcher: dispetcher,
	})
}

type getDispetchersResponse struct {
	Dispetchers []entity.Dispetcher `json:"dispetchers"`
}

func (h *Handler) getDispetchers(c *gin.Context) {
	dispetchers, err := h.services.Dispetcher.GetDispetchers()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getDispetchersResponse{
		Dispetchers: dispetchers,
	})
}

type updateDispetcherResponse struct {
	Dispetcher entity.Dispetcher `json:"update_dispetcher"`
}

func (h *Handler) updateDispetcher(c *gin.Context) {
	dispetcherId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var updateData entity.UpdateDispetcherInput

	if err := c.BindJSON(&updateData); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	updateDispetcher, err := h.services.Dispetcher.UpdateDispetcher(dispetcherId, updateData)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, updateDispetcherResponse{
		Dispetcher: updateDispetcher,
	})
}

func (h *Handler) deleteDispetcher(c *gin.Context) {
	dispetcherId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	err = h.services.Dispetcher.DeleteDispetcher(dispetcherId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
