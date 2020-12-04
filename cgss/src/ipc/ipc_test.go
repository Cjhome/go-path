package ipc

import (
	"testing"
)

type EchoServer struct {
	Server
}

func (server *EchoServer)Handle(request string) string {
	return "ECHO:" + request
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})
	client1 := NewIpClient(server)
	client2 := NewIpClient(server)

	resp1, _ := client1.Call("method1", "params1")
	resp2, _ := client2.Call("method2", "params2")

	if resp1.Body != "ECHO:From Client1" || resp2.Body != "ECHO:From Client2" {
		t.Error("IpcClient.Call failed, resp1:", resp1, "resp2:", resp2)
	}
	client1.Close()
	client2.Close()
}
