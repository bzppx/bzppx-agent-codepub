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
			go func() {
				defer func() {
					e := recover()
					if e != nil {
						Log.Error(e)
						Tasks.End(task.TaskLogId, Task_Failed, "goroutine runtime error", "")
					}
				}()
				// start publish code
				commitId, err := utils.NewGitX().Publish(task.GitX)
				if err != nil {
					Log.Error("agent task "+task.TaskLogId+" publish faild: "+err.Error())
					Tasks.End(task.TaskLogId, Task_Failed, err.Error(), commitId)
				}else {
					Log.Info("agent task "+task.TaskLogId+" publish success, commit_id: "+commitId)
					Tasks.End(task.TaskLogId, Task_Success, "success", commitId)
				}
			}()
		}

		time.Sleep(2 * time.Second)
	}
}