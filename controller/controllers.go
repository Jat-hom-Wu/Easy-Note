package controller

import (
	"easy_note/model"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetToDoList(c *gin.Context) {
	r := model.GetAllData()
	c.JSON(http.StatusOK, r)
}

func CreateToDoList(c *gin.Context) {
	var todo model.ToDo
	c.BindJSON(&todo)
	fmt.Println(todo)
	r := model.CreateDataList(&todo)
	if r == 1 {
		c.JSON(http.StatusOK, todo)
	} else {
		c.JSON(http.StatusOK, gin.H{"error": "create data error"})
	}
}

func UpdateToDoList(c *gin.Context) {
	//直接修改就好了... (似乎不需要查询？)
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	//前端逻辑怪怪的...
	//原逻辑：先找出id，再通过save的api修改

	var todo model.ToDo
	c.BindJSON(&todo)
	//前端发过来直接赋值
	_status := todo.Status
	_id,_ := strconv.Atoi(id)
	model.UpdateOneData(_id, _status)
	c.JSON(http.StatusOK, todo)
}

func DeleteToDoList(c *gin.Context) {
	//直接删除就好了...
	id, ok := c.Params.Get("id")
	_id, _ := strconv.Atoi(id)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}
	model.DeleteOneData(_id)
	c.JSON(http.StatusOK, gin.H{id: "deleted"})
}
