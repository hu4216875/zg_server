package gate

import (
	"battleScene/game"
	"battleScene/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.RequestInterServerInfo{}, game.ChanRPC)
}
