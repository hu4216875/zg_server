package game

import (
	"client/msg"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func ResponseLoadTasks(data []byte) {
	temp := &msg.ResponseLoadTasks{}
	if err := proto.Unmarshal(data[4:], temp); err == nil {
		fmt.Printf("ResponseLoadTasks err:%v ", err)
	}
	fmt.Printf("ResponseLoadTasks result:%v\n", temp.Result)
}
