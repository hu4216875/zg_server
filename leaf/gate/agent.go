package gate

import (
	"net"
)

type Agent interface {
	WriteMsg(msg interface{})
	WriteBinaryMsg(msgId uint32, data []byte)
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()
	UserData() interface{}
	SetUserData(data interface{})
}
