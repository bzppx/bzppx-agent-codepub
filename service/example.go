package service

// 一个 service 示例

type Example struct {

}

func NewExample() *Example {
	return &Example{}
}

// 相乘
func (math *Example) Multiply(args map[string]interface{}, reply *int) error {
	*reply = args["a"].(int) * args["b"].(int)
	return nil
}

// auto register
func init()  {
	Register(NewExample())
}