package tmiddleware

import (
	"bytes"
	"fmt"
	"htblog/internal/pkg/tginutils"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// MiddleRequestConfig request config
type MiddleRequestConfig struct {
	LogHeaders    bool
	LogAllCookies bool
}

func RegisterRequest(config *MiddleRequestConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		logContent := ""
		url := c.Request.URL.String()
		body, _ := c.GetRawData()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		reqStr := ""
		if len(body) > 0 {
			reqStr = "; data: " + string(body)
		}
		if config.LogAllCookies {
			reqStr = reqStr + "; cookies: " + c.GetHeader("cookie")
		} else { //log start_token
			startToken, err := c.Cookie("start_token")
			if err == nil || startToken != "" {
				reqStr = reqStr + "; start_token: " + startToken
			}
		}
		if config.LogHeaders {
			reqStr = fmt.Sprintf("%s;  headers: %+v", reqStr, c.Request.Header)
		}

		logContent = fmt.Sprintf("request[%s]: %s %s; ip: %s", c.Request.Method, url, reqStr, c.ClientIP())

		logger := tginutils.GetLog(c)
		if logContent != "" {
			logger.Info(logContent)
		}

		if strings.Contains(path, "//") {
			logger.Warn("Non-standard url: ", path)
		}

		timer := time.AfterFunc(20*time.Second, func() {
			logger.Error("request[", c.Request.Method, "]: execute timeout 20s")
		})
		start := time.Now()

		c.Next()
		timer.Stop()

		//计算耗时
		latency := time.Since(start)
		logger.Infof("request[%s]: %s cost time %v", c.Request.Method, url, latency)
	}
}
