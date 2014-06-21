package main

import (
	"fmt"
	"github.com/drone/routes"
	"net/http"
	"strconv"

	//"os"
	"rest/gen-go/rest"
	"thrift/lib/go/thrift"
)

const (
	NetworkAddr = "127.0.0.1:9090"
)

type RpcServiceImpl struct {
}

func (this *RpcServiceImpl) Config(client_id string) (r string, err error) {
	fmt.Println("Received request:%v", client_id)
	return client_id, nil

}

func (this *RpcServiceImpl) Add(num1 int32, num2 int32) (r int32, err error) {
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

func config(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got request:%v\n", r)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Success\n"))
}

func add(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got request:%v\n", r)

	num1 := r.FormValue("num1")
	fmt.Printf("got num1:%s\n", num1)
	num2 := r.FormValue("num2")
	fmt.Printf("got num2:%s\n", num2)
	i1, _ := strconv.Atoi(num1)
	i2, _ := strconv.Atoi(num2)
	i3 := i1 + i2

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(fmt.Sprintf("%d", i3)))
}

func main2() {
	mux := routes.New()
	mux.Post("/add", add)
	mux.Post("/config", config)
	http.Handle("/", mux)
	http.ListenAndServe(":9090", nil)

}
