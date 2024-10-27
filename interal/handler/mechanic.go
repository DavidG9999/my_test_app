package handler

import (
	"net/http"
	"strconv"

	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/gin-gonic/gin"
)

type createMechanicResponse struct {
	Mechanic entity.Mechanic `json:"mechanic"`
}

func (h *Handler) createMechanic(c *gin.Context) {
	var input entity.Mechanic
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	mechanic, err := h.services.Mechanic.CreateMechanic(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, createMechanicResponse{
		Mechanic: mechanic,
	})
}

type getMechanicsResponse struct {
	Mechanics []entity.Mechanic `json:"mechanics"`
}

func (h *Handler) getMechanics(c *gin.Context) {
	mechanics, err := h.services.Mechanic.GetMechanics()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getMechanicsResponse{
		Mechanics: mechanics,
	})
}

type updateMechanicResponse struct {
	Mechanic entity.Mechanic `json:"update_mechanic"`
}

func (h *Handler) updateMechanic(c *gin.Context) {
	mechanicId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var updateData entity.UpdateMechanicInput

	if err := c.BindJSON(&updateData); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	updateMechanic, err := h.services.Mechanic.UpdateMechanic(mechanicId, updateData)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, updateMechanicResponse{
		Mechanic: updateMechanic,
	})
}

func (h *Handler) deleteMechanic(c *gin.Context) {
	mechanicId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	err = h.services.Mechanic.DeleteMechanic(mechanicId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok"})
}
