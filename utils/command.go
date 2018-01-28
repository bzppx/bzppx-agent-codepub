package utils

import (
	"io/ioutil"
)

type Command struct {
	
}

// 同步执行
func (c *Command) SyncExec(command string, timeout int, user string) error {
	return nil
}

// 异步执行
func (c *Command) AsyExec(command string, timeout int, user string) error {
	return nil
}

// 创建临时的 shell 脚本文件
func (c *Command) createTmpShellFile(path string, content string) (tmpFile string, err error) {
	file, err := ioutil.TempFile(path, "tmp")
	defer file.Close()
	if err != nil {
		return
	}
	_, err = file.WriteString(content)
	if err != nil {
		return
	}
	return file.Name(), nil
}