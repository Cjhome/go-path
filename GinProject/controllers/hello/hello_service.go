package hello

import "fmt"

type HelloService struct {

}

// (h *HelloService) 定义结构的方法
func (h *HelloService)SayHello(username string) string {
	return fmt.Sprintf("Hello, %s", username)
}

// 新服务
type GreeterService struct {

}

func (h *GreeterService)SayHello(username string) string {
	return fmt.Sprintf("hello, %s, how are you!", username)
}