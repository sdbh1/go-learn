package handler

import (
	"log/slog"
	"net/http"
	database "sdbh/database/gorm"
	"sdbh/global"
	util "sdbh/util"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	COOKIE_LIFE    = 7 * 86400
	COOKIE_EXPIRED = -1
)

func Login(ctx *gin.Context) {

	var input struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindBodyWithJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	userName := input.UserName

	user := database.GetUserByName(userName, global.BlogDB)

	if user == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "user not found"})
		return
	}

	if user.Password != input.Password {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "password not match"})
		return
	}

	//登录成功，返回cookie
	header := util.DefautHeader

	payload := util.JwtPayload{ //payload以明文形式编码在token中，server用自己的密钥可以校验该信息是否被篡改过
		Issue:       global.Config.App.Name,
		IssueAt:     time.Now().Unix(),                                //因为每次的IssueAt不同，所以每次生成的token也不同
		Expiration:  time.Now().Add(COOKIE_LIFE * time.Second).Unix(), //7天后过期
		UserDefined: map[string]any{UID_IN_TOKEN: user.Id},            //用户自定义字段。如果token里包含敏感信息，请结合https使用
	}
	if token, err := util.GenJWT(header, payload, global.Config.JWT.Serect); err != nil {
		slog.Error("生成token失败", "error", err)
		ctx.String(http.StatusInternalServerError, "token生成失败")
	} else {
		//response header里会有一条 Set-Cookie: jwt=xxx; other_key=other_value，浏览器后续请求会自动把同域名下的cookie再放到request header里来，即request header里会有一条Cookie: jwt=xxx; other_key=other_value
		ctx.SetCookie(
			COOKIE_NAME,
			token,       //注意：受cookie本身的限制，这里的token不能超过4K
			COOKIE_LIFE, //maxAge，cookie的有效时间，时间单位秒。如果不设置过期时间，默认情况下关闭浏览器后cookie被删除
			"/",         //path，cookie存放目录
			"localhost", //cookie从属的域名,不区分协议和端口。如果不指定domain则默认为本host(如b.a.com)，如果指定的domain是一级域名(如a.com)，则二级域名(b.a.com)下也可以访问。访问登录页面时必须用http://localhost:5678/login，而不能用http://127.0.0.1:5678/login，否则浏览器不会保存这个cookie
			false,       //是否只能通过https访问
			true,        //设为false,允许js修改这个cookie（把它设为过期）,js就可以实现logout。如果为true，则需要由后端来重置过期时间
		)
	}
	ctx.JSON(http.StatusOK, gin.H{"token": 401, "msg": "password not match"})
}

func RegisterUser(ctx *gin.Context) {

	var input struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.RegisterUser(input.UserName, input.Password, global.BlogDB)
	if err != nil {
		slog.Error("用户注册失败", "name", input.UserName, "error", err)

		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": err.Error()})
		return
	}
	slog.Info("用户注册成功", "name", input.UserName)

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "注册成功"})
}

func LogOut(ctx *gin.Context) {
	loginUID, err := GetUIDByCtx(ctx)

	if err != nil {
		slog.Error("[DataBase][Out] user not already login :", loginUID)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": err.Error()})
		return
	}

	user := database.GetUserById(loginUID, global.BlogDB)

	if user == nil {
		slog.Error("[DataBase][Out] user not find :", "id", loginUID)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": err.Error()})
		return
	}

	ctx.SetCookie(
		COOKIE_NAME,
		"",
		COOKIE_EXPIRED,
		"/",
		"localhost",
		false,
		true,
	)
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "注销成功"})
}

func UpdatePassword(ctx *gin.Context) {
	var input struct {
		UserName string `json:"username"`
		OldVal   string `json:"old_val"`
		NewVal   string `json:"new_val"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginUID, err := GetUIDByCtx(ctx)

	if err != nil {
		slog.Error("[handler][UpdatePassword] user not already login :", loginUID)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": err.Error()})
		return
	}

	err = database.UpdateUserSingleField(loginUID, "password", input.OldVal, input.NewVal, global.BlogDB, false)

	if err != nil {
		slog.Info("[handle][UpdatePassword] fail", "error", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": err.Error()})
		return
	}

	ctx.HTML(http.StatusOK, "update_pass.html", nil)
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "Update success"})
}

func UpdateDisplayName(ctx *gin.Context) {
	var input struct {
		UserName string `json:"username"`
		OldVal   string `json:"old_val"`
		NewVal   string `json:"new_val"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginUID, err := GetUIDByCtx(ctx)

	if err != nil {
		slog.Error("[handler][UpdateDisplayName] user not already login :", loginUID)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": err.Error()})
		return
	}

	err = database.UpdateUserSingleField(loginUID, "display_name", input.OldVal, input.NewVal, global.BlogDB, false)

	if err != nil {
		slog.Info("[handle][UpdateDisplayName] fail", "error", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "Update success"})
}

func GetUserAllBriefPost(ctx *gin.Context) {
	loginUID, err := GetUIDByCtx(ctx)

	if err != nil {
		slog.Error("[handler][GetUserAllBriefPost] user not already login :", loginUID)
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "msg": err.Error()})
		return
	}

	posts, err := database.GetUserAllBriefPost(global.BlogDB, loginUID)

	if err != nil {
		slog.Error("[handler][GetUserAllBriefPost] fail", "error", err.Error())
	}

	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "Update success", "data": posts})
}
