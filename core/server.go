/**
    @Author:     ZonzeeLi
    @Project:    spikesystem
    @CreateDate: 12/21/2021
    @UpdateDate: 12/22/2021
    @Note:       服务启动监听
**/

package core

import (
	"fmt"
	"net/http"
	"spikesystem/global"
	"spikesystem/routers"
	"time"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type server interface {
	ListenAndServe() error
}

func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func RunServer() {
	router := routers.Routers()
	address := fmt.Sprintf(":%d", global.Spike_CONFIG.System.Port)
	s := initServer(address, router)
	global.Spike_LOG.Info("server run success on ", zap.String("address", address))
	global.Spike_LOG.Error(s.ListenAndServe().Error())
}
