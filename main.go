/**
    @Author:     ZonzeeLi
    @Project:    spikesystem
    @CreateDate: 12/21/2021
    @UpdateDate: xxx
    @Note:       spikeSystem

**/

package main

import (
	"fmt"
	"spikesystem/core"
	"spikesystem/dao/redis"

	"spikesystem/global"
)

func main() {
	fmt.Println(1)
	// 先写入配置文件viper
	global.Spike_VP = core.Viper("config.yaml")
	global.Spike_LOG = core.Zap()
	// 初始化redis
	redis.Redis()
	// 端口启动
	core.RunServer()
}
