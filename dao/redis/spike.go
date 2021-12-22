/**
    @Author:     ZonzeeLi
    @Project:    spikesystem
    @CreateDate: 12/21/2021
    @UpdateDate: 12/21/2021
    @Note:       Redis处理函数
**/

package redis

import (
	"fmt"
	"spikesystem/global"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type RemoteSpikeKeys struct {
	SpikeOrderHashKey  string //redis中秒杀订单hash结构key
	TotalInventoryKey  string //hash结构中总订单库存key
	QuantityOfOrderKey string //hash结构中已有订单数量key
}

const LuaScript = `
        local ticket_key = KEYS[1]
        local ticket_total_key = ARGV[1]
        local ticket_sold_key = ARGV[2]
        local ticket_total_nums = tonumber(redis.call('HGET', ticket_key, ticket_total_key))
        local ticket_sold_nums = tonumber(redis.call('HGET', ticket_key, ticket_sold_key))
		-- 查看是否还有余票,增加订单数量,返回结果值
        if(ticket_total_nums >= ticket_sold_nums) then
            return redis.call('HINCRBY', ticket_key, ticket_sold_key, 1)
        end
        return 0
`

//远端统一扣库存
func (RemoteSpikeKeys *RemoteSpikeKeys) RemoteDeductionStock(client *redis.Client) bool {
	fmt.Println(RemoteSpikeKeys.SpikeOrderHashKey, RemoteSpikeKeys.QuantityOfOrderKey, RemoteSpikeKeys.TotalInventoryKey)
	ctx := client.Context()
	lua := redis.NewScript(LuaScript)
	keys := []string{RemoteSpikeKeys.SpikeOrderHashKey}
	values := []interface{}{RemoteSpikeKeys.TotalInventoryKey, RemoteSpikeKeys.QuantityOfOrderKey}
	fmt.Println(1)
	fmt.Println(keys, values)
	if result, err := lua.Run(ctx, client, keys, values).Int(); err != nil {
		global.Spike_LOG.Error("lua.Run(ctx, client, keys, values).Int() failed", zap.Error(err))
	} else {
		return result != 0
	}
	return false
}
