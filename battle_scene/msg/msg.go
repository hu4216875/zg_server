package msg

import "github.com/name5566/leaf/network/protobuf"

var Processor = protobuf.NewProcessor()

func init() {
	Processor.RegisterWithMsgId(uint32(MsgId_ID_RequestInterServerInfo), &RequestInterServerInfo{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_ResponseInterServerInfo), &ResponseInterServerInfo{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_ReqeustInterGs2BattlHallEnter), &ReqeustInterGs2BattlHallEnter{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_ResponseInterGs2BattlHallEnter), &ResponseInterGs2BattlHallEnter{})
}
