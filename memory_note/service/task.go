package service

import (
	"GOproject/GIT/memory_note/model"
	"GOproject/GIT/memory_note/serializer"
	"time"
)

type DeleteTaskService struct {
}

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

type UpdateService struct {
	Title   string `json:"title"  form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status"  form:"status"`
}

type SearchService struct {
	Info     string `json:"info" form:"info" `
	PageNum  int    `json:"page_num" form:"page_num"`
	PageSize int    `json:"page_size" form:"page_size"`
}

func (d *DeleteTaskService) Delete(id string) serializer.Response {
	var task model.Task
	err := model.DB.Delete(&task, id).Error
	code := 200
	if err != nil {
		return serializer.Response{
			Status:  500,
			Message: "删除失败",
		}
	}
	return serializer.Response{
		Status:  code,
		Message: "删除成功",
	}
}

func (u *UpdateService) Update(id string) serializer.Response {
	var task model.Task
	model.DB.First(&task, id)
	code := 200
	task.Content = u.Content
	task.Title = u.Title
	task.Status = u.Status
	err := model.DB.Save(&task).Error
	if err != nil {
		return serializer.Response{
			Status:  500,
			Message: "保存错误",
		}
	}
	return serializer.Response{Status: code, Data: serializer.BuildTask(task), Message: "更新完成"}
}

func (l *ListTaskService) List(id uint) serializer.Response {
	var tasks []model.Task
	count := 0
	if l.PageSize == 0 {
		l.PageSize = 15
	}
	model.DB.Model(model.Task{}).Preload("User").Where("uid=?", id).Count(&count).
		Limit(l.PageSize).Offset((l.PageNum - 1) * l.PageSize).Find(&tasks)
	return serializer.BuildListResponse(serializer.BuildTasks(tasks), uint(count))

}

func (s *ShowTaskService) Show(id string) serializer.Response {
	var task model.Task
	code := 200
	err := model.DB.First(&task, id).Error
	if err != nil {
		return serializer.Response{
			Status:  500,
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
	code := 200
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
		return serializer.Response{
			Status:  500,
			Message: "创建失败",
		}
	}
	return serializer.Response{
		Status:  code,
		Message: "创建成功",
	}
}

func (s *SearchService) Search(id uint) serializer.Response {
	var task []model.Task
	count := 0
	if s.PageSize == 0 {
		s.PageSize = 15
	}
	model.DB.Model(model.Task{}).Preload("User").Where("uid=?", id).
		Where("title LIKE ? OR content LIKE ?", "%"+s.Info+"%", "%"+s.Info+"%").Count(&count).
		Limit(s.PageSize).Offset((s.PageNum - 1) * s.PageSize).Find(&task)
	return serializer.BuildListResponse(serializer.BuildTasks(task), uint(count))
}
