package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/user/list", userListHandler)
	engine.GET("/school/list", schoolListHandler)
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

type u struct {
	Name string `json:"name" gorm:"name"`
}

// schoolListHandler /hello请求处理函数
func schoolListHandler(c *gin.Context) {
	users := make([]*u, 0)
	z := DBClient.Table("user").Find(&users)
	if z.Error != nil {
		c.JSON(http.StatusOK, z.Error)
		return
	}
	c.JSON(http.StatusOK, users)
	return
}
