package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sidecar/internal/http/rest"
)

func SetupRouter(handler *rest.Handler, cfg *MainConfig) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(CORSMiddleware())

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, "not found")
	})

	//single routes
	r.GET("/health", handler.Health)
	r.GET(":id", handler.Id)

	//group routes
	company := r.Group("/company")
	{
		company.GET("/", handler.Company)
		company.GET("/:id", handler.CompanyId)
		company.GET("/account", handler.CompanyId)
	}

	//sample auth on one of domains
	if cfg.Server.AuthEnabled {
		company.Use(gin.BasicAuth(gin.Accounts{
			cfg.Server.User: cfg.Server.Pass,
		}))
	}

	account := r.Group("/account")
	{
		account.GET("/:id/user", handler.Account)
		account.GET("/:id", handler.AccountId)
		account.GET("/", handler.Company)
	}

	tenant := r.Group("/tenant")
	{
		tenant.GET("/account/blocked", handler.Company)
	}

	var AllowedRoutes = make(map[string]bool, 0)

	routes := r.Routes()
	for _, i := range routes {
		AllowedRoutes[i.Path] = true
	}

	return r
}
