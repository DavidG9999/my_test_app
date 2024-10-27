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
