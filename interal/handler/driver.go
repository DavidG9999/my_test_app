package handler

import (
	"net/http"
	"strconv"

	"github.com/DavidG9999/my_test_app/interal/entity"
	"github.com/gin-gonic/gin"
)

type createDriverResponse struct {
	Driver entity.Driver `json:"driver"`
}

// @Summary Create Driver
// @Security ApiKeyAuth
// @Tags drivers
// @Description creates and adds data about the driver
// @ID create-driver
// @Accept json
// @Produce json
// @Param input body entity.Driver true "driver info"
// @Success 200 {object} createDriverResponse "new driver"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /drivers [post]
func (h *Handler) createDriver(c *gin.Context) {
	var input entity.Driver
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	driver, err := h.services.Driver.CreateDriver(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, createDriverResponse{
		Driver: driver,
	})
}

type getDriversResponse struct {
	Drivers []entity.Driver `json:"drivers"`
}

// @Summary Get Drivers
// @Security ApiKeyAuth
// @Tags drivers
// @Description gets data about all drivers
// @ID get-drivers
// @Accept json
// @Produce json
// @Success 200 {object} getDriversResponse "drivers info"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /drivers [get]
func (h *Handler) getDrivers(c *gin.Context) {
	drivers, err := h.services.Driver.GetDrivers()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getDriversResponse{
		Drivers: drivers,
	})
}

type updateDriverResponse struct {
	Driver entity.Driver `json:"update_driver"`
}

// @Summary Update Driver
// @Security ApiKeyAuth
// @Tags drivers
// @Description updates data about the driver
// @ID update-driver
// @Accept json
// @Produce json
// @Param driver_id body integer true "driver id"
// @Param updateData body entity.UpdateDriverInput true "update driver info"
// @Success 200 {object} updateDriverResponse "updated driver"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /drivers/{driver_id} [put]
func (h *Handler) updateDriver(c *gin.Context) {
	driverId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var updateData entity.UpdateDriverInput

	if err := c.BindJSON(&updateData); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	updateDriver, err := h.services.Driver.UpdateDriver(driverId, updateData)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, updateDriverResponse{
		Driver: updateDriver,
	})
}

// @Summary Delete Driver
// @Security ApiKeyAuth
// @Tags drivers
// @Description deletes data about driver
// @ID delete-driver
// @Accept json
// @Produce json
// @Param driver_id body integer true "driver id"
// @Success 200 {object} statusResponse "ok"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /drivers/{driver_id} [delete]
func (h *Handler) deleteDriver(c *gin.Context) {
	driverId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	err = h.services.Driver.DeleteDriver(driverId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
