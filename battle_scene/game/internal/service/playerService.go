package service

import (
	"sync"
)

type PlayerData struct {
	AccountId int64
	GsId      uint32
}

type PlayerService struct {
	playerMapLock sync.RWMutex
	playerMap     map[int64]*PlayerData
}

func NewPlayerService() *PlayerService {
	return &PlayerService{}
}

func (p *PlayerService) OnInit() {
	p.playerMap = make(map[int64]*PlayerData)
}

func (p *PlayerService) OnDestory() {

}

func (p *PlayerService) GetPlayerData(accountId int64) *PlayerData {
	p.playerMapLock.RLock()
	defer p.playerMapLock.RUnlock()

	if data, ok := p.playerMap[accountId]; ok {
		return data
	}
	return nil
}

func (p *PlayerService) AddPlayerData(data *PlayerData) {
	p.playerMapLock.Lock()
	defer p.playerMapLock.Unlock()
	p.playerMap[data.AccountId] = data
}

func (p *PlayerService) RemovePlayerData(accountId int64) {
	p.playerMapLock.Lock()
	defer p.playerMapLock.Unlock()
	delete(p.playerMap, accountId)
}

func (p *PlayerService) RemoveGsPlayer(gsId uint32) {
	p.playerMapLock.Lock()
	defer p.playerMapLock.Unlock()
	for key, data := range p.playerMap {
		if data.GsId == gsId {
			delete(p.playerMap, key)
		}
	}
}
