package service

import (
	"GOproject/GIT/memory_note/model"
	"GOproject/GIT/memory_note/serializer"
	"time"
)

var code = 200

type CreateTaskService struct {
	Title   string `json:"title"  form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status"  form:"status"`
}
type ShowTaskService struct {
}
type ListTaskService struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}

func (l *ListTaskService) List(uid uint) serializer.Response {
	var tasks []model.Task
	count := 0
	if l.PageSize == 0 {
		l.PageSize = 15
	}
	model.DB.Model(model.Task{}).Preload("User").Where("uid = ?", uid).Count(&count).
		Limit(l.PageSize).Offset((l.PageNum - 1) * l.PageSize).Find(&tasks)
	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count))

}
func (s *ShowTaskService) Show(id string) serializer.Response {
	var task model.Task

	err := model.DB.First(&task, id).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status:  code,
			Message: "查询失败",
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildTask(task),
	}
}

func (c *CreateTaskService) Creat(id uint) serializer.Response {
	var user model.User
	model.DB.First(&user, id)
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     c.Title,
		Status:    0,
		Content:   c.Content,
		StartTime: time.Now().Unix(),
	}
	err := model.DB.Create(&task).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status:  code,
			Message: "创建失败",
		}
	}
	return serializer.Response{
		Status:  code,
		Message: "创建成功",
	}
}
