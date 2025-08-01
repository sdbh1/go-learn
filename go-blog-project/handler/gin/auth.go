package handler

import (
	"net/http"
	"sdbh/global"
	"sdbh/util"

	"github.com/gin-gonic/gin"
)

const (
	UID_IN_TOKEN = "uid"
	UID_IN_CTX   = "uid"
	COOKIE_NAME  = "jwt"
)

// 身份认证中间件，先确保是登录状态
func Auth(ctx *gin.Context) {
	loginUid := GetLoginUid(ctx)
	if loginUid <= 0 {
		ctx.Redirect(http.StatusTemporaryRedirect, "/login") //重定向到登录页面
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		ctx.Abort()
	} else {
		ctx.Set(UID_IN_CTX, loginUid) //把登录的uid放入ctx中
	}
}

// 从cookie里取出jwt，从而取出uid
func GetLoginUid(ctx *gin.Context) int {
	//依靠浏览器自动回传的cookie，提取出jwt token
	token := ""
	for _, cookie := range ctx.Request.Cookies() {
		if cookie.Name == COOKIE_NAME {
			token = cookie.Value
		}
	}
	return GetUidFromJwt(token)
}

// 从jwt里取出uid
func GetUidFromJwt(token string) int {
	_, payload, err := util.VerifyJwt(token, global.Config.JWT.Serect)
	if err != nil {
		return 0
	}
	for k, v := range payload.UserDefined {
		if k == UID_IN_TOKEN {
			return int(v.(float64))
		}
	}
	return 0
}
