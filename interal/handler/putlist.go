package handler

import (
	"net/http"
	"strconv"

	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/gin-gonic/gin"
)

type createPutlistResponse struct {
	Putlist entity.PutlistHeader `json:"putlist"`
}

// @Summary Create Putlist
// @Security ApiKeyAuth
// @Tags putlists
// @Description creates and adds data about the putlist
// @ID create-putlist
// @Accept json
// @Produce json
// @Param input body entity.PutlistHeader true "putlist info"
// @Success 200 {object} createPutlistResponse "new putlist"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /putlists [post]
func (h *Handler) createPutlistHeader(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input entity.PutlistHeader
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	putlist, err := h.services.Putlist.CreatePutlist(userId, input)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, createPutlistResponse{
		Putlist: putlist,
	})

}

type getPutlistsResponse struct {
	Putlists []entity.GetPutlistResponse `json:"putlists"`
}

// @Summary Get Putlists
// @Security ApiKeyAuth
// @Tags putlists
// @Description gets data about all putlist headers
// @ID get-putlists
// @Accept json
// @Produce json
// @Success 200 {object} getPutlistsResponse "putlist headers info"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /putlists [get]
func (h *Handler) getPutlistHeaders(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	putlists, err := h.services.Putlist.GetPutlistHeaders(userId)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getPutlistsResponse{
		Putlists: putlists,
	})

}

// @Summary Get Putlist By Number
// @Security ApiKeyAuth
// @Tags putlists
// @Description gets data about putlist
// @ID get-putlist-by-number
// @Accept json
// @Produce json
// @Param number path integer true "number"
// @Success 200 {object} getPutlistsByNumberResponse "putlist info"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /putlists/{number} [get]
type getPutlistsByNumberResponse struct {
	Putlist       entity.GetPutlistResponse       `json:"putlist"`
	PutlistBodies []entity.GetPutlistBodyResponse `json:"putlist_bodies"`
}

func (h *Handler) getPutlistByNumber(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	number, err := strconv.Atoi(c.Param("number"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid number param")
		return
	}

	putlist, putlistBodies, err := h.services.Putlist.GetPutlistByNumber(userId, number)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getPutlistsByNumberResponse{
		Putlist:       putlist,
		PutlistBodies: putlistBodies,
	})
}

type updatePutlistResponse struct {
	UpdatePutlist entity.PutlistHeader `json:"update_putlist"`
}

// @Summary Update Putlist
// @Security ApiKeyAuth
// @Tags putlists
// @Description updates data about the putlist header
// @ID update-putlist
// @Accept json
// @Produce json
// @Param number path integer true "number"
// @Param updateData body entity.UpdatePutlistHeaderInput true "update putlist info"
// @Success 200 {object} updatePutlistResponse "updated putlist"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /putlists/{number} [put]
func (h *Handler) updatePutlistHeader(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	number, err := strconv.Atoi(c.Param("number"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid number param")
		return
	}

	var updateData entity.UpdatePutlistHeaderInput
	if err := c.BindJSON(&updateData); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	updatePutlist, err := h.services.Putlist.UpdatePutlist(userId, number, updateData)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, updatePutlistResponse{
		UpdatePutlist: updatePutlist,
	})
}

// @Summary Delete Putlist
// @Security ApiKeyAuth
// @Tags putlists
// @Description deletes data about putlist
// @ID delete-putlist
// @Accept json
// @Produce json
// @Param number path integer true "number"
// @Success 200 {object} statusResponse "ok"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /putlists/{number} [delete]
func (h *Handler) deletePutlistHeader(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	number, err := strconv.Atoi(c.Param("number"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid number param")
		return
	}

	err = h.services.Putlist.DeletePutlist(userId, number)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

type createPutlistBodyResponse struct {
	PutlistBody entity.PutlistBody `json:"putlist_body"`
}

// @Summary Create Putlist Body
// @Security ApiKeyAuth
// @Tags putlists
// @Description creates and adds data about the putlist body
// @ID create-putlist-body
// @Accept json
// @Produce json
// @Param number path integer true "putlist number"
// @Param input body entity.PutlistBody true "putlist body info"
// @Success 200 {object} createPutlistBodyResponse "new putlist body"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /{number}/putlist_bodies [post]
func (h *Handler) createPutlistBody(c *gin.Context) {
	putlistNumber, err := strconv.Atoi(c.Param("number"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid putlist_number param")
	}

	var input entity.PutlistBody
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	putlistBody, err := h.services.Putlist.CreatePutlistBody(putlistNumber, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, createPutlistBodyResponse{
		PutlistBody: putlistBody,
	})
}

type getPutlistBodiesResponse struct {
	PutlistBodies []entity.GetPutlistBodyResponse `json:"putlist_bodies"`
}

// @Summary Get Putlist Bodies
// @Security ApiKeyAuth
// @Tags putlists
// @Description gets data about all putlist bodies
// @ID get-putlist-bodies
// @Accept json
// @Produce json
// @Param number path integer true "putlist number"
// @Success 200 {object} getPutlistBodiesResponse "putlist bodies info"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /{number}/putlist_bodies [get]
func (h *Handler) getPutlistBodies(c *gin.Context) {
	putlistNumber, err := strconv.Atoi(c.Param("number"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid putlist_number param")
	}
	putlistBodies, err := h.services.Putlist.GetPutlistBodies(putlistNumber)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getPutlistBodiesResponse{
		PutlistBodies: putlistBodies,
	})
}

type updatePutlistBodyResponse struct {
	PutlistBody entity.PutlistBody `json:"update_putlist_body"`
}

// @Summary Update Putlist Body
// @Security ApiKeyAuth
// @Tags putlists
// @Description updates data about the putlist body
// @ID update-putlist-body
// @Accept json
// @Produce json
// @Param number path integer true "putlist number"
// @Param putlist_body_id path integer true "putlist body id"
// @Param updateData body entity.UpdatePutlistBodyInput true "update putlist body info"
// @Success 200 {object} updatePutlistBodyResponse "updated putlist body"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /{number}/putlist_bodies/{putlist_body_id} [put]
func (h *Handler) updatePutlistBody(c *gin.Context) {
	putlistBodyId, err := strconv.Atoi(c.Param("putlist_body_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid putlist_body_number param")
	}

	var updateData entity.UpdatePutlistBodyInput
	if err := c.BindJSON(&updateData); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	updatePutlistBody, err := h.services.Putlist.UpdatePutlistBody(putlistBodyId, updateData)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, updatePutlistBodyResponse{
		PutlistBody: updatePutlistBody,
	})
}

// @Summary Delete Putlist Body
// @Security ApiKeyAuth
// @Tags putlists
// @Description deletes data about putlist body
// @ID delete-putlist-body
// @Accept json
// @Produce json
// @Param number path integer true "putlist number"
// @Param putlist_body_id path integer true "putlist body id"
// @Success 200 {object} statusResponse "ok"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /{number}/putlist_bodies/{putlist_body_id} [delete]
func (h *Handler) deletePutlistBody(c *gin.Context) {
	putlistBodyId, err := strconv.Atoi(c.Param("putlist_body_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid putlist_body_number param")
	}

	err = h.services.Putlist.DeletePutlistBody(putlistBodyId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
