package router

import (
	"easy_note/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoutersInit() *gin.Engine{
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", index)

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
	}
	return r
}

func index(c *gin.Context){
	c.HTML(http.StatusOK, "index_view.html",nil)
}