package service

import "battleScene/game/internal/common"

var (
	hallCli *common.TcpClient
)

func RemoveHallConn() {
	hallCli = nil
}

func getBattleHallClient() *common.TcpClient {
	return hallCli
}
