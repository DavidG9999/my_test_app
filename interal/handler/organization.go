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

// @Summary Create Organization
// @Security ApiKeyAuth
// @Tags organizations
// @Description creates and adds data about the organization
// @ID create-organization
// @Accept json
// @Produce json
// @Param input body entity.Organization true "organization info"
// @Success 200 {object} createOrganizationResponse "new organization"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /organizations [post]
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

// @Summary Get Organizations
// @Security ApiKeyAuth
// @Tags organizations
// @Description gets data about all organizations
// @ID get-organizations
// @Accept json
// @Produce json
// @Success 200 {object} getOrganizationsResponse "organizations info"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /organizations [get]
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

// @Summary Update Organizaion
// @Security ApiKeyAuth
// @Tags organizations
// @Description updates data about the organization
// @ID update-organization
// @Accept json
// @Produce json
// @Param organization_id body integer true "organization id"
// @Param updateData body entity.UpdateOrganizationInput true "update organization info"
// @Success 200 {object} updateOrganizationResponse "updated organization"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /organizations/{organization_id} [put]
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

// @Summary Delete Organization
// @Security ApiKeyAuth
// @Tags organizations
// @Description deletes data about organization
// @ID delete-organization
// @Accept json
// @Produce json
// @Param organization_id body integer true "organization id"
// @Success 200 {object} statusResponse "ok"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /organizations/{organization_id} [delete]
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
