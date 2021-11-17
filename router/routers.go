package router

import (
	"easy_note/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoutersInit() *gin.Engine{
	r := gin.Default()
	controller.Rhelper = r
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/index", index)

	v1Group := r.Group("v1")
	{
		//get
		v1Group.GET("/todo", controller.GetToDoList)
		//create
		v1Group.POST("/todo",controller.CreateToDoList)
		//update
		v1Group.PUT("/todo/:id",controller.UpdateToDoList)
		//delete
		v1Group.DELETE("/todo/:id", controller.DeleteToDoList)

		//起始页
		v1Group.GET("/", controller.JudgePage)	
		// 注册
		v1Group.POST("/0", controller.RegisterPage)
		// 登录
		v1Group.POST("/1", controller.LogPage)
		// 登录检测
		v1Group.GET("/2", controller.LogErrorPage)
		// 注册检测
		v1Group.GET("/3", controller.RegisterErrorPage)
		// 用户登录处理
		v1Group.POST("/2CGISQL.cgi", controller.LogInHandler)
		//用户注册处理
		v1Group.POST("3CGISQL.cgi", controller.RegisterHandler)
	}


	return r
}

func index(c *gin.Context){
	c.HTML(http.StatusOK, "index_view.html",nil)
}