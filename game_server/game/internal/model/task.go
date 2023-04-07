package model

import "server/util"

type AccountTask struct {
	AccountId             int64
	NextDailyRefreshTime  uint32
	NextWeeklyRefreshTime uint32
	DailyTasks            []*Task
	WeeklyTasks           []*Task
	ChallengeTasks        []*Task
}

type Task struct {
	TaskId     uint32
	TaskValue  uint32
	TaskState  uint32
	CreateTime uint32
	UpdateTime uint32
}

func NewTask(taskId uint32) *Task {
	curTime := util.GetCurTime()
	return &Task{
		TaskId:     taskId,
		CreateTime: curTime,
		UpdateTime: curTime,
	}
}

func NewAccountTask(accountId int64) *AccountTask {
	ret := &AccountTask{
		AccountId: accountId,
	}
	ret.DailyTasks = make([]*Task, 0, 0)
	ret.WeeklyTasks = make([]*Task, 0, 0)
	ret.ChallengeTasks = make([]*Task, 0, 0)
	return ret
}
