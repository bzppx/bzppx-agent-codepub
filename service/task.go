package service

import (
	"bzppx-agent-codepub/utils"
	"errors"
	"bzppx-agent-codepub/containers"
	"strconv"
)

type ServiceTask struct {

}

func NewServiceTask() *ServiceTask {
	return &ServiceTask{}
}

// 验证参数
func (t *ServiceTask) validateParams(args map[string]interface{}) (gitX utils.GitXParams, err error) {
	if _, ok := args["task_id"]; !ok {
		return gitX, errors.New("args params task_id requied")
	}
	if _, ok := args["url"]; !ok {
		return gitX, errors.New("args params url requied")
	}
	if _, ok := args["ssh_key"]; !ok {
		return gitX, errors.New("args params ssh_key requied")
	}
	if _, ok := args["ssh_key_salt"]; !ok {
		return gitX, errors.New("args params ssh_key_salt requied")
	}
	if _, ok := args["path"]; !ok {
		return gitX, errors.New("args params path requied")
	}
	if _, ok := args["branch"]; !ok {
		return gitX, errors.New("args params branch requied")
	}
	if _, ok := args["username"]; !ok {
		return gitX, errors.New("args params username requied")
	}
	if _, ok := args["password"]; !ok {
		return gitX, errors.New("args params password requied")
	}

	return utils.GitXParams {
		Url: args["url"].(string),
		SshKey: args["ssh_key"].(string),
		SshKeySalt: args["ssh_key_salt"].(string),
		Path: args["path"].(string),
		Branch: args["branch"].(string),
		Username: args["username"].(string),
		Password: args["password"].(string),
	}, nil
}

// 创建发布任务
func (t *ServiceTask) Publish(args map[string]interface{}, reply *interface{}) error {
	gitParams, err := t.validateParams(args)
	if err != nil {
		return err
	}

	taskId := args["task_id"].(string)
	path := args["path"].(string)
	err = containers.Tasks.Add(taskId, path, gitParams)
	if err != nil {
		return err
	}

	return nil
}

// 获取发布任务执行结果
func (g *ServiceTask) Status(args map[string]interface{}, reply *interface{}) error {
	_, err := g.validateParams(args)
	if err != nil {
		return err
	}

	taskId := args["task_id"].(string)

	taskMessage, err := containers.Tasks.GetTask(taskId)
	if err != nil {
		return err
	}

	*reply = map[string]string{
		"status": strconv.Itoa(taskMessage.Status),
		"is_success": strconv.Itoa(taskMessage.IsSuccess),
		"result": taskMessage.Result,
	}

	return nil
}

// 确认完成，删除任务记录
func (g *ServiceTask) Delete(args map[string]interface{}, reply *interface{}) error {
	_, err := g.validateParams(args)
	if err != nil {
		return err
	}

	taskId := args["task_id"].(string)
	containers.Tasks.Delete(taskId)

	return nil
}

// auto register
func init()  {
	Register(NewServiceTask())
}