package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func userHandler(c *gin.Context) {
	c.String(http.StatusOK, "userHandler")
}

func userInfoHandler(c *gin.Context) {
	c.String(http.StatusOK, "userInfoHandler")
}
