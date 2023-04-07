package service

import (
	"server/conf"
	"server/game/internal/common"
	"server/msg"
	"server/publicconst"
)

type BattleService struct {
	IService
}

func NewBattleService() *BattleService {
	return &BattleService{}
}

func (b *BattleService) OnInit() {
}

func (b *BattleService) OnDestory() {

}

func (b *BattleService) RequestEnterBattle(accountId int64) int32 {
	playerData := common.PlayerMgr.FindPlayerData(accountId)
	if playerData.SceneServerId > 0 {
		return int32(msg.ErrCode_USER_IN_BATTLE)
	}

	cli := getBattleHallClient()
	if cli == nil {
		return int32(msg.ErrCode_SYSTEM_ERROR)
	}

	reqMsg := &msg.ReqeustInterGs2BattlHallEnter{
		AccountId:  accountId,
		GsServerId: conf.Server.ServerId,
	}
	cli.SendMessage(uint32(msg.MsgId_ID_ReqeustInterGs2BattlHallEnter), reqMsg)
	return int32(msg.ErrCode_SUCC)
}

func (b *BattleService) RequestLeaveBattle(accountId int64) int32 {
	playerData := common.PlayerMgr.FindPlayerData(accountId)
	if playerData.SceneServerId == 0 {
		return int32(msg.ErrCode_USER_NOT_IN_BATTLE)
	}
	return b.LeaveBattleHall(playerData)
}

func (b *BattleService) LeaveBattleHall(playerData *common.PlayerData) int32 {
	cli := getBattleHallClient()
	if cli == nil {
		return int32(msg.ErrCode_SYSTEM_ERROR)
	}
	reqMsg := &msg.ReqeustInterGs2BattlHallLeave{
		AccountId:     playerData.AccountInfo.AccountId,
		GsServerId:    conf.Server.ServerId,
		SceneServerId: playerData.SceneServerId,
	}
	playerData.SceneServerId = 0
	cli.SendMessage(uint32(msg.MsgId_ID_ReqeustInterGs2BattlHallLeave), reqMsg)
	return int32(msg.ErrCode_SUCC)
}

func OnSceneClose(sceneId uint32) {
	players := common.PlayerMgr.GetScenePlayer(sceneId)
	retMsg := &msg.ResponseLeaveBattle{Result: int32(msg.ErrCode_SUCC)}
	for i := 0; i < len(players); i++ {
		players[i].SceneServerId = 0
		if players[i].State == publicconst.Online {
			players[i].PlayerAgent.WriteMsg(retMsg)
		}
	}
}
