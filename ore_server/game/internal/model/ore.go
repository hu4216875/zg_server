package model

import (
	"ore_server/util"
	"time"
)

// OrePlayer 矿区玩家
type OrePlayer struct {
	AccountId int64
	ServerId  uint32
	Speed     uint32
	StartTime uint32
}

type OreInfo struct {
	OreDistId  uint32 // 矿区Id
	EndTime    uint32 // 矿洞资源挖完的时间
	Total      uint32 // 当前总量
	Players    []*OrePlayer
	CreateTime uint32
	UpdateTime uint32
}

type OreLog struct {
	OreId   uint32
	Records []*OreRecord
}

type OreRecord struct {
	AccountId  int64
	Num        uint32
	CreateTime string
}

func NewOreInfo(id, total uint32) *OreInfo {
	ret := &OreInfo{
		OreDistId: id,
		Total:     total,
	}
	ret.CreateTime = uint32(time.Now().Unix())
	ret.Players = make([]*OrePlayer, 0, 0)
	return ret
}

func NewOrePlayer(accountId int64, serverId, speed, startTime uint32) *OrePlayer {
	return &OrePlayer{
		AccountId: accountId,
		ServerId:  serverId,
		Speed:     speed,
		StartTime: startTime,
	}
}

func NewOreLog(oreId uint32) *OreLog {
	ret := &OreLog{
		OreId: oreId,
	}
	ret.Records = make([]*OreRecord, 0, 0)
	return ret
}

func NewOreRecord(accountId int64, num uint32) *OreRecord {
	return &OreRecord{
		AccountId:  accountId,
		Num:        num,
		CreateTime: util.GetLocalTimeStr(),
	}
}
