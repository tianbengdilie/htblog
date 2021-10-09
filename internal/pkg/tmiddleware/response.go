package tmiddleware

import (
	"bytes"
	"fmt"
	"htblog/internal/pkg/tginutils"

	"github.com/gin-gonic/gin"
)

// MiddleResponseConfig response config
type MiddleResponseConfig struct {
	MaxResponseLogLength int
}

// mRespBodyLogWriter gin response write
type mRespBodyLogWriter struct {
	ctx *gin.Context
	gin.ResponseWriter
	body   *bytes.Buffer
	config *MiddleResponseConfig
}

// Write gin write response
func (w mRespBodyLogWriter) Write(b []byte) (int, error) {
	xlog := tginutils.GetLog(w.ctx)
	logContent := ""
	statusCode := w.ResponseWriter.Status()
	respStr := ""
	if w.config.MaxResponseLogLength > 0 && len(b) > w.config.MaxResponseLogLength {
		respStr = string(b[0:w.config.MaxResponseLogLength]) + "..."
	} else {
		respStr = string(b)
	}
	logContent = fmt.Sprintf("response: %s, status: %d", respStr, statusCode)

	if logContent != "" {
		xlog.Info(logContent)
	}
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *mRespBodyLogWriter) WriteString(s string) (int, error) {
	return w.Write([]byte(s))
}

func RegisterResponse(config *MiddleResponseConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &mRespBodyLogWriter{
			ctx:            c,
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
			config:         config,
		}
		c.Writer = blw
		c.Next()
	}
}
