package protobuf

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/log"
	"reflect"
)

// -------------------------
// | id | protobuf message |
// -------------------------
type Processor struct {
	littleEndian bool
	//	msgInfo      []*MsgInfo
	msgInfoMap map[uint32]*MsgInfo
	msgID      map[reflect.Type]uint32
}

type MsgInfo struct {
	msgType       reflect.Type
	msgRouter     *chanrpc.Server
	msgHandler    MsgHandler
	msgRawHandler MsgHandler
}

type MsgHandler func([]interface{})

type MsgRaw struct {
	msgID      uint32
	msgRawData []byte
}

func NewProcessor() *Processor {
	p := new(Processor)
	p.littleEndian = false
	p.msgID = make(map[reflect.Type]uint32)
	p.msgInfoMap = make(map[uint32]*MsgInfo)
	return p
}

// It's dangerous to call the method on routing or marshaling (unmarshaling)
func (p *Processor) SetByteOrder(littleEndian bool) {
	p.littleEndian = littleEndian
}

// It's dangerous to call the method on routing or marshaling (unmarshaling)
func (p *Processor) Register(msg proto.Message) uint16 {
	//msgType := reflect.TypeOf(msg)
	//if msgType == nil || msgType.Kind() != reflect.Ptr {
	//	log.Fatal("protobuf message pointer required")
	//}
	//if _, ok := p.msgID[msgType]; ok {
	//	log.Fatal("message %s is already registered", msgType)
	//}
	//
	//i := new(MsgInfo)
	//i.msgType = msgType
	//p.msgInfo = append(p.msgInfo, i)
	//id := uint16(len(p.msgInfo) - 1)
	//p.msgID[msgType] = uint32(id)
	return 0
}

func (p *Processor) RegisterWithMsgId(msgId uint32, msg proto.Message) uint32 {
	msgType := reflect.TypeOf(msg)
	if msgType == nil || msgType.Kind() != reflect.Ptr {
		log.Fatal("protobuf message pointer required")
	}
	if _, ok := p.msgID[msgType]; ok {
		log.Fatal("message %s is already registered", msgType)
	}

	i := new(MsgInfo)
	i.msgType = msgType
	p.msgInfoMap[msgId] = i
	p.msgID[msgType] = msgId
	return msgId
}

// It's dangerous to call the method on routing or marshaling (unmarshaling)
func (p *Processor) SetRouter(msg proto.Message, msgRouter *chanrpc.Server) {
	msgType := reflect.TypeOf(msg)
	id, ok := p.msgID[msgType]
	if !ok {
		log.Fatal("message %s not registered", msgType)
	}

	p.msgInfoMap[id].msgRouter = msgRouter
}

// It's dangerous to call the method on routing or marshaling (unmarshaling)
func (p *Processor) SetHandler(msg proto.Message, msgHandler MsgHandler) {
	msgType := reflect.TypeOf(msg)
	id, ok := p.msgID[msgType]
	if !ok {
		log.Fatal("message %s not registered", msgType)
	}

	p.msgInfoMap[id].msgHandler = msgHandler
}

// It's dangerous to call the method on routing or marshaling (unmarshaling)
func (p *Processor) SetRawHandler(id uint32, msgRawHandler MsgHandler) {
	p.msgInfoMap[id].msgRawHandler = msgRawHandler
}

// goroutine safe
func (p *Processor) Route(msg interface{}, userData interface{}) error {
	// raw
	if msgRaw, ok := msg.(MsgRaw); ok {
		//if msgRaw.msgID >= uint16(len(p.msgInfo)) {
		//	return fmt.Errorf("message id %v not registered", msgRaw.msgID)
		//}
		i := p.msgInfoMap[msgRaw.msgID]
		if i.msgRawHandler != nil {
			i.msgRawHandler([]interface{}{msgRaw.msgID, msgRaw.msgRawData, userData})
		}
		return nil
	}

	// protobuf
	msgType := reflect.TypeOf(msg)
	id, ok := p.msgID[msgType]
	if !ok {
		return fmt.Errorf("message %s not registered", msgType)
	}
	i, ok := p.msgInfoMap[id]
	if ok {
		if i.msgHandler != nil {
			i.msgHandler([]interface{}{msg, userData})
		}
		if i.msgRouter != nil {
			i.msgRouter.Go(msgType, msg, userData)
		}
	}

	return nil
}

// goroutine safe
func (p *Processor) Unmarshal(data []byte) (interface{}, error) {
	if len(data) < 4 {
		return nil, errors.New("protobuf data too short")
	}

	// id
	var id uint32
	if p.littleEndian {
		id = binary.LittleEndian.Uint32(data)
	} else {
		id = binary.BigEndian.Uint32(data)
	}

	// msg
	i, ok := p.msgInfoMap[id]
	if ok {
		if i.msgRawHandler != nil {
			return MsgRaw{id, data[4:]}, nil
		} else {
			msg := reflect.New(i.msgType.Elem()).Interface()
			return msg, proto.UnmarshalMerge(data[4:], msg.(proto.Message))
		}
	}
	return nil, errors.New(fmt.Sprintf("msgId:%v handle not found", id))
}

// goroutine safe
func (p *Processor) Marshal(msg interface{}) ([][]byte, error) {
	msgType := reflect.TypeOf(msg)

	// id
	_id, ok := p.msgID[msgType]
	if !ok {
		err := fmt.Errorf("message %s not registered", msgType)
		return nil, err
	}

	id := make([]byte, 4)
	if p.littleEndian {
		binary.LittleEndian.PutUint32(id, _id)
	} else {
		binary.BigEndian.PutUint32(id, _id)
	}

	// data
	data, err := proto.Marshal(msg.(proto.Message))
	return [][]byte{id, data}, err
}

func (p *Processor) MarshalBinary(id uint32, data []byte) ([][]byte, error) {
	idBuf := make([]byte, 2)
	if p.littleEndian {
		binary.LittleEndian.PutUint32(idBuf, id)
	} else {
		binary.BigEndian.PutUint32(idBuf, id)
	}
	return [][]byte{idBuf, data}, nil
}

// goroutine safe
func (p *Processor) Range(f func(id uint32, t reflect.Type)) {
	for id, i := range p.msgInfoMap {
		f(id, i.msgType)
	}
}
