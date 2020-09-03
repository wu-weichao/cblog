package main

import (
	"cblog/pkg/setting"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	_ "cblog/models"
	"cblog/routers"
)

// @title cblog
// @version 1.0
// @description 基于gin的...
func main() {
	// debug or release
	gin.SetMode(setting.ServerSetting.RunMode)

	// 使用自定义路由
	r := routers.InitRouter()
	addr := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	// 使用 go 默认的 server 设置超时
	server := &http.Server{
		Addr:           addr,
		Handler:        r,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	log.Printf("[info] start http server listening %s", addr)
	server.ListenAndServe()

	//r.Run(host) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
