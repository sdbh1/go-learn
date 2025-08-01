package handler

import (
	"sdbh/logger"

	"github.com/gin-gonic/gin"
)

func GetUIDByCtx(ctx *gin.Context) (uint, error) {
	var uid uint = 0
	var error error = nil
	var ok bool

	value, exists := ctx.Get(UID_IN_CTX)
	if !exists {
		error = logger.Error("[util][GetUIDByCtx] fail no exists")
		return uid, error
	}

	uid, ok = value.(uint)

	if !ok {
		error = logger.Error("[util][GetUIDByCtx] fail to convert value")
		return uid, error
	}
	return uid, error
}
