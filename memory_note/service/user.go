package service

import (
	"GOproject/GIT/memory_note/model"
	"GOproject/GIT/memory_note/pkg/utils"
	"GOproject/GIT/memory_note/serializer"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=8"`
	Password string `form:"password"  json:"password" binding:"required,min=5,max=11"`
}

func (Useas *UserService) Login() serializer.Response {
	var user model.User
	if err := model.DB.Where("user_name=?", Useas.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serializer.Response{
				Status:  400,
				Message: "用户不存在",
			}
		}
		return serializer.Response{
			Status:  500,
			Message: "数据库错误",
		}
	}
	if user.CheckPassword(Useas.Password) == false {
		return serializer.Response{
			Status:  400,
			Message: "密码错误",
		}
	}
	token, err := utils.GenerateToken(user.ID, Useas.UserName)
	if err != nil {
		return serializer.Response{
			Status:  500,
			Message: "TOKEN签发错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Data: serializer.TokenData{
			User:  serializer.BuildUser(user),
			Token: token,
		},
		Message: "登陆成功",
	}
}
func (Useas *UserService) Register() serializer.Response {
	var user model.User
	var count int
	model.DB.Model(&model.User{}).Where("user_name=?", Useas.UserName).First(&user).Count(&count)
	if count == 1 {
		return serializer.Response{
			Status:  400,
			Message: "此用户已存在",
		}
	}
	user.UserName = Useas.UserName
	if err := user.SetPassword(Useas.Password); err != nil {
		return serializer.Response{
			Status:  400,
			Message: err.Error(),
		}
	}
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Status:  500,
			Message: "数据库错误",
		}
	}
	return serializer.Response{
		Status:  200,
		Message: "用户注册成功",
	}
}
