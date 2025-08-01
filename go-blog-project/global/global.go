package global

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	ProjectRootPath string
	Redis           *redis.Client
	BlogDB          *gorm.DB
	Engine          *gin.Engine
	Config          *BlogServerConfig
)

type BlogServerConfig struct {
	App struct {
		Name    string
		Port    string
		LogPath string
	}
	Database struct {
		Dsn          string
		MaxIdleConns int
		MaxOpenConns int
		LogPath      string
	}
	Redis struct {
		Addr     string
		DB       int
		Password string
	}
	JWT struct {
		Serect string
	}
}
