/**
    @Author:     ZonzeeLi
    @Project:    spikesystem
    @CreateDate: 12/21/2021
    @UpdateDate: 12/22/2021
    @Note:       服务配置信息
**/

package config

type Server struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
}
