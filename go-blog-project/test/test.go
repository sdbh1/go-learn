package test

import (
	"sdbh/config"
	database "sdbh/database/gorm"
	"sdbh/global"
	"sdbh/redis"
	"sdbh/router"
)

func InitAll() {
	global.Config = config.InitConfig()
	global.Redis = redis.InitRedis()
	global.BlogDB = database.InitBlogDB()
	global.Engine = router.InitRouter()
	// // 启动服务
	// if err := global.Engine.Run("127.0.0.1:5678"); err != nil {
	// 	panic("启动服务失败: " + err.Error())
	// }
}
