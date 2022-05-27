package main

import (
	"github.com/gorilla/mux"
	"github.com/nagcloudlab/todos-service/pkg/routes"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	routes.RegisterTodoRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:3000", r))

}
