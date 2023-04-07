package msg

import "github.com/name5566/leaf/network/protobuf"

var Processor = protobuf.NewProcessor()

func init() {
	Processor.RegisterWithMsgId(uint32(MsgId_ID_RequestInterServerInfo), &RequestInterServerInfo{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_ResponseInterServerInfo), &ResponseInterServerInfo{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_ReqeustInterGs2BattlHallEnter), &ReqeustInterGs2BattlHallEnter{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_ResponseInterGs2BattlHallEnter), &ResponseInterGs2BattlHallEnter{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_RequestInterBattlHall2SceneEnter), &RequestInterBattlHall2SceneEnter{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_ResponseInterBattlHall2SceneEnter), &ResponseInterBattlHall2SceneEnter{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_ReqeustInterGs2BattlHallLeave), &ReqeustInterGs2BattlHallLeave{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_ResponseInterGs2BattlHallLeave), &ResponseInterGs2BattlHallLeave{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_RequestInterBattlHall2SceneLeave), &RequestInterBattlHall2SceneLeave{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_ResponseInterBattlHall2SceneLeave), &ResponseInterBattlHall2SceneLeave{})

}
