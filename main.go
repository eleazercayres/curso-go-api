package main

import (
	"fmt"
	"github.com/filipefernandes007/curso-go-api/api/handlers"
	"net/http"
)

const (
	HttpPort = "9000"
)

func main() {
	CreateRoutes()

	fmt.Println("Start http server at port " + HttpPort)
	err := http.ListenAndServe(fmt.Sprintf(":%s", HttpPort), nil)
	if err!= nil {
		panic(err)
	}
}

func CreateRoutes() {
	http.HandleFunc("/api/healthy", handlers.HealthyHandler)
	http.HandleFunc("/api/create-client", handlers.CreateClientHandler)
	http.HandleFunc("/api/list-client", handlers.ListClientHandler)
	http.HandleFunc("/api/get-client", handlers.GetClientHandler)

}
