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

func (h *Handler) getPutlists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	putlists, err := h.services.Putlist.GetPutlists(userId)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getPutlistsResponse{
		Putlists: putlists,
	})

}

type getPutlistsByNumberResponse struct {
	Putlist       entity.GetPutlistResponse       `json:"putlist"`
	PutlistBodies []entity.GetPutlistBodyResponse `json:"putlist_bodies"`
}

func (h *Handler) getPutlistHeaderByNumber(c *gin.Context) {
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

func (h *Handler) createPutlistBody(c *gin.Context) {
	putlistId, err := strconv.Atoi(c.Param("number"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid putlist_number param")
	}

	var input entity.PutlistBody
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	putlistBody, err := h.services.Putlist.CreatePutlistBody(putlistId, input)
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

func (h *Handler) getPutlistBodies(c *gin.Context) {
	putlistId, err := strconv.Atoi(c.Param("number"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid putlist_number param")
	}
	putlistBodies, err := h.services.Putlist.GetPutlistBodies(putlistId)
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
