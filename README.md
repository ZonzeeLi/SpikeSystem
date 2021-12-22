### SpikeSytem 

本项目是在 https://github.com/GuoZhaoran/spikeSystem 开源项目的基础上，主要进行一些web 开发的脚手架补充([GVA](https://www.gin-vue-admin.com/) 的脚手架方法)，内容并不多，因原项目的思想非常值得了解，所以自己才简易地二次开发了一下。

主要添加为gin框架、viper读配置、zap自定义日志、go-redis、go-wrk。原作者的思路在[README_GuoZhaoran](./README_GuoZhaoran.md)

### 操作流程

启动redis并且建立hash结构
```启动redis建立基本hash结构
hmset ticket_hash_key "ticket_total_nums" 10000 "ticket_sold_nums" 0
```
配置文件中的内容可自行修改，并且可以根据再次开发进行补充其他内容
```修改配置文件
修改config.yaml的内容，可自定义
```
启动项目
```启动项目
go run main.go
```
并发测试
```并发测试
cd go-wrk
go-wrk -c=100  -n=10000 http://127.0.0.1:3005/buy/ticket
```