package model

type ServerInfo struct {
	ServerId    uint32            `json:"serverId"`    //
	Online      []*GsServerOnline `json:"online"`      // 在线人数
	OnlineLimit uint32            `json:"onlineLimit"` // 在线人数上限
}

type GsServerOnline struct {
	GsId      uint32 `json:"serverId"`
	OnlineNum int32  `json:"onlineNum"`
}

func NewGsServerOnlie(gsId uint32) *GsServerOnline {
	return &GsServerOnline{GsId: gsId, OnlineNum: 1}
}

func NewServerInfo(srvId uint32, limit uint32) *ServerInfo {
	ret := &ServerInfo{
		ServerId:    srvId,
		OnlineLimit: uint32(limit),
	}

	ret.Online = make([]*GsServerOnline, 0, 0)
	return ret
}
