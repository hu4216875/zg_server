package service

import (
	"server/conf"
	"server/game/internal/common"
	"server/game/internal/dao"
	"server/game/internal/model"
	"server/msg"
	"server/publicconst"
	"server/template"
	"server/util"
)

type OreService struct {
	IService

	ore oreInfo
}

type oreInfo struct {
	oreId      uint32
	total      uint32
	endTime    uint32
	updateTime uint32
}

func NewOreService() *OreService {
	return &OreService{}
}

func (o *OreService) OnInit() {
}

func (o *OreService) OnDestory() {

}

// StartOre 开始挖矿
func (o *OreService) StartOre(accountId int64) msg.ErrCode {
	playerData := common.PlayerMgr.FindPlayerData(accountId)
	oreInfo := playerData.AccountInfo.OreInfo
	if oreInfo != nil && oreInfo.StartTime > 0 {
		return msg.ErrCode_HAS_START_ORE
	}

	oreId := template.GetSystemItemTemplate().GetOreId()
	speed := template.GetSystemItemTemplate().GetOreSpeed()
	interMsg := &msg.RequestInterAddOrePlayer{
		ServerId:  conf.Server.ServerId,
		AccountId: accountId,
		OreId:     oreId,
		Speed:     speed,
	}
	oreCli.SendMessage(uint32(msg.MsgId_ID_RequestInterAddOrePlayer), interMsg)
	return msg.ErrCode_SUCC
}

// EndOre 结束挖矿
func (o *OreService) EndOre(accountId int64) msg.ErrCode {
	playerData := common.PlayerMgr.FindPlayerData(accountId)
	oreInfo := playerData.AccountInfo.OreInfo
	if oreInfo.StartTime == 0 {
		return msg.ErrCode_NO_START_ORE
	}

	o.SettleOre(accountId, oreInfo)
	return msg.ErrCode_SUCC
}

// UpgradeOreSpeed 升级挖矿速度
func (o *OreService) UpgradeOreSpeed(accountId int64) msg.ErrCode {
	oreCli := getOreClient()
	if oreCli == nil {
		return msg.ErrCode_SYSTEM_ERROR
	}

	playerData := common.PlayerMgr.FindPlayerData(accountId)
	oreInfo := playerData.AccountInfo.OreInfo
	if oreInfo == nil || oreInfo.StartTime == 0 {
		return msg.ErrCode_NO_START_ORE
	}

	costItems := template.GetSystemItemTemplate().GetOreUpgradeSpeedCostItem()
	for i := 0; i < len(costItems); i++ {
		if !ServMgr.GetItemService().EnoughItem(accountId, costItems[i].ItemId, costItems[i].ItemNum) {
			return msg.ErrCode_NO_ENOUGH_ITEM
		}
	}

	oreId := template.GetSystemItemTemplate().GetOreId()
	addSpeed := template.GetSystemItemTemplate().GetOreAddSpeed()

	reqMsg := &msg.RequestInterUpdateOrePlayer{
		OreId:     oreId,
		AccountId: playerData.AccountInfo.AccountId,
		Speed:     playerData.AccountInfo.OreInfo.Speed + addSpeed,
	}
	oreCli.SendMessage(uint32(msg.MsgId_ID_RequestInterUpdateOrePlayer), reqMsg)
	return msg.ErrCode_SUCC
}

// GetOreTotal 获取矿洞总量
func (o *OreService) GetOreTotal(playerData *common.PlayerData) {
	retMsg := &msg.ResponseOreTotal{}
	oreCli := getOreClient()
	if oreCli == nil {
		retMsg.Result = int32(msg.ErrCode_SYSTEM_ERROR)
		playerData.PlayerAgent.WriteMsg(retMsg)
		return
	}

	curTime := util.GetCurTime()
	if int(curTime-o.ore.updateTime) < publicconst.MAX_UPDATE_ORE_TOTAL_TIME {
		retMsg.Result = int32(msg.ErrCode_SUCC)
		retMsg.Total = o.ore.total
		playerData.PlayerAgent.WriteMsg(retMsg)
		return
	}

	oreId := template.GetSystemItemTemplate().GetOreId()
	reqMsg := &msg.RequestInterGetOreInfo{
		ServerId:  conf.Server.ServerId,
		AccountId: playerData.AccountInfo.AccountId,
		OreId:     oreId,
	}
	oreCli.SendMessage(uint32(msg.MsgId_ID_RequestInterGetOreInfo), reqMsg)
}

// SettleOre 结算挖矿
func (o *OreService) SettleOre(accountId int64, oreInfo *model.OreInfo) msg.ErrCode {
	oreInfo.StartTime = 0
	oreInfo.Speed = 0
	dao.OreInfoDao.UpdateOreInfo(accountId, oreInfo)

	oreId := template.GetSystemItemTemplate().GetOreId()
	reqMsg := &msg.RequestInterSettleOrePlayer{
		AccountId: accountId,
		OreId:     oreId,
	}
	oreCli.SendMessage(uint32(msg.MsgId_ID_RequestInterSettleOrePlayer), reqMsg)
	return msg.ErrCode_SUCC
}

func (o *OreService) updateOreInfo(oreId, total, endtime uint32) {
	o.ore.oreId = oreId
	o.ore.total = total
	o.ore.endTime = endtime
	o.ore.updateTime = util.GetCurTime()
}

func (o *OreService) GetOreEndTime() uint32 {
	return o.ore.endTime
}
