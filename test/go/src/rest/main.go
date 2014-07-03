package main

import (
	"fmt"

	//"os"
	"thrift/lib/go/thrift"
	"thrift/test/go/src/rest/gen-go/rest"
)

const (
	NetworkAddr = "127.0.0.1:9090"
)

type RpcServiceImpl struct {
}

func (this *RpcServiceImpl) ConfigGet(client_id string) (r string, err error) {
	fmt.Println("Received request:%v", client_id)
	return client_id, nil

}

func (this *RpcServiceImpl) AddPost(num1 int32, num2 int32) (r int32, err error) {
	fmt.Printf("Executing add:%d+%d\n", num1, num2)
	return num1 + num2, nil
}

func main() {
	//transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	//transportFactory := thrift.NewTTransportFactory()

	protocol := thrift.NewTHTTPServerProtocolTransport()
	//protocolFactory := thrift.NewTCompactProtocolFactory()

	// serverTransport, err := thrift.NewTServerSocket(NetworkAddr)
	// if err != nil {
	// 	fmt.Println("Error!", err)
	// 	os.Exit(1)
	// }

	handler := &RpcServiceImpl{}
	processor := rest.NewRestProcessor(handler)

	server := thrift.NewTHTTPServer4(NetworkAddr, processor, protocol, protocol)
	fmt.Println("thrift server in ", NetworkAddr)
	server.Serve()

}
