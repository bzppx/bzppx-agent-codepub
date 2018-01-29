package utils

import (
	"io/ioutil"
	"os/exec"
	"os"
)

func NewCommandX() *CommandX {
	return &CommandX{}
}

type CommandX struct {

}

const Command_ExecType_SyncErrorStop = 1; // 同步执行，遇到错误停止
const Command_ExecType_SyncErrorAccess = 2; // 同步执行，遇到错误继续
const Command_ExecType_Asy = 3; // 异步执行

type CommandXParams struct {
	Path string
	Command string
	CommandExecType int
	CommandExecTimeout int
	ExecUser string
}

// 执行命令
func (c *CommandX) Exec(commandXParams CommandXParams) error {
	if commandXParams.Command == "" {
		return nil
	}
	if commandXParams.CommandExecType == Command_ExecType_Asy {
		c.asyExec(commandXParams)
	}else {
		c.syncExec(commandXParams)
	}
	return nil
}

// 同步执行
func (c *CommandX) syncExec(commandXParams CommandXParams) (err error) {

	fileName := ""
	if commandXParams.CommandExecType == Command_ExecType_SyncErrorStop {
		fileName, err = c.createTmpShellFile(commandXParams.Path, commandXParams.Command, true)
	} else {
		fileName, err = c.createTmpShellFile(commandXParams.Path, commandXParams.Command, false)
	}
	if err != nil {
		return
	}
	cmd := exec.Command("/bin/bash", fileName)

	err = cmd.Run()
	return
}

// 异步执行
func (c *CommandX) asyExec(commandXParams CommandXParams) (err error) {
	fileName, err := c.createTmpShellFile(commandXParams.Path, commandXParams.Command, false)
	if err != nil {
		return
	}
	cmd := exec.Command("/bin/bash", fileName)

	err = cmd.Start()
	return
}

// 创建临时的 shell 脚本文件
// path 脚本执行目录
// content 创建的脚本内容
// isErrorStop 是否遇到错误停止
func (c *CommandX) createTmpShellFile(path string, content string, isErrorStop bool) (tmpFile string, err error) {

	ok, err := NewFile().PathIsExists("/tmp/.codepub")
	if err != nil {
		return
	}
	if ok == false {
		os.Mkdir("/tmp/.codepub", 0777)
	}
	file, err := ioutil.TempFile("/tmp/.codepub", "tmp")
	if err != nil {
		return
	}
	defer file.Close()

	file.Chmod(0777)
	file.WriteString("#!/bin/bash\n")
	if isErrorStop {
		// 遇到错误继续执行
		file.WriteString("set -e\n")
	}
	file.WriteString("cd "+path+" \n")
	_, err = file.WriteString(content)
	if err != nil {
		return
	}

	return file.Name(), nil
}