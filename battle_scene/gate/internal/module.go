package internal

import (
	"battleScene/conf"
	"battleScene/game"
	"battleScene/msg"
	"github.com/name5566/leaf/gate"
)

type Module struct {
	*gate.Gate
}

func (m *Module) OnInit() {
	m.Gate = &gate.Gate{
		MaxConnNum:      conf.Server.MaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		MaxMsgLen:       conf.MaxMsgLen,
		HTTPTimeout:     conf.HTTPTimeout,
		TCPAddr:         conf.Server.TCPAddr,
		LenMsgLen:       conf.LenMsgLen,
		LittleEndian:    conf.LittleEndian,
		Processor:       msg.Processor,
		AgentChanRPC:    game.ChanRPC,
	}
}
