package service

import (
	"GOproject/GIT/memory_note/model"
	"GOproject/GIT/memory_note/serializer"
	"time"
)

type CreateTaskService struct {
	Title   string `json:"title"  form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status"  form:"status"`
}

func (s *CreateTaskService) Creat(id uint) serializer.Response {
	var user model.User
	code := 200
	model.DB.First(&user)
	task := model.Task{
		User:      user,
		Uid:       user.ID,
		Title:     s.Title,
		Status:    0,
		Content:   s.Content,
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
