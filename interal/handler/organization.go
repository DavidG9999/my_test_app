package handler

import (
	"net/http"
	"strconv"

	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/gin-gonic/gin"
)

type createOrganizationResponse struct {
	Organization entity.Organization `json:"organization"`
}

func (h *Handler) createOrganization(c *gin.Context) {
	var input entity.Organization
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	organization, err := h.services.Organization.CreateOrganization(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, createOrganizationResponse{
		Organization: organization,
	})
}

type getOrganizationsResponse struct {
	Organizations []entity.Organization `json:"organizations"`
}

func (h *Handler) getOrganizations(c *gin.Context) {
	organizations, err := h.services.Organization.GetOrganizations()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getOrganizationsResponse{
		Organizations: organizations,
	})
}

type updateOrganizationResponse struct {
	Organization entity.Organization `json:"update_organization"`
}

func (h *Handler) updateOrganization(c *gin.Context) {
	organiationId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var updateData entity.UpdateOrganizationInput

	if err := c.BindJSON(&updateData); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	updateOrg, err := h.services.Organization.UpdateOrganization(organiationId, updateData)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, updateOrganizationResponse{
		Organization: updateOrg,
	})
}

func (h *Handler) deleteOrganization(c *gin.Context) {
	organiationId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	err = h.services.Organization.DeleteOrganization(organiationId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
