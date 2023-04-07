package service

import (
	"battleHall/msg"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"sync"
)

type SerrverInfo struct {
	ServerId   uint32
	Agent      gate.Agent
	ServerType msg.ServerType
}

var (
	serverMapLock sync.RWMutex
	serverMap     = make(map[uint32]*SerrverInfo)
)

func AddServer(serverId uint32, agent gate.Agent, servType, onlineLimit uint32) {
	serverMapLock.Lock()
	defer serverMapLock.Unlock()
	serverMap[serverId] = &SerrverInfo{
		ServerId:   serverId,
		Agent:      agent,
		ServerType: msg.ServerType(servType),
	}

	if servType == uint32(msg.ServerType_SERVER_TYPE_BATTLE_SCENE) {
		addSceneServerData(serverId, onlineLimit)
	}

	log.Debug("AddServer serverId:%v", serverId)
}

func RemoveServer(serverId uint32) {
	serverInfo := FindServer(serverId)
	if serverInfo != nil {
		// gs 掉线
		if serverInfo.ServerType == msg.ServerType_SERVER_TYPE_GS {
			onGsClose(serverId)
		} else if serverInfo.ServerType == msg.ServerType_SERVER_TYPE_BATTLE_SCENE {
			onSceneClose(serverId)
		}
	}
	serverMapLock.Lock()
	defer serverMapLock.Unlock()
	delete(serverMap, serverId)
	log.Debug("RemoveServer serverId:%v", serverId)
}

func FindServer(serverId uint32) *SerrverInfo {
	serverMapLock.RLock()
	defer serverMapLock.RUnlock()
	if data, ok := serverMap[serverId]; ok {
		return data
	}
	return nil
}

func getAllSceneServer() []uint32 {
	serverMapLock.RLock()
	defer serverMapLock.RUnlock()

	var ret []uint32
	for sceneServerId, serverInfo := range serverMap {
		if serverInfo.ServerType == msg.ServerType_SERVER_TYPE_BATTLE_SCENE {
			ret = append(ret, sceneServerId)
		}
	}
	return ret
}

func onGsClose(gsId uint32) {
	removeSceneServerGsData(gsId)
}

func onSceneClose(sceneId uint32) {
	removeSceneServerData(sceneId)
}
