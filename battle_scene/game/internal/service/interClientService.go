package service

import (
	"battleScene/base"
	"battleScene/conf"
	"battleScene/game/internal/common"
	"github.com/name5566/leaf/log"
	"net"
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
	discover := base.NewServiceDiscover(conf.Server.EtcdServerAddr)
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
	if strings.Contains(key, common.BATTLE_HALL_SERVER_PREFIX) {
		serverAddr := getServerAddr(key)
		serverId, err := strconv.Atoi(value)
		if err != nil {
			log.Error("listServer err:%v", err)
		}
		NewHallConn(uint32(serverId), serverAddr)
	}
}

func updateServer(key, value string) {
	if strings.Contains(key, common.BATTLE_HALL_SERVER_PREFIX) {
		serverAddr := getServerAddr(key)
		serverId, err := strconv.Atoi(value)
		if err != nil {
			log.Error("listServer err:%v", err)
		}
		NewHallConn(uint32(serverId), serverAddr)
	}
}

func removeServer(key string) {
	if strings.Contains(key, common.BATTLE_HALL_SERVER_PREFIX) {
		RemoveHallConn()
	}
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
