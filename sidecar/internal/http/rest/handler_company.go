package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Company(c *gin.Context) {

	c.JSON(http.StatusOK, "ok")
}
