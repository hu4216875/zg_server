// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: msgtype.proto

package msg

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MsgId int32

const (
	MsgId_ID_RequestGMCommand                  MsgId = 0
	MsgId_ID_ResponseGMCommand                 MsgId = 1
	MsgId_ID_RequestRegist                     MsgId = 2
	MsgId_ID_ResponseRegist                    MsgId = 3
	MsgId_ID_RequestLogin                      MsgId = 4
	MsgId_ID_ResponseLogin                     MsgId = 5
	MsgId_ID_RequestLogout                     MsgId = 6
	MsgId_ID_ResponseLogout                    MsgId = 7
	MsgId_ID_ResponseKickOut                   MsgId = 8
	MsgId_ID_RequestLoadItems                  MsgId = 9
	MsgId_ID_ResponseLoadItems                 MsgId = 10
	MsgId_ID_NotifyUpdateItem                  MsgId = 11
	MsgId_ID_RequestClientHeart                MsgId = 12
	MsgId_ID_ResponseClientHert                MsgId = 13
	MsgId_ID_RequestOreTotal                   MsgId = 14
	MsgId_ID_ResponseOreTotal                  MsgId = 15
	MsgId_ID_RequestStartOre                   MsgId = 16
	MsgId_ID_ResponseStartOre                  MsgId = 17
	MsgId_ID_RequestEndOre                     MsgId = 18
	MsgId_ID_ResponseEndOre                    MsgId = 19
	MsgId_ID_RequestUpgradeOreSpeed            MsgId = 20
	MsgId_ID_ResponseUpgradeOreSpeed           MsgId = 21
	MsgId_ID_RequestOreInfo                    MsgId = 22
	MsgId_ID_ResponseOreInfo                   MsgId = 23
	MsgId_ID_RequestEnterBattle                MsgId = 24
	MsgId_ID_ResponseEnterBattle               MsgId = 25
	MsgId_ID_RequestLeaveBattle                MsgId = 26
	MsgId_ID_ResponseLeaveBattle               MsgId = 27
	MsgId_ID_RequestLoadTasks                  MsgId = 28
	MsgId_ID_ResponseLoadTasks                 MsgId = 29
	MsgId_ID_NotifyUpdateTask                  MsgId = 30
	MsgId_ID_RequestInterServerInfo            MsgId = 10000
	MsgId_ID_ResponseInterServerInfo           MsgId = 10001
	MsgId_ID_ReqeustInterGs2BattlHallEnter     MsgId = 10002
	MsgId_ID_ResponseInterGs2BattlHallEnter    MsgId = 10003
	MsgId_ID_RequestInterBattlHall2SceneEnter  MsgId = 10004
	MsgId_ID_ResponseInterBattlHall2SceneEnter MsgId = 10005
	MsgId_ID_ReqeustInterGs2BattlHallLeave     MsgId = 10006
	MsgId_ID_ResponseInterGs2BattlHallLeave    MsgId = 10007
	MsgId_ID_RequestInterBattlHall2SceneLeave  MsgId = 10008
	MsgId_ID_ResponseInterBattlHall2SceneLeave MsgId = 10009
	MsgId_ID_RequestAddFri                     MsgId = 11000
	MsgId_ID_ResponseAddFri                    MsgId = 11001
	MsgId_ID_RequestInterGetOreInfo            MsgId = 12000
	MsgId_ID_ResponseInterGetOreInfo           MsgId = 12001
	MsgId_ID_RequestInterAddOrePlayer          MsgId = 12002
	MsgId_ID_ResponseInterAddOrePlayer         MsgId = 12003
	MsgId_ID_RequestInterUpdateOrePlayer       MsgId = 12004
	MsgId_ID_ResponseInterUpdateOrePlayer      MsgId = 12005
	MsgId_ID_RequestInterSettleOrePlayer       MsgId = 12006
	MsgId_ID_ResponseInterSettleOrePlayer      MsgId = 12007
	MsgId_ID_RequestInterTransMessage          MsgId = 12008
	MsgId_ID_RequestInterAddFri                MsgId = 12009
)

