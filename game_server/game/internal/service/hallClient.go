package service

import (
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/log"
	"server/conf"
	"server/game/internal/common"
	"server/msg"
)

func NewHallConn(serverId uint32, serverAddr string) *common.TcpClient {
	cli := common.NewTcpClient(serverId, serverAddr, nil)
	if err := cli.Conn(); err != nil {
		log.Error("serverid:%v serverAddr:%v err:%v", serverId, serverAddr, err)
		return nil
	}
	initHallHandler(cli)

	// 发送服务器信息
	serverInfo := &msg.RequestInterServerInfo{
		ServerId:   conf.Server.ServerId,
		ServerType: uint32(msg.ServerType_SERVER_TYPE_GS),
	}
	cli.SendMessage(uint32(msg.MsgId_ID_RequestInterServerInfo), serverInfo)
	log.Debug("NewHallConn serverId:%v\n", serverId)
	return cli
}

func initHallHandler(cli *common.TcpClient) {
	cli.AddMsgHandler(uint32(msg.MsgId_ID_ResponseInterServerInfo), responseInterServerInfo)
	cli.AddMsgHandler(uint32(msg.MsgId_ID_ResponseInterGs2BattlHallEnter), responseInterGs2BattlHallEnter)
	cli.AddMsgHandler(uint32(msg.MsgId_ID_ResponseInterGs2BattlHallLeave), responseInterGs2BattlHallLeave)
}

func responseInterServerInfo(data []byte) {
	res := &msg.ResponseInterServerInfo{}
	if err := proto.Unmarshal(data, res); err != nil {
		log.Error("responseInterServerInfo err:%v", err)
		return
	}
	log.Debug("responseInterServerInfo result:%v", res.Result)
}

func responseInterGs2BattlHallEnter(data []byte) {
	res := &msg.ResponseInterGs2BattlHallEnter{}
	if err := proto.Unmarshal(data, res); err != nil {
		log.Error("responseInterBattleHall2SceneEnter err:%v", err)
		return
	}

	if playerData := common.PlayerMgr.FindPlayerData(res.AccountId); playerData != nil {
		if res.Result == int32(msg.ErrCode_SUCC) {
			playerData.SceneServerId = res.SceneServerId
		}

		resBattle := &msg.ResponseEnterBattle{
			Result: res.Result,
		}
		playerData.PlayerAgent.WriteMsg(resBattle)
		log.Debug("AccountId:%v Enter Scene:%v err:%v\n", res.AccountId, res.SceneServerId, res.Result)
	}
	log.Debug("responseInterBattleHall2SceneEnter result:%v", res.Result)
}

func responseInterGs2BattlHallLeave(data []byte) {
	res := &msg.ResponseInterGs2BattlHallLeave{}
	if err := proto.Unmarshal(data, res); err != nil {
		log.Error("responseInterGs2BattlHallLeave err:%v", err)
		return
	}

	if playerData := common.PlayerMgr.FindPlayerData(res.AccountId); playerData != nil {
		resBattle := &msg.ResponseLeaveBattle{
			Result: res.Result,
		}
		playerData.PlayerAgent.WriteMsg(resBattle)
		log.Debug("AccountId:%v Leave sceneId:%v err:%v\n", res.AccountId, res.SceneServerID, resBattle.Result)
	}
	log.Debug("responseInterGs2BattlHallLeave result:%v", res.Result)
}
