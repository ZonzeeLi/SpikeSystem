/**
    @Author:     ZonzeeLi
    @Project:    spikesystem
    @CreateDate: 12/21/2021
    @UpdateDate: 12/21/2021
    @Note:       Redis初始化
**/

package redis

import (
	"context"
	"fmt"
	"spikesystem/global"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.Spike_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			redisCfg.Host,
			redisCfg.Port,
		),
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
		PoolSize: redisCfg.PoolSize,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.Spike_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.Spike_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.Spike_REDIS = client
	}
}
