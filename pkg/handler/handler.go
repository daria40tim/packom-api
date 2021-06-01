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
	}

	api := router.Group("/api", h.orgIdentity)
	{
		orgs := api.Group("/orgs")
		{
			//orgs.POST("/", h.createOrg)
			orgs.GET("/", h.getAllOrgs)
			orgs.GET("/:id", h.getOrgById)
			orgs.PUT("/", h.updateOrgById)
			orgs.PUT("/:id", h.addOrgById)
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
