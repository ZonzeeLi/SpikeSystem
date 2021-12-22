/**
    @Author:     ZonzeeLi
    @Project:    spikesystem
    @CreateDate: 12/21/2021
    @UpdateDate: 12/21/2021
    @Note:       全局变量
**/

package global

import (
	"spikesystem/config"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	Spike_VP     *viper.Viper
	Spike_REDIS  *redis.Client
	Spike_LOG    *zap.Logger
	Spike_CONFIG config.Server
)
