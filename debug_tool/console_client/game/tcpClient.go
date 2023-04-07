package game

import (
	"client/msg"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
)

type MsgHandler func(data []byte)

type OnClose func()

type TcpClient struct {
	servAddr string
	conn     net.Conn

	msgHandlerMap map[msg.MsgId]func(data []byte)
	onCloseFunc   OnClose
}

func NewTcpClient(sAddr string, f OnClose) *TcpClient {
	cli := &TcpClient{
		servAddr:    sAddr,
		onCloseFunc: f,
	}
	cli.msgHandlerMap = make(map[msg.MsgId]func(data []byte))
	return cli
}

func (t *TcpClient) GetServerAddr() string {
	return t.servAddr
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
				if t.onCloseFunc != nil {
					t.onCloseFunc()
				}
				return
			}
			if err != nil {
				if t.onCloseFunc != nil {
					t.onCloseFunc()
				}
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
						if t.onCloseFunc != nil {
							t.onCloseFunc()
						}
						return
					}
					if err2 != nil {
						if t.onCloseFunc != nil {
							t.onCloseFunc()
						}
						fmt.Println("err:%v", err2)
						break
					}
					readIndex += temp
					if readIndex == int(bufLen) {
						msgId := binary.BigEndian.Uint32(msgBuf)
						if handler, ok := t.msgHandlerMap[msg.MsgId(msgId)]; ok {
							handler(msgBuf[4:])
						} else {
							fmt.Printf("msgId:%v no handler \n", msgId)
						}
						if msgId != uint32(msg.MsgId_ID_ResponseClientHert) {
							ShowMenu()
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

func (t *TcpClient) AddMsgHandler(msgId msg.MsgId, handler MsgHandler) {
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
		fmt.Println("SendMessage msgId:%v err:%v", msgId, err)
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
