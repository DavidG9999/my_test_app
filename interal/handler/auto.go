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
