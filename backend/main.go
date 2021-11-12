package main

import (
	"fmt"
	"net/http"

	"github.com/railers/railing/routes"
	"github.com/railers/railing/stops"
)

func main() {
	fmt.Println("Starting api...")

	InitRoutes()
	go stops.GetStops()
	go routes.GetRoutes()
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
