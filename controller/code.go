/**
    @Author:     ZonzeeLi
    @Project:    spikesystem
    @CreateDate: 12/21/2021
    @UpdateDate: 12/22/2021
    @Note:       状态码信息
**/

package controller

type ResCode int64

const (
	CodeFailed ResCode = iota
	CodeSuccess
	CodeServerBusy = 404
)

var codeMsgMap = map[ResCode]string{
	CodeFailed:     "已售磐",
	CodeSuccess:    "抢票成功",
	CodeServerBusy: "服务繁忙",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
