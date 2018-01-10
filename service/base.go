package service

type BaseService struct {

}

var RegisterServices = []interface{}{}

func Register(service interface{})  {
	RegisterServices = append(RegisterServices, service)
}
