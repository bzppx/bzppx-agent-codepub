package service

type ServiceGit struct {

}

func NewServiceGit() *ServiceGit {
	return &ServiceGit{}
}

// 获取发布代码的 commit id
func (git *ServiceGit) GetCommitId(args map[string]interface{}, reply *interface{}) error {
	return nil
}

// 发布代码操作
func (git *ServiceGit) Publish(args map[string]interface{}, reply *interface{}) error {
	return nil
}

// 获取发布执行结果
func (git *ServiceGit) Status(args map[string]interface{}, reply *interface{}) error {
	return nil
}

// auto register
func init()  {
	Register(NewServiceGit())
}