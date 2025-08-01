package main

import (
	"log/slog"
	"os"
	"os/signal"
	"sdbh/config"
	database "sdbh/database/gorm"
	"sdbh/global"
	"sdbh/redis"
	"sdbh/router"
	"syscall"
)

func main() {
	InitAll()
	go ListenTermSignal() //监听信号
}

func InitAll() {
	global.Config = config.InitConfig()
	global.Redis = redis.InitRedis()
	global.BlogDB = database.InitBlogDB()
	global.Engine = router.InitRouter()

	slog.Info("初始化成功")

	// 启动服务
	if err := global.Engine.Run("127.0.0.1:5678"); err != nil {
		panic("启动服务失败: " + err.Error())
	}
	// gin.SetMode(gin.ReleaseMode) //GIN线上发布模式
	// gin.DefaultWriter = io.Discard //禁用GIN日志
}

// 通过kill -2 PID或kill -15 PID杀死进程时，做一些收尾工作
func ListenTermSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM) //注册信号2和15。Ctrl+C对应SIGINT信号
	sig := <-c                                        //阻塞，直到信号的到来
	slog.Info("receive term signal " + sig.String() + ", going to exit")
	database.CloseBlogDB() //关闭数据库连接
	os.Exit(0)             //进程退出
}
