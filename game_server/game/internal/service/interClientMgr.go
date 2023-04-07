package service

import (
	"server/game/internal/common"
	"sync"
)

var (
	hallCli *common.TcpClient

	sceneCliMapLock sync.RWMutex
	sceneCliMap     = make(map[uint32]*common.TcpClient)

	oreCli *common.TcpClient

	messageCli *common.TcpClient
)

func removeBattlHallClient() {
	hallCli = nil
}

func setBattleHallClient(cli *common.TcpClient) {
	hallCli = cli
}

func getBattleHallClient() *common.TcpClient {
	return hallCli
}

func removeOreClient() {
	oreCli = nil
}

func setOreClient(cli *common.TcpClient) {
	oreCli = cli
}

func getOreClient() *common.TcpClient {
	return oreCli
}

func addSceneClient(cli *common.TcpClient) {
	sceneCliMapLock.Lock()
	defer sceneCliMapLock.Unlock()

	sceneCliMap[cli.GetServerId()] = cli
}

func removeSceneClient(sceneId uint32) {
	sceneCliMapLock.Lock()
	defer sceneCliMapLock.Unlock()
	if _, ok := sceneCliMap[sceneId]; ok {
		delete(sceneCliMap, sceneId)
	}
}

func removeSceneClientByAddr(sceneAddr string) {
	sceneCliMapLock.Lock()
	defer sceneCliMapLock.Unlock()

	for serverId, data := range sceneCliMap {
		if data.GetServerAddr() == sceneAddr {
			delete(sceneCliMap, serverId)
			break
		}
	}
}

func findSceneClient(sceneId uint32) *common.TcpClient {
	sceneCliMapLock.RLock()
	defer sceneCliMapLock.RUnlock()
	if data, ok := sceneCliMap[sceneId]; ok {
		return data
	}
	return nil
}

func removeMessageClient() {
	messageCli = nil
}

func setMessageClient(cli *common.TcpClient) {
	messageCli = cli
}

func getMessageClient() *common.TcpClient {
	return messageCli
}
