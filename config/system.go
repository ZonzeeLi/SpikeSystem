/**
    @Author:     ZonzeeLi
    @Project:    spikesystem
    @CreateDate: 12/21/2021
    @UpdateDate: 12/22/2021
    @Note:       端口配置信息
**/

package config

type System struct {
	Name string `mapstructrue:"name" json:"name" yaml:"name"`
	Port int    `mapstructrue:"port" json:"port" yaml:"port"`
}
