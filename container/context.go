package container

import (
	"bzppx-agent-codepub/message"
)

var Ctx = NewContext()

func NewContext() *Context {
	return &Context{
		Tasks: message.NewTask(),
	}
}

type Context struct {

	// task
	Tasks *message.Task
}