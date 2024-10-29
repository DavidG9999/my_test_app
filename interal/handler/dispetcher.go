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

// @Summary Create Dispetcher
// @Security ApiKeyAuth
// @Tags dispetchers
// @Description creates and adds data about the dispetcher
// @ID create-dispetcher
// @Accept json
// @Produce json
// @Param input body entity.Dispetcher true "dispetcher info"
// @Success 200 {object} createDispetcherResponse "new dispetcher"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /dispetchers [post]
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

// @Summary Get Dispetchers
// @Security ApiKeyAuth
// @Tags dispetchers
// @Description gets data about all dispetchers
// @ID get-dispetchers
// @Accept json
// @Produce json
// @Success 200 {object} getDispetchersResponse "dispetchers info"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /dispetchers [get]
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

// @Summary Update Dispetcher
// @Security ApiKeyAuth
// @Tags dispetchers
// @Description updates data about the dispetcher
// @ID update-dispetcher
// @Accept json
// @Produce json
// @Param dispetcher_id  body integer true "dispetcher id"
// @Param updateData body entity.UpdateDispetcherInput true "update dispetcher info"
// @Success 200 {object} updateDispetcherResponse "updated dispetcher"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /dispetchers/{dispetcher_id} [put]
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

// @Summary Delete Dispetcher
// @Security ApiKeyAuth
// @Tags dispetchers
// @Description deletes data about dispetcher
// @ID delete-dispetcher
// @Accept json
// @Produce json
// @Param dispetcher_id body integer true "dispetcher id"
// @Success 200 {object} statusResponse "ok"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /dispetchers/{dispetcher_id} [delete]
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
