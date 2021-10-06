package main

import (
	"github.com/polarismesh/polaris-go/api"
	"io"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}
func main() {

	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":md"+
		"12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	provider, err := api.NewProviderAPI()
	if nil != err {
		log.Fatal(err)
	}
	defer provider.Destroy()

	request := &api.InstanceRegisterRequest{}
	request.Namespace = "default"
	request.Service = "dummy"
	request.ServiceToken = "token"
	request.Host = "192.168.25.105"
	request.Port = 8091
	request.SetTTL(2)
	resp, err := provider.Register(request)
	if nil != err {
		log.Fatalf("fail to register instance, err %v", err)
	}
	log.Printf("success to register instance, id is %s", resp.InstanceID)

}
