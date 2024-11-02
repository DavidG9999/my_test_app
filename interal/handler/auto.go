package handler

import (
	"net/http"
	"strconv"

	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/gin-gonic/gin"
)

type createAutoResponse struct {
	Auto entity.Auto `json:"auto"`
}

// @Summary Create Auto
// @Security ApiKeyAuth
// @Tags autos
// @Description creates and adds new data about the car
// @ID create-auto
// @Accept json
// @Produce json
// @Param input body entity.Auto true "auto info"
// @Success 200 {object} createAutoResponse "new auto"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /autos [post]
func (h *Handler) createAuto(c *gin.Context) {
	var input entity.Auto
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	auto, err := h.services.Auto.CreateAuto(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, createAutoResponse{
		Auto: auto,
	})

}

type getAutosResponse struct {
	Autos []entity.Auto `json:"autos,omitempty"`
}

// @Summary Get Autos
// @Security ApiKeyAuth
// @Tags autos
// @Description get all data about the cars
// @ID get-autos
// @Accept json
// @Produce json
// @Success 200 {object} getAutosResponse "autos"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /autos [get]
func (h *Handler) getAutos(c *gin.Context) {
	autos, err := h.services.Auto.GetAutos()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAutosResponse{
		Autos: autos,
	})
}

type updateAutoResponse struct {
	Auto entity.Auto
}

// @Summary Update Auto
// @Security ApiKeyAuth
// @Tags autos
// @Description updates data about the car
// @ID update-auto
// @Accept json
// @Produce json
// @Param auto_id path integer true "auto id"
// @Param input body entity.UpdateAutoInput true "update auto info"
// @Success 200 {object} updateAutoResponse "updated auto"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /autos/{auto_id} [put]
func (h *Handler) updateAuto(c *gin.Context) {

	autoId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input entity.UpdateAutoInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	updateAuto, err := h.services.Auto.UpdateAuto(autoId, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, updateAutoResponse{
		Auto: updateAuto,
	})

}

// @Summary Delete Auto
// @Security ApiKeyAuth
// @Tags autos
// @Description deletes data about the car
// @ID delete-auto
// @Accept json
// @Produce json
// @Param auto_id path integer true "auto id"
// @Success 200 {object} statusResponse "ok"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /autos/{auto_id} [delete]
func (h *Handler) deleteAuto(c *gin.Context) {
	autoId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err := h.services.Auto.DeleteAuto(autoId); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}
