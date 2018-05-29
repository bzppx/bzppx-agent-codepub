package service

import (
	"encoding/json"
	"bzppx-agent-codepub/app"
)

type ServiceSystem struct {

}

func NewServiceSystem() *ServiceSystem {
	return &ServiceSystem{}
}

// ping
func (s *ServiceSystem) Ping(args map[string]interface{}, reply *string) error {

	resMap := map[string]string {
		"version": app.Version,
	}

	resByte, _ := json.Marshal(resMap)
	*reply = string(resByte)
	return nil
}

// auto register
func init()  {
	Register(NewServiceSystem())
}