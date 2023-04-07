package internal

import (
	"battleScene/base"
	"battleScene/conf"
	"battleScene/game/internal/common"
	"battleScene/game/internal/service"
	"fmt"
	"github.com/name5566/leaf/module"
	"strconv"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	service.ServMgr.InitService()
	registService(int64(common.MAX_SERVICE_TIME))
}

func (m *Module) OnDestroy() {

}

func registService(ttl int64) {
	endpoint := conf.Server.EtcdServerAddr
	serviceKey := common.BATTLE_SCENE_SERVER_PREFIX + conf.Server.TCPAddr
	strServerId := strconv.Itoa(int(conf.Server.ServerId))
	register := base.NewServiceRegister(endpoint, common.BATTLE_SCENE_SERVER_PREFIX, serviceKey, strServerId)
	err := register.Register(ttl)
	if err != nil {
		fmt.Printf("registService err : %v\n", err)
	}
}
