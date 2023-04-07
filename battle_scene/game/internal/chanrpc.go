package internal

import (
	"battleScene/game/internal/service"
	"github.com/name5566/leaf/gate"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	if userData := a.UserData(); userData != nil {
		serverId := userData.(uint32)
		service.ServMgr.GetPlayerService().RemoveGsPlayer(serverId)
		service.ServMgr.GetServerService().RemoveServer(serverId)
	}
}