// Enum value maps for MsgId.
var (
	MsgId_name = map[int32]string{
		0:     "ID_RequestGMCommand",
		1:     "ID_ResponseGMCommand",
		2:     "ID_RequestRegist",
		3:     "ID_ResponseRegist",
		4:     "ID_RequestLogin",
		5:     "ID_ResponseLogin",
		6:     "ID_RequestLogout",
		7:     "ID_ResponseLogout",
		8:     "ID_ResponseKickOut",
		9:     "ID_RequestLoadItems",
		10:    "ID_ResponseLoadItems",
		11:    "ID_NotifyUpdateItem",
		12:    "ID_RequestClientHeart",
		13:    "ID_ResponseClientHert",
		14:    "ID_RequestOreTotal",
		15:    "ID_ResponseOreTotal",
		16:    "ID_RequestStartOre",
		17:    "ID_ResponseStartOre",
		18:    "ID_RequestEndOre",
		19:    "ID_ResponseEndOre",
		20:    "ID_RequestUpgradeOreSpeed",
		21:    "ID_ResponseUpgradeOreSpeed",
		22:    "ID_RequestOreInfo",
		23:    "ID_ResponseOreInfo",
		24:    "ID_RequestEnterBattle",
		25:    "ID_ResponseEnterBattle",
		26:    "ID_RequestLeaveBattle",
		27:    "ID_ResponseLeaveBattle",
		28:    "ID_RequestLoadTasks",
		29:    "ID_ResponseLoadTasks",
		30:    "ID_NotifyUpdateTask",
		10000: "ID_RequestInterServerInfo",
		10001: "ID_ResponseInterServerInfo",
		10002: "ID_ReqeustInterGs2BattlHallEnter",
		10003: "ID_ResponseInterGs2BattlHallEnter",
		10004: "ID_RequestInterBattlHall2SceneEnter",
		10005: "ID_ResponseInterBattlHall2SceneEnter",
		10006: "ID_ReqeustInterGs2BattlHallLeave",
		10007: "ID_ResponseInterGs2BattlHallLeave",
		10008: "ID_RequestInterBattlHall2SceneLeave",
		10009: "ID_ResponseInterBattlHall2SceneLeave",
		11000: "ID_RequestAddFri",
		11001: "ID_ResponseAddFri",
		12000: "ID_RequestInterGetOreInfo",
		12001: "ID_ResponseInterGetOreInfo",
		12002: "ID_RequestInterAddOrePlayer",
		12003: "ID_ResponseInterAddOrePlayer",
		12004: "ID_RequestInterUpdateOrePlayer",
		12005: "ID_ResponseInterUpdateOrePlayer",
		12006: "ID_RequestInterSettleOrePlayer",
		12007: "ID_ResponseInterSettleOrePlayer",
		12008: "ID_RequestInterTransMessage",
		12009: "ID_RequestInterAddFri",
	}
	MsgId_value = map[string]int32{
		"ID_RequestGMCommand":                  0,
		"ID_ResponseGMCommand":                 1,
		"ID_RequestRegist":                     2,
		"ID_ResponseRegist":                    3,
		"ID_RequestLogin":                      4,
		"ID_ResponseLogin":                     5,
		"ID_RequestLogout":                     6,
		"ID_ResponseLogout":                    7,
		"ID_ResponseKickOut":                   8,
		"ID_RequestLoadItems":                  9,
		"ID_ResponseLoadItems":                 10,
		"ID_NotifyUpdateItem":                  11,
		"ID_RequestClientHeart":                12,
		"ID_ResponseClientHert":                13,
		"ID_RequestOreTotal":                   14,
		"ID_ResponseOreTotal":                  15,
		"ID_RequestStartOre":                   16,
		"ID_ResponseStartOre":                  17,
		"ID_RequestEndOre":                     18,
		"ID_ResponseEndOre":                    19,
		"ID_RequestUpgradeOreSpeed":            20,
		"ID_ResponseUpgradeOreSpeed":           21,
		"ID_RequestOreInfo":                    22,
		"ID_ResponseOreInfo":                   23,
		"ID_RequestEnterBattle":                24,
		"ID_ResponseEnterBattle":               25,
		"ID_RequestLeaveBattle":                26,
		"ID_ResponseLeaveBattle":               27,
		"ID_RequestLoadTasks":                  28,
		"ID_ResponseLoadTasks":                 29,
		"ID_NotifyUpdateTask":                  30,
		"ID_RequestInterServerInfo":            10000,
		"ID_ResponseInterServerInfo":           10001,
		"ID_ReqeustInterGs2BattlHallEnter":     10002,
		"ID_ResponseInterGs2BattlHallEnter":    10003,
		"ID_RequestInterBattlHall2SceneEnter":  10004,
		"ID_ResponseInterBattlHall2SceneEnter": 10005,
		"ID_ReqeustInterGs2BattlHallLeave":     10006,
		"ID_ResponseInterGs2BattlHallLeave":    10007,
		"ID_RequestInterBattlHall2SceneLeave":  10008,
		"ID_ResponseInterBattlHall2SceneLeave": 10009,
		"ID_RequestAddFri":                     11000,
		"ID_ResponseAddFri":                    11001,
		"ID_RequestInterGetOreInfo":            12000,
		"ID_ResponseInterGetOreInfo":           12001,
		"ID_RequestInterAddOrePlayer":          12002,
		"ID_ResponseInterAddOrePlayer":         12003,
		"ID_RequestInterUpdateOrePlayer":       12004,
		"ID_ResponseInterUpdateOrePlayer":      12005,
		"ID_RequestInterSettleOrePlayer":       12006,
		"ID_ResponseInterSettleOrePlayer":      12007,
		"ID_RequestInterTransMessage":          12008,
		"ID_RequestInterAddFri":                12009,
	}
)

