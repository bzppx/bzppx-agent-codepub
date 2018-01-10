package containers

import "bzppx-agent-codepub/utils"

var Workers = Worker{}

type Worker struct {

}

func (w *Worker) Task() {
	for {
		tasks := Tasks.GetDefaultTasks()
		if len(tasks) == 0 {
			return
		}
		for _, task := range tasks {
			pathIsHave := Tasks.PathIsHaveTask(task.Path)
			if pathIsHave {
				continue
			}
			go func() {
				// start publish code
				err := utils.NewGitX().Publish(&task.GitX)
				if err != nil {
					Tasks.End(task.TaskId, Task_Failed)
				}else {
					Tasks.End(task.TaskId, Task_Success)
				}
			}()
		}
	}
}

func init()  {
	go Workers.Task()
}
