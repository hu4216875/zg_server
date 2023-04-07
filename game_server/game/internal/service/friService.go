package service

import (
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/log"
	"server/conf"
	"server/msg"
)

type FriService struct {
	IService
}

func NewFriService() *FriService {
	return &FriService{}
}

func (f *FriService) OnInit() {
}

func (f *FriService) OnDestory() {

}

func (f *FriService) AddFriend(accountId, friAccountId int64, friGsId uint32) msg.ErrCode {
	cli := getMessageClient()
	if cli == nil {
		log.Error("AddFriend accountId:%v messageRoute not found ", accountId)
		return msg.ErrCode_SYSTEM_ERROR
	}

	//if conf.Server.ServerId != friGsId {
	interMsg := &msg.RequestInterTransMessage{
		SrcGsId:  conf.Server.ServerId,
		DestGsId: friGsId,
	}
	interMsg.MsgId = uint32(msg.MsgId_ID_RequestInterAddFri)

	friMsg := &msg.RequestInterAddFri{
		SrcAccountId:  accountId,
		SrcGsId:       conf.Server.ServerId,
		DestAccountId: friAccountId,
		DestGsId:      friGsId,
	}

	if data, err := proto.Marshal(friMsg); err == nil {
		interMsg.Data = data
	} else {
		return msg.ErrCode_SYSTEM_ERROR
	}
	cli.SendMessage(uint32(msg.MsgId_ID_RequestInterTransMessage), interMsg)
	//} else {

	//}

	return msg.ErrCode_SUCC
}
