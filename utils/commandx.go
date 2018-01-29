package utils

import (
	"io/ioutil"
	"os/exec"
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
	cmd := exec.Command("/bin/bash", "./"+fileName)
	err = cmd.Run()
	return
}

// 异步执行
func (c *CommandX) asyExec(commandXParams CommandXParams) (err error) {
	fileName, err := c.createTmpShellFile(commandXParams.Path, commandXParams.Command, false)
	if err != nil {
		return
	}
	cmd := exec.Command("/bin/bash", "./"+fileName)
	err = cmd.Start()
	return
}

// 创建临时的 shell 脚本文件
// path 创建脚本目录
// content 创建的脚本内容
// isErrorStop 是否遇到错误停止
func (c *CommandX) createTmpShellFile(path string, content string, isErrorStop bool) (tmpFile string, err error) {
	file, err := ioutil.TempFile(path, "tmp")
	if err != nil {
		return
	}
	defer file.Close()

	file.Chmod(777)
	file.WriteString("#!/bin/bash\r\n")
	if isErrorStop {
		// 遇到错误继续执行
		file.WriteString("set -e\r\n")
	}
	file.WriteString("cd "+path+" \r\n")
	_, err = file.WriteString(content)
	if err != nil {
		return
	}
	return file.Name(), nil
}