func (x MsgId) Enum() *MsgId {
	p := new(MsgId)
	*p = x
	return p
}

func (x MsgId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgId) Descriptor() protoreflect.EnumDescriptor {
	return file_msgtype_proto_enumTypes[0].Descriptor()
}

func (MsgId) Type() protoreflect.EnumType {
	return &file_msgtype_proto_enumTypes[0]
}

func (x MsgId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgId.Descriptor instead.
func (MsgId) EnumDescriptor() ([]byte, []int) {
	return file_msgtype_proto_rawDescGZIP(), []int{0}
}

type GMCommand int32

const (
	GMCommand_Command_NONE GMCommand = 0
	GMCommand_AddItem      GMCommand = 1 //添加道具 数据格式 [{"itemId":1, "num":2}, {"itemId":2, "num":1}]
	GMCommand_ReloadConfig GMCommand = 2 //重新加载配置文件["item","system"]
)

// Enum value maps for GMCommand.
var (
	GMCommand_name = map[int32]string{
		0: "Command_NONE",
		1: "AddItem",
		2: "ReloadConfig",
	}
	GMCommand_value = map[string]int32{
		"Command_NONE": 0,
		"AddItem":      1,
		"ReloadConfig": 2,
	}
)

func (x GMCommand) Enum() *GMCommand {
	p := new(GMCommand)
	*p = x
	return p
}

func (x GMCommand) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GMCommand) Descriptor() protoreflect.EnumDescriptor {
	return file_msgtype_proto_enumTypes[1].Descriptor()
}

func (GMCommand) Type() protoreflect.EnumType {
	return &file_msgtype_proto_enumTypes[1]
}

func (x GMCommand) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GMCommand.Descriptor instead.
func (GMCommand) EnumDescriptor() ([]byte, []int) {
	return file_msgtype_proto_rawDescGZIP(), []int{1}
}

type ErrCode int32

