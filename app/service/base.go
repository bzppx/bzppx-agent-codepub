package service

import "net/rpc"

type BaseService struct {

}

var RegisterServices = []interface{}{}

func Register(service interface{})  {
	RegisterServices = append(RegisterServices, service)
}

func RegisterRpc()  {
	for _, ser := range RegisterServices {
		rpc.Register(ser)
	}
}