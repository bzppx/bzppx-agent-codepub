package service

import (
	"bzppx-agent-codepub/utils"
	"errors"
	"bzppx-agent-codepub/containers"
	"strconv"
	"encoding/json"
	"github.com/snail007/mini-logger"
)

type ServiceTask struct {
	gitXParams utils.GitXParams
	preCommandXParams utils.CommandXParams
	postCommandXParams utils.CommandXParams
}

var taskParams = []string{
	"task_log_id", // task_log_id
	"url", // git 仓库地址
	"ssh_key", //ssh_key
	"ssh_key_salt", // ssh_key_salt
	"path", // 代码目录
	"branch", // 发布分支或 commit_id
	"username", // 用户名
	"password", // 密码
	"pre_command", // 前置命令
	"pre_command_exec_type",// 前置命令执行方式
	"pre_command_exec_timeout", // 前置命令超时时间
	"post_command", // 后置命令
	"post_command_exec_type", // 后置命令执行方式
	"post_command_exec_timeout", // 后置命令超时时间
	"exec_user", // 执行命令用户
}

func NewServiceTask() *ServiceTask {
	return &ServiceTask{
		gitXParams: utils.GitXParams{},
		preCommandXParams: utils.CommandXParams{},
		postCommandXParams: utils.CommandXParams{},
	}
}

// 验证参数
func (t *ServiceTask) validateParams(args map[string]interface{}) error {

	for _, taskParam := range taskParams {
		if _, ok := args[taskParam]; !ok {
			return errors.New("args params "+taskParam+" requied")
		}
	}
	preCommandType, _ := strconv.Atoi(args["pre_command_exec_type"].(string))
	preCommandTimeout, _ := strconv.Atoi(args["pre_command_exec_timeout"].(string))
	postCommandType, _ := strconv.Atoi(args["post_command_exec_type"].(string))
	postCommandTimeout, _ := strconv.Atoi(args["post_command_exec_timeout"].(string))

	t.gitXParams = utils.GitXParams {
		Url: args["url"].(string),
		SshKey: args["ssh_key"].(string),
		SshKeySalt: args["ssh_key_salt"].(string),
		Path: args["path"].(string),
		Branch: args["branch"].(string),
		Username: args["username"].(string),
		Password: args["password"].(string),
	}

	t.preCommandXParams = utils.CommandXParams {
		Path: args["path"].(string),
		Command: args["pre_command"].(string),
		CommandExecType: preCommandType,
		CommandExecTimeout: preCommandTimeout,
		ExecUser: args["exec_user"].(string),
	}

	t.postCommandXParams = utils.CommandXParams {
		Path: args["path"].(string),
		Command: args["post_command"].(string),
		CommandExecType: postCommandType,
		CommandExecTimeout: postCommandTimeout,
		ExecUser: args["exec_user"].(string),
	}

	return nil
}

// 创建发布任务
func (t *ServiceTask) Publish(args map[string]interface{}, reply *string) error {
	log := containers.Log
	err := t.validateParams(args)
	if err != nil {
		log.Error("agent task service add task error: "+err.Error())
		return err
	}

	taskLogId := args["task_log_id"].(string)
	path := args["path"].(string)
	err = containers.Tasks.Add(taskLogId, path, t.gitXParams, t.preCommandXParams, t.postCommandXParams)
	if err != nil {
		log.Error("agent task service add task error: "+err.Error())
		return err
	}

	log.Info("agent task service add task "+taskLogId+" success")
	return nil
}

// 获取发布任务执行结果
func (g *ServiceTask) Status(args map[string]interface{}, reply *string) error {
	log := containers.Log
	err := g.validateParams(args)
	if err != nil {
		log.Error("agent service task status error: "+err.Error())
		return err
	}

	taskLogId := args["task_log_id"].(string)

	taskMessage, err := containers.Tasks.GetTask(taskLogId)
	if err != nil {
		log.Error("agent task service status error: "+err.Error())
		return err
	}

	resMap := map[string]string {
		"status": strconv.Itoa(taskMessage.Status),
		"is_success": strconv.Itoa(taskMessage.IsSuccess),
		"result": taskMessage.Result,
		"commit_id": taskMessage.CommitId,
	}

	resByte, _ := json.Marshal(resMap)
	*reply = string(resByte)

	log.With(logger.Fields(resMap)).Infof("agent task %s status", taskLogId)

	return nil
}

// 确认完成，删除任务记录
func (g *ServiceTask) Delete(args map[string]interface{}, reply *string) error {
	err := g.validateParams(args)
	if err != nil {
		containers.Log.Error("agent task service delete error: "+err.Error())
		return err
	}

	taskLogId := args["task_log_id"].(string)
	containers.Tasks.Delete(taskLogId)

	containers.Log.Info("agent task service detele task "+taskLogId+" success")

	return nil
}

// auto register
func init()  {
	Register(NewServiceTask())
}