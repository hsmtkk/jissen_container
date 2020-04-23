package main

import (
	"log"
	"net/http"

	"github.com/hsmtkk/jissen_container/pkg/smpwebsrv"
)

func main() {
	http.HandleFunc("/", smpwebsrv.HelloDocker)
	log.Println("start server")
	server := &http.Server{Addr: ":8080"}
	log.Fatal(server.ListenAndServe())
}
