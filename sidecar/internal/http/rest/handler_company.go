package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func (h *Handler) Company(c *gin.Context) {
	fmt.Println(c.Param("action"))
	c.JSON(http.StatusOK, "ok")
}
func (h *Handler) CompanyId(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
func (h *Handler) CompanyIdTest(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}

func (h *Handler) CompanyAccount(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
func (h *Handler) Tenant(c *gin.Context) {

	c.JSON(http.StatusOK, "ok")
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
	c.JSON(http.StatusOK, "ok")
}
func (h *Handler) Id(c *gin.Context) {

	c.JSON(http.StatusOK, "ok")
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

	c.JSON(http.StatusOK, "ok")
}
