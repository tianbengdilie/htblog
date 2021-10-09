package main

import (
	"context"
	"fmt"
	"htblog/internal/api"
	"htblog/internal/config"
	"htblog/internal/pkg/tconfig"
	"htblog/internal/pkg/tlog"
	"htblog/internal/pkg/tmiddleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	config.Init()
	tlog.Init(config.Config().GetStringSlice("log.hooks"))
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			tlog.Tag("app.recovery").Error("Error: ", err, "; stack: ", string(debug.Stack()))
		}
	}()

	server := initServer()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Fatal("srv.Shutdown", err) // failure/timeout shutting down the server gracefully
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func initServer() *http.Server {
	// Set gin mode.
	gin.SetMode(config.Config().GetString("ginmode"))

	r := gin.New()
	// r.Use(gin.RecoveryWithWriter(&logWriter{
	// 	monitor: config.GetPanicMonitor(),
	// }))

	host := config.Config().GetString("apiaddr")
	fmt.Println("******* WebServer Starting: " + host + " *******")
	server := &http.Server{
		Addr:         host,
		Handler:      r,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	if config.Config().GetString("APPENV") != tconfig.EnvProduction {
		r.Use(tmiddleware.RegisterCors())
	}
	r.Use(tmiddleware.RegisterRequest(&tmiddleware.MiddleRequestConfig{
		LogHeaders:    true,
		LogAllCookies: true,
	}))
	r.Use(tmiddleware.RegisterResponse(&tmiddleware.MiddleResponseConfig{
		MaxResponseLogLength: 1000,
	}))
	// r.Use(tmiddleware.GinRecovery())

	//初始化API路由
	api.Init(r)

	// 初始化前端页面 发布的时候才能解开下面的注释
	// r.LoadHTMLGlob("./frontend/dist/*.html")
	// r.StaticFile("/", "./frontend/dist/index.html")

	return server
}
