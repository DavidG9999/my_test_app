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
			user.GET("/id", h.getUserId)
			user.GET("/", h.getUser)
			user.PUT("/", h.updateUser)
			user.DELETE("/", h.deleteUser)
		}
		autos := api.Group("/autos")
		{
			autos.POST("/", h.createAuto)
			autos.GET("/", h.getAutos)
			autos.GET("/:id", h.getAutoById)
			autos.PUT("/:id", h.updateAuto)
			autos.DELETE("/:id", h.deleteAuto)
		}
		contragents := api.Group("/contragents")
		{
			contragents.POST("/", h.createContragent)
			contragents.GET("/", h.getContragents)
			contragents.GET("/:id", h.getContragentById)
			contragents.PUT("/:id", h.updateContragent)
			contragents.DELETE("/:id", h.deleteContragent)

			items := contragents.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getItems)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}
		dispetchers := api.Group("/dispetchers")
		{
			dispetchers.POST("/", h.createDispetcher)
			dispetchers.GET("/", h.getDispetchers)
			dispetchers.GET("/:id", h.getDispetcherById)
			dispetchers.PUT("/:id", h.updateDispetcher)
			dispetchers.DELETE("/:id", h.deleteDispetcher)
		}
		drivers := api.Group("/drivers")
		{
			drivers.POST("/", h.createDriver)
			drivers.GET("/", h.getDrivers)
			drivers.GET("/:id", h.getDriverById)
			drivers.PUT("/:id", h.updateDriver)
			drivers.DELETE("/:id", h.deleteDriver)
		}
		mehanics := api.Group("/mehanics")
		{
			mehanics.POST("/", h.createMechanic)
			mehanics.GET("/", h.getMechanics)
			mehanics.GET("/:id", h.getMechanicById)
			mehanics.PUT("/:id", h.updateMechanic)
			mehanics.DELETE("/:id", h.deleteMechanic)
		}
		organiations := api.Group("/organiations")
		{
			organiations.POST("/", h.createOrganization)
			organiations.GET("/", h.getOrganizations)
			organiations.GET("/:id", h.getOrganizationById)
			organiations.PUT("/:id", h.updateOrganization)
			organiations.DELETE("/:id", h.deleteOrganization)

			accounts := organiations.Group(":id/accounts")
			{
				accounts.POST("/", h.createAccount)
				accounts.GET("/", h.getAccounts)
				accounts.GET("/:account_id", h.getAccountById)
				accounts.PUT("/:account_id", h.updateAccount)
				accounts.DELETE("/:account_id", h.deleteAccount)
			}
		}
		putlists := api.Group("/putlists")
		{
			putlists.POST("/", h.createPutlistHeader)
			putlists.GET("/", h.getPutlistHeaders)
			putlists.GET("/:id", h.getPutlistHeaderById)
			putlists.PUT("/:id", h.updatePutlistHeader)
			putlists.DELETE("/:id", h.deletePutlistHeader)

			putlist_bodies := putlists.Group(":id/putlists")
			{
				putlist_bodies.POST("/", h.createPutlistBody)
				putlist_bodies.GET("/", h.getPutlistBodies)
				putlist_bodies.GET("/:putlist_body_id", h.getPutlistBodyById)
				putlist_bodies.PUT("/:putlist_body_id", h.updatePutlistBody)
				putlist_bodies.DELETE("/:putlist_body_id", h.deletePutlistBody)

			}
		}
	}
	return router
}
