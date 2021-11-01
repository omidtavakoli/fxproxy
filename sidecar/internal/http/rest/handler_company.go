package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func (h *Handler) Company(c *gin.Context) {
	resp, err := h.ProxyService.Company()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "company func err")
	}
	c.JSON(http.StatusOK, resp)
}
func (h *Handler) CompanyId(c *gin.Context) {
	c.JSON(http.StatusOK, "CompanyId ok")
}
func (h *Handler) CompanyIdTest(c *gin.Context) {
	c.JSON(http.StatusOK, "CompanyIdTest ok")
}

func (h *Handler) CompanyAccount(c *gin.Context) {
	c.JSON(http.StatusOK, "CompanyAccount ok")
}
func (h *Handler) Tenant(c *gin.Context) {
	c.JSON(http.StatusOK, "Tenant ok")
}
func (h *Handler) Account(c *gin.Context) {
	param := c.Param("id")
	account, err := regexp.Compile("^acc*.")
	if err != nil {
		c.JSON(http.StatusInternalServerError, "regex err")
	}
	validId := account.MatchString(param)
	if !validId {
		c.JSON(http.StatusNotFound, "not found")
	}
	c.JSON(http.StatusOK, "Account ok")
}

func (h *Handler) Id(c *gin.Context) {
	c.JSON(http.StatusOK, "Id ok")
}
func (h *Handler) AccountId(c *gin.Context) {
	param := c.Param("id")
	account, err := regexp.Compile("^acc*.")
	if err != nil {
		c.JSON(http.StatusInternalServerError, "regex err")
	}
	validId := account.MatchString(param)
	if !validId {
		c.JSON(http.StatusNotFound, "not found")
	}
	c.JSON(http.StatusOK, "AccountId ok")
}
