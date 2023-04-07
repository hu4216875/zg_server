package gate

import (
	"battleHall/game"
	"battleHall/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.RequestInterServerInfo{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.ReqeustInterGs2BattlHallEnter{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.ReqeustInterGs2BattlHallLeave{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.ResponseInterBattlHall2SceneEnter{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.ResponseInterBattlHall2SceneLeave{}, game.ChanRPC)
}
