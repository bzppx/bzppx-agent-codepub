package containers

import (
	"bzppx-agent-codepub/utils"
	"time"
)

var Workers = Worker{}

type Worker struct {

}

func (w *Worker) Task() {
	for {
		tasks := Tasks.GetDefaultTasks()
		if len(tasks) == 0 {
			time.Sleep(2 * time.Second)
			continue
		}
		for _, task := range tasks {
			pathIsHave := Tasks.PathIsHaveTask(task.Path)
			if pathIsHave {
				continue
			}
			err := Tasks.Start(task.TaskLogId)
			if err != nil {
				Log.Error(err)
				continue
			}
			go func(task *TaskMessage) {
				defer func() {
					e := recover()
					if e != nil {
						Log.Error(e)
						Tasks.End(task.TaskLogId, Task_Failed, "goroutine runtime error", "")
					}
				}()

				Log.Info("agent task "+task.TaskLogId+" publish start")

				// start exec pre_command
				if task.PreCommandX.Command != "" {
					err := utils.NewCommandX().Exec(task.PreCommandX)
					if err != nil {
						Log.Error("agent task "+task.TaskLogId+" exec pre_command faild: "+err.Error())
						if task.PreCommandX.CommandExecType == utils.Command_ExecType_SyncErrorStop {
							Tasks.End(task.TaskLogId, Task_Failed, "exec pre_command error: "+err.Error(), "")
							return
						}
					} else {
						Log.Info("agent task "+task.TaskLogId+" exec pre_command success")
					}
				}

				// start publish code
				commitId, err := utils.NewGitX().Publish(task.GitX)
				if err != nil {
					Log.Error("agent task "+task.TaskLogId+" publish faild: "+err.Error())
					Tasks.End(task.TaskLogId, Task_Failed, "publish code error: "+err.Error(), commitId)
					return
				}
				Log.Info("agent task "+task.TaskLogId+" publish code success, commit_id: "+commitId)

				// start exec post_command
				if task.PostCommandX.Command != "" {
					err = utils.NewCommandX().Exec(task.PostCommandX)
					if err != nil {
						Log.Error("agent task "+task.TaskLogId+" exec post_command faild: "+err.Error())
						if task.PostCommandX.CommandExecType == utils.Command_ExecType_SyncErrorStop {
							Tasks.End(task.TaskLogId, Task_Failed, "exec post_command error: "+err.Error(), "")
							return
						}
					}else {
						Log.Info("agent task "+task.TaskLogId+" exec post_command success")
					}
				}

				Tasks.End(task.TaskLogId, Task_Success, "success", commitId)

				Log.Info("agent task "+task.TaskLogId+" publish end")
			}(task)
		}

		time.Sleep(2 * time.Second)
	}
}