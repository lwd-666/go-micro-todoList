package handlers

import (
	"api-gateway/pkg/utils"
	"api-gateway/services"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

//用户注册
func UserRegister(c *gin.Context) {
	var userReg services.UserRequest
	PanicIfUserError(c.Bind(&userReg))
	//从gin.Keys中取出服务实例
	userService := c.Keys["userService"].(services.UserService)
	userResp, err := userService.UserRegister(context.Background(), &userReg)

	PanicIfUserError(err)
	c.JSON(http.StatusOK, gin.H{
		"data": userResp,
	})
}

//用户登录
func UserLogin(c *gin.Context) {
	var userReg services.UserRequest
	PanicIfUserError(c.Bind(&userReg))
	//从gin.Keys中取出服务实例
	userService := c.Keys["userService"].(services.UserService)
	userResp, err := userService.UserLogin(context.Background(), &userReg)
	PanicIfUserError(err)
	token, err := utils.GenerateToken(uint(userResp.UserDetail.ID))
	if err != nil {

	}
	c.JSON(http.StatusOK, gin.H{
		"code": userResp.Code,
		"msg":  "成功",
		"data": gin.H{
			"user":  userResp.UserDetail,
			"token": token,
		},
	})
}
