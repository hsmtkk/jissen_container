package smpwebsrv

import (
	"fmt"
	"log"
	"net/http"
)

func HelloDocker(w http.ResponseWriter, r *http.Request) {
	log.Println("received request")
	fmt.Fprintf(w, "Hello Docker!")
}
