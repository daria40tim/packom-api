package handler

import (
	"time"

	"github.com/daria40tim/packom/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func CORSM() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	//router.Use(cors.Default())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 120 * time.Hour,
	}))

	//router.Use(CORSM())

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/countries", h.getSelectCountries)
		auth.GET("/login", h.getSelectLogin)
	}

	api := router.Group("/api", h.orgIdentity)
	{
		orgs := api.Group("/orgs")
		{
			orgs.GET("/", h.getAllOrgs)
			orgs.GET("/:id", h.getOrgById)
			orgs.PUT("/", h.updateOrgById)
			orgs.PUT("/:id", h.addOrgById)
			orgs.GET("/select", h.getSelectSpecs)
			orgs.POST("/docs", h.getDocsById)
			orgs.GET("/doc/:name/:id", h.getDoc)
			orgs.GET("/delete/:id", h.deleteTrusted)
			orgs.GET("/filter", h.getFilterData)
			orgs.POST("/job/", h.getFilteredOrgs)
		}
		techs := api.Group("/techs")
		{
			techs.POST("/", h.createTech)
			techs.GET("/", h.getAllTechs)
			techs.GET("/:id", h.getTechById)
			techs.POST("/:id", h.updateTechById)
			techs.GET("/select", h.getSelect)
			techs.POST("/delete_cal", h.deleteCal)
			techs.POST("/delete_cst", h.deleteCst)
			techs.POST("/docs/:id", h.getTechDocsById)
			techs.GET("/doc/:name/:id", h.getTechDoc)
			techs.GET("/filter", h.getTechFilterData)
			techs.POST("/job/", h.getFilteredTechs)
		}
		cps := api.Group("/cps")
		{
			cps.POST("/", h.createCP)
			cps.GET("/", h.getAllCPs)
			cps.GET("/:id", h.getCPById)
			cps.PUT("/:id", h.updateCPById)
			cps.POST("/delete_cal", h.cpDeleteCal)
			cps.POST("/delete_cst", h.cpDeleteCst)
			cps.POST("/docs/:id", h.getCpDocsById)
			cps.GET("/doc/:name/:id", h.getCpDoc)
		}
		tenders := api.Group("/tenders")
		{
			tenders.POST("/", h.createTender)
			tenders.GET("/", h.getAllTenders)
			tenders.GET("/:id", h.getTenderById)
			tenders.PUT("/decide", h.updateTenderById)
			tenders.GET("/min", h.getMinandMax)
		}
	}

	/*router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		/*		AllowMethods:     []string{"PUT", "POST", "GET"},
				AllowHeaders:     []string{"Origin"},
				ExposeHeaders:    []string{"Content-Length"},
				AllowCredentials: true,
				MaxAge:           12 * time.Hour,
	}))*/

	return router
}
