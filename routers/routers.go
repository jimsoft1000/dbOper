package routers

import (
	"dbOper/controller"
	"dbOper/util"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	gin.SetMode(util.ApplicationConfig.GetString("server.AppMode"))
	r := gin.Default()
	//r.Use(middleware.CORSMiddleware())

	/**
	用户User路由组
	*/
	userGroup := r.Group("user")
	{
		//用户登录
		userGroup.POST("/login", controller.Login)
		//增加用户User
		userGroup.POST("/users", controller.CreateUser)
		//查看所有的User
		userGroup.GET("/users", controller.GetUserList)
		//修改某个User
		userGroup.PUT("/users/:id", controller.UpdateUser)
		//删除某个User
		userGroup.DELETE("/users/:id", controller.DeleteUserById)
	}

	return r
}
