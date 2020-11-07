package main

import (
	"github.com/micro/go-micro/v2/web"
	"net/http"
)

func main() {
	server:=web.NewService(web.Address(":8000"))
	server.HandleFunc("/", func(write http.ResponseWriter, request *http.Request){
		write.Write([]byte("hello World"))
	})
	server.Run()
}
