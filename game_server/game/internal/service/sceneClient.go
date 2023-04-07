package service

import (
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/log"
	"server/conf"
	"server/game/internal/common"
	"server/msg"
)

func NewSceneClient(serverId uint32, serverAddr string) *common.TcpClient {
	cli := common.NewTcpClient(serverId, serverAddr, OnSceneClose)
	if err := cli.Conn(); err != nil {
		log.Error("serverid:%v serverAddr:%v err:%v", serverId, serverAddr, err)
		return nil
	}
	initSceneHandle(cli)

	// 发送服务器信息
	serverInfo := &msg.RequestInterServerInfo{
		ServerId:   conf.Server.ServerId,
		ServerType: uint32(msg.ServerType_SERVER_TYPE_GS),
	}
	cli.SendMessage(uint32(msg.MsgId_ID_RequestInterServerInfo), serverInfo)
	return cli
}

func initSceneHandle(cli *common.TcpClient) {
	cli.AddMsgHandler(uint32(msg.MsgId_ID_ResponseInterServerInfo), responseSceneInterServerInfo)
}

func responseSceneInterServerInfo(data []byte) {
	res := &msg.ResponseInterServerInfo{}
	if err := proto.Unmarshal(data, res); err != nil {
		log.Error("responseInterServerInfo err:%v", err)
		return
	}
	log.Debug("responseSceneInterServerInfo result:%v", res.Result)
}
