package service

type Base struct {

}

var RegisterServices = []interface{}{}

func Register(service interface{})  {
	RegisterServices = append(RegisterServices, service)
}
