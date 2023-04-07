package common

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/log"
	"net"
)

type MsgHandler func(data []byte)

type TcpClient struct {
	serverId uint32
	servAddr string
	conn     net.Conn

	msgHandlerMap map[uint32]func(data []byte)
}

func NewTcpClient(sId uint32, sAddr string) *TcpClient {
	cli := &TcpClient{
		serverId: sId,
		servAddr: sAddr,
	}
	cli.msgHandlerMap = make(map[uint32]func(data []byte))
	return cli
}

func (t *TcpClient) GetServerId() uint32 {
	return t.serverId
}

func (t *TcpClient) Conn() error {
	var err error
	t.conn, err = net.Dial("tcp", t.servAddr)
	if err != nil {
		msgContent := fmt.Sprintf("Conn servAddr:%v err:%v", t.servAddr, err)
		return errors.New(msgContent)
	}

	go func() {
		headBuf := make([]byte, 4)
		headIndex := 0
		for {
			readLen, err := t.conn.Read(headBuf[headIndex:])
			if readLen == 0 {
				return
			}
			if err != nil {
				fmt.Println("err:%v\n", err)
				return
			}
			headIndex += readLen
			if headIndex == 4 {
				bufLen := binary.BigEndian.Uint32(headBuf)
				msgBuf := make([]byte, bufLen)
				readIndex := 0
				for {
					temp, err2 := t.conn.Read(msgBuf)
					if temp == 0 {
						return
					}
					if err2 != nil {
						fmt.Println("err:%v", err2)
						break
					}
					readIndex += temp
					if readIndex == int(bufLen) {
						msgId := binary.BigEndian.Uint32(msgBuf)
						if handler, ok := t.msgHandlerMap[msgId]; ok {
							handler(msgBuf[4:])
						} else {
							log.Error("msgId:%v no handler", msgId)
						}
						headIndex = 0
						break
					}
				}
			}
		}
	}()
	return nil
}

func (t *TcpClient) AddMsgHandler(msgId uint32, handler MsgHandler) {
	t.msgHandlerMap[msgId] = handler
}

func (t *TcpClient) OnDestory() {
	if t.conn != nil {
		t.conn.Close()
	}
}

func (t *TcpClient) SendMessage(msgId uint32, m proto.Message) {
	// 进行编码
	data, err := proto.Marshal(m)
	if err != nil {
		log.Error("SendMessage msgId:%v err:%v", msgId, err)
		return
	}

	// len + data
	out := make([]byte, 4+4+len(data))
	// 默认使用大端序
	binary.BigEndian.PutUint32(out, uint32(4+len(data)))
	binary.BigEndian.PutUint32(out[4:], msgId)
	copy(out[8:], data)
	t.conn.Write(out)
}
