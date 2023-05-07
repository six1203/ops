package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义一个UserController结构体
type UserController struct {
}

// 把如下方法绑定给UserController结构体

func (u UserController) UserHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "hello",
	})
}

func (u UserController) UserListHandler(c *gin.Context) {
	c.String(http.StatusOK, "userListHandler")
}
func (u UserController) UpdateUserHandler(c *gin.Context) {
	c.String(http.StatusOK, "updateUserHandler")
}
