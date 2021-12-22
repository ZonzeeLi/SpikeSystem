/**
    @Author:     ZonzeeLi
    @Project:    spikesystem
    @CreateDate: 12/21/2021
    @UpdateDate: 12/22/2021
    @Note:       viper读配置
**/

package core

import (
	"fmt"
	"spikesystem/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper(filePath string) *viper.Viper {
	vp := viper.New()
	vp.SetConfigFile(filePath)
	err := vp.ReadInConfig() // 读取配置信息
	if err != nil {
		// 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 把读取到的配置信息反序列化到 Conf 变量中
	if err := vp.Unmarshal(&global.Spike_CONFIG); err != nil {
		fmt.Printf("vp.Unmarshal failed, err: %v\n", err)
	}
	vp.WatchConfig()
	vp.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了...")
		if err := vp.Unmarshal(global.Spike_CONFIG); err != nil {
			fmt.Printf("vp.Unmarshal failed, err:%v\n", err)
		}
	})
	return vp
}