const (
	ErrCode_ERR_NONE               ErrCode = 0
	ErrCode_SUCC                   ErrCode = 1
	ErrCode_SYSTEM_ERROR           ErrCode = 2  // 系统错误
	ErrCode_USER_ID_EXIST          ErrCode = 3  // uid 已经存在
	ErrCode_USERID_EMPTY           ErrCode = 4  //uid 空
	ErrCode_USERID_OVER_MAX_LEN    ErrCode = 5  //超过uid最大长度
	ErrCode_FORBIDDEN_USER         ErrCode = 6  // 禁止用户登录
	ErrCode_ISLOGINING             ErrCode = 7  // 登录中
	ErrCode_OTHER_LOGIN            ErrCode = 8  // 其他地方登录
	ErrCode_USER_NOT_LOGIN         ErrCode = 10 //玩家没有登录
	ErrCode_HAS_START_ORE          ErrCode = 11 // 已经开始挖矿了
	ErrCode_NO_START_ORE           ErrCode = 12 // 没有开始挖矿
	ErrCode_NO_ENOUGH_ITEM         ErrCode = 13 // 没有足够的道具
	ErrCode_NO_ORE_RESOURCE        ErrCode = 14 // 没有矿洞资源了
	ErrCode_USER_IN_BATTLE         ErrCode = 15 // 玩家在战斗中了
	ErrCode_USER_NOT_IN_BATTLE     ErrCode = 16 // 玩家不在战斗中
	ErrCode_NO_BATTLE_SCENE        ErrCode = 17 // 没有战斗场景服务器了
	ErrCode_SCENE_SERVER_NOT_EXIST ErrCode = 18 // 场景服务器不存在
)

// Enum value maps for ErrCode.
var (
	ErrCode_name = map[int32]string{
		0:  "ERR_NONE",
		1:  "SUCC",
		2:  "SYSTEM_ERROR",
		3:  "USER_ID_EXIST",
		4:  "USERID_EMPTY",
		5:  "USERID_OVER_MAX_LEN",
		6:  "FORBIDDEN_USER",
		7:  "ISLOGINING",
		8:  "OTHER_LOGIN",
		10: "USER_NOT_LOGIN",
		11: "HAS_START_ORE",
		12: "NO_START_ORE",
		13: "NO_ENOUGH_ITEM",
		14: "NO_ORE_RESOURCE",
		15: "USER_IN_BATTLE",
		16: "USER_NOT_IN_BATTLE",
		17: "NO_BATTLE_SCENE",
		18: "SCENE_SERVER_NOT_EXIST",
	}
	ErrCode_value = map[string]int32{
		"ERR_NONE":               0,
		"SUCC":                   1,
		"SYSTEM_ERROR":           2,
		"USER_ID_EXIST":          3,
		"USERID_EMPTY":           4,
		"USERID_OVER_MAX_LEN":    5,
		"FORBIDDEN_USER":         6,
		"ISLOGINING":             7,
		"OTHER_LOGIN":            8,
		"USER_NOT_LOGIN":         10,
		"HAS_START_ORE":          11,
		"NO_START_ORE":           12,
		"NO_ENOUGH_ITEM":         13,
		"NO_ORE_RESOURCE":        14,
		"USER_IN_BATTLE":         15,
		"USER_NOT_IN_BATTLE":     16,
		"NO_BATTLE_SCENE":        17,
		"SCENE_SERVER_NOT_EXIST": 18,
	}
)

func (x ErrCode) Enum() *ErrCode {
	p := new(ErrCode)
	*p = x
	return p
}

func (x ErrCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrCode) Descriptor() protoreflect.EnumDescriptor {
	return file_msgtype_proto_enumTypes[2].Descriptor()
}

func (ErrCode) Type() protoreflect.EnumType {
	return &file_msgtype_proto_enumTypes[2]
}

func (x ErrCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrCode.Descriptor instead.
func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return file_msgtype_proto_rawDescGZIP(), []int{2}
}

type ServerType int32

const (
	ServerType_SERVER_TYPE_NONE         ServerType = 0
	ServerType_SERVER_TYPE_GS           ServerType = 1 //gs 服务器
	ServerType_SERVER_TYPE_BATTLE_HALL  ServerType = 2 // 战斗大厅服务器
	ServerType_SERVER_TYPE_BATTLE_SCENE ServerType = 3 // 战斗场景服务器
)

// Enum value maps for ServerType.
var (
	ServerType_name = map[int32]string{
		0: "SERVER_TYPE_NONE",
		1: "SERVER_TYPE_GS",
		2: "SERVER_TYPE_BATTLE_HALL",
		3: "SERVER_TYPE_BATTLE_SCENE",
	}
	ServerType_value = map[string]int32{
		"SERVER_TYPE_NONE":         0,
		"SERVER_TYPE_GS":           1,
		"SERVER_TYPE_BATTLE_HALL":  2,
		"SERVER_TYPE_BATTLE_SCENE": 3,
	}
)

