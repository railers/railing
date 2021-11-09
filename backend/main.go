package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting api...")

	InitRoutes()
	StartServer()
}

func InitRoutes() {
	http.HandleFunc("/test", Testing)
}

func StartServer() {
	server := http.Server{
		Addr: "127.0.0.1:5000",
	}
	server.ListenAndServe()
}

func Testing(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
