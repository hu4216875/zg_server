package internal

import (
	"battleHall/game/internal/service"
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
	agent := args[0].(gate.Agent)
	userData := agent.UserData()
	if userData != nil {
		serverId := userData.(uint32)
		service.RemoveServer(serverId)
		agent.Close()
	}
}
