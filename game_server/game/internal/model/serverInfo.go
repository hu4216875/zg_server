package model

import (
	"server/conf"
	"time"
)

type ServerInfo struct {
	ServerId   uint32
	RegistNum  int64
	UpdateTime uint32
	CreateTime uint32
	LimitNum   uint32
}

func NewServerInfo(serverId uint32) *ServerInfo {
	curTime := uint32(time.Now().Unix())
	return &ServerInfo{
		ServerId:   serverId,
		UpdateTime: curTime,
		CreateTime: curTime,
		LimitNum: conf.Server.MaxRegistNum,
	}
}
