package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web-chat/auth"
	"web-chat/model"
	"web-chat/utils"
)

//var rdb = database.InitDB()

// 注册
func RegisterHandler(ctx *gin.Context) {
	user := model.UserInfo{}
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}
	exist, _ := model.CheckUserIfExist(user)
	if !exist {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2009,
			"msg":  "该用户名已存在，请勿重复注册",
		})
		return
	}
	// 加密
	user.Password = utils.Base64Encode([]byte(user.Password))
	id, err := model.AddUser(user)
	if id > 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "注册成功",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 2002,
		"msg":  "注册失败",
	})
	return
}

// 登录认证
func AuthHandler(ctx *gin.Context) {
	var user model.UserInfo
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}
	// 检查该用户是否已经存在
	exist, _ := model.CheckUserIfExist(user)
	if exist {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2006,
			"msg":  "该用户不存在",
		})
		return
	}
	// 检查密码是否正确
	user.Password = utils.Base64Encode([]byte(user.Password))
	e2 := model.CheckPassword(user)
	if e2 != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2010,
			"msg":  "密码错误",
		})
		return
	}
	// 生成token
	token, _ := auth.GetToken(user.Username)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "登录成功",
		"data": gin.H{"token": token},
	})
	return
}
