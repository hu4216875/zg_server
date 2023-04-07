package internal

import (
	"battleScene/game/internal/service"
	"battleScene/msg"
	"github.com/name5566/leaf/gate"
	"reflect"
)

func init() {
	handler(&msg.RequestInterServerInfo{}, handleRequestInterServerInfo)

}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleRequestInterServerInfo(args []interface{}) {
	req := args[0].(*msg.RequestInterServerInfo)
	agent := args[1].(gate.Agent)

	agent.SetUserData(req.ServerId)

	service.ServMgr.GetServerService().AddServer(req.ServerId, agent)
	res := msg.ResponseInterServerInfo{
		Result: int32(msg.ErrCode_SUCC),
	}
	agent.WriteMsg(&res)
}
