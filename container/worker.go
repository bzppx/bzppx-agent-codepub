package container

import (
	"bzppx-agent-codepub/app"
	"bzppx-agent-codepub/utils"
	"time"
	"bzppx-agent-codepub/message"
)

func NewWorker() *Worker {
	return &Worker{}
}

type Worker struct {}

func (w *Worker) StartTask() {
	for {
		tasks := Ctx.Tasks.GetDefaultTasks()
		if len(tasks) == 0 {
			time.Sleep(2 * time.Second)
			continue
		}
		for _, task := range tasks {
			pathIsHave := Ctx.Tasks.PathIsHaveTask(task.Path)
			if pathIsHave {
				continue
			}
			err := Ctx.Tasks.Start(task.TaskLogId)
			if err != nil {
				app.Log.Error(err.Error())
				continue
			}
			go func(task *message.TaskMessage) {
				defer func() {
					e := recover()
					if e != nil {
						app.Log.Errorf("task handle crash, %v", e)
						Ctx.Tasks.End(task.TaskLogId, message.Task_Failed, "task runtime crash", "")
					}
				}()

				app.Log.Info("agent task "+task.TaskLogId+" publish start")

				// start exec pre_command
				if task.PreCommandX.Command != "" {
					err := utils.NewCommandX().Exec(task.PreCommandX)
					if err != nil {
						app.Log.Error("agent task "+task.TaskLogId+" exec pre_command faild: "+err.Error())
						if task.PreCommandX.CommandExecType == utils.Command_ExecType_SyncErrorStop {
							Ctx.Tasks.End(task.TaskLogId, message.Task_Failed, "exec pre_command error: "+err.Error(), "")
							return
						}
					} else {
						app.Log.Info("agent task "+task.TaskLogId+" exec pre_command success")
					}
				}

				// start publish code
				commitId, err := utils.NewGitX().Publish(task.GitX)
				if err != nil {
					app.Log.Error("agent task "+task.TaskLogId+" publish faild: "+err.Error())
					Ctx.Tasks.End(task.TaskLogId, message.Task_Failed, "publish code error: "+err.Error(), commitId)
					return
				}
				app.Log.Info("agent task "+task.TaskLogId+" publish code success, commit_id: "+commitId)

				// start exec post_command
				if task.PostCommandX.Command != "" {
					err = utils.NewCommandX().Exec(task.PostCommandX)
					if err != nil {
						app.Log.Error("agent task "+task.TaskLogId+" exec post_command faild: "+err.Error())
						if task.PostCommandX.CommandExecType == utils.Command_ExecType_SyncErrorStop {
							Ctx.Tasks.End(task.TaskLogId, message.Task_Failed, "exec post_command error: "+err.Error(), "")
							return
						}
					}else {
						app.Log.Info("agent task "+task.TaskLogId+" exec post_command success")
					}
				}

				Ctx.Tasks.End(task.TaskLogId, message.Task_Success, "success", commitId)

				app.Log.Info("agent task "+task.TaskLogId+" publish end")
			}(task)
		}

		time.Sleep(2 * time.Second)
	}
}