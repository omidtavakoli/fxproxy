package main

import (
	"app/internal/http/rest"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter(handler *rest.Handler, cfg *MainConfig) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(CORSMiddleware())

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, "not found")
	})

	r.GET("/health", handler.Health)

	v1 := r.Group("/v1")
	if cfg.Server.AuthEnabled {
		v1.Use(gin.BasicAuth(gin.Accounts{
			cfg.Server.User: cfg.Server.Pass,
		}))
	}

	{
		//v1.GET("/company/", handler.Company)
	}

	var AllowedRoutes = make(map[string]bool, 0)

	routes := r.Routes()
	for _, i := range routes {
		AllowedRoutes[i.Path] = true
	}

	return r
}
