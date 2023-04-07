package game

import (
	"client/msg"
	"github.com/golang/protobuf/proto"
	"log"
)

func sendLoadTasks() {
	reqMsg := &msg.RequestLoadTasks{}
	// 进行编码
	data, err := proto.Marshal(reqMsg)
	if err != nil {
		log.Fatal("marshaling error: %v", err)
	}
	sendMsg(uint32(msg.MsgId_ID_RequestLoadTasks), data)
}
