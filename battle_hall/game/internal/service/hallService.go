package service

import (
	"battleHall/base"
	"battleHall/common"
	"battleHall/conf"
	"battleHall/game/internal/dao"
	"battleHall/model"
	"github.com/name5566/leaf/log"
	"strconv"
	"sync"
)

var (
	sceneDataMapLock sync.RWMutex
	sceneDataMap     map[uint32]*model.ServerInfo
)

type sceneServer struct {
	serverId  uint32
	remainNum uint32
}

type sceneServerSlice struct {
	Data []*sceneServer
	By   func(q, p *sceneServer) bool
}

func InitHallService() {
	sceneDataMap = make(map[uint32]*model.ServerInfo)
	registService(int64(common.MAX_SERVICE_TIME))
	loadServerInfo()
}

func registService(ttl int64) {
	endpoint := conf.Server.EtcdServerAddr
	serviceKey := common.BATTLE_HALL_SERVER_PREFIX + conf.Server.TCPAddr
	strServerId := strconv.Itoa(int(conf.Server.ServerId))
	register := base.NewServiceRegister(endpoint, common.BATTLE_HALL_SERVER_PREFIX, serviceKey, strServerId)
	err := register.Register(ttl)
	if err != nil {
		log.Debug("registService %v\n", err)
	}
}

func loadServerInfo() {
	lst := dao.ServerInfoDao.LoadAll()
	for i := 0; i < len(lst); i++ {
		lst[i].Online = make([]*model.GsServerOnline, 0, 0)
		sceneDataMap[lst[i].ServerId] = lst[i]
	}
}

func addSceneServerData(sceneSrvId, onlineLimit uint32) {
	sceneDataMapLock.Lock()
	defer sceneDataMapLock.Unlock()
	if data, ok := sceneDataMap[sceneSrvId]; ok {
		data.Online = data.Online[0:0]
	} else {
		sceneDataMap[sceneSrvId] = model.NewServerInfo(sceneSrvId, onlineLimit)
	}
}

func removeSceneServerData(sceneId uint32) {
	sceneDataMapLock.Lock()
	defer sceneDataMapLock.Unlock()
	delete(sceneDataMap, sceneId)
}

func removeSceneServerGsData(gsId uint32) {
	sceneDataMapLock.Lock()
	defer sceneDataMapLock.Unlock()

	for _, sceneData := range sceneDataMap {
		for k := 0; k < len(sceneData.Online); k++ {
			if sceneData.Online[k].GsId == gsId {
				sceneData.Online[k].OnlineNum = 0
			}
		}
	}
}

func getAllScene() []*model.ServerInfo {
	var ret []*model.ServerInfo
	serverIds := getAllSceneServer()
	for i := 0; i < len(serverIds); i++ {
		if data := getSceneData(serverIds[i]); data != nil {
			ret = append(ret, data)
		}
	}
	return ret
}

func updateSceneLimit(serverId, onlineLimit uint32) {
	sceneDataMapLock.Lock()
	defer sceneDataMapLock.Unlock()

	if data, ok := sceneDataMap[serverId]; ok {
		data.OnlineLimit = onlineLimit
		dao.ServerInfoDao.InsertOrUpdateServerInfo(data)
	}
}

func getSceneData(serverId uint32) *model.ServerInfo {
	sceneDataMapLock.RLock()
	defer sceneDataMapLock.RUnlock()
	if data, ok := sceneDataMap[serverId]; ok {
		return data
	}
	return nil
}

func UpdateSceneOnlineNum(sceneId, gsId uint32, delta int32) {
	sceneDataMapLock.Lock()
	defer sceneDataMapLock.Unlock()
	if data, ok := sceneDataMap[sceneId]; ok {
		var findGs = false
		for i := 0; i < len(data.Online); i++ {
			if data.Online[i].GsId == gsId {
				data.Online[i].OnlineNum += delta
				if data.Online[i].OnlineNum < 0 {
					data.Online[i].OnlineNum = 0
				}
				findGs = true
				break
			}
		}

		if !findGs && delta > 0 {
			data.Online = append(data.Online, model.NewGsServerOnlie(gsId))
		}
	}
}

func SelectOneSceneServer() uint32 {
	sceneServers := getAllSceneServer()
	sceneDataMapLock.RLock()
	defer sceneDataMapLock.RUnlock()

	var servers []*sceneServer

	for i := 0; i < len(sceneServers); i++ {
		sceneServerId := sceneServers[i]
		if sceneData, ok := sceneDataMap[sceneServerId]; ok {
			var num int32 = 0
			for k := 0; k < len(sceneData.Online); k++ {
				num += sceneData.Online[k].OnlineNum
			}

			if num < int32(sceneData.OnlineLimit) {
				servers = append(servers, &sceneServer{
					serverId:  sceneServerId,
					remainNum: sceneData.OnlineLimit - uint32(num),
				})
			}
		}
	}

	if len(servers) == 0 {
		return 0
	}

	return servers[0].serverId
}

func (d sceneServerSlice) Len() int {
	return len(d.Data)
}

func (d sceneServerSlice) Swap(i, j int) {
	d.Data[i], d.Data[j] = d.Data[j], d.Data[i]
}

func (d sceneServerSlice) Less(i, j int) bool {
	return d.By(d.Data[i], d.Data[j])
}
