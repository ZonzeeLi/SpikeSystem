/**
    @Author:     ZonzeeLi
    @Project:    spikesystem
    @CreateDate: 12/21/2021
    @UpdateDate: 12/22/2021
    @Note:       路由创建
**/

package routers

import (
	"spikesystem/controller"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.Default()
	//r.Use(core.GinLogger(), core.GinRecovery(true))
	r.GET("/buy/ticket", controller.Spike)
	return r
}
