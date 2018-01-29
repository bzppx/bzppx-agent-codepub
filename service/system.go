package service

type ServiceSystem struct {

}

func NewServiceSystem() *ServiceSystem {
	return &ServiceSystem{}
}

// ping
func (s *ServiceSystem) Ping(args map[string]interface{}, reply *string) error {
	*reply = "ok"
	return nil
}

// auto register
func init()  {
	Register(NewServiceSystem())
}