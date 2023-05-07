package shop

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func shopHandler(c *gin.Context) {
	c.String(http.StatusOK, "shopHandler")
}

func shopInfoHandler(c *gin.Context) {
	c.String(http.StatusOK, "shopInfoHandler")
}
