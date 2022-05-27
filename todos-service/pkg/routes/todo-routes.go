package routes

import (
	"github.com/gorilla/mux"
	"github.com/nagcloudlab/todos-service/pkg/controllers"
)

var RegisterTodoRoutes = func(router *mux.Router) {

	router.HandleFunc("/api/v1/todos", controllers.GetTodos).Methods("GET")
	router.HandleFunc("/api/v1/todos/{id}", controllers.GetTodo).Methods("GET")

}
