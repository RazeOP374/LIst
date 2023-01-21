package api

import (
	"GOproject/GIT/memory_note/service"
	"github.com/gin-gonic/gin"
)

func UserRegister(e *gin.Context) {
	var userRegiseter service.UserService
	err := e.ShouldBind(&userRegiseter)
	if err == nil {
		res := userRegiseter.Register()
		e.JSON(200, res)
	} else {
		e.JSON(400, ErrorResponse(err))
	}
}
func UserLogin(e *gin.Context) {
	var userLoginSer service.UserService
	err := e.ShouldBind(&userLoginSer)
	if err == nil {
		res := userLoginSer.Login()
		e.JSON(200, res)
	} else {
		e.JSON(400, ErrorResponse(err))
	}

}
