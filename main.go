package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/user/list", userListHandler)
	engine.Run(":8080")
}

// userListHandler /hello请求处理函数
func userListHandler(c *gin.Context) {
	userList, err := getUserList()
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
		return
	}
	c.JSON(http.StatusOK, userList)
	return
}
