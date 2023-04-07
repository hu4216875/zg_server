package service

import (
	"github.com/name5566/leaf/log"
	"server/game/internal/common"
	"server/game/internal/dao"
	"server/game/internal/model"
	"server/msg"
)

type TaskService struct {
	IService
}

func NewTaskService() *TaskService {
	return &TaskService{}
}

func (t *TaskService) OnInit() {
}

func (t *TaskService) OnDestory() {

}

func (t *TaskService) LoadTasks(accountId int64) (msg.ErrCode, []*model.Task) {
	playerData := common.PlayerMgr.FindPlayerData(accountId)
	if playerData == nil {
		log.Error("Account:%v LoadTasks player not find", accountId)
		return msg.ErrCode_SYSTEM_ERROR, nil
	}

	if len(playerData.Tasks) == 0 {
		err, accountTask := dao.TaskDao.LoadTasks(accountId)
		playerData.Tasks = accountTask.Tasks
		return err, accountTask.Tasks
	}
	return msg.ErrCode_SUCC, nil
}

func (t *TaskService) GetTaskReward(taskId uint32) msg.ErrCode {
	return msg.ErrCode_SUCC
}

func (t *TaskService) AcceptTaskByTaskType() {

}

func (t *TaskService) AcceptTaskByTaskId() {

}

func ToProtocolITask(data *model.Task) *msg.Task {
	return &msg.Task{
		TaskId:      data.TaskId,
		TaskState:   data.TaskState,
		TaskValue:   data.TaskValue,
		ComplteTime: data.UpdateTime,
	}
}
func ToProtocolITasks(datas []*model.Task) []*msg.Task {
	var ret []*msg.Task
	for i := 0; i < len(datas); i++ {
		ret = append(ret, ToProtocolITask(datas[i]))
	}
	return ret
}

func (t *TaskService) updateTask(playData *common.PlayerData, task *model.Task) {
	if task == nil {
		return
	}

	if t.getTask(playData, task.TaskId) != nil {
		dao.TaskDao.UpdateTask(playData.AccountInfo.AccountId, task)
		return
	}
}

func (t *TaskService) addTask(playData *common.PlayerData, taskId uint32) *model.Task {
	if t.getTask(playData, taskId) == nil {
		return nil
	}

	task := model.NewTask(taskId)
	playData.Tasks = append(playData.Tasks, task)
	dao.TaskDao.AddTask(playData.AccountInfo.AccountId, task)
	return task
}

func (t *TaskService) getTask(playData *common.PlayerData, taskId uint32) *model.Task {
	for i := 0; i < len(playData.Tasks); i++ {
		if playData.Tasks[i].TaskId == taskId {
			return playData.Tasks[i]
		}
	}
	return nil
}
