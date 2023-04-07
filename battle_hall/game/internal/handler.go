package internal

import (
	"battleHall/game/internal/service"
	"battleHall/msg"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"reflect"
)

func init() {
	handler(&msg.RequestInterServerInfo{}, handleRequestInterServerInfo)
	handler(&msg.ReqeustInterGs2BattlHallEnter{}, handleReqeustInterGs2BattlHallEnter)
	handler(&msg.ReqeustInterGs2BattlHallLeave{}, handleReqeustInterGs2BattlHallLeave)
	handler(&msg.ResponseInterBattlHall2SceneEnter{}, handleResponseInterBattlHall2SceneEnter)
	handler(&msg.ResponseInterBattlHall2SceneLeave{}, handleResponseInterBattlHall2SceneLeave)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleRequestInterServerInfo(args []interface{}) {
	req := args[0].(*msg.RequestInterServerInfo)
	agent := args[1].(gate.Agent)

	agent.SetUserData(req.ServerId)
	service.AddServer(req.ServerId, agent, req.ServerType, req.OnLineNum)

	res := msg.ResponseInterServerInfo{
		Result: int32(msg.ErrCode_SUCC),
	}
	agent.WriteMsg(&res)
}

func handleReqeustInterGs2BattlHallEnter(args []interface{}) {
	req := args[0].(*msg.ReqeustInterGs2BattlHallEnter)
	agent := args[1].(gate.Agent)

	serverId := service.SelectOneSceneServer()
	if serverId == 0 {
		res := msg.ResponseInterGs2BattlHallEnter{
			Result: int32(msg.ErrCode_NO_BATTLE_SCENE),
		}
		agent.WriteMsg(&res)
	} else {
		server := service.FindServer(serverId)
		if server == nil {
			res := msg.ResponseInterGs2BattlHallEnter{
				Result: int32(msg.ErrCode_NO_BATTLE_SCENE),
			}
			agent.WriteMsg(&res)
			log.Error("handleReqeustInterGs2BattlHallEnter serverId:%v not exist", serverId)
		} else {
			service.UpdateSceneOnlineNum(serverId, req.GsServerId, 1)
			sceneMsg := &msg.RequestInterBattlHall2SceneEnter{
				AccountId:  req.AccountId,
				GsServerId: req.GsServerId,
			}
			server.Agent.WriteMsg(sceneMsg)
		}
	}
}

func handleReqeustInterGs2BattlHallLeave(args []interface{}) {
	req := args[0].(*msg.ReqeustInterGs2BattlHallLeave)
	agent := args[1].(gate.Agent)
	serverInfo := service.FindServer(req.SceneServerId)
	if serverInfo != nil {
		service.UpdateSceneOnlineNum(req.SceneServerId, req.GsServerId, -1)
		sceneMsg := &msg.RequestInterBattlHall2SceneLeave{
			AccountId:  req.AccountId,
			GsServerId: req.GsServerId,
		}
		serverInfo.Agent.WriteMsg(sceneMsg)
	} else {
		res := msg.ResponseInterGs2BattlHallLeave{
			Result: int32(msg.ErrCode_SCENE_SERVER_NOT_EXIST),
		}
		agent.WriteMsg(&res)
		log.Error("handleReqeustInterGs2BattlHallLeave serverId:%v not exist", req.SceneServerId)
	}
}

func handleResponseInterBattlHall2SceneEnter(args []interface{}) {
	res := args[0].(*msg.ResponseInterBattlHall2SceneEnter)
	if res.Result != int32(msg.ErrCode_SUCC) {
		service.UpdateSceneOnlineNum(res.SceneId, res.GsServerId, -1)
	}

	server := service.FindServer(res.GsServerId)
	if server != nil {
		gsRes := msg.ResponseInterGs2BattlHallEnter{
			Result:        res.Result,
			SceneServerId: res.SceneId,
			AccountId:     res.AccountId,
		}
		server.Agent.WriteMsg(&gsRes)
	} else {
		log.Error("handleResponseInterBattlHall2SceneEnter serverId:%v not exist", res.GsServerId)
	}
}

func handleResponseInterBattlHall2SceneLeave(args []interface{}) {
	res := args[0].(*msg.ResponseInterBattlHall2SceneLeave)
	if res.Result != int32(msg.ErrCode_SUCC) {
		service.UpdateSceneOnlineNum(res.SceneServerId, res.GsServerId, 1)
	}

	server := service.FindServer(res.GsServerId)
	if server != nil {
		gsRes := msg.ResponseInterGs2BattlHallLeave{
			Result:        res.Result,
			AccountId:     res.AccountId,
			SceneServerID: res.SceneServerId,
		}
		server.Agent.WriteMsg(&gsRes)
	} else {
		log.Error("handleResponseInterBattlHall2SceneLeave serverId:%v not exist", res.GsServerId)
	}
}
