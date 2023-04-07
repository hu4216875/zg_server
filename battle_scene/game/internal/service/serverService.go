package service

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"sync"
)

var (
	serverMapLock sync.RWMutex
	serverMap     map[uint32]*ServerInfo
)

type ServerInfo struct {
	ServerId uint32
	Agent    gate.Agent
}

type ServerService struct {
	IService
}

func NewServerService() *ServerService {
	return &ServerService{}
}

func (s *ServerService) OnInit() {
	serverMap = make(map[uint32]*ServerInfo)
}

func (s *ServerService) OnDestory() {

}

func (s *ServerService) GetServer(serverId uint32) *ServerInfo {
	serverMapLock.RLock()
	defer serverMapLock.RUnlock()

	if data, ok := serverMap[serverId]; ok {
		return data
	}
	return nil
}

func (s *ServerService) AddServer(serverId uint32, agent gate.Agent) {
	serverMapLock.Lock()
	defer serverMapLock.Unlock()

	serverMap[serverId] = &ServerInfo{
		ServerId: serverId,
		Agent:    agent,
	}
	log.Debug("AddServer serverId:%v", serverId)
}

func (s *ServerService) RemoveServer(serverId uint32) {
	serverMapLock.Lock()
	defer serverMapLock.Unlock()
	delete(serverMap, serverId)

	log.Debug("RemoveServer serverId:%v", serverId)
}
