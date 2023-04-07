package msg

import (
	"github.com/name5566/leaf/network/protobuf"
)

var Processor = protobuf.NewProcessor()

func init() {
	Processor.RegisterWithMsgId(uint32(MsgId_ID_RequestGMCommand), &RequestGMCommand{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_ResponseGMCommand), &ResponseGMCommand{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_RequestRegist), &RequestRegist{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_ResponseRegist), &ResponseRegist{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_RequestLogin), &RequestLogin{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_ResponseLogin), &ResponseLogin{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_RequestLoadItems), &RequestLoadItems{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_ResponseLoadItems), &ResponseLoadItems{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_RequestLogout), &RequestLogout{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_ResponseLogout), &ResponseLogout{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_ResponseKickOut), &ResponseKickOut{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_NotifyUpdateItem), &NotifyUpdateItem{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_RequestClientHeart), &RequestClientHeart{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_ResponseClientHert), &ResponseClientHert{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_RequestOreTotal), &RequestOreTotal{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_ResponseOreTotal), &ResponseOreTotal{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_RequestStartOre), &RequestStartOre{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_ResponseStartOre), &ResponseStartOre{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_RequestEndOre), &RequestEndOre{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_ResponseEndOre), &ResponseEndOre{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_RequestUpgradeOreSpeed), &RequestUpgradeOreSpeed{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_ResponseUpgradeOreSpeed), &ResponseUpgradeOreSpeed{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_RequestOreInfo), &RequestOreInfo{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_ResponseOreInfo), &ResponseOreInfo{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_RequestEnterBattle), &RequestEnterBattle{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_ResponseEnterBattle), &ResponseEnterBattle{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_RequestLeaveBattle), &RequestLeaveBattle{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_ResponseLeaveBattle), &ResponseLeaveBattle{})

	Processor.RegisterWithMsgId(uint32(MsgId_ID_RequestLoadTasks), &RequestLoadTasks{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_ResponseLoadTasks), &ResponseLoadTasks{})

	Processor.RegisterWithMsgId(uint32(MsgId_ID_RequestAddFri), &RequestAddFri{})
	Processor.RegisterWithMsgId(uint32(MsgId_ID_ResponseAddFri), &ResponseAddFri{})
}
