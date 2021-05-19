package handler

import (
	"github.com/daria40tim/packom/pkg/service"
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
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.orgIdentity)
	{
		orgs := api.Group("/orgs")
		{
			//orgs.POST("/", h.createOrg)
			orgs.GET("/", h.getAllOrgs)
			orgs.GET("/:id", h.getOrgById)
			orgs.PUT("/", h.updateOrgById)
		}
		techs := api.Group("/techs")
		{
			techs.POST("/", h.createTech)
			techs.GET("/", h.getAllTechs)
			techs.GET("/:id", h.getTechById)
			techs.PUT("/:id", h.updateTechById)
		}
		cps := api.Group("/cps")
		{
			cps.POST("/", h.createCP)
			cps.GET("/", h.getAllCPs)
			cps.GET("/:id", h.getCPById)
			cps.PUT("/:id", h.updateCPById)
		}
		tenders := api.Group("/tenders")
		{
			tenders.POST("/", h.createTender)
			tenders.GET("/", h.getAllTenders)
			tenders.GET("/:id", h.getTenderById)
			tenders.PUT("/:id", h.updateTenderById)
		}
	}
	return router
}
