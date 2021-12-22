/**
    @Author:     ZonzeeLi
    @Project:    spikesystem
    @CreateDate: 12/21/2021
    @UpdateDate: 12/22/2021
    @Note:       Redis配置信息
**/

package config

type Redis struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                      // redis的哪个数据库
	Host     string `mapstructure:"host" json:"host" yaml:"host"`                // 服务器地址
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`                // 端口
	Password string `mapstructure:"password" json:"password" yaml:"password"`    // 密码
	PoolSize int    `mapstructure:"pool_size" json:"pool_size" yaml:"pool_size"` // 连接池
}