func (x ServerType) Enum() *ServerType {
	p := new(ServerType)
	*p = x
	return p
}

func (x ServerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServerType) Descriptor() protoreflect.EnumDescriptor {
	return file_msgtype_proto_enumTypes[3].Descriptor()
}

func (ServerType) Type() protoreflect.EnumType {
	return &file_msgtype_proto_enumTypes[3]
}

func (x ServerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServerType.Descriptor instead.
func (ServerType) EnumDescriptor() ([]byte, []int) {
	return file_msgtype_proto_rawDescGZIP(), []int{3}
}

var File_msgtype_proto protoreflect.FileDescriptor

var file_msgtype_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x73, 0x67, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x6d, 0x73, 0x67, 0x2a, 0xa2, 0x0c, 0x0a, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x13, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x4d, 0x43, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x49, 0x44, 0x5f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x47, 0x4d, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x10,
	0x01, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x44, 0x5f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x10, 0x03, 0x12, 0x13,
	0x0a, 0x0f, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x44, 0x5f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x10, 0x06, 0x12,
	0x15, 0x0a, 0x11, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x6f,
	0x67, 0x6f, 0x75, 0x74, 0x10, 0x07, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4b, 0x69, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x10, 0x08, 0x12, 0x17,
	0x0a, 0x13, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x61, 0x64,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x10, 0x09, 0x12, 0x18, 0x0a, 0x14, 0x49, 0x44, 0x5f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x6f, 0x61, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x10,
	0x0a, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x44, 0x5f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x10, 0x0b, 0x12, 0x19, 0x0a, 0x15, 0x49, 0x44,
	0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x10, 0x0c, 0x12, 0x19, 0x0a, 0x15, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x48, 0x65, 0x72, 0x74, 0x10, 0x0d,
	0x12, 0x16, 0x0a, 0x12, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x72,
	0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x10, 0x0e, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x44, 0x5f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4f, 0x72, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x10,
	0x0f, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x4f, 0x72, 0x65, 0x10, 0x10, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x44, 0x5f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4f, 0x72, 0x65,
	0x10, 0x11, 0x12, 0x14, 0x0a, 0x10, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x45, 0x6e, 0x64, 0x4f, 0x72, 0x65, 0x10, 0x12, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x44, 0x5f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x6e, 0x64, 0x4f, 0x72, 0x65, 0x10, 0x13, 0x12,
	0x1d, 0x0a, 0x19, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x70, 0x67,
	0x72, 0x61, 0x64, 0x65, 0x4f, 0x72, 0x65, 0x53, 0x70, 0x65, 0x65, 0x64, 0x10, 0x14, 0x12, 0x1e,
	0x0a, 0x1a, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x55, 0x70, 0x67,
	0x72, 0x61, 0x64, 0x65, 0x4f, 0x72, 0x65, 0x53, 0x70, 0x65, 0x65, 0x64, 0x10, 0x15, 0x12, 0x15,
	0x0a, 0x11, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x72, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x10, 0x16, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x4f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x10, 0x17, 0x12, 0x19, 0x0a,
	0x15, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x65, 0x72,
	0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x10, 0x18, 0x12, 0x1a, 0x0a, 0x16, 0x49, 0x44, 0x5f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x42, 0x61, 0x74, 0x74,
	0x6c, 0x65, 0x10, 0x19, 0x12, 0x19, 0x0a, 0x15, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x10, 0x1a, 0x12,
	0x1a, 0x0a, 0x16, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x65,
	0x61, 0x76, 0x65, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x10, 0x1b, 0x12, 0x17, 0x0a, 0x13, 0x49,
	0x44, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x61, 0x64, 0x54, 0x61, 0x73,
	0x6b, 0x73, 0x10, 0x1c, 0x12, 0x18, 0x0a, 0x14, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x4c, 0x6f, 0x61, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x10, 0x1d, 0x12, 0x17,
	0x0a, 0x13, 0x49, 0x44, 0x5f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x61, 0x73, 0x6b, 0x10, 0x1e, 0x12, 0x1e, 0x0a, 0x19, 0x49, 0x44, 0x5f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x10, 0x90, 0x4e, 0x12, 0x1f, 0x0a, 0x1a, 0x49, 0x44, 0x5f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x10, 0x91, 0x4e, 0x12, 0x25, 0x0a, 0x20, 0x49, 0x44, 0x5f, 0x52,
	0x65, 0x71, 0x65, 0x75, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x47, 0x73, 0x32, 0x42, 0x61,
	0x74, 0x74, 0x6c, 0x48, 0x61, 0x6c, 0x6c, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x10, 0x92, 0x4e, 0x12,
	0x26, 0x0a, 0x21, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x47, 0x73, 0x32, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x48, 0x61, 0x6c, 0x6c, 0x45,
	0x6e, 0x74, 0x65, 0x72, 0x10, 0x93, 0x4e, 0x12, 0x28, 0x0a, 0x23, 0x49, 0x44, 0x5f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x48,
	0x61, 0x6c, 0x6c, 0x32, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x10, 0x94,
	0x4e, 0x12, 0x29, 0x0a, 0x24, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x48, 0x61, 0x6c, 0x6c, 0x32, 0x53,
	0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x10, 0x95, 0x4e, 0x12, 0x25, 0x0a, 0x20,
	0x49, 0x44, 0x5f, 0x52, 0x65, 0x71, 0x65, 0x75, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x47,
	0x73, 0x32, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x48, 0x61, 0x6c, 0x6c, 0x4c, 0x65, 0x61, 0x76, 0x65,
	0x10, 0x96, 0x4e, 0x12, 0x26, 0x0a, 0x21, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x47, 0x73, 0x32, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x48,
	0x61, 0x6c, 0x6c, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x10, 0x97, 0x4e, 0x12, 0x28, 0x0a, 0x23, 0x49,
	0x44, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x42, 0x61,
	0x74, 0x74, 0x6c, 0x48, 0x61, 0x6c, 0x6c, 0x32, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4c, 0x65, 0x61,
	0x76, 0x65, 0x10, 0x98, 0x4e, 0x12, 0x29, 0x0a, 0x24, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x48, 0x61,
	0x6c, 0x6c, 0x32, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x10, 0x99, 0x4e,
	0x12, 0x15, 0x0a, 0x10, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x41, 0x64,
	0x64, 0x46, 0x72, 0x69, 0x10, 0xf8, 0x55, 0x12, 0x16, 0x0a, 0x11, 0x49, 0x44, 0x5f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x10, 0xf9, 0x55, 0x12,
	0x1e, 0x0a, 0x19, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x10, 0xe0, 0x5d, 0x12,
	0x1f, 0x0a, 0x1a, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x10, 0xe1, 0x5d,
	0x12, 0x20, 0x0a, 0x1b, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x10,
	0xe2, 0x5d, 0x12, 0x21, 0x0a, 0x1c, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x65, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x10, 0xe3, 0x5d, 0x12, 0x23, 0x0a, 0x1e, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x10, 0xe4, 0x5d, 0x12, 0x24, 0x0a, 0x1f, 0x49, 0x44,
	0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x10, 0xe5, 0x5d,
	0x12, 0x23, 0x0a, 0x1e, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4f, 0x72, 0x65, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x10, 0xe6, 0x5d, 0x12, 0x24, 0x0a, 0x1f, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4f,
	0x72, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x10, 0xe7, 0x5d, 0x12, 0x20, 0x0a, 0x1b, 0x49,
	0x44, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x10, 0xe8, 0x5d, 0x12, 0x1a, 0x0a,
	0x15, 0x49, 0x44, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x41, 0x64, 0x64, 0x46, 0x72, 0x69, 0x10, 0xe9, 0x5d, 0x2a, 0x3c, 0x0a, 0x09, 0x47, 0x4d, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x49,
	0x74, 0x65, 0x6d, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x10, 0x02, 0x2a, 0xe5, 0x02, 0x0a, 0x07, 0x45, 0x72, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x52, 0x52, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x55, 0x43, 0x43, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53,
	0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x11, 0x0a,
	0x0d, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0x03,
	0x12, 0x10, 0x0a, 0x0c, 0x55, 0x53, 0x45, 0x52, 0x49, 0x44, 0x5f, 0x45, 0x4d, 0x50, 0x54, 0x59,
	0x10, 0x04, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x53, 0x45, 0x52, 0x49, 0x44, 0x5f, 0x4f, 0x56, 0x45,
	0x52, 0x5f, 0x4d, 0x41, 0x58, 0x5f, 0x4c, 0x45, 0x4e, 0x10, 0x05, 0x12, 0x12, 0x0a, 0x0e, 0x46,
	0x4f, 0x52, 0x42, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0x06, 0x12,
	0x0e, 0x0a, 0x0a, 0x49, 0x53, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x07, 0x12,
	0x0f, 0x0a, 0x0b, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x10, 0x08,
	0x12, 0x12, 0x0a, 0x0e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4c, 0x4f, 0x47,
	0x49, 0x4e, 0x10, 0x0a, 0x12, 0x11, 0x0a, 0x0d, 0x48, 0x41, 0x53, 0x5f, 0x53, 0x54, 0x41, 0x52,
	0x54, 0x5f, 0x4f, 0x52, 0x45, 0x10, 0x0b, 0x12, 0x10, 0x0a, 0x0c, 0x4e, 0x4f, 0x5f, 0x53, 0x54,
	0x41, 0x52, 0x54, 0x5f, 0x4f, 0x52, 0x45, 0x10, 0x0c, 0x12, 0x12, 0x0a, 0x0e, 0x4e, 0x4f, 0x5f,
	0x45, 0x4e, 0x4f, 0x55, 0x47, 0x48, 0x5f, 0x49, 0x54, 0x45, 0x4d, 0x10, 0x0d, 0x12, 0x13, 0x0a,
	0x0f, 0x4e, 0x4f, 0x5f, 0x4f, 0x52, 0x45, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x10, 0x0e, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x5f, 0x42, 0x41,
	0x54, 0x54, 0x4c, 0x45, 0x10, 0x0f, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x49, 0x4e, 0x5f, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x10, 0x10, 0x12, 0x13,
	0x0a, 0x0f, 0x4e, 0x4f, 0x5f, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x53, 0x43, 0x45, 0x4e,
	0x45, 0x10, 0x11, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x43, 0x45, 0x4e, 0x45, 0x5f, 0x53, 0x45, 0x52,
	0x56, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0x12, 0x2a,
	0x71, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x10, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x47, 0x53, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x53, 0x45, 0x52, 0x56, 0x45,
	0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x48, 0x41,
	0x4c, 0x4c, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x53, 0x43, 0x45, 0x4e, 0x45,
	0x10, 0x03, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_msgtype_proto_rawDescOnce sync.Once
	file_msgtype_proto_rawDescData = file_msgtype_proto_rawDesc
)

func file_msgtype_proto_rawDescGZIP() []byte {
	file_msgtype_proto_rawDescOnce.Do(func() {
		file_msgtype_proto_rawDescData = protoimpl.X.CompressGZIP(file_msgtype_proto_rawDescData)
	})
	return file_msgtype_proto_rawDescData
}

var file_msgtype_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_msgtype_proto_goTypes = []interface{}{
	(MsgId)(0),      // 0: msg.MsgId
	(GMCommand)(0),  // 1: msg.GMCommand
	(ErrCode)(0),    // 2: msg.ErrCode
	(ServerType)(0), // 3: msg.ServerType
}
var file_msgtype_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_msgtype_proto_init() }
func file_msgtype_proto_init() {
	if File_msgtype_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_msgtype_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msgtype_proto_goTypes,
		DependencyIndexes: file_msgtype_proto_depIdxs,
		EnumInfos:         file_msgtype_proto_enumTypes,
	}.Build()
	File_msgtype_proto = out.File
	file_msgtype_proto_rawDesc = nil
	file_msgtype_proto_goTypes = nil
	file_msgtype_proto_depIdxs = nil
}
