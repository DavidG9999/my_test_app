package handler

import (
	"github.com/DavidG9999/my_test_app/interal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}
	api := router.Group("/", h.userIdentity)
	{
		user := api.Group("/user")
		{
			user.GET("/", h.getUser)
			user.PUT("/name", h.updateName)
			user.PUT("/password", h.updatePassword)
			user.DELETE("/", h.deleteUser)
		}
		autos := api.Group("/autos")
		{
			autos.POST("/", h.createAuto)
			autos.GET("/", h.getAutos)
			autos.PUT("/:id", h.updateAuto)
			autos.DELETE("/:id", h.deleteAuto)
		}
		contragents := api.Group("/contragents")
		{
			contragents.POST("/", h.createContragent)
			contragents.GET("/", h.getContragents)
			contragents.PUT("/:id", h.updateContragent)
			contragents.DELETE("/:id", h.deleteContragent)

		}
		dispetchers := api.Group("/dispetchers")
		{
			dispetchers.POST("/", h.createDispetcher)
			dispetchers.GET("/", h.getDispetchers)
			dispetchers.PUT("/:id", h.updateDispetcher)
			dispetchers.DELETE("/:id", h.deleteDispetcher)
		}
		drivers := api.Group("/drivers")
		{
			drivers.POST("/", h.createDriver)
			drivers.GET("/", h.getDrivers)
			drivers.PUT("/:id", h.updateDriver)
			drivers.DELETE("/:id", h.deleteDriver)
		}
		mehanics := api.Group("/mehanics")
		{
			mehanics.POST("/", h.createMechanic)
			mehanics.GET("/", h.getMechanics)
			mehanics.PUT("/:id", h.updateMechanic)
			mehanics.DELETE("/:id", h.deleteMechanic)
		}
		organiations := api.Group("/organizations")
		{
			organiations.POST("/", h.createOrganization)
			organiations.GET("/", h.getOrganizations)
			organiations.PUT("/:id", h.updateOrganization)
			organiations.DELETE("/:id", h.deleteOrganization)

			accounts := organiations.Group(":id/accounts")
			{
				accounts.POST("/", h.createAccount)
				accounts.GET("/", h.getAccounts)
				accounts.PUT("/:account_id", h.updateAccount)
				accounts.DELETE("/:account_id", h.deleteAccount)
			}
		}
		putlists := api.Group("/putlists")
		{
			putlists.POST("/", h.createPutlistHeader)
			putlists.GET("/", h.getPutlists)
			putlists.GET("/:number", h.getPutlistHeaderByNumber)
			putlists.PUT("/:number", h.updatePutlistHeader)
			putlists.DELETE("/:number", h.deletePutlistHeader)

			putlist_bodies := putlists.Group(":number/putlist_bodies")
			{
				putlist_bodies.POST("/", h.createPutlistBody)
				putlist_bodies.GET("/", h.getPutlistBodies)
				putlist_bodies.PUT("/:putlist_body_id", h.updatePutlistBody)
				putlist_bodies.DELETE("/:putlist_body_id", h.deletePutlistBody)

			}
		}
	}
	return router
}
