package service

import (
	"battleScene/game/internal/common"
)

var (
	ServMgr ServiceMgr
)

type IService interface {
	OnInit()
	OnDestory()
}

type ServiceMgr struct {
	serviceMap map[common.ServiceId]IService
}

func (m *ServiceMgr) InitService() {
	m.serviceMap = make(map[common.ServiceId]IService)
	m.registService(common.InterClientSerivce, NewInterClientService())
	m.registService(common.SceneService, NewSceneService())
	m.registService(common.PlayerServgice, NewPlayerService())
	m.registService(common.ServerService, NewServerService())

	for _, service := range m.serviceMap {
		service.OnInit()
	}
}

func (m *ServiceMgr) Destory() {
	for _, service := range m.serviceMap {
		service.OnDestory()
	}
}

func (m *ServiceMgr) registService(serviceId common.ServiceId, service IService) {
	m.serviceMap[serviceId] = service
}

func (m *ServiceMgr) GetSceneService() *SceneService {
	if service, ok := m.serviceMap[common.SceneService]; ok {
		return service.(*SceneService)
	}
	return nil
}

func (m *ServiceMgr) GetPlayerService() *PlayerService {
	if service, ok := m.serviceMap[common.PlayerServgice]; ok {
		return service.(*PlayerService)
	}
	return nil
}

func (m *ServiceMgr) GetServerService() *ServerService {
	if service, ok := m.serviceMap[common.ServerService]; ok {
		return service.(*ServerService)
	}
	return nil
}
