package handle

import (
	"server/game/internal/common"
	"server/game/internal/service"
	"server/msg"
)

// RequestOreInfoHandle 请求挖矿信息
func RequestOreInfoHandle(args interface{}, playerData *common.PlayerData) {
	retMsg := &msg.ResponseOreInfo{
		Result: int32(msg.ErrCode_SUCC),
	}

	oreInfo := playerData.AccountInfo.OreInfo
	if oreInfo != nil {
		retMsg.OreId = oreInfo.OreId
		retMsg.StartTime = oreInfo.StartTime
		retMsg.Speed = oreInfo.Speed
	}
	playerData.PlayerAgent.WriteMsg(retMsg)
}

// RequestOreTotalHandle 请求矿洞总量
func RequestOreTotalHandle(args interface{}, playerData *common.PlayerData) {
	service.ServMgr.GetOreService().GetOreTotal(playerData)
}

// RequestStartOreHandle 开始挖矿
func RequestStartOreHandle(args interface{}, playerData *common.PlayerData) {
	if res := service.ServMgr.GetOreService().StartOre(playerData.AccountInfo.AccountId); res != msg.ErrCode_SUCC {
		retMsg := &msg.ResponseStartOre{
			Result: int32(res),
		}
		playerData.PlayerAgent.WriteMsg(retMsg)
	}
}

// RequestEndOreHandle 结束挖矿
func RequestEndOreHandle(args interface{}, playerData *common.PlayerData) {
	retMsg := &msg.ResponseEndOre{
		Result: int32(service.ServMgr.GetOreService().EndOre(playerData.AccountInfo.AccountId)),
	}
	playerData.PlayerAgent.WriteMsg(retMsg)
}

// RequestUpgradeOreSpeedHandle 升级挖矿速度
func RequestUpgradeOreSpeedHandle(args interface{}, playerData *common.PlayerData) {
	retMsg := &msg.ResponseUpgradeOreSpeed{}
	ret := service.ServMgr.GetOreService().UpgradeOreSpeed(playerData.AccountInfo.AccountId)
	if ret != msg.ErrCode_SUCC {
		retMsg.Result = int32(ret)
		playerData.PlayerAgent.WriteMsg(retMsg)
	}
}
