package config

import (
	"fmt"
	"sdbh/global"
	log "sdbh/logger"
	"sdbh/util"

	"github.com/spf13/viper"
)

var BlogConfig *global.BlogServerConfig

func InitConfig() *global.BlogServerConfig {

	path, err := util.FindProjectRoot()
	if err != nil {
		fmt.Println(err) // 打印错误信息
	}

	global.ProjectRootPath = path

	viper.AddConfigPath(global.ProjectRootPath + "/config")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error reading config file: ", "error", err.Error())
	}

	BlogConfig = &global.BlogServerConfig{}

	if err := viper.Unmarshal(BlogConfig); err != nil {
		fmt.Println("Unable to decode into struct: ", "error", err.Error())
	}
	log.InitSlog(global.ProjectRootPath + BlogConfig.App.LogPath)
	return BlogConfig
}
