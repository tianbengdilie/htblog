package helper

import (
	"htblog/internal/pkg/tlog"
	"htblog/internal/pkg/utils"

	"github.com/gin-gonic/gin"
)

func GetLog(ctx *gin.Context) *tlog.Logger {
	val := ctx.Value("logger")
	if val == nil {
		val = tlog.Tag(utils.GetRandomStr(12))
		ctx.Set("logger", val)
	}
	return val.(*tlog.Logger)
}
