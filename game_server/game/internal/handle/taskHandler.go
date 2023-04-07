package handle

import (
	"server/game/internal/common"
	"server/game/internal/service"
	"server/msg"
)

// RequestLoadTasksHandle 请求加载任务
func RequestLoadTasksHandle(args interface{}, playerData *common.PlayerData) {
	err, tasks := service.ServMgr.GetTaskService().LoadTasks(playerData.AccountInfo.AccountId)
	retMsg := &msg.ResponseLoadTasks{
		Result: int32(err),
	}
	if len(tasks) > 0 {
		retMsg.Tasks = append(retMsg.Tasks, service.ToProtocolITasks(tasks)...)
	}
	playerData.PlayerAgent.WriteMsg(retMsg)
}
