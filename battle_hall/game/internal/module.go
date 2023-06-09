package internal

import (
	"battleHall/base"
	"battleHall/game/internal/service"
	"github.com/name5566/leaf/module"
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
	service.InitHallService()
	service.InitHttpService()
}

func (m *Module) OnDestroy() {

}
