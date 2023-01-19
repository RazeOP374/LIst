package api

import (
	"GOproject/GIT/memory_note/service"
	"github.com/gin-gonic/gin"
)

func UserRegister(e *gin.Context) {
	var userRegiseter service.UserService
	if err := e.ShouldBind(&userRegiseter); err == nil {
		res := userRegiseter.Register()
		e.JSON(200, res)
	} else {
		e.JSON(400, err)
	}
}
func UserLogin(e *gin.Context) {
	var userLoginSer service.UserService
	if err := e.ShouldBind(&userLoginSer); err == nil {
		res := userLoginSer.Login()
		e.JSON(200, res)
	} else {
		e.JSON(400, err)
	}

}
