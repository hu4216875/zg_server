package template

import "sync"

type Task struct {
	Id       uint32
	TaskType uint32
	Value0   uint32
}

type TaskTemplate struct {
	IDataTempalte
	mutex sync.RWMutex
	data  map[uint32]*Task
}

func NewTaskTemplate() *TaskTemplate {
	return &TaskTemplate{}
}

func (t *TaskTemplate) OnLoad() {
	t.data = make(map[uint32]*Task)
	rf := readRf(Task{})
	for k := 0; k < rf.NumRecord(); k++ {
		task := rf.Record(k).(*Task)
		t.data[task.Id] = task
	}
}

func (t *TaskTemplate) OnInit() {

}

func (t *TaskTemplate) OnReload() {
	data := make(map[uint32]*Task)
	rf := readRf(Item{})
	for k := 0; k < rf.NumRecord(); k++ {
		task := rf.Record(k).(*Task)
		data[task.Id] = task
	}

	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.data = data
	t.OnInit()
}

func (t *TaskTemplate) OnDestory() {

}

// GetTask 获取任务
func (t *TaskTemplate) GetTask(id uint32) *Task {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	if ret, ok := t.data[id]; ok {
		return ret
	}
	return nil
}

func (t *TaskTemplate) GetTaskByType(taskType uint32) []*Task {
	t.mutex.RLock()
	t.mutex.RUnlock()
	var ret []*Task
	for _, data := range t.data {
		if data.TaskType == taskType {
			ret = append(ret, data)
		}
	}
	return ret
}
