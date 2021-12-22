/**
    @Author:     ZonzeeLi
    @Project:    spikesystem
    @CreateDate: 12/21/2021
    @UpdateDate: 12/22/2021
    @Note:       抢票系统函数，初始化设置本地库存量为150，已卖出为0，并用channel做锁来防止资源竞争。
**/

package controller

import (
	"spikesystem/dao/redis"
	"spikesystem/global"
	"spikesystem/localSpike"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	LocalSpike  localSpike.LocalSpike
	RemoteSpike redis.RemoteSpikeKeys
	done        chan int
)

func init() {
	LocalSpike = localSpike.LocalSpike{
		LocalInStock:     150,
		LocalSalesVolume: 0,
	}
	RemoteSpike = redis.RemoteSpikeKeys{
		SpikeOrderHashKey:  "ticket_hash_key",
		TotalInventoryKey:  "ticket_total_nums",
		QuantityOfOrderKey: "ticket_sold_nums",
	}
	done = make(chan int, 1)
	done <- 1
}

func Spike(c *gin.Context) {
	// 加锁
	<-done
	// 判断
	if LocalSpike.LocalDeductionStock() && RemoteSpike.RemoteDeductionStock(global.Spike_REDIS) {
		// 回复抢票成功
		global.Spike_LOG.Info("抢票成功! result:1,localSales:" + strconv.FormatInt(LocalSpike.LocalSalesVolume, 10))
		ResponseSuccess(c, CodeSuccess)
	} else {
		// 回复抢票失败
		global.Spike_LOG.Info("已售磐! result:0,localSales:" + strconv.FormatInt(LocalSpike.LocalSalesVolume, 10))
		ResponseError(c, CodeFailed)

	}
	done <- 1
	// 将消息写入日志

}
