package api

import (
	"GOproject/GIT/memory_note/pkg/utils"
	"GOproject/GIT/memory_note/service"
	"github.com/gin-gonic/gin"
)

func SearchTask(c *gin.Context) {
	var searchTask service.SearchService
	chaim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	err := c.ShouldBind(&searchTask)
	if err == nil {
		res := searchTask.Search(chaim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}

}

func UpdateTask(c *gin.Context) {
	var updateTask service.UpdateService
	err := c.ShouldBind(&updateTask)
	if err == nil {
		res := updateTask.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

func CreateTask(c *gin.Context) {
	var createTask service.CreateTaskService
	chaim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	err := c.ShouldBind(&createTask)
	if err == nil {
		res := createTask.Creat(chaim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

func ShowTask(c *gin.Context) {
	var showTask service.ShowTaskService
	err := c.ShouldBind(&showTask)
	if err == nil {
		res := showTask.Show(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

func ListTask(c *gin.Context) {
	var listTask service.ListTaskService
	chaim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	err := c.ShouldBind(&listTask)
	if err == nil {
		res := listTask.List(chaim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}
}

func DeleteTask(c *gin.Context) {
	var deleteTask service.DeleteTaskService
	err := c.ShouldBind(&deleteTask)
	if err == nil {
		res := deleteTask.Delete(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(400, ErrorResponse(err))
	}

}
