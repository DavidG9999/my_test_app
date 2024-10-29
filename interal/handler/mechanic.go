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

// @Summary Create Mehanic
// @Security ApiKeyAuth
// @Tags mehanics
// @Description creates and adds data about the mehanic
// @ID create-mehanic
// @Accept json
// @Produce json
// @Param input body entity.Mechanic true "mehanic info"
// @Success 200 {object} createMechanicResponse "new mehanic"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /mehanics [post]
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

// @Summary Get Mehanics
// @Security ApiKeyAuth
// @Tags mehanics
// @Description gets data about all mehanics
// @ID get-mehanics
// @Accept json
// @Produce json
// @Success 200 {object} getMechanicsResponse "mehanics info"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /mehanics [get]
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

// @Summary Update Mehanic
// @Security ApiKeyAuth
// @Tags mehanics
// @Description updates data about the mehanic
// @ID update-mehanic
// @Accept json
// @Produce json
// @Param mehanic_id body integer true "mehanic id"
// @Param updateData body entity.UpdateMechanicInput true "update mehanic info"
// @Success 200 {object} updateMechanicResponse "updated mehanic"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /mehanics/{mehanic_id} [put]
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

// @Summary Delete Mehanic
// @Security ApiKeyAuth
// @Tags mehanics
// @Description deletes data about mehanic
// @ID delete-mehanic
// @Accept json
// @Produce json
// @Param mehanic_id body integer true "mehanic id"
// @Success 200 {object} statusResponse "ok"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /mehanics/{mehanic_id} [delete]
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
