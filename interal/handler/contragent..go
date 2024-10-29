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

// @Summary Create Contragent
// @Security ApiKeyAuth
// @Tags contragents
// @Description creates data about the contragent
// @ID create-contragent
// @Accept json
// @Produce json
// @Param input body entity.Contragent true "contragent info"
// @Success 200 {object} createContragentResponse "new account"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /contragent [post]
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

// @Summary Get Contragents
// @Security ApiKeyAuth
// @Tags contragents
// @Description gets data about all contragents
// @ID get-contragents
// @Accept json
// @Produce json
// @Success 200 {object} getContragentsResponse "contragents info"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /contragents [get]
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

// @Summary Update Contragent
// @Security ApiKeyAuth
// @Tags contragents
// @Description updates data about the contragent
// @ID update-contragent
// @Accept json
// @Produce json
// @Param contragent_id body integer true "contragent id"
// @Param updateData body entity.UpdateContragentInput true "update contragent info"
// @Success 200 {object} updateContragentResponse "updated contragent"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /contragents/{contragent_id} [put]
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

// @Summary Delete Contragent
// @Security ApiKeyAuth
// @Tags contragents
// @Description deletes data about contragent
// @ID delete-contragent
// @Accept json
// @Produce json
// @Param contragent_id path integer true "contragent id"
// @Success 200 {object} statusResponse "ok"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /contragents/{contragent_id} [delete]
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
