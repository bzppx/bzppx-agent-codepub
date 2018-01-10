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

const Task_Failed = 0 // 执行成功
const Task_Success = 1 // 执行失败

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
	TaskId string
	Path string
	Status int
	IsSuccess int
	GitX utils.GitXParams
	Result string
}

// add a task message
func (t *Task) Add(taskId string, path string, gitX utils.GitXParams) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	return t.add(taskId, path, gitX)
}

// add task
func (t *Task) add(taskId string, path string, gitX utils.GitXParams) error {

	taskMessages := t.TaskMessages
	for _, taskMessage := range taskMessages {
		if taskMessage.TaskId == taskId {
			return errors.New("The task already exists")
		}
	}
	taskMsg := &TaskMessage{
		TaskId: taskId,
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
func (t *Task) Delete(taskId string) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	return t.del(taskId)
}

// delete task
func (t *Task) del(taskId string) error {

	taskMessages := []*TaskMessage{}
	for _, taskMessage := range t.TaskMessages {
		if taskMessage.TaskId == taskId {
			continue
		}
		taskMessages = append(taskMessages, taskMessage)
	}

	t.TaskMessages = taskMessages
	return nil
}

// task start
func (t *Task) Start(taskId string) (err error) {
	t.lock.Lock()
	defer t.lock.Unlock()
	
	taskMessage, err := t.GetTask(taskId)
	if err != nil {
		return
	}
	
	newTaskMessage := &TaskMessage{
		TaskId: taskId,
		Path: taskMessage.Path,
		Status: Task_Status_Starting,
		IsSuccess: taskMessage.IsSuccess,
		GitX: taskMessage.GitX,
		Result: taskMessage.Result,
	}
	
	return t.update(newTaskMessage)
}

// task start
func (t *Task) End(taskId string, isSuccess int) (err error) {
	t.lock.Lock()
	defer t.lock.Unlock()
	
	taskMessage, err := t.GetTask(taskId)
	if err != nil {
		return
	}
	
	newTaskMessage := &TaskMessage{
		TaskId: taskId,
		Path: taskMessage.Path,
		Status:Task_Status_End,
		IsSuccess: isSuccess,
		GitX: taskMessage.GitX,
		Result: taskMessage.Result,
	}
	
	return t.update(newTaskMessage)
}

// get a task
func (t *Task) GetTask(taskId string) (*TaskMessage, error)  {
	taskMessages := t.TaskMessages
	for _, taskMessage := range taskMessages {
		if taskMessage.TaskId == taskId {
			return taskMessage, nil
		}
	}
	
	return nil, errors.New("The task not exists")
}

// update by taskId
func (t *Task) update(task *TaskMessage) error {

	isExists := false
	taskMessages := []*TaskMessage{}
	for _, taskMessage := range t.TaskMessages {
		if taskMessage.TaskId == task.TaskId {
			isExists = true
			taskMessages = append(taskMessages, task)
		}else {
			taskMessages = append(taskMessages, taskMessage)
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