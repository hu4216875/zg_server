package service

import (
	"github.com/name5566/leaf/log"
	"net"
	"server/base"
	"server/conf"
	"server/publicconst"
	"strconv"
	"strings"
)

var (
	hallClient net.Conn
)

type InterClientService struct {
	IService
}

func NewInterClientService() *InterClientService {
	return &InterClientService{}
}

func (i *InterClientService) OnInit() {
	discover := base.NewServiceDiscover(conf.Server.EtcdServer)
	discover.ListFunc = listServer
	discover.UpdateFunc = updateServer
	discover.RemoveFunc = removeServer
	if err := discover.DiscoverService("/server"); err != nil {
		log.Error("InterClientService err:%v", err)
	}
}

func (i *InterClientService) OnDestory() {

}

func listServer(key, value string) {
	serverAddr := getServerAddr(key)
	serverId, err := strconv.Atoi(value)
	if err != nil {
		log.Error("listServer err:%v", err)
	}
	if strings.Contains(key, publicconst.BATTLE_HALL_SERVER_PREFIX) {
		if cli := NewHallConn(uint32(serverId), serverAddr); cli != nil {
			setBattleHallClient(cli)
		}
	} else if strings.Contains(key, publicconst.BATTLE_SCENE_SERVER_PREFIX) {
		if cli := NewSceneClient(uint32(serverId), serverAddr); cli != nil {
			addSceneClient(cli)
		}
	} else if strings.Contains(key, publicconst.ORE_SERVER_PREFIX) {
		if cli := NewOreConn(uint32(serverId), serverAddr); cli != nil {
			setOreClient(cli)
		}
	} else if strings.Contains(key, publicconst.MESSAGE_SERVER_PREFIX) {
		if cli := NewMessageConn(uint32(serverId), serverAddr); cli != nil {
			setMessageClient(cli)
		}
	}
	log.Debug("listServer key:%v value:%v", key, value)
}

func updateServer(key, value string) {
	serverAddr := getServerAddr(key)
	serverId, err := strconv.Atoi(value)
	if err != nil {
		log.Error("updateServer err:%v", err)
	}

	if strings.Contains(key, publicconst.BATTLE_HALL_SERVER_PREFIX) {
		if cli := NewHallConn(uint32(serverId), serverAddr); cli != nil {
			setBattleHallClient(cli)
		}
	} else if strings.Contains(key, publicconst.BATTLE_SCENE_SERVER_PREFIX) {
		if cli := NewSceneClient(uint32(serverId), serverAddr); cli != nil {
			addSceneClient(cli)
		}
	} else if strings.Contains(key, publicconst.ORE_SERVER_PREFIX) {
		if cli := NewOreConn(uint32(serverId), serverAddr); cli != nil {
			setOreClient(cli)
		}
	} else if strings.Contains(key, publicconst.MESSAGE_SERVER_PREFIX) {
		if cli := NewMessageConn(uint32(serverId), serverAddr); cli != nil {
			setMessageClient(cli)
		}
	}
	log.Debug("updateServer key:%v value:%v", key, value)
}

func removeServer(key string) {
	if strings.Contains(key, publicconst.BATTLE_HALL_SERVER_PREFIX) {
		removeBattlHallClient()
	} else if strings.Contains(key, publicconst.BATTLE_SCENE_SERVER_PREFIX) {
		removeSceneClientByAddr(getServerAddr(key))
	} else if strings.Contains(key, publicconst.ORE_SERVER_PREFIX) {
		removeOreClient()
	} else if strings.Contains(key, publicconst.MESSAGE_SERVER_PREFIX) {
		removeMessageClient()
	}
	log.Debug("removeServer key:%v", key)
}

func getServerAddr(key string) string {
	pos := strings.LastIndex(key, "/")
	if pos > 0 {
		addr := key[pos+1:]
		addr = strings.Trim(addr, " ")
		return addr
	}
	return ""
}
