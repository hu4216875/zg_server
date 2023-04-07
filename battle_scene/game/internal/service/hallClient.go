package service

import (
	"battleScene/conf"
	"battleScene/game/internal/common"
	"battleScene/msg"
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/log"
)

func NewHallConn(serverId uint32, serverAddr string) {
	cli := common.NewTcpClient(serverId, serverAddr)
	if err := cli.Conn(); err != nil {
		log.Error("serverid:%v serverAddr:%v err:%v", serverId, serverAddr, err)
		return
	}
	initHandler(cli)
	hallCli = cli

	// 发送服务器信息
	serverInfo := &msg.RequestInterServerInfo{
		ServerId:   conf.Server.ServerId,
		ServerType: uint32(msg.ServerType_SERVER_TYPE_BATTLE_SCENE),
		OnLineNum:  conf.Server.MaxOnlineNum,
	}
	cli.SendMessage(uint32(msg.MsgId_ID_RequestInterServerInfo), serverInfo)
}

func initHandler(cli *common.TcpClient) {
	cli.AddMsgHandler(uint32(msg.MsgId_ID_ResponseInterServerInfo), responseInterServerInfo)
	cli.AddMsgHandler(uint32(msg.MsgId_ID_RequestInterBattlHall2SceneEnter), requestInterBattlHall2SceneEnter)
	cli.AddMsgHandler(uint32(msg.MsgId_ID_RequestInterBattlHall2SceneLeave), requestInterBattlHall2SceneLeave)

}

func responseInterServerInfo(data []byte) {

}

func requestInterBattlHall2SceneEnter(data []byte) {
	req := &msg.RequestInterBattlHall2SceneEnter{}
	if err := proto.Unmarshal(data, req); err != nil {
		log.Error("requestInterBattlHall2SceneEnter err:%v", err)
		return
	}

	res := msg.ResponseInterBattlHall2SceneEnter{
		Result:     int32(msg.ErrCode_SUCC),
		AccountId:  req.AccountId,
		SceneId:    conf.Server.ServerId,
		GsServerId: req.GsServerId,
	}

	playerData := ServMgr.GetPlayerService().GetPlayerData(req.AccountId)
	if playerData != nil {
		res.Result = int32(msg.ErrCode_USER_IN_BATTLE)
	} else {
		playerData := &PlayerData{AccountId: req.AccountId, GsId: req.GsServerId}
		ServMgr.GetPlayerService().AddPlayerData(playerData)
		log.Debug("AccountId:%v Enter scene", playerData.AccountId)
	}

	hallClient := getBattleHallClient()
	if hallClient != nil {
		hallClient.SendMessage(uint32(msg.MsgId_ID_ResponseInterBattlHall2SceneEnter), &res)
	} else {
		log.Debug("requestInterBattlHall2SceneEnter hallClient not exist")
	}
}

func requestInterBattlHall2SceneLeave(data []byte) {
	req := &msg.RequestInterBattlHall2SceneLeave{}
	if err := proto.Unmarshal(data, req); err != nil {
		log.Error("requestInterBattlHall2SceneLeave err:%v", err)
		return
	}

	res := msg.ResponseInterBattlHall2SceneLeave{
		Result:        int32(msg.ErrCode_SUCC),
		AccountId:     req.AccountId,
		SceneServerId: conf.Server.ServerId,
		GsServerId:    req.GsServerId,
	}

	playerData := ServMgr.GetPlayerService().GetPlayerData(req.AccountId)
	if playerData != nil {
		ServMgr.GetPlayerService().RemovePlayerData(req.AccountId)
		log.Debug("AccountId:%v leave scene", playerData.AccountId)
	} else {
		res.Result = int32(msg.ErrCode_USER_NOT_IN_BATTLE)
	}

	hallClient := getBattleHallClient()
	if hallClient != nil {
		hallClient.SendMessage(uint32(msg.MsgId_ID_ResponseInterBattlHall2SceneLeave), &res)
	} else {
		log.Debug("requestInterBattlHall2SceneLeave hallClient not exist")
	}

}
