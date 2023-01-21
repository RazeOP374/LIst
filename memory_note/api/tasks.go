package api

import (
	"GOproject/GIT/memory_note/pkg/utils"
	"GOproject/GIT/memory_note/service"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var createTask service.CreateTaskService
	chaim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	err := c.ShouldBind(&createTask)
	if err == nil {
		res := createTask.Creat(chaim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
