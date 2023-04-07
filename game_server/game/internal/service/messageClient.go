package service

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/log"
	"server/conf"
	"server/game/internal/common"
	"server/msg"
)

var (
	interMsgHandle map[uint32]func(*msg.RequestInterTransMessage)
)

func NewMessageConn(serverId uint32, serverAddr string) *common.TcpClient {
	cli := common.NewTcpClient(serverId, serverAddr, nil)
	if err := cli.Conn(); err != nil {
		log.Error("serverid:%v serverAddr:%v err:%v", serverId, serverAddr, err)
		return nil
	}
	initMessageHandler(cli)

	reqMsg := &msg.RequestInterServerInfo{
		ServerId:   conf.Server.ServerId,
		ServerType: uint32(msg.ServerType_SERVER_TYPE_GS),
	}

	cli.SendMessage(uint32(msg.MsgId_ID_RequestInterServerInfo), reqMsg)
	log.Debug("NewMessageConn serverId:%v\n", serverId)
	return cli
}

func initMessageHandler(cli *common.TcpClient) {
	cli.AddMsgHandler(uint32(msg.MsgId_ID_ResponseInterServerInfo), responseInterMessageInfo)
	cli.AddMsgHandler(uint32(msg.MsgId_ID_RequestInterTransMessage), requestInterTransMessage)

	interMsgHandle = make(map[uint32]func(*msg.RequestInterTransMessage))
	interMsgHandle[uint32(msg.MsgId_ID_RequestInterAddFri)] = interFridMsg
}

func responseInterMessageInfo(data []byte) {

}

func requestInterTransMessage(data []byte) {
	req := &msg.RequestInterTransMessage{}
	if err := proto.Unmarshal(data, req); err != nil {
		log.Error("requestInterTransMessage err:%v", err)
		return
	}

	if f, ok := interMsgHandle[req.MsgId]; ok {
		f(req)
	} else {
		log.Error("requestInterTransMessage msgId:%v not found handle", req.MsgId)
	}
}

func interFridMsg(req *msg.RequestInterTransMessage) {
	interMsg := msg.RequestInterAddFri{}
	if err := proto.Unmarshal(req.Data, &interMsg); err != nil {
		return
	}

	fmt.Println(interMsg.DestAccountId)
}
