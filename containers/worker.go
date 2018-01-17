package containers

import (
	"time"
	"bzppx-agent-codepub/utils"
	"log"
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
			go func() {
				defer func() {
					e := recover()
					if e != nil {
						log.Println(e)
					}
				}()
				// start publish code
				err := utils.NewGitX().Publish(task.GitX)
				if err != nil {
					Tasks.End(task.TaskLogId, Task_Failed, err.Error())
				}else {
					Tasks.End(task.TaskLogId, Task_Success, "success")
				}
			}()
		}
		time.Sleep(2 * time.Second)
	}
}

func init()  {
	go Workers.Task()
}