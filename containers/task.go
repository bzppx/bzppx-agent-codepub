package containers

import (
	"sync"
	"errors"
	"bzppx-agent-codepub/utils"
)

var Tasks = NewTask()

const Task_Status_Default = 0 // 任务未开始
const Task_Status_Starting = 1 // 任务开始
const Task_Status_End = 2 // 任务完成

const Task_Failed = 0 // 执行失败
const Task_Success = 1 // 执行成功

func NewTask() Task {
	return Task{
		lock: sync.Mutex{},
		TaskMessages: []*TaskMessage{},
	}
}

type Task struct {
	lock sync.Mutex
	TaskMessages []*TaskMessage
}

type TaskMessage struct {
	TaskLogId string
	Path string
	Status int
	IsSuccess int
	GitX utils.GitXParams
	Result string
}

// add a task message
func (t *Task) Add(taskLogId string, path string, gitX utils.GitXParams) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	return t.add(taskLogId, path, gitX)
}

// add task
func (t *Task) add(taskLogId string, path string, gitX utils.GitXParams) error {

	taskMessages := t.TaskMessages
	for _, taskMessage := range taskMessages {
		if taskMessage.TaskLogId == taskLogId {
			return errors.New("The task already exists")
		}
	}
	taskMsg := &TaskMessage{
		TaskLogId: taskLogId,
		Path: path,
		Status: Task_Status_Default,
		IsSuccess: Task_Failed,
		GitX: gitX,
		Result: "",
	}

	t.TaskMessages = append(t.TaskMessages, taskMsg)
	return nil
}

// delete a task
func (t *Task) Delete(taskLogId string) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	return t.del(taskLogId)
}

// delete task
func (t *Task) del(taskLogId string) error {

	taskMessages := []*TaskMessage{}
	for _, taskMessage := range t.TaskMessages {
		if taskMessage.TaskLogId == taskLogId {
			continue
		}
		taskMessages = append(taskMessages, taskMessage)
	}

	t.TaskMessages = taskMessages
	return nil
}

// task start
func (t *Task) Start(taskLogId string) (err error) {
	t.lock.Lock()
	defer t.lock.Unlock()
	
	taskMessage, err := t.GetTask(taskLogId)
	if err != nil {
		return
	}
	
	newTaskMessage := &TaskMessage{
		TaskLogId: taskLogId,
		Path: taskMessage.Path,
		Status: Task_Status_Starting,
		IsSuccess: taskMessage.IsSuccess,
		GitX: taskMessage.GitX,
		Result: taskMessage.Result,
	}
	
	return t.update(newTaskMessage)
}

// task start
func (t *Task) End(taskLogId string, isSuccess int, Result string) (err error) {
	t.lock.Lock()
	defer t.lock.Unlock()
	
	taskMessage, err := t.GetTask(taskLogId)
	if err != nil {
		return
	}
	
	newTaskMessage := &TaskMessage{
		TaskLogId: taskLogId,
		Path: taskMessage.Path,
		Status:Task_Status_End,
		IsSuccess: isSuccess,
		GitX: taskMessage.GitX,
		Result: Result,
	}
	
	return t.update(newTaskMessage)
}

// get a task
func (t *Task) GetTask(taskLogId string) (*TaskMessage, error)  {
	taskMessages := t.TaskMessages
	for _, taskMessage := range taskMessages {
		if taskMessage.TaskLogId == taskLogId {
			return taskMessage, nil
		}
	}
	
	return nil, errors.New("The task not exists")
}

// update by taskLogId
func (t *Task) update(task *TaskMessage) error {

	isExists := false
	for _, taskMessage := range t.TaskMessages {
		if taskMessage.TaskLogId == task.TaskLogId {
			isExists = true
			taskMessage.Path = task.Path
			taskMessage.Status = task.Status
			taskMessage.IsSuccess = task.IsSuccess
			taskMessage.GitX = task.GitX
			taskMessage.Result = task.Result
		}
	}
	if !isExists {
		return errors.New("The task not exists")
	}

	return nil
}

// path is have staring task
func (t *Task) PathIsHaveTask(path string) bool {
	taskMessages := t.TaskMessages
	for _, taskMessage := range taskMessages {
		if (taskMessage.Path == path) &&
			(taskMessage.Status == Task_Status_Starting){
			return true
		}
	}
	return false
}

// get a task
func (t *Task) GetDefaultTasks() ([]*TaskMessage)  {
	taskMessages := []*TaskMessage{}
	for _, taskMessage := range t.TaskMessages {
		if taskMessage.Status == Task_Status_Default {
			taskMessages = append(taskMessages, taskMessage)
		}
	}
	return taskMessages
}